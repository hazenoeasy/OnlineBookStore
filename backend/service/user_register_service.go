package service

// 主题： 提供用户注册相关的服务
// 作者： 章星明
// 版本： v0.0.1
// 时间： 2020-4-8

import (
	"DuckyGo/model"
	"DuckyGo/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	UserName	string 	`form:"username" json:"username" binding:"required"`
	Password  	string 	`form:"password" json:"password" binding:"required"`
}

type UserName string
// Valid 验证用户名是否被注册过了
func (n UserName) Valid() (err *serializer.Response) {
	// 检查名称是否已经注册过了
	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ?", n).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 	serializer.UserNameRepeat,
			Msg:	"用户已注册",
		}
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		UserName: service.UserName,
		Password: service.Password,
	}

	// 验证用户是否已经注册过了
	if err := UserName(service.UserName).Valid(); err != nil {
		return *err
	}
	// 备注：
	// 此处没有将《查询用户名是否已经注册》和《创建用户》放进一个事务中，
	// 这是因为：数据库的username字段为UNIQUE，当相同的username试图插入时，
	// 	数据库会终止操作。
	// 如果出现这种情况，我们直接返回错误响应报文就行了

	//// 加密密码
	//if err := user.SetPassword(service.Password); err != nil {
	//	return user, &serializer.Response{
	//		Code: serializer.ServerPanicError,
	//		Msg:  "密码加密失败",
	//	}
	//}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Code: serializer.DBWriteErr,
			Msg:  "注册失败",
		}
	}

	return serializer.Response{
		Code: serializer.OpSuccess,
		Msg:  "注册成功",
	}
}
