package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 结构体
// GORM 会默认寻找名为 "users" (复数) 的表
// 字段名 ID 默认对应数据库的 id 列，Name 对应 name 列...
type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"` // 标记这是主键
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 接收转账请求的结构体
type TransferRequest struct {
	FromID uint `json:"from_id"`
	ToID   uint `json:"to_id"`
}

// 全局 DB 对象，现在它是 *gorm.DB 类型，不是 *sql.DB 了！
var db *gorm.DB

func main() {
	// ==========================================
	// 1. GORM 连接数据库
	// ==========================================
	dsn := "root:@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	// 注意：这里用 gorm.Open
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("GORM 连接失败: ", err)
	}
	fmt.Println("🎉 GORM 连接成功！")

	// ==========================================
	// 2. 启动 Gin
	// ==========================================
	r := gin.Default()

	// 📌 接口 A: 获取所有用户 (对比一下以前多简单！)
	r.GET("/users", func(c *gin.Context) {
		var users []User
		// Find(&users): 去数据库找所有用户，自动填满 users 切片
		// 不需要 rows.Scan，不需要循环，一行搞定！
		result := db.Find(&users)

		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, users)
	})

	// 📌 接口 B: 转账 (事务版)
	r.POST("/transfer", func(c *gin.Context) {
		var req TransferRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "参数不对"})
			return
		}

		// ✨ GORM 的自动事务闭包
		// 你不需要手动 Begin/Commit/Rollback
		// 只要函数返回 error，它自动回滚；返回 nil，它自动提交。
		err := db.Transaction(func(tx *gorm.DB) error {

			// 1. 扣减 (UpdateColumn 类似 SQL 的 Update)
			// gorm.Expr("age - ?", 10) 代表在原值基础上 -10
			if err := tx.Model(&User{}).Where("id = ?", req.FromID).UpdateColumn("age", gorm.Expr("age - ?", 10)).Error; err != nil {
				return err // 返回错误，自动回滚
			}

			// 2. 增加
			if err := tx.Model(&User{}).Where("id = ?", req.ToID).UpdateColumn("age", gorm.Expr("age + ?", 10)).Error; err != nil {
				return err // 返回错误，自动回滚
			}

			// 3. 返回 nil，代表一切正常，自动提交
			return nil
		})

		if err != nil {
			c.JSON(500, gin.H{"error": "转账失败", "detail": err.Error()})
			return
		}

		c.JSON(200, gin.H{"status": "success", "msg": "GORM 转账成功！"})
	})

	r.Run(":8080")
}
