package main

import (
	"day3/readfile"
	"fmt"
)

type RowNumber struct {
	Number   string
	StartPow int
}

func getRowNumbers(input string) []RowNumber {
	result := []RowNumber{}

	parsingNumber := ""
	parsingNumberStart := -1

	for n, char := range input {

		// actively parsing a number
		if char >= '0' && char <= '9' {
			if parsingNumberStart == -1 {
				parsingNumberStart = n
			}
			parsingNumber += string(char)
			continue
		}

		// found invalid symbol
		if parsingNumberStart != -1 {
			rn := RowNumber{
				Number:   parsingNumber,
				StartPow: parsingNumberStart,
			}
			result = append(result, rn)
			parsingNumberStart = -1
			parsingNumber = ""
		}
	}

	if parsingNumberStart != -1 {
		rn := RowNumber{
			Number:   parsingNumber,
			StartPow: parsingNumberStart,
		}
		result = append(result, rn)
		parsingNumberStart = -1
		parsingNumber = ""
	}

	return result
}

func main() {
	lines, err := readfile.ReadLines("test_input.txt")
	if err != nil {
		panic(err)
	}

	rows := []string{}

	for line := range lines {
		rows = append(rows, line)
	}

	for _, row := range rows {
		rowNumbers := getRowNumbers(row)
		fmt.Println(row)
		fmt.Println(rowNumbers)
	}

	// fmt.Println(rows)
}
