package service

import (
    "github.com/smartwalle/alipay/v3"
)

// this global variable is initialized in conf.conf.go
var aliClient *alipay.Client

func SetAliClient(c *alipay.Client)  {
    if c != nil {
        aliClient = c
    } else {
        panic("SetAliClient(v): v is nil")
    }
}

func GetAliClient() *alipay.Client {
    if aliClient != nil {
        return aliClient
    } else {
        panic("GetAliClient(): return nil")
    }
}

type tokenInfo struct {
    UserId       string `json:"user_id"`
    AccessToken  string `json:"access_token"`
    ExpiresIn    int64  `json:"expires_in"`
    RefreshToken string `json:"refresh_token"`
    ReExpiresIn  int64  `json:"re_expires_in"`
}
var aliToken *tokenInfo

func SetAliToken(rsp *alipay.SystemOauthTokenRsp)  {
    if rsp == nil {
        panic("SetAliToken(rsp): rsp is nil")
    } else {
        if aliToken == nil {
            aliToken = new(tokenInfo)
        }
        aliToken.UserId         = rsp.Content.UserId
        aliToken.AccessToken    = rsp.Content.AccessToken
        aliToken.ExpiresIn      = rsp.Content.ExpiresIn
        aliToken.RefreshToken   = rsp.Content.RefreshToken
        aliToken.ReExpiresIn    = rsp.Content.ReExpiresIn
    }
}

func GetAliToken() *tokenInfo  {
    // TODO: 处理token过期的问题
    if aliToken == nil {
        panic("GetAliToken: return nil")
    } else {
        return aliToken
    }
}