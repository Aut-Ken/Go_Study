package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 厨师：只负责往管道里【发送】数据
// chan<- string 表示这个管道在这个函数里只能写，不能读（单向管道，为了安全）
func chef(ch chan<- string) {
	var sb strings.Builder
	for i := 1; i <= 3; i++ {
		sb.Reset()
		sb.WriteString("招牌红烧肉")
		sb.WriteString("#")
		sb.WriteString(strconv.Itoa(i))
		dish := sb.String()
		fmt.Println("👨‍🍳 厨师: 正在全力爆炒...", dish)
		time.Sleep(time.Second * 1)
		fmt.Println("👨‍🍳 厨师: 菜做好了，放在窗口等待取餐 ->", dish)
		ch <- dish
	}
	close(ch)
	fmt.Println("👨‍🍳 厨师: 也就是个锅铲把子，下班！")
}

func waiter(ch <-chan string) {
	fmt.Println("💁 服务员: 准备接客...")
	for dish := range ch {
		fmt.Println("💁 服务员: 拿到菜了 ->", dish)
		time.Sleep(2 * time.Second) // 模拟端菜走的耗时
		fmt.Println("💁 服务员: 客人吃完了")
	}
	fmt.Println("💁 服务员: 厨师下班了，我也收工。")
}

func main() {
	kitchenChannel := make(chan string)
	done := make(chan bool)
	go chef(kitchenChannel)
	go func() {
		waiter(kitchenChannel)
		done <- true
	}()
	<-done
	fmt.Println("🏫 老板: 全部结束，关门打烊！")
}
