package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type SubUserAddAddressService struct {
    RealName    string      `form:"realname" json:"realname" binding:"required"`
    Address     string      `form:"address" json:"address" binding:"required"`
    Phone       string      `form:"phone" json:"phone" binding:"required"`
}

// UserAddAddressService 增加收货人地址的服务
type UserAddAddressService struct {
    Header  UserHeader
    Body    SubUserAddAddressService
}

func (u *UserAddAddressService) AddAddress() serializer.Response {
    recv := model.RecvInfo{
        UserId:     u.Header.UserId,
        RecverName: u.Body.RealName,
        Address:    u.Body.Address,
        Phone:      u.Body.Phone,
    }
    if err := model.DB.Create(&recv).Error; err != nil {
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: nil,
            Msg:  "收货地址添加失败",
        }
    }
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: nil,
        Msg:  "收货地址添加成功",
    }
}


