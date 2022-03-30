package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Create Json")
	EncodeJson()
}

func EncodeJson() {
	Courses := []course{
		{"React", 299, "learnCode", "password", []string{"js", "web-dev"}},
		{"Machine Learning", 250, "learnCode", "password", nil},
	}

	// package this data  into json data
	finalJson, err := json.MarshalIndent(Courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{

		}
	`)

	var Course course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &Course)
		fmt.Printf("%#v", Course)
	} else {
		fmt.Println("Json was not valid")
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v, value is %v and type is %T\n", k, v, v)
	}
}
