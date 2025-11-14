package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 表示一个可执行的任务（函数）
type Task func()

func schedule(tasks []Task) []time.Duration {

	var wg sync.WaitGroup
	// 等待所有任务执行完毕
	wg.Add(len(tasks))

	durations := make([]time.Duration, len(tasks))

	for i, t := range tasks {
		i, task := i, t // 避免闭包引用问题
		go func() {
			defer wg.Done()
			start := time.Now()
			task()                           // 执行任务
			durations[i] = time.Since(start) // 记录耗时
		}()
	}

	wg.Wait()
	return durations
}

func main() {

	var tasks []Task = []Task{
		func() {
			fmt.Println("我是任务一：-------")
			time.Sleep(1 * time.Second)
		},
		func() {
			fmt.Println("我是任务二：-------")
			time.Sleep(2 * time.Second)
		},
		func() {
			fmt.Println("我是任务三：-------")
			time.Sleep(3 * time.Second)
		},
	}

	result := schedule(tasks)

	for i, v := range result {
		fmt.Println("任务 ", i, " 的执行时间为 ", v)
	}

}
