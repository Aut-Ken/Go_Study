package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 爷爷 (root)
	ctxRoot := context.Background()

	// 爸爸 (parent): 设了 2 秒超时
	ctxParent, cancel := context.WithTimeout(ctxRoot, 2*time.Second)
	defer cancel()

	// 儿子 (child): 它是基于 parent 创建的
	// 虽然儿子自己没设超时，但因为它的根是 parent
	// 所以 parent 一挂，child 也会立刻挂
	ctxChild := context.WithValue(ctxParent, "name", "小明")

	go doTask(ctxChild)

	// 等待结果
	select {
	case <-ctxParent.Done():
		fmt.Println("👴 主程: 时间到了，所有人停手！")
	}
	time.Sleep(1 * time.Second)
}

func doTask(ctx context.Context) {
	fmt.Println("👶 儿子: 开始干活...")
	for {
		select {
		case <-ctx.Done():
			// ctx.Err() 会告诉你到底是因为超时(DeadlineExceeded)还是被取消(Canceled)
			fmt.Printf("👶 儿子: 我挂了！原因: %v\n", ctx.Err())
			return
		default:
			fmt.Println("👶 儿子: 还在干活...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
