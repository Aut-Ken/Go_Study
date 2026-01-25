package database

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

// 全局 Redis 客户端
var RDB *redis.Client

// 上下文 (Redis 操作需要这个)
var Ctx = context.Background()

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // 刚才安装的默认没密码，留空
		DB:       0,                // 默认使用 0 号数据库
	})

	// 测试连接
	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Redis 连接失败: ", err)
	}
	fmt.Println("⚡ Redis 连接成功！")
}
