package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"is_done"`
}

var todoList []Todo
var nextID = 1

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

		newTodo.ID = nextID
		nextID += 1

		todoList = append(todoList, newTodo)
		c.JSON(http.StatusOK, gin.H{
			"message": "任务添加成功",
			"data":    newTodo,
		})
	})

	r.GET("/todo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": todoList,
		})
	})

	r.PUT("/todo/:id", func(c *gin.Context) {
		strid := c.Param("id")
		targetID, err := strconv.Atoi(strid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var foundTodo *Todo
		for i := 0; i < len(todoList); i++ {
			if todoList[i].ID == targetID {
				foundTodo = &todoList[i]
				break
			}
		}

		if foundTodo == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "找不到这个 ID 所对应的任务",
			})
			return
		}

		foundTodo.IsDone = !foundTodo.IsDone
		c.JSON(http.StatusOK, gin.H{
			"message": "状态已经更新完成",
			"data":    foundTodo,
		})
	})

	r.DELETE("/todo/:id", func(c *gin.Context) {
		// A. 解析 ID (跟 PUT 一模一样)
		idStr := c.Param("id")
		targetID, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		index := -1
		for i := 0; i < len(todoList); i++ {
			if todoList[i].ID == targetID {
				index = i
				break
			}
		}

		if index == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "找不到这个 ID"})
			return
		}

		todoList = append(todoList[:index], todoList[index+1:]...)
		c.JSON(http.StatusOK, gin.H{
			"message": "删除成功",
			"count":   len(todoList), // 看看剩下的数量
		})
	})

	r.Run(":8080")
}
