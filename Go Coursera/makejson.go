package main 

import (
	"encoding/json"
	"fmt"
)

func main () {
	person := map[string]string{
		"name": "joe", 
		"addr": "a st. NYC",
	}

	barr, err := json.Marshal(person)
	if err != nil{
		fmt.Println("Error marshaling JSON: ", err)
		return
	}

	fmt.Println(string(barr))
}