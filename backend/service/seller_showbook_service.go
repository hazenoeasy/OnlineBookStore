package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type SellerShowBookService struct {
    UserHeader
    Page    int     `header:"page" binding:"required"`
    Items   int     `header:"items" binding:"required"`
}

func (p *SellerShowBookService) ShowBook() serializer.Response {
    var (
    	books []model.Book
    	num int
    )
    model.DB.Model(&model.Book{}).Where("salesman_id = ?", p.UserId).Count(&num)
    if err := model.DB.Where("salesman_id = ?", p.UserId).Order("salesnum DESC").
        Offset((p.Page - 1) * p.Items).Limit(p.Items).Find(&books).Error; err == nil {
        return serializer.Response{
            Code: serializer.OpSuccess,
            Data: serializer.NewShowBookRespData(books, (num + p.Items - 1) / p.Items),
            Msg:  "ok",
        }
    } else {
        return serializer.Response{
            Code: serializer.DBReadErr,
            Data: err.Error(),
            Msg:  "数据库：图书查询失败",
        }
    }
}