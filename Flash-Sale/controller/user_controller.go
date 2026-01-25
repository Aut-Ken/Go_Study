package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// 👇 注意：这里的 "flash-sale-demo" 必须改成你 go.mod 第一行写的那个名字！
	"flash-sale/models"
	"flash-sale/service"
	"flash-sale/utils"
)

// ================= 用户相关 =================

// Login 登录接口
func Login(c *gin.Context) {
	// 定义临时的请求结构体
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 模拟数据库校验 (以后可以换成 service.CheckUser(req.Username, req.Password))
	if req.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	// 假设这是 ID=1 的用户
	realUserID := uint(1)

	// 生成 Token
	token, err := utils.GenerateToken(realUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成Token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
	})
}

// Transfer 转账接口
func Transfer(c *gin.Context) {
	var req struct {
		FromID uint `json:"from_id"`
		ToID   uint `json:"to_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不对"})
		return
	}

	err := service.TransferLife(req.FromID, req.ToID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "msg": "转账成功"})
}

// QueryUsers 查询接口
func QueryUsers(c *gin.Context) {
	var params models.UserSearchParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": "参数格式不对"})
		return
	}

	users, err := service.GetUsers(params)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": users, "count": len(users)})
}
