// There are three levels of scopes --> block, package, universal

package main

import "fmt"

func main() {
	// block scope
	{
		var integer = 2
		fmt.Println(integer)
	}
	var integer = 1
	fmt.Println("hello")
	fmt.Println(integer)
}
