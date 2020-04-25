package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type PayOrderService struct {
    Header  UserHeader
    Body    BodyPayOrderService
}

type BodyPayOrderService struct {
    OrderId     int     `form:"order_id" json:"order_id" binding:"required"`
}

func (this *PayOrderService) PayOrder() serializer.Response {
    order := model.Order{OrderId: this.Body.OrderId}
    if err := model.DB.Model(&order).Update("status", serializer.OrderUndelivered).Error;
        err != nil {
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: err,
            Msg:  "PayOrder(): 修改订单状态失败",
        }
    }
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: nil,
        Msg:  "ok",
    }
}