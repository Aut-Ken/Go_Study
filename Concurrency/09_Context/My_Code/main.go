package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var ctxRoot context.Context = context.Background()
	ctxParent, cancel := context.WithTimeout(ctxRoot, time.Second*2)
	defer cancel()
	ctxChild := context.WithValue(ctxParent, "name", "XiaoMing")

	go func(ctx context.Context) {
		fmt.Println("👶 儿子: 开始干活...")
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("👶 %v: 我挂了！原因: %v\n", ctx.Value("name"), ctx.Err())
				return
			default:
				fmt.Printf("👶 %v: 还在干活...\n", ctx.Value("name"))
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctxChild)

	select {
	case <-ctxParent.Done():
		fmt.Println("👴 主程: 时间到了，所有人停手！")
	}
	time.Sleep(1 * time.Second)
}
