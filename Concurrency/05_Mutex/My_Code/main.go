package main

import (
	"fmt"
	"sync"
)

var money int = 0
var lock sync.Mutex

func deposit(wg *sync.WaitGroup) {
	for i := 0; i < 100000; i++ {
		lock.Lock()
		money += 1
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
	fmt.Printf("💰 最终余额: %d\n", money)
}
