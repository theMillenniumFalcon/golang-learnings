package main

import "fmt"

func main() {
	fmt.Println("Hello")

	// var ptr *int
	// fmt.Println("Value of pointer is: ", ptr)

	myNumber := 23

	var ptr = &myNumber // creating a pointer that is referencing to some memory

	fmt.Println("Value of pointer is: ", ptr)
	fmt.Println("Value of pointer is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of pointer is: ", myNumber) // the original value is changed not the duplicate value
}
