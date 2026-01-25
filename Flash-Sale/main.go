package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"flash-sale/controller"
	"flash-sale/database"
	"flash-sale/middleware"
	"flash-sale/models"
	"flash-sale/service"
)

func main() {
	database.InitDB()
	database.InitRedis()
	database.InitRabbitMQ()

	var countProducts int64
	database.DB.Model(&models.Product{}).Count(&countProducts)
	if countProducts == 0 {
		database.DB.Create(&models.Product{
			Name:   "iPhone 16 Pro Max",
			Price:  9999.00,
			Stock:  50,
			Status: 1,
		})
		fmt.Println("✨ 已初始化测试商品：iPhone 16 Pro Max (库存 50)")
	}

	var countUsers int64
	database.DB.Model(&models.User{}).Count(&countUsers)
	if countUsers == 0 {
		users := []models.User{
			{Name: "赵艺凯", Age: 21}, // ID 1
			{Name: "张三", Age: 25},  // ID 2
			{Name: "李四", Age: 30},  // ID 3
			{Name: "王五", Age: 18},  // ID 4
			{Name: "赵六", Age: 99},  // ID 5
		}
		result := database.DB.Create(&users)
		if result.Error == nil {
			fmt.Printf("✨ 成功初始化 %d 个用户！\n", result.RowsAffected)
		}
	}

	var product models.Product
	database.DB.First(&product, 1)

	database.RDB.Set(database.Ctx, "product:1:stock", product.Stock, 0)
	fmt.Printf("🔥 库存已同步到 Redis: %d\n", product.Stock)

	go service.StartConsumer()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.POST("/login", controller.Login)
	r.POST("/users/search", controller.QueryUsers)
	r.POST("/transfer", controller.Transfer)
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// 只有带着 Token 的人才能访问这个 /buy
		authorized.POST("/buy", controller.Buy)
	}
	r.Run(":8080")
}
