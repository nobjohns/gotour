package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice of %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v byte long\n", v, len(v))
	default:
		fmt.Printf("Unknown type %v\n", v)

	}
}

func main() {
	do(42)
	do("Hello")
	do(true)
}
