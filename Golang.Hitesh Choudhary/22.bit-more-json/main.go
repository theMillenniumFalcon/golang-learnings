package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Bit more JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcocourses := []course{
		{"React bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev_1", "js", ""}},
		{"Angular bootcamp", 299, "LearnCodeOnline.in", "def123", []string{"web-dev_2", "js", ""}},
		{"Vue bootcamp", 299, "LearnCodeOnline.in", "ghi123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcocourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursenames": "ReactJs bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev", "js"]
	}
	`)

	var lcoCourse course
	checkvalid := json.Valid(jsonDataFromWeb) // checking the json data is valid or not

	if checkvalid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
