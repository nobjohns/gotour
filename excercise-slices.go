package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for x := range a {
		a[x] = make([]uint8, dx)
		for y := range a[x] {
			a[x][y] = (uint8(x) + uint8(y)) / 2
		}
	}
	return a
}

func main() {
	fmt.Println(Pic(4, 5))
	pic.Show(Pic)
}
