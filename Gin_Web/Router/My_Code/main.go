package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

func main() {
	r := gin.Default()

	r.POST("/todo", func(c *gin.Context) {
		var newTodo Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Printf("收到新任务: %s (完成状态: %v)\n", newTodo.Title, newTodo.IsDone)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "任务已经添加",
			"data":    newTodo,
		})
	})
	r.Run(":8080")
}
