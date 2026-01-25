package models

import "gorm.io/gorm"

// Order 订单表
type Order struct {
	gorm.Model
	UserID    uint   `json:"user_id"`    // 谁买的
	ProductID uint   `json:"product_id"` // 买了啥
	OrderNum  string `json:"order_num"`  // 订单号 (比如 202601250001)
	Status    int    `json:"status"`     // 1: 待支付, 2: 已支付, 3: 已发货
}

// 注意：真实生产环境通常会把 UserID 和 ProductID 建一个联合唯一索引
// 保证：一个用户对同一个秒杀商品只能下一单
