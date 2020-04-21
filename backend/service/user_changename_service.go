package service

// 主题： 提供更改用户名相关的服务
// 作者： 章星明
// 版本： v0.0.1
// 时间： 2020-4-8

import (
    "DuckyGo/model"
    "DuckyGo/serializer"
)

type SubUserChangeNameService struct {
    NewName     string  `form:"username_new" json:"username_new" binding:"required"`
}

// UserChangeNameService 管理修改用户名的服务
type UserChangeNameService struct {
    Header      UserHeader
    Body        SubUserChangeNameService
}

// ChangeName 更改用户名
// 返回JSON格式的响应报文
// 可重入
func (s *UserChangeNameService) ChangeName() serializer.Response {
    // 由于使用了middleware/jwt.go/JwtRequired()中间件，
    // 此处不再需要重新验证token

    // 检测新用户名是否已经存在
    if err := UserName(s.Body.NewName).Valid(); err != nil {
        return *err
    }
    // 备注：
    // 此处没有将《查询用户名是否已经注册》和《修改用户名》放进一个事务中，
    // 这是因为：数据库的username字段为UNIQUE，当相同的username试图插入时，
    // 	数据库会终止操作。
    // 如果出现这种情况，我们直接返回错误响应报文就行了

    // 修改用户名
    user := model.User{ UserId: s.Header.UserId, }
    if err := model.DB.Model(&user).Update("user_name", s.Body.NewName).Error; err != nil {
        return serializer.Response{
            Code: serializer.DBWriteErr,
            Data: err,
            Msg:  "更新用户名失败",
        }
    }
    return serializer.Response{
        Code: serializer.OpSuccess,
        Data: s.Body.NewName,
        Msg:  "用户名修改成功",
    }
}