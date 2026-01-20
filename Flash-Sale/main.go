package main

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Stock int64
}

var (
	iphone     = Product{Stock: 10}
	orderQueue = make(chan string, 100)
)

func worker(workerID int) {
	fmt.Printf("👷 工人 %d 启动，准备接单...\n", workerID)
	for orderID := range orderQueue {
		time.Sleep(time.Second * 1)
		fmt.Printf("✅ 工人 %d: 完成订单 %s，剩余真实库存 %d\n",
			workerID, orderID, atomic.LoadInt64(&iphone.Stock))
	}
}

func main() {
	r := gin.Default()

	for i := 1; i <= 3; i++ {
		go worker(i)
	}

	r.GET("/stock", func(c *gin.Context) {
		currentStock := atomic.LoadInt64(&iphone.Stock)
		c.JSON(http.StatusOK, gin.H{
			"stock": currentStock,
			"msg":   "速来抢购",
		})
	})

	r.POST("/buy", func(c *gin.Context) {
		leftBound := atomic.AddInt64(&iphone.Stock, -1)
		if leftBound < 0 {
			c.JSON(200, gin.H{"status": "fail", "msg": "手慢无, 商品售罄了"})
			return
		}

		orderID := fmt.Sprintf("ORD-%d", time.Now().UnixNano())
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
		defer cancel()

		select {
		case orderQueue <- orderID:
			c.JSON(200, gin.H{
				"status": "success",
				"msg":    "抢购成功啦！订单ID：" + orderID,
			})
		case <-ctx.Done():
			atomic.AddInt64(&iphone.Stock, 1)
			c.JSON(503, gin.H{
				"status": "fail",
				"msg":    "排队人数太多，系统繁忙",
			})
		}
	})

	r.Run(":8080")
}
