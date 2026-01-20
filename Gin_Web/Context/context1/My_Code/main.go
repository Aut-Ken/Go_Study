package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

var todoList []Todo

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
		todoList = append(todoList, newTodo)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "任务已经保存到内存！",
			"count":   len(todoList),
		})
	})

	r.GET("/todo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data":  todoList,
			"count": len(todoList),
		})
	})

	r.Run(":8080")
}
