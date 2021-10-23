package main // this is an indicator that this is the main package

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(3.000001))
	fmt.Println(math.Floor(3.000001))
	fmt.Println(math.Min(5, 8))
	fmt.Println(math.Max(1, 0))
	fmt.Println(math.Abs(3.000001))
	fmt.Println(math.Sqrt(3.000001))
	fmt.Println(math.Pow(2, 3))

	// complex numbers
	fmt.Println(complex(5, 7))
}
