package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type GetCommodityService struct {
    Header  UserHeader
    Body    BodyGetCommodityService
}

type BodyGetCommodityService struct {
    OrderId     int     `form:"order_id" json:"order_id" binding:"required"`
}

func (this *GetCommodityService) OrderGetCommodity() serializer.Response {
    order := model.Order{OrderId: this.Body.OrderId}
    if err := model.DB.Model(&order).Update("status", serializer.OrderClosed).Error;
        err != nil {
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: err,
            Msg:  "完成订单收货: 修改订单状态失败",
        }
    }
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: nil,
        Msg:  "ok",
    }
}