package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"flash-sale/database"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 定义一个简单的消息结构体，用来打包数据
type SeckillMessage struct {
	UserID    uint
	ProductID uint
}

// BuyProduct 升级版：Redis + RabbitMQ 异步下单
func BuyProduct(userID uint, productID uint) error {
	// 1. Redis 拦截 (完全不变)
	stockKey := fmt.Sprintf("product:%d:stock", productID)
	result, err := database.RDB.Decr(database.Ctx, stockKey).Result()
	if err != nil {
		return err
	}
	if result < 0 {
		return errors.New("Redis: 手慢了！商品已抢光")
	}

	// 2. 发送消息到 RabbitMQ (不再直接操作数据库)

	// A. 把要传的数据打包成 JSON
	msg := SeckillMessage{UserID: userID, ProductID: productID}
	msgBody, _ := json.Marshal(msg)

	// B. 发送！
	err = database.MQChannel.Publish(
		"",              // exchange: 交换机 (空字符串表示使用默认交换机)
		"seckill_queue", // routing_key: 路由键 (这就写我们的队列名)
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msgBody, // 这里的 Body 就是上面的 JSON 数据
		},
	)

	if err != nil {
		log.Printf("❌ 消息发送失败: %v", err)
		// 如果消息发失败了，理论上应该把 Redis 的库存加回去 (这里省略，为了简化)
		return err
	}

	log.Printf("✅ 用户 %d 的请求已入队", userID)
	return nil
}
