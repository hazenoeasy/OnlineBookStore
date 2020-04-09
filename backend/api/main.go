package api

import (
	"DuckyGo/conf"
	"DuckyGo/serializer"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// Index 主页
func Index(c *gin.Context) {
	c.String(http.StatusOK, "这是主页")
}

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, serializer.Response{
		Msg: "Pong",
	})
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field()))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))
			return serializer.Response{
				Code:  	serializer.RequestParamErr,
				Msg:   	fmt.Sprintf("%s%s", field, tag),
				Data: 	fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Code:  	serializer.RequestParamErr,
			Msg:   	"JSON类型不匹配",
			Data: 	fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Code:  	serializer.RequestParamErr,
		Msg:   	"参数错误",
		Data: 	fmt.Sprint(err),
	}
}
