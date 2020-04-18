package model

// ShoppingCart 一条购物车记录
type ShoppingCart struct {
    UserId      int     `gorm:"column:user_id;foreignkey:user_id;association_foreignkey:user_id;not null"`
    BookId      int     `gorm:"column:book_id;foreignkey:book_id;association_foreignkey:book_id;not null"`
    Name        string  `gorm:"column:name;type:varchar(128);not null"`
    CoverUrl    string  `gorm:"column:cover_url;type:varchar(128)"`
    Price       int     `gorm:"column:price;not null"`
    Amount      int     `gorm:"column:amount;not null"`
}