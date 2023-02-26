package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	firstRoutine := make(chan int)
	secRoutine := make(chan int)

	rand.Seed(time.Now().UnixNano())

	go func() {
		r := rand.Intn(100)
		time.Sleep(time.Microsecond * time.Duration(r))
		firstRoutine <- r
	}()

	go func() {
		r := rand.Intn(100)
		time.Sleep(time.Microsecond * time.Duration(r))
		secRoutine <- r
	}()

	select {
	case f := <-firstRoutine:
		fmt.Println(f)
		return
	case s := <-secRoutine:
		fmt.Println(s)
		return
	}

}
