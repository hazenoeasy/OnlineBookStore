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

		// 用户相关接口:
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
			// TODO: 注销
			// * 查看收货地址
			v1.GET("/user/address", api.UserShowAddress)
			// * 设置收货地址
			v1.POST("/user/address", api.UserAddAddress)
			// * 删除收货地址
			v1.DELETE("/user/address", api.UserDelAddress)

			// 上传书籍
			v1.POST("/user/books", api.SellerAddBook)
			// 查看我的卖书
			v1.GET("/user/books", api.SellerShowBook)
			//// 修改卖书信息
			//v1.PUT("/user/books", api.SellerUpdateBook)
			//// 删除我的卖书
			//v1.DELETE("/user/books", api.SellerDelBook)
		}

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
	}

	return r
}
