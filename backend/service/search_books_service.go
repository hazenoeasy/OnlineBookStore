package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type SearchBooksService struct {
    Name    string  `form:"name" json:"name" binding:"required"`
    Page    int     `form:"page" json:"page" binding:"required"`
    Items   int     `form:"items" json:"items" binding:"required"`
}

// XXX：此函数内部使用了两次全文搜索（索引失效），因此性能较差，有待优化
func (this *SearchBooksService) Search() serializer.Response {
    var (
    	books   []model.Book
    	num     int
    )
    if err := model.DB.Select("book_id,title,author,kind,num,cover_url,price,salesnum,descp_url").
        Where("title LIKE ?", "%" + this.Name + "%").
        Order("salesnum DESC").Offset((this.Page - 1) * this.Items).
        Limit(this.Items).Find(&books).Error; err != nil {
        return serializer.Response{
            Code: serializer.DBReadErr,
            Data: err,
            Msg:  "数据库：数据搜索错误",
        }
    }
    // 由于GORM的count必须在链式操作的最后一个操作，所以我们单独计算数量
    model.DB.Model(&model.Book{}).Where("title LIKE ?", "%"+this.Name+"%").Count(&num)
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: serializer.NewSearchResult(books, (num + this.Items - 1) / this.Items),
        Msg:  "ok",
    }
}