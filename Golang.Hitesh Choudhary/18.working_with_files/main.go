package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files in golang")
	content := "This needs to go in a file"

	file, err := os.Create("./mylocalfile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content) // to write some content, we use the io package

	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylocalfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is: \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
