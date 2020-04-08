package model

// RecvInfo 收货人信息
type RecvInfo struct {
    RecvInfoId  int     `gorm:"column:recv_info_id;primary_key;auto_increment" json:"id"`
    UserId      int     `gorm:"column:user_id;foreignkey:user_id;association_foreignkey:user_id;not null" json:"user_id"`
    RecverName  string  `gorm:"column:receiver_name;type:varchar(64);not null" json:"realname"`
    Address     string  `gorm:"column:address;type:varchar(128);not null" json:"address"`
    Phone       string  `gorm:"column:phone;type:varchar(16);not null" json:"phone"`
}