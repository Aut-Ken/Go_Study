package main

import (
	"fmt"
	"sync"
	"time"
)

var rwLock sync.RWMutex
var notice string = "初始公告：明天放假"

func readNotice(id int, wg *sync.WaitGroup) {
	rwLock.RLock()
	fmt.Printf("👀 读者 %d: 正在看公告 -> %s\n", id, notice)
	time.Sleep(time.Second * 1)
	fmt.Printf("👋 读者 %d: 看完了\n", id)
	rwLock.RUnlock()
	wg.Done()
}

func writeNotice(newContent string, wg *sync.WaitGroup) {
	rwLock.Lock()
	fmt.Println("✍️  小编: 正在修改公告，闲人避让...")
	time.Sleep(1 * time.Second)
	notice = newContent
	fmt.Println("✅ 小编: 修改完毕！")
	rwLock.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("--- 场景一：大家都在读 ---")
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go readNotice(i, &wg)
	}
	wg.Wait()
	fmt.Println("\n--- 场景二：有人要写 ---")
	wg.Add(5)
	for i := 1; i < 3; i++ {
		go readNotice(i, &wg)
	}

	time.Sleep(time.Second * 1)
	go writeNotice("紧急通知：明天不放假了！", &wg)
	for i := 3; i < 5; i++ {
		go readNotice(i, &wg)
	}

	wg.Wait()
}
