package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer fmt.Println("go routine drop out")

		fmt.Println("start a go routine")
		time.Sleep(time.Second)
	}()

	fmt.Println("wait a goroutine")
	wg.Wait()
}
