package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Int64
	ops.Store(10)
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
		}()
	}

	wg.Wait()
	val := ops.Load()
	fmt.Println("🚀 最终计数:", val)
}
