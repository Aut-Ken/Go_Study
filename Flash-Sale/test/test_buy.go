package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	// 👇 这里的 "flash-sale" 必须和你 go.mod 里的名字保持一致！
	// 如果报错找不到包，请检查 go.mod 第一行
	"flash-sale/utils"
)

// 新的请求结构体 (不需要 UserID 了，只要 ProductID)
type BuyRequest struct {
	ProductID uint `json:"product_id"`
}

func main() {
	// 设置并发人数：1000 人抢 50 台
	const peopleCount = 1000

	fmt.Printf("🔥 开始模拟 %d 个持票用户同时抢购 iPhone ...\n", peopleCount)

	var wg sync.WaitGroup
	startTime := time.Now()

	// 创建一个 HTTP 客户端 (复用连接，性能更好)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for i := 1; i <= peopleCount; i++ {
		wg.Add(1)

		// 启动协程模拟用户
		go func(uid int) {
			defer wg.Done()

			// 1. 【现场造票】直接调用工具函数生成 Token
			// 模拟这就是 UserID = uid 的用户
			token, err := utils.GenerateToken(uint(uid))
			if err != nil {
				fmt.Printf("用户 %d 生成Token失败: %v\n", uid, err)
				return
			}

			// 2. 准备请求数据
			reqBody := BuyRequest{
				ProductID: 1,
			}
			jsonData, _ := json.Marshal(reqBody)

			// 3. 创建请求对象 (必须用 NewRequest 才能设置 Header)
			req, err := http.NewRequest("POST", "http://localhost:8080/buy", bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("创建请求失败:", err)
				return
			}

			// 4. 【关键步骤】把 Token 塞进 Header
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token) // 注意 Bearer 后面有空格

			// 5. 发送请求
			resp, err := client.Do(req)
			if err != nil {
				// 网络层面的错误（比如连接被拒绝），不打印详细日志以免刷屏
				// fmt.Printf("请求发送失败: %v\n", err)
				return
			}
			defer resp.Body.Close()

			// 6. 简单的结果检查
			// 只有 200 OK 且 body 里包含 "success":true 才算真正进入队列
			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyString := string(bodyBytes)

			if resp.StatusCode == 200 {
				// 只有成功的时候打印一下，证明通了
				// fmt.Printf("用户 %d 请求响应: %s\n", uid, bodyString)
			} else {
				// 如果是 401 说明 Token 没带对
				fmt.Printf("用户 %d 鉴权失败 (%d): %s\n", uid, resp.StatusCode, bodyString)
			}

		}(i)
	}

	wg.Wait()
	elapsed := time.Since(startTime)
	fmt.Printf("🏁 抢购结束！耗时: %v\n", elapsed)
}
