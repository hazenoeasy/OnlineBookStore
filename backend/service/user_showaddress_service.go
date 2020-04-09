package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

// UserShowAddressService 展示用户所有的收货地址
type UserShowAddressService struct {
    UserHeader
}

func (u *UserShowAddressService) ShowAddresses() serializer.Response {
    var recvInfos []model.RecvInfo
    if err := model.DB.Where("user_id = ?", u.UserId).Find(&recvInfos).Error;
        err != nil {
            return serializer.Response{
                Code: serializer.DBReadErr,
                Data: serializer.NewAddressRespData(recvInfos),
                Msg:  "查询收货地址错误",
            }
    }
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: serializer.NewAddressRespData(recvInfos),
        Msg:  "收货地址查询成功",
    }
}