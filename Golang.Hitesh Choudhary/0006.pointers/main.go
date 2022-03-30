package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println(ptr)

	myNumbber := 23
	var ptr = &myNumbber
	fmt.Println(ptr, *ptr)
}
