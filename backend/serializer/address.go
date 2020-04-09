package serializer

import "DuckyGo/model"

// 用户查看所有收货地址时，返回给前端的数据格式
type AddressesRespData struct{
    Num     int                 `json:"num"`
    Infos   []model.RecvInfo    `json:"infos"`
}

func NewAddressRespData(infos []model.RecvInfo) AddressesRespData {
    return AddressesRespData{
        Num: len(infos),
        Infos: infos,
    }
}
