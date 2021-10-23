package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	// const five int = 5
	// fmt.Println(five)

	// const (
	// 	a = 1
	// 	b
	// 	c = 3
	// 	d
	// )
	// fmt.Println(a, b, c, d)

	// const (
	// 	zero int = iota + 5
	// 	one
	// 	two
	// )
	// fmt.Println(zero, one, two)

	const (
		_  = 1 << (10 * iota)
		kb // all of these values are yntyped constants
		mb
		gb
		tb
		pb
		eb
		zb
		yb
	)
	fmt.Println(kb, mb, gb, tb, pb, eb)
	// bigger integer are stored as an array of unsigned integers
	// const num = math.pow(2, 10)

	fmt.Println(math.Pi)
	fmt.Println(time.February)
	fmt.Println(time.Second)
	fmt.Println(time.UTC)
	fmt.Println(big.MaxExp)
	fmt.Println(big.MinExp)
}
