package main

import "fmt"

func main() {
	fmt.Println("Hello")

	var fruitList [4]string

	fruitList[0] = "Mango"
	fruitList[1] = "Tomato"
	fruitList[2] = "Peach"

	fmt.Println("FritList: ", fruitList)
	fmt.Println("FritList: ", len(fruitList)) // 4

	var vegList = [3]string{"potato", "tomato", "mushroom"}
	fmt.Println("VegList: ", vegList)
	fmt.Println("VegList: ", len(vegList))
}
