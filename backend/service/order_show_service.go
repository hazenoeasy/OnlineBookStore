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
    if err := model.DB.Select("order_id,status,title,price,amount,express_id").
        Where("user_id = ?", this.UserId).Order("order_time DESC").
        Offset((this.Page-1)*this.Items).Limit(this.Items).
        Find(&orders).Error; err != nil {
            return serializer.Response{
                Code: serializer.DBReadErr,
                Data: err,
                Msg:  "查看我的购物订单：数据库查询订单失败",
            }
    }
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: serializer.NewRespOrderShow(orders),
        Msg:  "ok",
    }
}
