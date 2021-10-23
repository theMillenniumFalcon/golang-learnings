package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("handling web requests in go")
	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("Response: is of type: %T\n", response)
	defer response.Body.Close() // caller's responsibility to close the connection
	databytes, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	content := string(databytes)
	fmt.Printf(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
