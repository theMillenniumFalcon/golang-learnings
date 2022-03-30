package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and values is %v\n", index, day)
	}

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto sample
		}

		if rougueValue == 5 {
			break
		}

		if rougueValue == 6 {
			rougueValue++
			continue
		}

		fmt.Println("value is: ", rougueValue)
		rougueValue++
	}

sample:
	fmt.Println("Jump to this")

}
