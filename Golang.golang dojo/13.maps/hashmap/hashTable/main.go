package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m == nil) // true --> as zero value of maps is nil
	m = map[string]string{}
	fmt.Println(m == nil)
	fmt.Println(m)
	fmt.Println(len(m))
	m = make(map[string]string, 5)
	fmt.Println(len(m))
	m = map[string]string{"wallace": "programmer"}
	fmt.Println(m)
	fmt.Println(len(m))
	m["Johnny"] = "Designer"
	fmt.Println(m)
	// a default map is not ordered
	m["Johnny"] = "To be a programmer"
	fmt.Println(m)
	fmt.Println(m["Johnny"]) // map value where key is Johnny

	delete(m, "Johnny")
	fmt.Println(m)
	m["Wallace"] += " Ninja"
	fmt.Println(m)
	m["Johnny"] = "To be a programmer"

	// maps are iterable data structures
	for name, title := range m {
		fmt.Println(name, title)
	}

	fmt.Println(m["Johnny"])
	m["Johnny"] = "To be a programmer"

	title, ok := m["Johnny"]
	if ok {
		fmt.Println(title)
	} else {
		fmt.Println("Didn't find Johnny")
	}
	title, _ = m["Johnny"]
	title = m["Johnny"]

}
