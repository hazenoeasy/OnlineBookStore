package server

import (
	"DuckyGo/api"
	"DuckyGo/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	// 启动Redis的情况下将切换成Redis保存Session.
	if os.Getenv("RIM") == "use" {
		r.Use(middleware.SessionRedis(os.Getenv("SESSION_SECRET")))
	} else {
		r.Use(middleware.SessionCookie(os.Getenv("SESSION_SECRET")))
	}
	r.Use(middleware.Cors())

	// 主页
	r.GET("/", api.Index)

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		// 主页部分的接口：
		v1.GET("books/search", api.SearchBook)

		// * 注册
		v1.POST("user/register", api.UserRegister)
		// * 登录
		v1.POST("user/login", api.UserLogin)
		v1.Use(middleware.JwtRequired())
		{
			// * 修改用户名
			v1.PUT("user/name", api.UserChangeName)
			// * 修改密码
			v1.PUT("user/password", api.UserChangePwd)
			// * 查看收货地址
			v1.GET("/user/address", api.UserShowAddress)
			// * 设置收货地址
			v1.POST("/user/address", api.UserAddAddress)
			// * 删除收货地址
			v1.DELETE("/user/address", api.UserDelAddress)

			// 添加到购物车中
			v1.POST("/user/cart", api.UserAddCart)
			// 查看购物车内容
			v1.GET("/user/cart", api.UserShowCart)
			// 从购物车中删除
			v1.DELETE("/user/cart", api.UserDelCart)
			// 购物车结算
			v1.POST("/user/order", api.UserCreateOrder)
			// 付款
			v1.PUT("/user/order", api.UserPayOrder)

			// 用户查看购物订单
			v1.GET("/user/order", api.UserShowOrder)
			// 用户确认收货
			v1.PUT("/user/commodity", api.UserGetCommodity)

			// 卖家相关接口
			// 上传书籍
			v1.POST("/user/books", api.SellerAddBook)
			// 查看我的卖书
			v1.GET("/user/books", api.SellerShowBook)
			// 修改卖书信息
			v1.PUT("/user/books", api.SellerUpdateBook)
			//// 删除我的卖书
			v1.DELETE("/user/books", api.SellerDelBook)

			// 查看订单
			v1.GET("/user/seller/order", api.SellerShowOrder)
			// 订单发货
			v1.POST("/user/commodity", api.SellerPostCommodity)
		}
	}

	return r
}
