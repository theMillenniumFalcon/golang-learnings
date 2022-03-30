package main

import "fmt"

func main() {
	fmt.Println("Methods")

	peter := User{"Peter", "peter@peter.com", true, 14}
	peter.GetStatus()
	peter.NewEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println(u.Name)
}

func (u User) NewEmail() {
	u.Email = "random@gmail.com"
	fmt.Println(u.Email)
}
