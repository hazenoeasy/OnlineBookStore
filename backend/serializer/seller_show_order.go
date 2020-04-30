package serializer

import "DuckyGo/model"

type RespSellerShowOrder struct {
    Items   int                     `json:"items"`
    Item    []RespSellerOrderItem   `json:"item"`
}

type RespSellerOrderItem struct {
    Id      int     `json:"id"`
    Status  int     `json:"status"`
    Title   string  `json:"title"`
    Price   int     `json:"price"`
    Num     int     `json:"num"`
    Express string  `json:"express"`
    ConsumerName    string  `json:"consumer_name"`
    ConsumerAddr    string  `json:"consumer_addr"`
    ConsumerPhone   string  `json:"consumer_phone"`
}

func NewRespSellerShowOrder(orders []model.Order, infos []model.RecvInfo) RespSellerShowOrder {
    resp := RespSellerShowOrder{
        Items: len(orders),
        Item:  make([]RespSellerOrderItem, len(orders)),
    }
    for i := 0; i < len(orders); i++ {
        resp.Item[i].Id         =   orders[i].OrderId
        resp.Item[i].Status     =   orders[i].Status
        resp.Item[i].Title      =   orders[i].Title
        resp.Item[i].Price      =   orders[i].Price
        resp.Item[i].Num        =   orders[i].Amount
        resp.Item[i].Express    =   orders[i].ExpId
        resp.Item[i].ConsumerName   =   infos[i].RecverName
        resp.Item[i].ConsumerAddr   =   infos[i].Address
        resp.Item[i].ConsumerPhone  =   infos[i].Phone
    }
    return resp
}