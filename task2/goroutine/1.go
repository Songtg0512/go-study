package main

import "sync"

var wait sync.WaitGroup

func printOddNumber() {
	defer wait.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			println("奇数：", i)
		}
	}
}

func printEvenNumber() {
	defer wait.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			println("偶数：", i)
		}
	}
}

func main() {

	wait.Add(2)

	go printOddNumber()

	go printEvenNumber()

	wait.Wait()

}
