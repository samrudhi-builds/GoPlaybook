package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Printf("Enter file name : ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	var names []Name

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		nameParts := strings.Split(line, " ")
		if len(nameParts) == 2 {
			firstName := strings.TrimSpace(nameParts[0])
			lastName := strings.TrimSpace(nameParts[1])

			name := Name{fname: firstName, lname: lastName}
			names = append(names, name)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	fmt.Println("\nNames found in file: ")
	for _, name := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", name.fname, name.lname)
	}
}
