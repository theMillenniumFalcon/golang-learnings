package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	// arrays are non mutable in go
	fixed := [...]int{1, 2, 3, 4}
	fmt.Println(fixed)
	// fixed = [...]int{1, 2, 3, 4, 5} // --> error
	a := []int{0, 1}
	fmt.Println(a)
	fmt.Println(len(a))
	a = []int{1, 2, 3, 4}
	fmt.Println(a)
	fmt.Println(len(a))
	a = append(a, 4, 5, 6)
	fmt.Println(a)
	fmt.Println(cap(a), len(a))

	b := make([]int, 5)
	fmt.Println(b)
	fmt.Println()

	// a slice of a slice
	fmt.Println(a)
	a = a
	a = a[0:7]
	fmt.Println(a)
	a = a[0:8]
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = a[0:7]
	fmt.Println(a)
	a = a[0:]
	fmt.Println(a)
	a = a[:7]
	fmt.Println(a)
	fmt.Println()

	if a == nil {
	}
	var c []int           // this will asign a zero value to the slice
	fmt.Println(c == nil) // true
}
