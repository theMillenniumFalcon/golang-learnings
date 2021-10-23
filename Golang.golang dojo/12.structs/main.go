package main

import "fmt"

func main() {
	type ninja struct {
		name    string
		weapons []string
		level   int
	}

	wallace := ninja{name: "wallace"}
	wallace = ninja{"Wallace", []string{"Ninja-star"}, 1}
	fmt.Println(wallace)
	fmt.Println(wallace.name)
	fmt.Println(wallace.weapons)
	fmt.Println(wallace.level)
	wallace.level++
	wallace.weapons = append(wallace.weapons, "Ninja-sword")
	fmt.Println(wallace)

	type dojo struct {
		name  string
		ninja ninja
	}

	golangDojo := dojo{"Golang dojo", wallace}

	fmt.Println(golangDojo)
	fmt.Println(golangDojo.ninja.level)
	golangDojo.ninja.level = 3
	fmt.Println(golangDojo.ninja.level)
	fmt.Println(wallace.level)

	type adderessDojo struct {
		name  string
		ninja *ninja
	}

	adderessed := adderessDojo{"Adderessed golangDojo", &wallace}
	fmt.Println(adderessed)
	fmt.Println(*adderessed.ninja)
	adderessed.ninja.level = 4
	fmt.Println(adderessed.ninja.level)
	fmt.Println(wallace.level)

	jonnyPointer := new(ninja) // using the new keyword, it will be returned as a pointer
	fmt.Println(jonnyPointer)
	fmt.Println(*jonnyPointer)
	jonnyPointer.name = "Jonny"
	jonnyPointer.weapons = []string{"Ninja star"}
	jonnyPointer.level = 1
	fmt.Println(*jonnyPointer)

	intern := ninjaIntern{"Intern", 1}
	intern.sayHello()
	intern.sayName()

}

type ninjaIntern struct {
	name  string
	level int
}

// a function that belongs to the ninjaIntern
func (ninjaIntern) sayHello() {
	fmt.Println("Hello")
}

func (n ninjaIntern) sayName() {
	fmt.Println(n.name)
}
