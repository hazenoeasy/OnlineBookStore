package service

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type UserDelAddressService struct {
    Id  int `form:"id" json:"id" binding:"required"`
}

func (u *UserDelAddressService) DelAddress() serializer.Response {
    var address model.RecvInfo
    address.RecvInfoId = u.Id
    if err := model.DB.Delete(&address).Error; err == nil {
        return serializer.Response{
            Code: serializer.OpSuccess,
            Data: nil,
            Msg:  "ok",
        }
    } else {
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: nil,
            Msg:  "数据库删除时发生错误",
        }
    }
}