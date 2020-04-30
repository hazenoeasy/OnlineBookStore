package serializer

import "DuckyGo/model"

type RespOrderShowService struct {
    Items   int             `json:"items"`
    Item    []RespOrderItem `json:"item"`
}

type RespOrderItem struct {
    Id      int     `json:"id"`
    Price   int     `json:"price"`
    Books   []RespOrderBookItem   `json:"books"`
    Status  int     `json:"status"`
}

type RespOrderBookItem struct {
    Title   string  `json:"title"`
    Num     int     `json:"num"`
}

// order：多个订单
// books：<订单号，订单购买的所有书籍>
func NewRespOrderShow(order []model.Order, books map[int][]model.OrderCommodity) RespOrderShowService {
    resp := RespOrderShowService{
        Items: len(order),
        Item:  make([]RespOrderItem, len(order)),
    }
    for i := range resp.Item {
        orderId := order[i].OrderId
        resp.Item[i].Id     = orderId
        resp.Item[i].Price  = order[i].TotalPrice
        resp.Item[i].Status = order[i].Status
        resp.Item[i].Books  = make([]RespOrderBookItem, len(books[orderId]))
        for j := range resp.Item[i].Books {
            resp.Item[i].Books[j].Title = books[orderId][j].Title
            resp.Item[i].Books[j].Num   = books[orderId][j].Amount
        }
    }
    return resp
}