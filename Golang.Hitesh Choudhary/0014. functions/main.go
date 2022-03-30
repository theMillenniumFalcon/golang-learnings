package main

import "fmt"

func main() {
	fmt.Println("Functions")
	fmt.Println(adder(4, 3))
	fmt.Println(proAdder(2, 3, 4, 6))
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hello world"
}
