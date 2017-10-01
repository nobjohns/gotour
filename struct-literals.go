package main

import "fmt"

type vertex struct {
	x, y int
}

var (
	v1 = vertex{1, 2}
	p  = &vertex{1, 2}
	v2 = vertex{x: 1}
	v3 = vertex{}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
