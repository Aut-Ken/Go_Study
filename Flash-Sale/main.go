package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建一个默认的路由引擎
	r := gin.Default()

	// 2. 配置一个简单的路由：当访问 /ping 时
	r.GET("/ping", func(c *gin.Context) {
		// 3. 返回 JSON 数据
		// gin.H 本质上就是一个 map[string]interface{}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "success",
		})
	})

	// 3. 启动服务，默认监听 8080 端口
	// 你可以在这里看到 "Listening and serving HTTP on :8080"
	r.Run(":8080")
}
