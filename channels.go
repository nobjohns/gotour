//The order of output is wierd here. It should print in FIFO order. Just adding a Sleep correct the orer
//https://stackoverflow.com/questions/25795131/do-golang-channels-maintain-order

package main

import "fmt"

//import "time"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	//time.Sleep(100 * time.Microsecond)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
