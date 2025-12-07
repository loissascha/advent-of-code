package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")

	lines := strings.Split(input, "\n")
	fmt.Println("first line:", lines[0])

	lineSlices := [][][]string{}
	maxLen := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		lineSlice := parseLine(line)
		maxLen = len(lineSlice)
		lineSlices = append(lineSlices, lineSlice)
	}

	fmt.Println("max len:", maxLen)
	fmt.Println("line slices:", lineSlices)

	result := 0
	for pos := range maxLen {
		fmt.Println("pos:", pos)

		firsts := ""
		seconds := ""
		thirds := ""
		symbol := ""

		for _, lineSlice := range lineSlices {
			ls := lineSlice[pos]
			fmt.Println(ls)

			first := ls[0]
			second := ls[1]
			third := ls[2]

			if first == "*" {
				symbol = "*"
				continue
			}
			if first == "+" {
				symbol = "+"
				continue
			}

			firsts = fmt.Sprintf("%s%s", firsts, first)
			seconds = fmt.Sprintf("%s%s", seconds, second)
			thirds = fmt.Sprintf("%s%s", thirds, third)
		}

		fmt.Println(firsts, seconds, thirds, symbol)
		firstN, _ := strconv.Atoi(strings.TrimSpace(firsts))
		secondN, _ := strconv.Atoi(strings.TrimSpace(seconds))
		thirdN, _ := strconv.Atoi(strings.TrimSpace(thirds))
		fmt.Println(firstN, secondN, thirdN, symbol)

		switch symbol {
		case "*":
			result += (firstN * secondN * thirdN)
		case "+":
			result += (firstN + secondN + thirdN)
		default:
			panic("Symbold not found!: " + symbol)
		}
	}

	fmt.Println("result:", result)
}

func parseLine(line string) [][]string {
	fmt.Println("parse line:", line)

	lineSlice := [][]string{}

	parsing := 0
	chars := make([]string, 0, 3)
	for _, char := range line {
		if parsing < 3 {
			chars = append(chars, string(char))
		}

		parsing++

		switch parsing {
		case 3: // parsed 3 chars -> put them in the slice
			lineSlice = append(lineSlice, chars)
			chars = make([]string, 0, 3)
		case 4:
			// one free space before the next parsing begins
			parsing = 0
		}
	}

	fmt.Println("parsed line:", lineSlice)
	return lineSlice
}

func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}
