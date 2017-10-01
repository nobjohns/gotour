package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func do() error {
	return &MyError{time.Now(), "It didn't work"}
}
func main() {
	if err := do(); err != nil {
		fmt.Println(err)
	}
}
