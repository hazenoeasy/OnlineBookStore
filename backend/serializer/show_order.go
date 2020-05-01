package serializer

import "DuckyGo/model"

type RespOrderShowService struct {
    Items   int             `json:"items"`
    Item    []RespOrderItem `json:"item"`
}

type RespOrderItem struct {
    Id      int     `json:"id"`
    Status  int     `json:"status"`
    Title   string  `json:"title"`
    Price   int     `json:"price"`
    Num     int     `json:"num"`
    Express string  `json:"express"`
}

// order：多个订单
// books：<订单号，订单购买的所有书籍>
func NewRespOrderShow(order []model.Order) RespOrderShowService {
    resp := RespOrderShowService{
        Items: len(order),
        Item:  make([]RespOrderItem, len(order)),
    }
    for i := 0; i < len(order); i++ {
        resp.Item[i].Id     =   order[i].OrderId
        resp.Item[i].Status =   order[i].Status
        resp.Item[i].Title  =   order[i].Title
        resp.Item[i].Price  =   order[i].Price
        resp.Item[i].Num    =   order[i].Amount
        resp.Item[i].Express=   order[i].ExpId
    }
    return resp
}