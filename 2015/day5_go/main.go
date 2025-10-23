package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputByte, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputStr := string(inputByte)
	splits := strings.SplitSeq(inputStr, "\n")
	howManyNices := 0
	for line := range splits {
		fmt.Println(line)
		isNice := isLineNice(line)
		if isNice {
			howManyNices++
		}
	}
	fmt.Println("how many nices:", howManyNices)

	// isNice := isLineNice("jchzalrnumimnmhp")
	// fmt.Println("isNice: ", isNice)
}

var vowels []string = []string{"a", "e", "i", "o", "u"}
var forbiddenStrings []string = []string{"ab", "cd", "pq", "xy"}

func isLineNice(input string) bool {

	// has at least 3 vowels
	vowelCount := 0
	for _, vowel := range vowels {
		vowelCount += strings.Count(input, vowel)
	}
	if vowelCount < 3 {
		return false
	}

	// does not contain the not allowed string
	for _, forbiddenString := range forbiddenStrings {
		if strings.Contains(input, forbiddenString) {
			return false
		}
	}

	// at least one letter that appears twice in a row
	for n, char := range input {
		if n+2 > len(input) {
			continue
		}
		nextChar := input[n+1 : n+2]
		if string(char) == nextChar {
			fmt.Println("has double letter", string(char), n, nextChar, n+1)
			return true
		}
	}
	return false
}
