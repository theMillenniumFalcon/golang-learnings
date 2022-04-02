package main

import "fmt"

func main() {
	fmt.Println("Channels")

	myCh := make(chan int)

	myCh <- 5
	fmt.Println(<-myCh)
}
