package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter string starting with 'i', containing 'a' & ending with 'n' : ")
	fmt.Scan(&input)

	input = strings.ToLower(input)

	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
