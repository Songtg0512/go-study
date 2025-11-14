package main

import (
	"fmt"
	"time"
)

func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

func sendOnly(ch chan<- int) {
	for i := 1; i < 10; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch)
}

func main() {

	ch := make(chan int, 3)
	go sendOnly(ch)

	go receiveOnly(ch)

	time.Sleep(5 * time.Second)
}
