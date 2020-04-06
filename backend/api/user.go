package api

import (
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"net/http"
	"singo/service"
)

// 支付宝授权(网页版)
func AliAuth(c *gin.Context)  {
	if c.Query("auth_code") == "" {
		// first come here
		if url, err := service.GetAliClient().PublicAppAuthorize([]string{"auth_user"},
			"http://zzzyh.top/api/v1/auth",
			""); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg" : "can't invoke alipay authorize!",
				"data" : nil,
				"err" : err.Error(),
			})
		} else {
			c.Redirect(http.StatusTemporaryRedirect, url.String())
		}
	} else {
		// routine with auth_code come here for getting token
		p := alipay.SystemOauthToken{GrantType: "authorization_code", Code: c.Query("auth_code")}
		if rsp, err := service.GetAliClient().SystemOauthToken(p); err != nil {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"msg" : "can't get auth token",
				"data" : nil,
				"err" : err.Error(),
			})
		} else {
			service.SetAliToken(rsp)
			c.JSON(http.StatusOK, gin.H{
				"msg" : "got token!",
				"data" : rsp.Content,
				"err" : nil,
			})
		}
	}

}

// 展示user info
func ShowUserInfo(c *gin.Context)  {
	info := alipay.UserInfoShare{
		AuthToken:    service.GetAliToken().AccessToken,
	}
	if rsp, err := service.GetAliClient().UserInfoShare(info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "can't get user info",
			"data": nil,
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user info here",
			"data": rsp.Content,
			"err": err.Error(),
		})
	}
}


/************************Remove these api we don't need*************************/
//// UserRegister 用户注册接口
//func UserRegister(c *gin.Context) {
//	var service service.UserRegisterService
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.Register()
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//	}
//}
//
//// UserLogin 用户登录接口
//func UserLogin(c *gin.Context) {
//	var service service.UserLoginService
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.Login(c)
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//	}
//}
//
//// UserMe 用户详情
//func UserMe(c *gin.Context) {
//	user := CurrentUser(c)
//	res := serializer.BuildUserResponse(*user)
//	c.JSON(200, res)
//}
//
//// UserLogout 用户登出
//func UserLogout(c *gin.Context) {
//	s := sessions.Default(c)
//	s.Clear()
//	s.Save()
//	c.JSON(200, serializer.Response{
//		Code: 0,
//		Msg:  "登出成功",
//	})
//}
