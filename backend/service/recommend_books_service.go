package service

import (
    "DuckyGo/cache"
    "DuckyGo/cache/hotRank"
    "DuckyGo/serializer"
    "os"
)

type RecommendBooksService struct {
    Page    int     `header:"page" binding:"required"`
    Items   int     `header:"items" binding:"required"`
}

func (this *RecommendBooksService) Recommend() serializer.Response {
    // 查看排行榜中的书籍bookID
    var (
        start   = int64((this.Page - 1) * this.Items)
        end     = int64(this.Page * this.Items - 1)
    )
    books := hotRank.HOT.View(start, end)
    // XXX: 随着数据的增大，一般情况下，items===100
    //      此处求items的值，只是为了当测试数据不足的时候让前端可以正确渲染页数
    items := int(cache.RedisClient.ZCard(os.Getenv("HOT_RANK_NAME")).Val())
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: serializer.NewSearchResult(books, (items+this.Items-1)/this.Items),
        Msg:  "ok",
    }
}
