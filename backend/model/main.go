package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func InitDB(connString string) {
	db, err := gorm.Open("mysql", connString)
	// 开启Shell显示SQL
	db.LogMode(true)
	// Error
	if err != nil {
		panic(fmt.Sprintf("连接数据库出现异常: %v", err))
	}
	// 设置连接池
	// 空闲
	db.DB().SetMaxIdleConns(50)
	// 打开
	db.DB().SetMaxOpenConns(100)
	// 超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
	// 自动迁移
	DB.AutoMigrate(&Book{}, &Order{}, &RecvInfo{}, &ShoppingCart{}, &User{})
	DB.Model(&RecvInfo{}).AddForeignKey("user_id", "users(user_id)",
		"CASCADE", "CASCADE")
	DB.Model(&Order{}).AddForeignKey("user_id", "users(user_id)",
		"CASCADE", "CASCADE")
	DB.Model(&Order{}).AddForeignKey("recv_info_id", "recv_infos(recv_info_id)",
		"SET NULL", "CASCADE")
	DB.Model(&ShoppingCart{}).AddForeignKey("book_id", "books(book_id)",
		"CASCADE", "CASCADE")
	DB.Model(&ShoppingCart{}).AddForeignKey("user_id", "users(user_id)",
		"CASCADE", "CASCADE")
}
