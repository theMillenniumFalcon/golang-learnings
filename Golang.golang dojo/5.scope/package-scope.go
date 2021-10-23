// host/userOrganization/project/(dir)/package

package main

import "fmt"

var five = 5

func main() {
	// block scope
	{
		var integer = 2
		fmt.Println(integer)
		fmt.Println(five)
	}
	var integer = 1
	fmt.Println("hello")
	fmt.Println(integer)
	// fmt.Println(three)
}

func nonMain() {
	fmt.Println(five)
}
