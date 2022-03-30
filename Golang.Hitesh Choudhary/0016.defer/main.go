package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Defer")
}
