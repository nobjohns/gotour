package main

import (
	"fmt"
	"math"
)

const delta = 1e-8

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z, t := 1.0, 0.0
	for {
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(z-t) < delta {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
