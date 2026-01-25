package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// 👇 确保这里的 import 路径和你 go.mod 里的名字一致
	"flash-sale/service"
)

// ================= 商品/秒杀相关 =================

// Buy 秒杀接口 (JWT 改造版)
// 之前这个文件里不小心放了 Login，现在我们要把它改回 Buy
func Buy(c *gin.Context) {
	// 1. 从 Context 中获取 UserID (由中间件 AuthMiddleware 注入)
	// 这是一个安全操作，只有通过了中间件的请求才会带有 userID
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权，请先登录"})
		return
	}
	userID := userIDRaw.(uint) // 类型断言：把 interface{} 转成 uint

	// 2. 解析商品 ID
	var req struct {
		ProductID uint `json:"product_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 3. 调用 Service 进行异步秒杀
	// 注意：这里不再需要传 userID 给前端，而是我们自己从 token 解析出来的
	err := service.BuyProduct(userID, req.ProductID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "已进入队列排队..."})
}
