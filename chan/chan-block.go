package main

import (
	"fmt"
	"time"
)

func main() {
	forever := make(chan int)

	go func() {
		fmt.Println("go routine")
		time.Sleep(time.Second)

		forever <- 1
	}()

	fmt.Println("before exit")
	<-forever
	fmt.Println("after exit")
}
