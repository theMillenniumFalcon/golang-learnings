package main

import "fmt"

func index() {
	// how to remove a value from slice based on index
	fmt.Println("Print-indexes")

	var courses = []string{"react", "vue", "angular", "javascript"}
	fmt.Println(courses)
	var index int = 2
	// we can use append to remove the values from array
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
