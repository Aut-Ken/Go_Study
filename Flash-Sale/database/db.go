package database

import (
	"flash-sale/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局 DB 对象，首字母大写，意味着别的包也能用
var DB *gorm.DB

func InitDB() {
	// ... 连接代码 ...
	dsn := "root:@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("GORM 连接失败: ", err)
	}

	// 👇 新增：自动迁移 (建表)
	// GORM 会检测 User, Product, Order 结构体，自动在数据库创建对应的表
	// ⚠️ 注意：引入 models 包时，确保路径是对的
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	fmt.Println("🎉 数据库表结构同步完成！")
}
