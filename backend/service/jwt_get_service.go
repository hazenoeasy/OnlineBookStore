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