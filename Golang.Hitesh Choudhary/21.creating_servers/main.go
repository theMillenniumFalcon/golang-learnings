package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Creating servers")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:5000/get"
	response, err := http.Get(url)
	checkNilErr(err)
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("byteCount is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:5000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's start learning golang",
			"price": 0,
			"platform": "Code2Change.in"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	checkNilErr(err)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:5000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "nishank")
	data.Add("lastname", "priydarshi")
	data.Add("email", "nishankpr435@gmail.com")

	response, err := http.PostForm(myurl, data)
	checkNilErr(err)

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
