package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Taking input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number: ")

	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
