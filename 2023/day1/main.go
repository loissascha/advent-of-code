package main

import (
	"day1/readfile"
	"fmt"
	"strconv"
	"strings"
)

var letterMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	inputFile := "input.txt"
	lines, err := readfile.ReadLines(inputFile)
	if err != nil {
		panic(err)
	}
	sum := 0
	for line := range lines {
		res := getLineValue([]byte(line))
		fmt.Println(res)
		sum += res
	}
	fmt.Println("Sum:", sum)
}

func findTextNumber(line string) string {
	found := false
	foundIdx := len(line) + 1
	foundValue := ""
	for k, v := range letterMap {
		idx := strings.Index(line, k)
		if idx > -1 && idx < foundIdx {
			foundIdx = idx
			foundValue = v
			found = true
		}
	}
	if found {
		return foundValue
	}
	return ""
}

func getLineValue(line []byte) int {
	resStr := ""

	firstNum := ""
	secondNum := ""

	// find first number
	for i, ch := range line {
		lineStr := string(line)
		shortStr := lineStr[:i]
		strNum := findTextNumber(shortStr)
		if strNum != "" {
			firstNum = strNum
			break
		}
		if ch >= byte('0') && ch <= byte('9') {
			firstNum = string(ch)
			break
		}
	}

	// find last number
	for i, _ := range line {
		idx := len(line) - i - 1
		ch := line[idx]
		lineStr := string(line)
		shortStr := lineStr[idx:]
		strNum := findTextNumber(shortStr)
		if strNum != "" {
			secondNum = strNum
			break
		}
		if ch >= byte('0') && ch <= byte('9') {
			secondNum = string(ch)
			break
		}
	}

	resStr = firstNum + secondNum

	// old way
	// for _, ch := range line {
	// 	if ch >= byte('0') && ch <= byte('9') {
	// 		resStr += string(ch)
	// 	}
	// }

	if len(resStr) == 0 {
		return 0
	}

	if len(resStr) == 1 {
		resStr = resStr + resStr
	}

	if len(resStr) > 2 {
		resStr = resStr[:1] + resStr[len(resStr)-1:]
	}

	res, err := strconv.Atoi(resStr)
	if err != nil {
		panic(err)
	}
	return res
}
