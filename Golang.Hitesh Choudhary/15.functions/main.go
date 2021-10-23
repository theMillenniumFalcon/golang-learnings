package main

import "fmt"

func main() {
	fmt.Println("functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 4, 5)
	fmt.Println("Result is: ", proRes)
	fmt.Println("My message is: ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hello World!"
}

func greeter() {
	fmt.Println("Hello from golang")
}
