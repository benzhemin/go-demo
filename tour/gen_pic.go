package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var a [][]uint8 = make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	// Do something.
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			a[i][j] = uint8(i * j)
		}
	}

	return a
}

func main() {
	pic.Show(Pic)
}

/*
import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
*/
