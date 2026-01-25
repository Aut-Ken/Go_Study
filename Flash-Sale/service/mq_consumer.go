package service

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"flash-sale/database"
	"flash-sale/models"

	"gorm.io/gorm"
)

// StartConsumer 开启消费者 (这个函数要在 main 里用 go 启动)
func StartConsumer() {
	// 1. 告诉 RabbitMQ：我要开始从 "seckill_queue" 拿东西了
	msgs, err := database.MQChannel.Consume(
		"seckill_queue", // 队列名
		"",              // consumer名 (留空自动生成)
		true,            // auto-ack: 自动确认收到 (true表示我拿到了就算成功，不用专门回复)
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		log.Fatal("消费者启动失败: ", err)
	}

	// 2. 开启一个循环，不断从通道里读消息
	// forever 是一个阻塞的 channel，为了让这个协程不退出
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// d.Body 就是我们就收到的 JSON 数据
			log.Printf("👷 收到消息: %s", d.Body)

			// A. 解析 JSON
			var msg SeckillMessage
			json.Unmarshal(d.Body, &msg)

			// B. 执行真正的下单逻辑 (直接搬运之前的数据库操作代码)
			err := createOrderInDB(msg.UserID, msg.ProductID)
			if err != nil {
				log.Printf("❌ 下单失败: %v", err)
			} else {
				log.Printf("🎉 下单成功: UserID=%d", msg.UserID)
			}
		}
	}()

	log.Println("🚀 消费者已启动，正在等待消息...")
	<-forever // 卡在这里，不让函数结束
}

// 具体的数据库操作逻辑 (私有函数)
func createOrderInDB(userID uint, productID uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 扣库存
		res := tx.Model(&models.Product{}).
			Where("id = ? AND stock > 0", productID).
			Update("stock", gorm.Expr("stock - ?", 1))

		if res.RowsAffected == 0 {
			return fmt.Errorf("库存不足")
		}

		// 2. 创建订单
		order := models.Order{
			UserID:    userID,
			ProductID: productID,
			OrderNum:  fmt.Sprintf("%d", time.Now().UnixNano()),
			Status:    1,
		}
		return tx.Create(&order).Error
	})
}
