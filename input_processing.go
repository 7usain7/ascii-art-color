package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func input(str string, substr string, color string) error {
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
	print_ascii(all_results, str, substr, color)
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
		for i := range len(substr) {
			indexes = append(indexes, actualIndex+i)
		}
		start = actualIndex + 1 // Allow overlaps by moving just 1 position
	}
	return indexes

}

func containts_index(num int, indexes []int) bool {
	return slices.Contains(indexes, num)
}

func color(color string) string {
	if strings.HasPrefix(color, "--color=") {
		color := color[8:]
		switch strings.ToLower(color) {
		case "red":
			return "\033[31m"
		case "green":
			return "\033[32m"
		case "yellow":
			return "\033[33m"
		case "blue":
			return "\033[34m"
		case "magneta":
			return "\033[35m"
		case "cyan":
			return "\033[36m"
		case "gray":
			return "\033[37m"
		case "orange":
			return "\033[38;5;208m"
		case "white":
			return "\033[97m"
		}
	}
	return ""
}

func print_ascii(all_results [][]string, str string, substr string, color_arg string) {
	empty := true
	reset := "\033[0m"
	color := color(color_arg)
	indexes := substr_indexes(str, substr)
	all_colored := false
	if len(os.Args[1:]) == 2 {
		all_colored = true
	}

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
					if containts_index(k, indexes) || all_colored {
						fmt.Print(color + result[j+k*9] + reset)
					} else {
						fmt.Print(result[j+k*9])
					}
				}
			}
			fmt.Println()
		}
	}
}
