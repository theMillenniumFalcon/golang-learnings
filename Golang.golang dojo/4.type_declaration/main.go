package main

import "fmt"

func main() {
	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0

	c = celsius((f - 32) * 5 / 9)

	fmt.Println(f, c)
}
