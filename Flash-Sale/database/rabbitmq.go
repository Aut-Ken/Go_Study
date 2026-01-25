package database

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 全局变量
var MQConn *amqp.Connection
var MQChannel *amqp.Channel

func InitRabbitMQ() {
	var err error
	// 1. 连接 RabbitMQ 服务
	// 格式: amqp://账号:密码@IP:端口/
	url := "amqp://guest:guest@localhost:5672/"
	MQConn, err = amqp.Dial(url)
	if err != nil {
		log.Fatal("❌ RabbitMQ 连接失败: ", err)
	}

	// 2. 创建一个通道 (Channel)
	// 我们的大部分操作（发消息、收消息）都是在 Channel 上进行的
	MQChannel, err = MQConn.Channel()
	if err != nil {
		log.Fatal("❌ RabbitMQ Channel 创建失败: ", err)
	}

	// 3. 声明一个队列 (Queue)
	// 这一步是为了保证队列存在，如果不存在会自动创建
	// 名字叫 "seckill_queue"
	_, err = MQChannel.QueueDeclare(
		"seckill_queue", // 队列名字
		true,            // durable: 是否持久化 (重启还在吗？true=在)
		false,           // autoDelete: 没人用时是否自动删除
		false,           // exclusive: 是否由当前连接独占
		false,           // noWait: 是否非阻塞
		nil,             // args: 其他参数
	)
	if err != nil {
		log.Fatal("❌ 队列声明失败: ", err)
	}

	log.Println("🐰 RabbitMQ 连接并初始化成功！")
}