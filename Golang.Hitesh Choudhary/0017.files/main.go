package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files")
	content := "My Content"

	file, err := os.Create("./myFile.txt")
	checkNil(err)

	length, err := io.WriteString(file, content)
	checkNil(err)

	fmt.Println(length)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNil(err)
	fmt.Println(string(dataByte))
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
