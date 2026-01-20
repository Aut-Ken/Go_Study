package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

// 1. 【新增】定义一个全局切片（黑板）
// []Todo 表示这是一个这就装着 Todo 结构体的列表
// var todoList = []Todo{} 也可以
var todoList []Todo

func main() {
	r := gin.Default()

	// --- 接口 1: 添加任务 (POST) ---
	r.POST("/todo", func(c *gin.Context) {
		var newTodo Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 2. 【修改】把新任务“追加”到黑板上
		// append(原列表, 新元素) 会返回一个新的列表
		todoList = append(todoList, newTodo)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "任务已保存到内存！",
			// 这里我们可以打印一下当前总共有多少个任务 (len函数)
			"count": len(todoList),
		})
	})

	// --- 接口 2: 【新增】查看所有任务 (GET) ---
	r.GET("/todo", func(c *gin.Context) {
		// 3. 直接把全局变量 todoList 返回给前端
		// Gin 会自动把切片转换成 JSON 数组 [{}, {}, {}]
		c.JSON(http.StatusOK, gin.H{
			"data":  todoList,
			"count": len(todoList),
		})
	})

	r.Run(":8080")
}
