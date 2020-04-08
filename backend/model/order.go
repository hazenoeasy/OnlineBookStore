package model

import (
    "time"
)

// Order 订单
type Order struct {
    OrderId     int     `gorm:"column:order_id;primary_key;auto_increment"`
    UserId      int     `gorm:"column:user_id;foreignkey:user_id;association_foreignkey:user_id;not null"`
    RecvInfoId  int     `gorm:"column:recv_info_id;foreignkey:recv_info_id;association_foreignkey:recv_info_id"`
    // 快递公司名称，由商家填写
    ExpCompName string  `gorm:"column:express_comp_name;type:varchar(16)"`
    // 快递单号，由商家填写
    ExpId       int     `gorm:"column:express_id"`
    OrderTime   time.Time
    Status      int     `gorm:"column:status;not null"`
    // 一份订单中多个book_id转换为字符串JSON序列化结果
    // e.g. {"books_id":[12345678,22334455,38383898]}
    BooksId     string  `gorm:"column:books_id;type:text;not null"`
    // 对应上一项中每个book_id对应的书籍的数量
    // e.g. {"amounts":[1,1,1]}
    Amounts     string  `gorm:"column:amounts;type:text;not null"`
    TotalPrice  int     `gorm:"column:total_price;not null"`
}
