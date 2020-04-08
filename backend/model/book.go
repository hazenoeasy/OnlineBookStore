package model

// Book 书籍信息
type Book struct {
    BookId  int     `gorm:"column:book_id;primary_key;auto_increment"`
    Title   string  `gorm:"column:title;type:varchar(128);not null"`
    Author  string  `gorm:"column:author;type:varchar(64);not null"`
    Price   int     `gorm:"column:price;not null"`
    SaleNum int     `gorm:"column:salesnum;default:0"`
    PicUrl  string  `gorm:"column:pic_url;type:varchar(64);not null"`
    Descpt  string  `gorm:"column:description;type:varchar(64);not null"` // 书籍详细信息页的URL
}
