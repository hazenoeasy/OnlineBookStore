package model

import "time"

// Order 订单
type Order struct {
    OrderId     int         `gorm:"column:order_id;primary_key;auto_increment"`
    UserId      int         `gorm:"column:user_id;foreignkey:user_id;association_foreignkey:user_id;not null"`
    SalesmanId  int         `gorm:"column:salesman_id;foreignkey:salesman_id;association_foreignkey:user_id"`
    RecvInfoId  int         `gorm:"column:recv_info_id;foreignkey:recv_info_id;association_foreignkey:recv_info_id"`
    BookId      int         `gorm:"column:book_id;foreignkey:book_id;association_foreignkey:book_id"`
    OrderTime   time.Time   `gorm:"column:order_time"`
    Status      int         `gorm:"column:status;not null;index:idx_status"`
    // 快递单号，由商家填写
    ExpId       string      `gorm:"column:express_id;type:varchar(32)"`
    Title       string      `gorm:"column:title;type:varchar(128);not null"`
    Price       int         `gorm:"column:price;not null"`
    Amount      int         `gorm:"column:amount;default:1"`
}
