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

	r.GET("", api.Welcome)

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)
		v1.GET("auth", api.AliAuth)
		v1.GET("user", api.ShowUserInfo)

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
