package main

import (
	"day3/readfile"
	"fmt"
	"strconv"
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

func getItemOnPos(rows []Row, x int, y int) string {
	if y < 0 {
		return "."
	}
	if y >= len(rows) {
		return "."
	}
	row := rows[y]
	if x < 0 {
		return "."
	}
	if x >= len(row.Row) {
		return "."
	}
	v := row.Row[x]
	return string(v)
}

func getRowsWithAdjacents(rows []Row) []string {
	numbersWithAdjacents := []string{}
	for y, r := range rows {
		for _, rn := range r.RowNumbers {
			hasAdjacent := false
			for i := 0; i < len(rn.Number); i++ {
				first := false
				if i == 0 {
					first = true
				}

				last := false
				if i == len(rn.Number)-1 {
					last = true
				}

				x := rn.StartPos + i

				working := rows[y].Row[x]

				fmt.Println("working on:", string(working))

				if first {
					left := getItemOnPos(rows, x-1, y)
					fmt.Println("left:", left)
					if left != "." {
						hasAdjacent = true
					}
				}
				topLeft := getItemOnPos(rows, x-1, y-1)
				topRight := getItemOnPos(rows, x+1, y-1)
				top := getItemOnPos(rows, x, y-1)
				bottom := getItemOnPos(rows, x, y-1)
				bottomLeft := getItemOnPos(rows, x-1, y+1)
				bottomRight := getItemOnPos(rows, x+1, y+1)
				fmt.Println("topLeft", topLeft)
				if topLeft != "." || topRight != "." || top != "." || bottom != "." || bottomLeft != "." || bottomRight != "." {
					hasAdjacent = true
				}

				if last {
					right := getItemOnPos(rows, x+1, y)
					fmt.Println("right:", right)
					if right != "." {
						hasAdjacent = true
					}
				}
			}
			if hasAdjacent {
				numbersWithAdjacents = append(numbersWithAdjacents, rn.Number)
			}
		}
	}
	return numbersWithAdjacents
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
	lines, err := readfile.ReadLines("input.txt")
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

	res := getRowsWithAdjacents(rows)
	fmt.Println(res)

	sum := getSum(res)
	fmt.Println("sum:", sum)

}

func getSum(input []string) int {
	sum := 0
	for _, v := range input {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		sum += i
	}
	return sum
}
