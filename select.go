// This is another version of select.go
// Here, fibonacci() running in another goroutine, and anonymous function is replaced with named function getFibonacci()

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func getFibonacci(n int, c, quit chan int) {
	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go fibonacci(c, quit)
	getFibonacci(10, c, quit)
}
