package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println(languages)
	fmt.Println(languages["js"])

	delete(languages, "rb")
	fmt.Println(languages)
}
