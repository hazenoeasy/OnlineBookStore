package serializer

import "DuckyGo/model"

type ShoppingCart struct {
    Items   int             `json:"items"`
    Item    []CartCommodity `json:"item"`
}

type CartCommodity struct {
    Id      int     `json:"id"`
    Name    string  `json:"name"`
    Price   int     `json:"price"`
    Cover   string  `json:"cover"`
    Num     int     `json:"num"`
}

func NewCartResponse(commodities []model.ShoppingCart) (resp ShoppingCart) {
    resp.Items  =   len(commodities)
    resp.Item   =   make([]CartCommodity, resp.Items)
    for i := range commodities {
        resp.Item[i].Id     =   commodities[i].BookId
        resp.Item[i].Name   =   commodities[i].Name
        resp.Item[i].Price  =   commodities[i].Price
        resp.Item[i].Cover  =   commodities[i].CoverUrl
        resp.Item[i].Num    =   commodities[i].Amount
    }
    return
}