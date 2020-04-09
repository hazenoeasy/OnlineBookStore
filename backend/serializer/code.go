package serializer

// 定义所有状态码

const (
	// 操作成功
	OpSuccess 			=	20000

	// 用户已存在
	UserNameRepeat		= 	40000

	// 用户名或密码错误
	UserNamePwdErr		=	40001

	// token授权到期
	TokenExpired		=	40002
)

const (
	// 严重的错误
	FatalErr 			= 	50000

	// 数据库写入错误
	DBWriteErr 			= 	50001

	// 数据库读取错误
	DBReadErr 			= 	50002

	// 请求格式错误
	RequestParamErr 	=	50003
)
