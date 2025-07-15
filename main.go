package main

import (
	"fmt"
	"os"
)

func isAscii(str string) bool {
	// runes := []rune(str)
	for _, e := range str {
		if e > 127 {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Invalid number of arguements.")
		return
	}
	sentence := args[0]
	if !isAscii(sentence) {
		fmt.Println("Only Ascii is accepted.")
		return
	}
	if sentence != "" {
		err := input(sentence)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
