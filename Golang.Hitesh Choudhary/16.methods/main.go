package main

import "fmt"

func main() {
	fmt.Println("methods in go")

	// nno inheritance in golang; No super or parent
	nishank := User{"Nishank", "nishank@go.dev", true, 19}
	fmt.Println(nishank)
	fmt.Printf("Nishank's details: %+v\n", nishank)
	fmt.Printf("Name is %v and email is %v\n", nishank.Name, nishank.Email)
	nishank.GetStatus()
	nishank.NewMail()
	fmt.Printf("Name is %v and email is %v\n", nishank.Name, nishank.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Eail of this user is: ", u.Email)
}
