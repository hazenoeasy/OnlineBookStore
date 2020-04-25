package model

// ShoppingCart 一条购物车记录
type ShoppingCart struct {
    UserId      int     `gorm:"primary_key;auto_increment:false;column:user_id;foreignkey:user_id;association_foreignkey:user_id;not null"`
    // FIXME: 怎加复合主键后，user_id的索引突然消失了，而book_id的索引还保留
    BookId      int     `gorm:"primary_key;auto_increment:false;column:book_id;foreignkey:book_id;association_foreignkey:book_id;not null"`
    Name        string  `gorm:"column:name;type:varchar(128);not null"`
    CoverUrl    string  `gorm:"column:cover_url;type:varchar(128)"`
    Price       int     `gorm:"column:price;not null"`
    Amount      int     `gorm:"column:amount;not null"`
}