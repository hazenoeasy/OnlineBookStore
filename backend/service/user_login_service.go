package service

// 主题：用户登录服务
// 作者： 章星明
// 版本： v0.0.1
// 时间： 2020-4-8

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
	"strconv"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Login 用户登录函数
func (service *UserLoginService) Login() serializer.Response {
	var user model.User

	// 检测用户是否存在
	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.Response{
			Code: serializer.UserNamePwdErr,
			Data: nil,
			Msg:  "用户不存在",
		}
	}

	// 校验密码是否正确
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Code: serializer.UserNamePwdErr,
			Data: nil,
			Msg:  "密码错误",
		}
	}

	// 生成token
	token, err := NewJwtToken(strconv.Itoa(user.UserId))
	if err != nil {
		return serializer.Response{
			Code: serializer.FatalErr,
			Data: err.Error(),
			Msg:  "登录失败（token生成错误）",
		}
	}
	return  serializer.Response{
		Code: serializer.OpSuccess,
		Data: serializer.UserLoginRespData{
			Id:       user.UserId,
			UserName: user.UserName,
			Token:    token,
		},
		Msg:  "登录成功",
	}
}