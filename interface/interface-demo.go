package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{2, 3}
	fmt.Println("area: %d", (&r).area())
	fmt.Println("perim: %d", r.perim())
}
