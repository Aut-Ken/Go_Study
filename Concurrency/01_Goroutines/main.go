package main

import (
	"fmt"
	"runtime"
	"time"
)

func sayHello(times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("sayHello: %d Hello Golang\n", i)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	go sayHello(10)
	for i := 0; i < 10; i++ {
		fmt.Printf("main: %d 你好\n", i)
		time.Sleep(time.Millisecond)
	}
	time.Sleep(time.Second)
	runtime.GOMAXPROCS(8)
	fmt.Println(runtime.NumCPU())
}
