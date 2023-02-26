package main

import "fmt"

func main() {

	var arr []int = make([]int, 10)
	for i, _ := range arr {
		arr[i] = i + 1
	}

	for i, _ := range arr {
		fmt.Printf("value %d, addr %p\n", arr[i], &arr[i])
	}
}
