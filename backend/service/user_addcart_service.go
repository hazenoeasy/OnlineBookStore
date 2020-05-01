package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type BodyAddCartService struct {
    Id      int     `form:"id" json:"id" binding:"required"`
    Name    string  `form:"name" json:"name" binding:"required"`
    Price   int     `form:"price" json:"price" binding:"required"`
    Cover   string  `form:"cover" json:"cover" binding:"required"`
    Num     int     `form:"num" json:"num" binding:"required"`
}

type UserAddCartService struct {
    Header  UserHeader
    Body    BodyAddCartService
}

func (this *UserAddCartService) AddCart() serializer.Response {
    shoppingCart := model.ShoppingCart{
        UserId:   this.Header.UserId,
        BookId:   this.Body.Id,
        Name:     this.Body.Name,
        CoverUrl: this.Body.Cover,
        Price:    this.Body.Price,
        Amount:   this.Body.Num,
    } 
    if err := model.DB.Create(&shoppingCart).Error; err != nil {
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: err,
            Msg:  "数据库：购物车数据插入失败",
        }
    }
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: nil,
        Msg:  "ok",
    }
}
