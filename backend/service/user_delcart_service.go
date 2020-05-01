package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type BodyDelCartService struct {
    BookId  int     `form:"book_id" json:"book_id" binding:"required"`
}

type DelCartService struct {
    Header  UserHeader
    Body    BodyDelCartService
}

func (this *DelCartService) DelCart() serializer.Response {
    commodity := model.ShoppingCart{
        UserId: this.Header.UserId,
        BookId: this.Body.BookId,
    }
    if err := model.DB.Delete(&commodity).Error; err != nil {
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: err,
            Msg:  "数据库：购物车商品删除失败",
        }
    }
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: nil,
        Msg:  "ok",
    }
}
