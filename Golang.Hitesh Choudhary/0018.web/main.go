package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)
	checkNil((err))
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	checkNil(err)
	content := string(dataBytes)
	fmt.Println(content)
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
