package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	source_text = flag.String("s", "", "this is the source text for analysis")
	source_file = flag.String("ss", "", "source file to check word count")
)

// analyse frewuency of each word
func AnalyseString(input string) (map[string]int, error) {
	//1. Loop over
	//Assuming that all words seperated by whitespace:
	words := strings.Split(input, " ")
	out := make(map[string]int)
	for _, word := range words {
		//word exits, increment count
		//doesn't exit, set 1
		// log.Println(word)
		out[word] += 1
		// log.Println(out[word])
	}
	return out, nil
}

func AnalyseFile(path string) (map[string]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	//file f is open and ready to analyse

	text, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return AnalyseString(string(text))
}

func main() {
	flag.Parse()
	
	if *source_file != "" {
		result, err := AnalyseFile(*source_file)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(result)
	} else if *source_text != "" {
		result, err := AnalyseString(*source_text)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(result)
	} else {
		fmt.Println("Please provide either -s (text) or -ss (file path)")
	}
}
