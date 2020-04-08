package service

// 主题： 提供更改用户密码的服务
// 作者： 章星明
// 版本： v0.0.1
// 时间： 2020-4-8

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

// UserChangePwdService 修改用户密码服务
type UserChangePwdService struct {
	Header	UserHeader
	Pwds 	struct {
		OldPassword	string		`form:"password_old" json:"password_old" binding:"required"`
		NewPassword	string		`form:"password_new" json:"password_new" binding:"required"`
	}
}

// ChangePassword 修改密码
func (serv *UserChangePwdService) ChangePassword() serializer.Response {
	// 先检测旧的密码是否正确
	var user model.User
	if err := model.DB.Where("user_id = ?", serv.Header.UserId).First(&user).Error;
		err != nil {
		return serializer.Response{
			Code: serializer.UserNamePwdErr,
			Data: err,
			Msg:  "用户不存在",
		}
	}
	if serv.Pwds.OldPassword != user.Password {
		return serializer.Response{
			Code: serializer.UserNamePwdErr,
			Data: nil,
			Msg:  "密码错误",
		}
	}
	// 如果新旧密码不同，则更新密码
	if serv.Pwds.NewPassword != user.Password {
		if err := model.DB.Model(&user).Update("password", serv.Pwds.NewPassword).Error;
			err != nil {
			return serializer.Response{
				Code: serializer.DBWriteErr,
				Data: nil,
				Msg:  "密码更新失败",
			}
		}
	}
	return serializer.Response{
		Code: serializer.OpSuccess,
		Data: nil,
		Msg:  "密码修改成功",
	}
}
