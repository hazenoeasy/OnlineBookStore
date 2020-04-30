package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type SellerShowOrderService struct {
    UserHeader
    Page    int     `header:"page"  binding:"required"`
    Items   int     `header:"items" binding:"required"`
    Status  int     `header:"status" binding:"required"`
}

func (this *SellerShowOrderService) ShowOrder() serializer.Response {
    // 查询卖家的所有购物订单
    var orders []model.Order
    if err := model.DB.Select("order_id,status,title,price,amount,express_id,recv_info_id").
        Where("salesman_id = ? AND status = ?", this.UserId, this.Status).
        Order("order_time DESC").Offset((this.Page-1)*this.Items).Limit(this.Items).
        Find(&orders).Error; err != nil {
        return serializer.Response{
            Code: serializer.DBReadErr,
            Data: err,
            Msg:  "卖家查看订单：数据库查询订单失败",
        }
    }
    // 查询买家地址信息
    var receivers []int
    for i := 0; i < len(orders); i++ {
        receivers = append(receivers, orders[i].RecvInfoId)
    }
    var addresses []model.RecvInfo
    if err := model.DB.Select("receiver_name,address,phone").
        Where("recv_info_id IN (?)", receivers).Find(&addresses).Error;
        err != nil {
        return serializer.Response{
            Code: serializer.DBReadErr,
            Data: err,
            Msg:  "卖家查看订单：数据库查询买家地址信息失败",
        }
    }
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: serializer.NewRespSellerShowOrder(orders, addresses),
        Msg:  "ok",
    }
}