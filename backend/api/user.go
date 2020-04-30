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
	// XXX: 如果form和header标签在一个结构体里面出现，对改结构体的成员的分别绑定会出现问题，最好导致只有部分的成员绑定成功（出现在结构体靠前位置的成员）
	if err := c.ShouldBindHeader(&userid); err == nil {
		if err := c.ShouldBind(&newname); err == nil {
			serv := service.UserChangeNameService{Header: userid, Body: newname}
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
			serv := service.UserAddAddressService{Header: userid, Body: recvinfo}
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

// UserDelAddress 删除用户的一个地址
// 在调用此api前，需要使用jwt中间件验证token
func UserDelAddress(c *gin.Context)  {
	var serv service.UserDelAddressService
	if err := c.ShouldBind(&serv); err == nil {
		c.JSON(http.StatusOK, serv.DelAddress())
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserAddCart 商品添加到购物车
func UserAddCart(ctx *gin.Context)  {
	var (
		header 	service.UserHeader
		body 	service.BodyAddCartService
	)
	if err := ctx.ShouldBindHeader(&header); err == nil {
		if err = ctx.ShouldBind(&body); err == nil {
			serv := service.UserAddCartService{Header: header, Body: body}
			ctx.JSON(http.StatusOK, serv.AddCart())
		} else {
			ctx.JSON(http.StatusOK, ErrorResponse(err))
		}
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserShowCart 查看购物车
func UserShowCart(ctx *gin.Context) {
	var serv service.ShowCartService
	if err := ctx.ShouldBindHeader(&serv); err == nil {
		ctx.JSON(http.StatusOK, serv.ShowCart())
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserDelCart 从购物车中删除货物
func UserDelCart(ctx *gin.Context)  {
	var (
		header 	service.UserHeader
		body 	service.BodyDelCartService
	)
	if err := ctx.ShouldBindHeader(&header); err == nil {
		if err = ctx.ShouldBind(&body); err == nil {
			serv := service.DelCartService{Header: header, Body: body}
			ctx.JSON(http.StatusOK, serv.DelCart())
		} else {
			ctx.JSON(http.StatusOK, ErrorResponse(err))
		}
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserCreateOrder 购物车结算
func UserCreateOrder(ctx *gin.Context)  {
	var (
		header 	service.UserHeader
		body 	service.BodyOrderCreateService
	)
	if err := ctx.ShouldBindHeader(&header); err == nil {
		if err = ctx.ShouldBind(&body); err == nil {
			serv := service.OrderCreateService{Header: header, Body: body}
			ctx.JSON(http.StatusOK, serv.CreateOrder())
		} else {
			ctx.JSON(http.StatusOK, ErrorResponse(err))
		}
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserPayOrder 付款
func UserPayOrder(ctx *gin.Context)  {
	var (
		header 	service.UserHeader
		body 	service.BodyPayOrderService
	)
	if err := ctx.ShouldBindHeader(&header); err == nil {
		if err = ctx.ShouldBind(&body); err == nil {
			serv := service.PayOrderService{Header: header, Body: body}
			ctx.JSON(http.StatusOK, serv.PayOrder())
		} else {
			ctx.JSON(http.StatusOK, ErrorResponse(err))
		}
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserShowOrder 查看购物订单
func UserShowOrder(ctx *gin.Context)  {
	var serv service.OrderShowService
	if err := ctx.ShouldBindHeader(&serv); err == nil {
		ctx.JSON(http.StatusOK, serv.ShowOrder())
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// UserGetCommodity 用户确认收货
func UserGetCommodity(ctx *gin.Context)  {
	var (
		header 	service.UserHeader
		body 	service.BodyGetCommodityService
	)
	if err := ctx.ShouldBindHeader(&header); err == nil {
		if err = ctx.ShouldBind(&body); err == nil {
			serv := service.GetCommodityService{Header: header, Body: body}
			ctx.JSON(http.StatusOK, serv.OrderGetCommodity())
		} else {
			ctx.JSON(http.StatusOK, ErrorResponse(err))
		}
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}