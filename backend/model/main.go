package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"singo/util"
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB = nil

// Database 在中间件中初始化mysql链接
func InitDB(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
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
	DB.Model(&RecvInfo{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	DB.Model(&Order{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	DB.Model(&Order{}).AddForeignKey("recv_info_id", "recv_infos(recv_info_id)", "SET NULL", "CASCADE")
	DB.Model(&ShoppingCart{}).AddForeignKey("book_id", "books(book_id)", "CASCADE", "CASCADE")
	DB.Model(&ShoppingCart{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
}
