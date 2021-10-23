package main

import "fmt"

const LoginToken string = "gtgt" // this is equivalent to  a public declaration

func main() {
	var username string = "nishank"
	fmt.Println(username)
	fmt.Printf("Variables is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type: %T \n", smallVal)

	// in case of float32, we get 5 values after the decimal, for more precision float64 is used
	var smallFloat float32 = 255.457327838
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable) // 0
	fmt.Printf("Variables is of type: %T \n", anotherVariable)

	// implicit way to declare the variables
	var website = "goole.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000 // walrus operator --> walrus operator can be used only inside a method
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}
