package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input(str string) error {
	file, err := os.Open("models/standard.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	var all_results [][]string
	lines := strings.Split(str, "\\n")
	for _, line := range lines {
		var result []string
		i := 0
		for _, rune := range line {
			file.Seek(0, 0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if int(rune)-32 == i/9 {
					line := scanner.Text()
					result = append(result, line)
				}
				i++
			}
			scanner_err := scanner.Err()
			if scanner_err != nil {
				return scanner_err
			}
			i = 0
		}
		if len(result) > 0 {
			all_results = append(all_results, result)
		} else {
			all_results = append(all_results, []string{})
		}
	}
	print_ascii(all_results, str)
	return nil
}

func substr_indexes(str string, substr string) []int {
	var indexes []int
	start := 0
	subLen := len(substr)

	for start <= len(str)-subLen {
		index := strings.Index(str[start:], substr)
		if index == -1 {
			break
		}
		actualIndex := start + index
		indexes = append(indexes, actualIndex)
		start = actualIndex + 1 // Allow overlaps by moving just 1 position
	}
	return indexes

}

func print_ascii(all_results [][]string, str string) {
	empty := true
	var reset = "\033[0m"
	var red = "\033[031m"

	// indexes := substr_indexes(str, "kit")

	for i, result := range all_results {
		if len(all_results)-1 == i && i != 0 && empty && len(result) == 0 {
			break
		}

		if len(result) == 0 {
			fmt.Println()
			continue
		} else {
			empty = false
		}

		for j := 1; j < 9; j++ {
			for k := range len(result) {
				if j+k*9 < len(result) {
					if string(str[k]) == "t" {
						fmt.Print(red + result[j+k*9] + reset)
					} else {
						fmt.Print(result[j+k*9])
					}
				}
			}
			fmt.Println()
		}
	}
}
