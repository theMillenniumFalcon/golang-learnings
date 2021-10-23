package main

import "fmt"

func main() {
	// for loop
	for i, j := 0, 1; i < 10 && j < 10; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	fmt.Println()

	//for-each loop
	s := "Hello world"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %c", i, s[i])
		fmt.Println()
	}
	fmt.Println()

	for i, c := range s {
		fmt.Printf("%d %c", i, c)
		fmt.Println()
	}
	fmt.Println()

	// break
	for _, c := range s {
		if c == ' ' {
			break
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

	// continue
	for _, c := range s {
		if c == ' ' {
			continue
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

	// label
iForLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break iForLoop
			}
			fmt.Printf("%d%d", i, j)
		}
		fmt.Println()
	}

}
