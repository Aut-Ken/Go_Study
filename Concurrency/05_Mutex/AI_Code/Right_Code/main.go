package main

import (
	"fmt"
	"sync"
)

var balance int = 0

// 1. 定义一把锁（Mutex = Mutual Exclusion 互斥锁）
var lock sync.Mutex

func deposit(wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		// 2. 进门前先上锁
		// 如果此时已经有人锁了，我会在这里卡住排队，直到他解锁
		lock.Lock()

		// --- 临界区 (Critical Section) 开始 ---
		balance = balance + 1
		// --- 临界区 结束 ---

		// 3. 办完事，一定要解锁！如果不解，后面的人就永远死锁了
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go deposit(&wg)
	go deposit(&wg)

	wg.Wait()
	fmt.Printf("💰 最终余额: %d\n", balance)
}
