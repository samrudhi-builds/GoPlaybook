package main

import (
	"fmt"
)

func main() {
	var float float64
	fmt.Println("Enter a floating point number: ")
	fmt.Scan(&float)
	truncatedInt := int(float)
	fmt.Printf("Truncated Integer : %d\n", truncatedInt)
}
