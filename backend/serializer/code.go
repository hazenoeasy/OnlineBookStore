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

	// 资源没有了
	ResourceEmptied 	=	40003
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

	// 文件保存失败
	FileSaveErr 		=	50004
)

// 定义订单的状态信息

const (
	OrderUnpaid			=	0	// 买家还未付款
	OrderUndelivered	=	1	// 买家还未发货
	OrderDelivered		=	2	// 买家已发货
	OrderClosed 		=	3	// 订单已经关闭
)
