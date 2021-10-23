package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")
	var fruitList = []string{"Apple", "Tomato", "Orange"}
	fmt.Println("Type of fruitList is %T", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 235
	highscore[2] = 236
	highscore[3] = 237
	// highscore[4] = 237 // error
	highscore = append(highscore, 555, 657, 542, 985)
	fmt.Println(highscore)
	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted((highscore)))
}
