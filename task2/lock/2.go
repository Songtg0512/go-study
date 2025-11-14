package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 = 0
	var wg sync.WaitGroup

	numGoroutine := 10
	incrementPerGoroutine := 1000

	wg.Add(numGoroutine)

	for i := 0; i < numGoroutine; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementPerGoroutine; j++ {
				atomic.AddInt64(&counter, 1) // 原子递增
			}
		}()
	}

	wg.Wait()

	fmt.Println("最终计数器的值:", counter)
}
