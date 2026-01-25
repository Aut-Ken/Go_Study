package models

import "gorm.io/gorm"

// Product 商品表
type Product struct {
	gorm.Model         // 自动加上 ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`  // 📦 库存！这是秒杀的核心
	Image      string  `json:"image"`  // 商品图片 URL
	Status     int     `json:"status"` // 1: 上架, 0: 下架
}
