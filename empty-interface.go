package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	i = map[string]string{"42": "43"}
	describe(i)

	i = "Hello"
	describe(i)
}
