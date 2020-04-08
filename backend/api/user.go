package api

import (
	"DuckyGo/service"
	"net/http"

	"github.com/gin-gonic/gin"
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
	var service service.UserChangeNameService
	if err := c.ShouldBindHeader(&(service.Header)); err == nil {
		if err := c.ShouldBind(&(service.NewName)); err == nil {
			c.JSON(http.StatusOK, service.ChangeName())
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
	// 在调用此api前，需要使用jwt中间件验证token
	var serv service.UserChangePwdService
	if err := c.ShouldBindHeader(&(serv.Header)); err == nil {
		if err := c.ShouldBind(&(serv.Pwds)); err == nil {
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
	var serv service.UserAddAddressService
	if err := c.ShouldBindHeader(&(serv.Header)); err == nil {
		if err := c.ShouldBind(&(serv.Body)); err == nil {
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
