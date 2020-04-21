package serializer

import (
    "DuckyGo/model"
)

type ShowBookData struct {
    Id          int     `json:"id"`
    Title       string  `json:"title"`
    Author      string  `json:"author"`
    Kind        string  `json:"kind"`
    Cover       string  `json:"cover"`
    Price       int     `json:"price"`
    SalesNum    int     `json:"sales_num"`
    Num         int     `json:"num"`
    Descp       string  `json:"descp"`
}

type ShowBookRespData struct {
    Pages   int             `json:"pages""`
    Items   int             `json:"items"`
    Item    []ShowBookData  `json:"item"`
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