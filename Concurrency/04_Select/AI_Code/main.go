package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 创建一个中止管道
	abort := make(chan string)

	// 2. 开启一个协程，模拟有人在第 2 秒的时候按下了“中止按钮”
	go func() {
		// 尝试修改这里：
		// 如果改成 time.Sleep(6 * time.Second)，也就是比发射时间晚，
		// 那么火箭就能成功发射！
		time.Sleep(2 * time.Second)

		fmt.Println("⚠️  警报：检测到异常，尝试发送中止信号...")
		abort <- "发动机着火啦！"
	}()

	fmt.Println("🚀 火箭发射倒计时开始... (目标：5秒后升空)")

	// 3. select 竞赛开始
	// 选手 A: abort 管道 (2秒收到数据)
	// 选手 B: time.After (5秒触发)
	select {
	case reason := <-abort:
		// 情况一：abort 管道先收到了消息
		fmt.Printf("❌ 发射中止！原因: %s\n", reason)
		return // 退出程序，不再往下执行

	case <-time.After(5 * time.Second):
		// 情况二：5秒钟过去了，没有任何人往 abort 里发消息
		fmt.Println("🎉 倒计时结束... 3, 2, 1... 点火！发射升空！")
	}
}
