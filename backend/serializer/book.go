package serializer

import (
    "DuckyGo/model"
)

type ShowBookData struct {
    Id          int     `form:"id" json:"id" binding:"required"`
    Title       string  `form:"title" json:"title" binding:"required"`
    Author      string  `form:"author" json:"author" binding:"required"`
    Kind        string  `form:"kind" json:"kind" binding:"required"`
    Cover       string  `form:"cover" json:"cover" binding:"required"`
    Price       int     `form:"price" json:"price" binding:"required"`
    SalesNum    int     `form:"salesnum" json:"sales_num" binding:"required"`
    Num         int     `form:"num" json:"num" binding:"required"`
    Descp       string  `form:"descp" json:"descp" binding:"required"`
}

type ShowBookRespData struct {
    Pages   int     `form:"pages" json:"pages" binding:"required"`
    Items   int     `form:"items" json:"items" binding:"required"`
    Item    []ShowBookData    `json:"item" json:"item" binding:"required"`
}

func NewShowBookRespData(books []model.Book, pages int) ShowBookRespData {
    rtv := make([]ShowBookData, len(books))
    for i := 0; i < len(books); i++ {
        rtv[i].Id   =   books[i].BookId
        rtv[i].Title    =   books[i].Title
        rtv[i].Author   =   books[i].Author
        rtv[i].Kind     =   books[i].Kind
        rtv[i].Cover    =   books[i].CoverUrl
        rtv[i].Price    =   books[i].Price
        rtv[i].SalesNum =   books[i].SalesNum
        rtv[i].Num      =   books[i].Num
        rtv[i].Descp    =   books[i].DescpUrl
    }
    return ShowBookRespData{
        Pages: pages,
        Items: len(rtv),
        Item:  rtv,
    }
}