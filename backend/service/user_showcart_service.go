package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type ShowCartService struct {
    UserHeader
}

func (this *ShowCartService) ShowCart() serializer.Response {
    var commodities []model.ShoppingCart
    if err := model.DB.Where("user_id = ?", this.UserId).Find(&commodities).Error;
        err != nil {
            return serializer.Response{
                Code: serializer.DBReadErr,
                Data: err,
                Msg:  "数据库：购物车数据查询错误",
            }
    }
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: serializer.NewCartResponse(commodities),
        Msg:  "ok",
    }
}
