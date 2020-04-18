package model

import (
    "time"
)

// Order 订单
type Order struct {
    OrderId     int         `gorm:"column:order_id;primary_key;auto_increment"`
    UserId      int         `gorm:"column:user_id;foreignkey:user_id;association_foreignkey:user_id;not null"`
    RecvInfoId  int         `gorm:"column:recv_info_id;foreignkey:recv_info_id;association_foreignkey:recv_info_id"`
    // 快递公司名称，由商家填写
    ExpCompName string      `gorm:"column:express_comp_name;type:varchar(16)"`
    // 快递单号，由商家填写
    ExpId       string      `gorm:"column:express_id;type:varchar(32)"`
    OrderTime   time.Time   `gorm:"column:order_time"`
    Status      int         `gorm:"column:status;not null;index:idx_status"`
    TotalPrice  int         `gorm:"column:total_price;not null"`
}
