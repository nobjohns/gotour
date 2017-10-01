package main

import (
	"fmt"
	"math"
)

const delta = 1e-8

func Sqrt(x float64) float64 {
	z, t := 1.0, 0.0
	for {
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(z-t) < delta {
			break
		}
	}
	return z
}

func main() {
	i := 2.0
	fmt.Println(Sqrt(i))
	fmt.Println(Sqrt(i) == math.Sqrt(i))
}
