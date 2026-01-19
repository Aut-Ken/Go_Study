package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1. 定义数据的“形状” (回顾 JSON 知识)
type Todo struct {
	// 这里的 tag 决定了前端传来的 JSON key 必须叫什么
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

func main() {
	// 2. 启动引擎
	r := gin.Default()

	// 3. 定义 POST 接口 (回顾 HTTP 知识: POST用于提交)
	// ★ 这里的 c 是 *gin.Context (回顾指针知识: 这是一个共享的遥控器)
	r.POST("/todo", func(c *gin.Context) {

		// --- 第一步：收货 (解析 JSON) ---

		var newTodo Todo

		// ★ 重点来了！
		// ShouldBindJSON 的作用就是 json.Unmarshal
		// ★ 必须传 &newTodo (指针)，因为我们要修改 newTodo 这个变量的值！
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			// 如果解析失败 (比如发来的不是 JSON)，回个 400 (Bad Request)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// --- 第二步：处理业务 ---

		// 这里我们就简单打印一下，假装存进数据库了
		fmt.Printf("收到新任务: %s (完成状态: %v)\n", newTodo.Title, newTodo.IsDone)

		// --- 第三步：发货 (返回 JSON) ---

		// ★ c.JSON 会自动把 map 序列化成 JSON 字符串
		// 并设置 Content-Type 为 application/json
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "任务已添加！",
			"data":    newTodo, // 把刚才收到的数据再发回去确认一遍
		})
	})

	// 4. 监听端口
	r.Run(":8080")
}
