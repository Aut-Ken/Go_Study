package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个读写锁
var rwLock sync.RWMutex
var notice string = "初始公告：明天放假"

// 读者：只读，不改
func readNotice(id int, wg *sync.WaitGroup) {
	// 【关键点】这里用 RLock (Read Lock)
	// RLock 允许其他人的 RLock 同时进来，不用排队！
	rwLock.RLock()
	fmt.Printf("👀 读者 %d: 正在看公告 -> %s\n", id, notice)

	time.Sleep(1 * time.Second) // 模拟看公告耗时 1 秒

	fmt.Printf("👋 读者 %d: 看完了\n", id)
	// 读完了解锁
	rwLock.RUnlock()

	wg.Done()
}

// 写者：修改数据
func writeNotice(newContent string, wg *sync.WaitGroup) {
	// 【关键点】这里用 Lock (Write Lock)
	// Lock 是霸道的，它加上后，任何 读锁 或 写锁 都进不来
	rwLock.Lock()
	fmt.Println("✍️  小编: 正在修改公告，闲人避让...")

	time.Sleep(1 * time.Second) // 模拟修改耗时
	notice = newContent

	fmt.Println("✅ 小编: 修改完毕！")
	// 修改完解锁
	rwLock.Unlock()

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// 1. 模拟 5 个读者同时来看
	// 预期效果：他们会【几乎同时】打印“正在看公告”，而不是隔1秒打印一个
	fmt.Println("--- 场景一：大家都在读 ---")
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go readNotice(i, &wg)
	}
	wg.Wait()

	fmt.Println("\n--- 场景二：有人要写 ---")
	wg.Add(2)
	// 2. 一个读者先来
	go readNotice(99, &wg)
	// 3. 小编紧接着要改
	go writeNotice("紧急通知：明天不放假了！", &wg)

	wg.Wait()
}
