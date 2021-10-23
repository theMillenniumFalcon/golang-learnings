package main

import (
	"fmt"
	"os"
	"text/template"
)

// creating a type to store the two peices of information
type secret struct {
	Name   string
	Secret string
}

func main() {
	fmt.Println("Hello")

	secret := secret{"Wallace", "I'm a ninja!"}

	templatePath := "F:/Study/Golang.golang dojo/16.golang-text-template/template.txt"

	t, err := template.New("template.txt").ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, secret)
	if err != nil {
		fmt.Println(err)
	}
}
