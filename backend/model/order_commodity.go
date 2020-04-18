package model

type OrderCommodity struct {
    OrderId     int     `gorm:"column:order_id;foreignkey:order_id;association_foreignkey:order_id;not null"`
    BookId      int     `gorm:"column:book_id;foreignkey:book_id;association_foreignkey:book_id;not null"`
    Title       string  `gorm:"column:title;type:varchar(128);not null"`
    Price       int     `gorm:"column:price;not null"`
    Amount      int     `gorm:"column:amount;not null"`
}