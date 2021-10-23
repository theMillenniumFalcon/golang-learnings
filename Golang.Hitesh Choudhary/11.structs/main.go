package main

import "fmt"

func main() {
	fmt.Println("structs")
	// there is no inheritance in golang, No super or parent

	nishank := User{"Nishank", "nishank@go.dev", true, 19}
	fmt.Println(nishank)
	fmt.Printf("Nishank's details: %+v\n", nishank)
	fmt.Printf("Name is %v and email is %v\n", nishank.Name, nishank.Email)

}

// User has to be accessed from anywhere, a convention to write uppercase
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
