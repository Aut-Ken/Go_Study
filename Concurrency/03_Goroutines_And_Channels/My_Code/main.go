package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("👷 柜员 %d: 开始办理业务 [客户ID: %d]\n", id, j)
		time.Sleep(time.Second * 1)
		fmt.Printf("👷 柜员 %d: 办完了 [客户ID: %d]\n", id, j)
		results <- 2 * j
	}
	fmt.Printf("👷 工人 %d: 没活了，下班！\n", id)
}

func manager(count int, jobs chan<- int) {
	fmt.Println("👨‍💼 经理: 我开始派单了，你们准备接单！")
	for i := 1; i <= count; i++ {
		fmt.Printf("👨‍💼 经理: 派发任务 #%d\n", i)
		jobs <- i
		time.Sleep(200 * time.Millisecond) // 模拟经理派单也要点时间
	}
	close(jobs)
	fmt.Println("👨‍💼 经理: 单派完了，我先下班喝茶去了。")
}

func main() {
	const jobNums = 5
	jobs := make(chan int, jobNums)
	results := make(chan int, jobNums)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	go manager(jobNums, jobs)

	time.Sleep(time.Second * 1)

	fmt.Println("🚪 主程: 所有人都在干活了，我在门口等结果...")

	for a := 1; a <= jobNums; a++ {
		res := <-results
		fmt.Printf("✅ 主程: 收到结果 -> %d\n", res)
	}
	fmt.Println("🎉 所有业务办理完成！")
}
