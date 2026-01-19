package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once
var config map[string]string

func loadConfig() {
	fmt.Println("♻️  正在加载繁重的配置文件... (只应该出现一次)")
	time.Sleep(time.Second * 1)

	config = map[string]string{
		"db":   "mysql",
		"port": "8800",
	}
	fmt.Println("✅ 配置文件加载完毕！")
}

func getConfig() {
	once.Do(func() {
		loadConfig()
	})
	fmt.Println("   ---> 获取到配置:", config["db"])
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getConfig()
		}()
	}
	wg.Wait()
}
