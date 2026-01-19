package main

import (
	"fmt"
	"time"
)

func main() {
	abort := make(chan string)

	go func() {
		time.Sleep(time.Second * 8)
		fmt.Println("⚠️  警报：检测到异常，尝试发送中止信号...")
		str := "发动机着火啦！"
		abort <- str
	}()

	fmt.Println("🚀 火箭发射倒计时开始... (目标：5秒后升空)")
LoopEnd:
	for {
		select {
		case reason := <-abort:
			fmt.Printf("❌ 发射中止！原因: %s\n", reason)
			break LoopEnd
		case <-time.After(5 * time.Second):
			fmt.Println("🎉 倒计时结束... 3, 2, 1... 点火！发射升空！")
		}
	}

	n := 4
LoopEnd2:
	for {
		for i := range n {
			if i == 1 {
				break LoopEnd2
			}
		}
	}
}
