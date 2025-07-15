package main

import (
	"fmt"
	"os"
	"strings"
)

func isAscii(str string) bool {
	for _, e := range str {
		if e > 127 {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 3 || len(args) == 2 {
		fmt.Println("Invalid number of arguements.")
		return
	}
	var sentence string
	var substr string
	if strings.HasPrefix(args[0], "--color=") {
		sentence = args[2]
		substr = args[1]
	} else {
		sentence = args[0]
	}
	if !isAscii(sentence) {
		fmt.Println("Only Ascii is accepted.")
		return
	}
	if sentence != "" {
		err := input(sentence, substr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
