package main

import "fmt"

type Vertex struct {
	long, lat float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		42.23456, -56.324678,
	}
	fmt.Println(m["Bell Labs"])
}
