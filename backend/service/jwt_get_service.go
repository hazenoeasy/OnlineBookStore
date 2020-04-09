package service

// 主题： 生成token
// 作者： 章星明
// 版本： v0.0.1
// 时间： 2020-4-8

import (
	"DuckyGo/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// UserHeader 用户授权信息
// token 由jwt中间件负责验证，因此我们只需要header里面的user_id属性
type UserHeader struct {
	UserId  int     `header:"user_id" binding:"required"`
}

// 返回一个token
// 如果生成token失败，则返回空字符串和错误
func NewJwtToken(audience string) (string, error) {
	claim := &jwt.StandardClaims{
		Audience:	audience,
		ExpiresAt: 	time.Now().Add(time.Minute * time.Duration(60)).Unix(),
		IssuedAt:  	time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	s, err := token.SignedString(conf.SigningKey)
	if err != nil {
		return "", err
	} else {
		return s, nil
	}
}