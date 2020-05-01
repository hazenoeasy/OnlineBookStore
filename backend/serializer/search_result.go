package serializer

import "DuckyGo/model"

type SearchResultItem struct {
    Id          int         `json:"id"`
    Title       string      `json:"title"`
    Author      string      `json:"author"`
    Cover       string      `json:"cover"`
    Price       int         `json:"price"`
    SalesNum    int         `json:"salesnum"`
    Descp       string      `json:"descp"`
}

type SearchResult struct {
    Pages   int                 `json:"pages"`
    Items   int                 `json:"items"`
    Item    []SearchResultItem  `json:"item"`
}

func NewSearchResult(books []model.Book, pages int) SearchResult {
    item := make([]SearchResultItem, len(books))
    for i := 0; i < len(books); i++ {
        item[i].Id      =   books[i].BookId
        item[i].Title   =   books[i].Title
        item[i].Author  =   books[i].Author
        item[i].Cover   =   books[i].CoverUrl
        item[i].Price   =   books[i].Price
        item[i].SalesNum=   books[i].SalesNum
        item[i].Descp   =   books[i].DescpUrl
    }
    return SearchResult{
        Pages: pages,
        Items: len(item),
        Item:  item,
    }
}