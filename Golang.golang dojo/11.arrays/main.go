package main

import "fmt"

func main() {
	fmt.Println("Hello")

	var a [3]int
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(len(a))
	fmt.Println()

	a[0] = 10
	fmt.Println(a)

	aCopy := a
	fmt.Println(aCopy == a)
	fmt.Println(aCopy)
	a[0] = 2
	fmt.Println(a)
	fmt.Println(aCopy)
	fmt.Println(aCopy == a)

	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)
	c := [...]int{1, 2}
	fmt.Println(c)
	d := [3]int{1, 2}
	fmt.Println(d)
	fmt.Println(d == a)
	fmt.Println()

	for _, v := range b {
		fmt.Println(v)
	}
	fmt.Println()

	array := [...]int{0: 1}
	fmt.Println(array)
	array2 := [...]int{1: 2, 3: 4}
	fmt.Println(array2)

	stringArray := [...]string{"string1", "string2"}
	fmt.Println(stringArray)

	array2d := [2][2]int{{1, 2}, {2, 3}}
	fmt.Println(array2d)
	array3d := [2][2][2]int{array2d, array2d}
	fmt.Println(array3d)
}
