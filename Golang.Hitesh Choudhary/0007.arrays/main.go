package main

import "fmt"

func main() {
	fmt.Println("Array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
}
