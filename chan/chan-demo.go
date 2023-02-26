package main

import "fmt"

func main() {
	var val = make(chan int)

	go func() {
		fmt.Println("input value 1")
		val <- 1
	}()

	go func() {
		fmt.Println("input value 2")
		val <- 2
	}()

	ans := []int{}
	for {
		ans = append(ans, <-val)
		fmt.Println(ans)
		if len(ans) >= 2 {
			break
		}
	}
}
