package middleware

import (
	"net/http"
	"strings"

	"flash-sale/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 是认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从请求头获取 Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请先登录 (缺少Token)"})
			c.Abort() // ⛔ 拦截！不准往下走
			return
		}

		// 2. 格式通常是 "Bearer xxxxxxx"
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token格式错误"})
			c.Abort()
			return
		}

		// 3. 解析 Token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token无效或已过期"})
			c.Abort()
			return
		}

		// 4. ✅ 验证成功！把 UserID 存到上下文里，供后面的 Controller 使用
		c.Set("userID", claims.UserID)

		// 放行，进入下一个环节
		c.Next()
	}
}
