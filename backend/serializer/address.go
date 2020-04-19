package serializer

import (
    "DuckyGo/model"
)


// 用户查看所有收货地址时，返回给前端的数据格式
type RecvAddressData struct {
    Id          int     `json:"id"`
    RealName    string  `json:"realname"`
    Address     string  `json:"address"`
    Phone       string  `json:"phone"`
}
type AddressesRespData struct{
    Num     int                 `json:"num"`
    Items   []RecvAddressData   `json:"items"`
}

func NewAddressRespData(infos []model.RecvInfo) AddressesRespData {
    items := make([]RecvAddressData, len(infos))
    for i := 0; i < len(infos); i++ {
        items[i].Id         =   infos[i].RecvInfoId
        items[i].Address    =   infos[i].Address
        items[i].RealName   =   infos[i].RecverName
        items[i].Phone      =   infos[i].Phone
    }
    return AddressesRespData{
        Num:    len(items),
        Items:  items,
    }
}
