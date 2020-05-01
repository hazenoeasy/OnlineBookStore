package service

import (
    "DuckyGo/cache/hotRank"
    "DuckyGo/model"
    "DuckyGo/serializer"
    "fmt"
    "strconv"
    "time"
)

type OrderCreateService struct {
    Header  UserHeader
    Body    BodyOrderCreateService
}

type BodyOrderCreateService struct {
    AddrId  int     `form:"addr_id" json:"addr_id" binding:"required"`
    Id      int     `form:"id" json:"id" binding:"required"`
    Num     int     `form:"num" json:"num" binding:"required"`
}

func (this *OrderCreateService) CreateOrder() serializer.Response {
    // 1、检查每件商品库存
    // 2、如果有存量则减少购买量，否则减库存行为失败，全部回滚，终止
    // 3、创建订单，并记录购买商品

    tx := model.DB.Begin()

    // 先查询商品库存
    book := model.Book{BookId: this.Body.Id}
    if err := tx.Select("book_id,salesman_id,title,author,kind,num,price,salesnum,cover_url,descp_url").
        Find(&book).Error; err != nil {
            tx.Rollback()
        return serializer.Response{
            Code: serializer.DBReadErr,
            Data: err,
            Msg:  "创建订单：查询商品库存失败",
        }
    }

    // 检查存量是否满足条件，如果满足则更新
    //  不满足则回滚
        buyNum  :=  this.Body.Num
        if book.Num < buyNum {
            // 库存不够了
            tx.Rollback()
            return serializer.Response{
                Code: serializer.ResourceEmptied,
                Data: fmt.Errorf("书籍id：%d，库存：%d，购买量：%d", book.BookId, book.Num, buyNum),
                Msg:  "库存不够",
            }
        } else {
            // 更新数据
            // FIXME: GORM只更新非零值，如果num==buynum（书卖完），则忽视更新
            data := map[string]interface{}{"num":book.Num-buyNum,"salesnum":book.SalesNum+buyNum}
            if err := model.DB.Model(&book).Updates(data).Error; err != nil {
                tx.Rollback()
                return serializer.Response{
                    Code: serializer.DBWriteErr,
                    Data: err,
                    Msg:  "创建订单：商品库存销量更新失败",
                }
            }
        }

    tx.Commit()

    tx = model.DB.Begin()

    // 创建订单，记录购买商品
    order := model.Order{
        UserId:      this.Header.UserId,
        SalesmanId:  book.SalesManId,
        BookId:      book.BookId,
        RecvInfoId:  this.Body.AddrId,
        OrderTime:   time.Now(),
        Status:      serializer.OrderUnpaid,
        Title:       book.Title,
        Price:       book.Price,
        Amount:      buyNum,
    }
    if err := tx.Create(&order).Error; err != nil {
        tx.Rollback()
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: err,
            Msg:  "创建订单：数据库插入订单失败",
        }
    }

    tx.Commit()

    // 更新热榜服务
    if hotRank.HOT.IsExisted(book.BookId) {
        hotRank.HOT.Update(strconv.Itoa(book.BookId), float64(buyNum))
    } else {
        hotRank.HOT.Add(&book)
    }

    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: nil,
        Msg:  "ok",
    }
}