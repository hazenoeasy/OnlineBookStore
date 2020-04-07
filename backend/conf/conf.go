package conf

import (
	"github.com/joho/godotenv"
	"github.com/smartwalle/alipay/v3"
	"os"
	"singo/cache"
	"singo/model"
	"singo/service"
	"singo/util"
)

// Init 初始化配置项
func Init() {

	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败:%v", err)
	}

	// 连接数据库
	model.InitDB(os.Getenv("MYSQL_DSN"))
	cache.InitRedis()

	// 初始化支付宝SDK
	aliClient, err := alipay.New(os.Getenv("ALI_APP_ID"), os.Getenv("ALI_MERCHANT_PRIVATE_KEY"), true)
	if err != nil {
		util.Log().Panic("支付宝SDK初始化失败:%v", err)
	}
	if err := aliClient.LoadAliPayPublicKey(os.Getenv("ALIPAY_PUBLIC_KEY")); err != nil {
		util.Log().Panic("支付宝初始化失败:%v", err)
	}
	service.SetAliClient(aliClient)
}
