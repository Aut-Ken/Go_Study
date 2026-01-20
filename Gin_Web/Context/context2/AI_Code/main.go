package main

import (
	"net/http"
	"strconv" // 【新】用来把字符串转成数字

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"` // 【新】增加 ID 字段
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

var todoList []Todo
var nextID = 1 // 【新】用来记录下一个 ID 该发多少号

func main() {
	r := gin.Default()

	// 1. 添加任务 (POST)
	r.POST("/todo", func(c *gin.Context) {
		var newTodo Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 【新】自动分配 ID，不是让用户传，而是服务器自己算
		newTodo.ID = nextID
		nextID += 1 // 下一次发号就要 +1

		todoList = append(todoList, newTodo)

		c.JSON(http.StatusOK, gin.H{"message": "添加成功", "data": newTodo})
	})

	// 2. 查看所有 (GET)
	r.GET("/todo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": todoList})
	})

	// 3. 【新接口】修改任务状态 (PUT)
	// :id 表示这是一个占位符，比如 /todo/1
	r.PUT("/todo/:id", func(c *gin.Context) {
		// A. 拿到 URL 里的 ID (它是字符串)
		idStr := c.Param("id")

		// B. 把字符串 "1" 变成 数字 1
		targetID, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID必须是数字"})
			return
		}

		// C. 在切片里“查户口”，找到那个 ID 对应的任务
		// 这是一个最基础的“遍历查找”算法
		var foundTodo *Todo // 定义一个指针，指向找到的任务
		for i := 0; i < len(todoList); i++ {
			if todoList[i].ID == targetID {
				foundTodo = &todoList[i] // 找到了！记下它的地址
				break
			}
		}

		// D. 如果没找到
		if foundTodo == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "找不到这个 ID 的任务"})
			return
		}

		// E. 找到了，修改它的状态 (取反：true变false，false变true)
		foundTodo.IsDone = !foundTodo.IsDone

		c.JSON(http.StatusOK, gin.H{
			"message": "状态已更新",
			"data":    foundTodo,
		})
	})

	r.Run(":8080")
}
