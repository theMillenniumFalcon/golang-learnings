package main

import "fmt"

func main() {
	// fmt.Println("Slices")

	// var fruitList = []string{"Apple", "Tomato", "Banana"}
	// fmt.Printf("Type is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Peach")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	// highScores := make([]int, 4)
	// highScores[0] = 234
	// highScores[1] = 744
	// highScores[2] = 203
	// highScores[3] = 981

	// highScores = append(highScores, 425, 321, 781)
	// append method reallocates the memory

	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))
	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"react", "javascript", "golang", "python", "design"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Print(courses)
}
