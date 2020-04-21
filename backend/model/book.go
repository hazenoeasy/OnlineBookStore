package model

// Book 书籍信息
type Book struct {
    BookId      int     `gorm:"column:book_id;primary_key;auto_increment"`
    SalesManId  int     `gorm:"column:salesman_id;foreignkey:salesman_id;association_foreignkey:user_id;not null"`
    Title       string  `gorm:"column:title;type:varchar(128);not null;index:idx_title"`
    Author      string  `gorm:"column:author;type:varchar(64);not null"`
    Price       int     `gorm:"column:price;not null"`
    Num         int     `gorm:"column:num;not null"`
    SalesNum    int     `gorm:"column:salesnum;default:0;index:idx_salesnum"`
    Kind        string  `gorm:"column:kind;type:varchar(32);index:idx_kind"`
    CoverUrl    string  `gorm:"column:cover_url;type:varchar(128);not null"`
    DescpUrl    string  `gorm:"column:descp_url;type:varchar(128);not null"`
}
