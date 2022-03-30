package main

import "fmt"

func main() {
	fmt.Println("Structs")

	nishank := User{"nishank", "nishank@nishank.com", true, 20}
	fmt.Println(nishank)
	fmt.Printf("nishank details are: %+v", nishank)
	fmt.Printf("name is %v and email is %v.", nishank.Name, nishank.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
