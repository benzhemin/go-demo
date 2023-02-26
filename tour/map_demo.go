package main

import "fmt"

type Vertix struct {
	x int
	y int
}

var m map[string]Vertix

func main() {
	m := make(map[string]Vertix)
	m["Bell"] = Vertix{1, 10}

	fmt.Println(m["Bell"])
}
