package main

import (
	"day3/readfile"
	"fmt"
)

type Row struct {
	Index      int
	Row        string
	RowNumbers []RowNumber
}

type RowNumber struct {
	Number   string
	StartPos int
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
				StartPos: parsingNumberStart,
			}
			result = append(result, rn)
			parsingNumberStart = -1
			parsingNumber = ""
		}
	}

	if parsingNumberStart != -1 {
		rn := RowNumber{
			Number:   parsingNumber,
			StartPos: parsingNumberStart,
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

	rows := []Row{}
	l := []string{}

	for line := range lines {
		l = append(l, line)
	}

	for n, row := range l {
		rows = append(rows, Row{
			Index:      n,
			Row:        row,
			RowNumbers: getRowNumbers(row),
		})
	}

	for _, row := range rows {
		fmt.Println(row.Row)
	}

}
