package server

import (
	"github.com/gin-gonic/gin"
	"singo/api"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// TODO: add middleware(like session, CORS) in future...
	// 中间件, 顺序不能改
	//r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	//r.Use(middleware.Cors())
	//r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		// 测试
		v1.POST("ping", api.Ping)
		v1.GET("auth", api.AliAuth)
		v1.GET("user", api.ShowUserInfo)

		// 用户相关接口:
		// * 注册
		// * 登录
		// * 修改用户名
		// * 修改密码
		// TODO: 注销
		// * 查看收货地址
		// * 设置收货地址

		// 订单相关接口:
		// * 创建订单
		// * 根据订单状态查询所有订单
		// * 支付
		// * 删除订单
		// * 查看购物车
		// * 添加到购物车
		// TODO: 从购物车中删除

		// 书籍相关接口:
		// * 书籍展示列表
		// * 查询书籍详细信息
		// * 新增书籍
		// * 删除书籍
		// * 修改书籍

		// TODO: we don't implement user info(need session) now
		// 需要登录保护的
		//auth := v1.Group("")
		//auth.Use(middleware.AuthRequired())
		//{
		//	// User Routing
		//	auth.GET("user/me", api.UserMe)
		//	auth.DELETE("user/logout", api.UserLogout)
		//}
	}
	//v2 := r.Group("/m.api/v1")
	//{
	//	// mobile version apis
	//}
	return r
}
