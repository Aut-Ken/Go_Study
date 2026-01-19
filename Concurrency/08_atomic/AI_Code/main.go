package main

import (
	"fmt"
	"sync"
	"sync/atomic" // 引入原子包
)

func main() {
	var ops int64 = 0 // 定义一个 64 位的整数计数器
	var wg sync.WaitGroup

	// 模拟 50 个协程，每人都要给 ops 加 1000 次
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				// ❌ 错误做法：ops++ (会有并发问题)
				// 🆗 笨重做法：lock.Lock(); ops++; lock.Unlock()

				// ✅ 原子做法：
				// 参数一：要改谁的地址（&ops）
				// 参数二：加多少（1）
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	wg.Wait()

	// ❌ 错误读取：fmt.Println(ops)
	// 虽然 Print 这一刻一般没事，但在高并发运行中，直接读变量也是不安全的

	// ✅ 原子读取：LoadInt64
	safeValue := atomic.LoadInt64(&ops)
	fmt.Println("🚀 最终计数:", safeValue)
}
