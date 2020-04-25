package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
    "errors"
    "fmt"
    "time"
)

type OrderCreateService struct {
    Header  UserHeader
    Body    BodyOrderCreateService
}

type BodyOrderCreateService struct {
    AddrId          int             `json:"addr_id"`
    CartItems       int             `json:"cart_items"`
    CartItem        []OrderCreateItem `json:"cart_item"`
}

type OrderCreateItem struct {
    Id      int     `json:"id"`
    Num     int     `json:"num"`
}

func (this *OrderCreateService) CreateOrder() serializer.Response {
    // getAmountById 用来通过@books的book_id去寻找@this.Body.CartItem的商品购买数量
    getAmountById := func() func(int)int {
        _m := make(map[int]int)
        for _, e := range this.Body.CartItem {
            _m[e.Id] = e.Num
        }
        return func(id int) int {
            if _, existed := _m[id]; existed {
                return _m[id]
            } else {
                // never coming here
                panic(errors.New("创建订单：前后端图书id无法对应"))
            }
        }
    }()

    // 1、检查每件商品库存
    // 2、如果有存量则减少购买量，否则减库存行为失败，全部回滚，终止
    // 3、创建订单，并记录购买商品

    tx := model.DB.Begin()

    // 先查询商品库存
    var books   []model.Book
    arrayOfBookid :=  make([]int, this.Body.CartItems)
    for i, e := range this.Body.CartItem {
        arrayOfBookid[i] = e.Id
    }
    if err := tx.Select("book_id , title, price, num, salesnum").Where("book_id IN (?)", arrayOfBookid).
        Find(&books).Error; err != nil {
            tx.Rollback()
        return serializer.Response{
            Code: serializer.DBReadErr,
            Data: err,
            Msg:  "创建订单：查询商品库存失败",
        }
    }

    totalPrice := 0 // 计算@books中的总价格
    // 检查存量是否满足条件，如果满足则更新
    //  不满足则回滚
    for _, e := range books {
        buyNum  :=  getAmountById(e.BookId)
        totalPrice += e.Price * buyNum
        if e.Num < buyNum {
            // 库存不够了
            tx.Rollback()
            return serializer.Response{
                Code: serializer.ResourceEmptied,
                Data: fmt.Errorf("书籍id：%d，库存：%d，购买量：%d", e.BookId, e.Num, getAmountById(e.BookId)),
                Msg:  "库存不够",
            }
        } else {
            // 更新数据
            if err := model.DB.Model(&e).Updates(model.Book{Num:e.Num - buyNum, SalesNum:e.SalesNum + buyNum}).Error;
                err != nil {
                tx.Rollback()
                return serializer.Response{
                    Code: serializer.DBWriteErr,
                    Data: err,
                    Msg:  "创建订单：商品库存销量更新失败",
                }
            }
        }
    }

    tx.Commit()

    tx = model.DB.Begin()

    // 创建订单，记录购买商品
    order := model.Order{
        UserId:      this.Header.UserId,
        RecvInfoId:  this.Body.AddrId,
        OrderTime:   time.Now(),
        Status:      serializer.OrderUnpaid,
        TotalPrice:  totalPrice,
    }
    if err := tx.Create(&order).Error; err != nil {
        tx.Rollback()
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: err,
            Msg:  "创建订单：数据库插入订单失败",
        }
    }
    // FIXME：此处假设在create后，@order的主键order_id被写入了
    for _, e := range books {
        commodity := model.OrderCommodity{
            OrderId: order.OrderId,
            BookId:  e.BookId,
            Title:   e.Title,
            Price:   e.Price,
            Amount:  getAmountById(e.BookId),
        }
        if err := tx.Create(&commodity).Error; err != nil {
            tx.Rollback()
            return serializer.Response{
                Code: serializer.DBWriteErr,
                Data: err,
                Msg:  "创建订单：订单关联的商品记录创建失败",
            }
        }
    }

    tx.Commit()

    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: nil,
        Msg:  "ok",
    }
}