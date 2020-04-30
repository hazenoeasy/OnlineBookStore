package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type OrderShowService struct {
    UserHeader
    Page    int     `header:"page"  binding:"required"`
    Items   int     `header:"items" binding:"required"`
}

func (this *OrderShowService) ShowOrder() serializer.Response {
    // 查询用用户的所有购物订单
    var orders []model.Order
    if err := model.DB.Select("order_id,total_price,status").
        Where("user_id = ?", this.UserId).Order("order_time DESC").
        Offset((this.Page-1)*this.Items).Limit(this.Items).
        Find(&orders).Error; err != nil {
            return serializer.Response{
                Code: serializer.DBReadErr,
                Data: err,
                Msg:  "查看我的购物订单：数据库查询订单失败",
            }
    }

    // 查找具体购买的书籍
    var orderIds []int
    for i := range orders {
        orderIds = append(orderIds, orders[i].OrderId)
    }
    var books []model.OrderCommodity
    if err := model.DB.Select("order_id,title,amount").
        Where("order_id IN (?)", orderIds).Find(&books).Error; err != nil {
        return serializer.Response{
            Code: serializer.DBReadErr,
            Data: err,
            Msg:  "查看我的购物订单：数据库查询相关购物书籍失败",
        }
    }
    var orderBooks = make(map[int][]model.OrderCommodity)
    // 将所有订单的书籍，按照订单号分类，相同订单号的放在一起
    for _, e := range books {
        orderBooks[e.OrderId] = append(orderBooks[e.OrderId], e)
    }

    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: serializer.NewRespOrderShow(orders, orderBooks),
        Msg:  "ok",
    }
}
