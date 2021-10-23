package main

import "fmt"

func main() {
	venues := []string{"Home", "Work", "School", "Bar", "Gym"}

	for _, venue := range venues {

		// switch statements
		switch venue {
		case "Home":
			greeting("Mom, I am hungry")
		case "Work", "School":
			greeting("Weather is great today")
		case "Bar":
			greeting("Hey I like your dress, but it's a little too blue")
		default:
			greeting("Sup bois")
		}

		// if statements
		// if venue == "Home" {
		// 	greeting("Mom, I am hungry")
		// } else if venue == "Work" || venue == "School" {
		// 	greeting("Weather is great today")
		// } else if venue == "Bar" {
		// 	greeting("Hey I like your dress, but it's a little too blue")
		// } else {
		// 	greeting("Sup bois")
		// }
	}
}

func greeting(greeting string) {
	fmt.Println(greeting)
}
