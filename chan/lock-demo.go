package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lock = sync.Mutex{}
	var wg = sync.WaitGroup{}

	val := 0

	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			lock.Lock()
			val++
			fmt.Println("goroutine1 val %d", val)
			lock.Unlock()
			time.Sleep(100)
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			lock.Lock()
			val--
			fmt.Println("goroutine2 val %d", val)
			lock.Unlock()
			time.Sleep(100)
		}
		wg.Done()
	}()

	wg.Wait()
}
