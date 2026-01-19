package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 定义一个 once 对象
var once sync.Once

// 模拟的配置数据
var config map[string]string

// 初始化函数（这是我们只想执行一次的任务）
func loadConfig() {
	fmt.Println("♻️  正在加载繁重的配置文件... (只应该出现一次)")
	time.Sleep(1 * time.Second) // 模拟耗时

	config = map[string]string{
		"db":   "mysql",
		"port": "8080",
	}
	fmt.Println("✅ 配置文件加载完毕！")
}

// 业务函数：获取配置
func getConfig() {
	// 【关键】无论在这个函数里调用多少次，loadConfig 只会跑一次
	// 这里的 func() 是一个闭包，把 loadConfig 包进去
	once.Do(func() {
		loadConfig()
	})

	// 只有等上面的 Do 执行完了（或者发现已经执行过了），才会走到这一行
	fmt.Println("   ---> 获取到配置:", config["db"])
}

func main() {
	var wg sync.WaitGroup

	// 模拟 10 个协程同时去获取配置
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getConfig()
		}()
	}

	wg.Wait()
}
