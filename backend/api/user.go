package api

import (
	"DuckyGo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		c.JSON(http.StatusOK, service.Register())
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		c.JSON(http.StatusOK, service.Login())
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserChangeName 修改用户名
// 在调用此api前，需要使用jwt中间件验证token
func UserChangeName(c *gin.Context) {
	var (
		userid 	service.UserHeader
		newname	service.SubUserChangeNameService
	)
	// TODO: 修bug了：目前测试后发现，如果form和header标签在一个结构体里面出现，对改结构体的成员的分别绑定会出现问题，最好导致只有部分的成员绑定成功（出现在结构体靠前位置的成员）
	if err := c.ShouldBindHeader(&userid); err == nil {
		if err := c.ShouldBind(&newname); err == nil {
			serv := service.UserChangeNameService{userid, newname}
			c.JSON(http.StatusOK, serv.ChangeName())
		} else {
			c.JSON(http.StatusOK, ErrorResponse(err))
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserChangePwd 修改密码
// 在调用此api前，需要使用jwt中间件验证token
func UserChangePwd(c *gin.Context) {
	var (
		userid 	service.UserHeader
		pwds	service.SubUserChangePwdService
	)
	if err := c.ShouldBindHeader(&userid); err == nil {
		if err := c.ShouldBind(&pwds); err == nil {
			serv := service.UserChangePwdService{userid, pwds}
			c.JSON(http.StatusOK, serv.ChangePassword())
		} else {
			c.JSON(http.StatusOK, ErrorResponse(err))
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserAddAddress 用户添加一个收货地址
// 在调用此api前，需要使用jwt中间件验证token
func UserAddAddress(c *gin.Context) {
	var (
		userid 		service.UserHeader
		recvinfo	service.SubUserAddAddressService
	)
	if err := c.ShouldBindHeader(&userid); err == nil {
		if err := c.ShouldBind(&recvinfo); err == nil {
			serv := service.UserAddAddressService{userid, recvinfo}
			c.JSON(http.StatusOK, serv.AddAddress())
		} else {
			c.JSON(http.StatusOK, ErrorResponse(err))
		}
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserShowAddress 查看用户的所有收货地址
// 在调用此api前，需要使用jwt中间件验证token
func UserShowAddress(c *gin.Context) {
	var serv service.UserShowAddressService
	if err := c.ShouldBindHeader(&serv); err == nil {
		c.JSON(http.StatusOK, serv.ShowAddresses())
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
