package main

import (
	"fmt"
	"sort"
)

func main() {
	var integers []int

	for {
		var input string
		fmt.Println("Enter an integer (or 'X' to quit): ")
		fmt.Scan(&input)

		if input == "X" || input == "x" {
			break
		}

		var integer int
		_, err := fmt.Sscanf(input, "%d", &integer)
		if err != nil {
			fmt.Println("Please enter a valid integer.")
			continue
		}
		integers = append(integers, integer)

		sort.Ints(integers)

		fmt.Println("Sorted Slice : ", integers)
	}
}
