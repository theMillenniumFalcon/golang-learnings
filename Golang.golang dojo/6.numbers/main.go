package main // this is an indicator that this is th emain package

import "fmt"

func main() {
	fmt.Println("Hello")
	var x = 1 // int 32 int64
	var x32 int32 = 1
	var x64 int64 = 1

	fmt.Println("%T, %T, %T", x, x32, x64)
	fmt.Println()

	x = int(x32)
	fmt.Println(x)
	x = int(x64)
	fmt.Println(x)

	var unsigned uint = 1
	fmt.Println(unsigned)

	y := 2
	fmt.Println(x + y)
	fmt.Println(x % y)
	fmt.Println(x == y)
	fmt.Println(x < y)

	// floating point numbers
	pi := 3.141
	fmt.Println(pi)
	fmt.Printf("%T", pi)
	fmt.Println()

	var pi32 float32 = 3.141
	fmt.Println("%T", pi32)

	pi = float64(3.141)
	fmt.Println(pi)

	ninjas := 1e100
	fmt.Println(ninjas)

	// conversion between the two
	a := 1
	b := 3.141
	fmt.Println(a)
	fmt.Println(b)

	b = float64(a)
	fmt.Println(b)

	a = int(b)
	fmt.Println(a)

}
