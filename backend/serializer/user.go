package serializer

// 用户登录后返回给前端的数据格式
// 用法: Responce{
// 	"code": xxx
// 	"data“: UserLoginRespData{ ... }
// 	"msg":	xxx
// }
type UserLoginRespData struct {
	Id 			int 	`json:"id"`
	UserName	string	`json:"username"`
	Token 		string	`json:"token"`
}