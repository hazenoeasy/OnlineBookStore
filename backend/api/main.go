package api

import (
	"github.com/gin-gonic/gin"
	"singo/model"
)

// 欢迎界面
func Welcome(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"code": 0,
		"data":	nil,
		"msg":  "Welcome!",
		"err": 	nil,
	})
}

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "Ping",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

//// ErrorResponse 返回错误消息
//func ErrorResponse(err error) serializer.Response {
//	if ve, ok := err.(validator.ValidationErrors); ok {
//		for _, e := range ve {
//			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
//			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
//			return serializer.ParamErr(
//				fmt.Sprintf("%s%s", field, tag),
//				err,
//			)
//		}
//	}
//	if _, ok := err.(*json.UnmarshalTypeError); ok {
//		return serializer.ParamErr("JSON类型不匹配", err)
//	}
//
//	return serializer.ParamErr("参数错误", err)
//}
