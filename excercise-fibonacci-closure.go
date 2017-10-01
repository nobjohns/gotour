package main

import "fmt"

func fibonacci() func() int {
	m, n := 1, 0
	return func() int {
		m, n = n, m+n
		return m
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
