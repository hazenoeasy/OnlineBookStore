package model

// User 用户模型
type User struct {
	UserId 		int		`gorm:"column:user_id;primary_key;auto_increment"`
	UserName	string	`gorm:"column:user_name;type:varchar(64);not null;unique"`
	Password 	string	`gorm:"column:password;type:varchar(128);not null"`
}

// 检测密码是否正确
// 如果正确，返回true，否则返回false
func (u *User) CheckPassword(pwd string) bool {
	return pwd == u.Password
}

//const (
//	// PassWordCost 密码加密难度
//	PassWordCost = 12
//	// Active 激活用户
//	Active string = "active"
//	// Inactive 未激活用户
//	Inactive string = "inactive"
//	// Suspend 被封禁用户
//	Suspend string = "suspend"
//)
//
//// GetUser 用ID获取用户
//func GetUser(ID interface{}) (User, error) {
//	var user User
//	result := DB.First(&user, ID)
//	return user, result.Error
//}
//
//// SetPassword 设置密码
//func (user *User) SetPassword(password string) error {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
//	if err != nil {
//		return err
//	}
//	user.PasswordDigest = string(bytes)
//	return nil
//}
//
//// CheckPassword 校验密码
//func (user *User) CheckPassword(password string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
//	return err == nil
//}
