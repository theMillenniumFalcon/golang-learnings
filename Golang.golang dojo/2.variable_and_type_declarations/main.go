package main

import "fmt"

func main() {

	// var name type = expression
	var integer int = 5
	fmt.Println(integer)
	// if we donot mention expression, integer will be given a 0 value by default
	var integer1, integer2 int
	fmt.Println(integer1, integer2)
	// we can also ommit type and write expressions directly
	var integer3, string = 45, "string"
	fmt.Println(integer3, string)
	// short declaration --> name:= expression
	integer4 := 76
	fmt.Println(integer4)
}
