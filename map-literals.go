package main

import "fmt"

type Vertex struct {
	lat, long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{1.234, -2.3456},
	"Google":    Vertex{3.2345, -5.876},
}

func main() {
	fmt.Println(m)
}
