package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 1. 定义密钥 (绝对不能泄露给别人！)
// 在真实项目中，这个应该从配置文件或环境变量里读
var jwtKey = []byte("my_secret_key_123456")

// 2. 定义 Token 里要存什么数据 (Claims)
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// 3. 核心功能：生成 Token (发票)
func GenerateToken(userID uint) (string, error) {
	// 设置有效期：24小时
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "flash-sale-demo",
		},
	}

	// 使用 HS256 算法进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成加密后的字符串
	return token.SignedString(jwtKey)
}

// 4. 核心功能：解析 Token (验票)
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
