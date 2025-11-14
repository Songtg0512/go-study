package main

import "sync"

func main() {

	var sn sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	var i = 10

	go func() {
		defer wg.Done()
		for j := 0; j < 100; j++ {
			sn.Lock()
			i += 1
			sn.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for j := 0; j < 200; j++ {
			sn.Lock()
			i += 1
			sn.Unlock()
		}
	}()

	wg.Wait()
	println(i)

}
