package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "3113322113"
	// input := "1"

	for n := range 50 {
		fmt.Println("Input (", n, "):", len(input))
		input = runCycle(input)
		fmt.Println("Output (", n, "):", len(input))
	}

	fmt.Println("result len:", len(input))
}

func runCycle(input string) string {
	var output strings.Builder

	var currentRune rune
	currentRuneCount := 0
	for _, c := range input {
		if currentRune == c {
			currentRuneCount++
		} else {
			if currentRuneCount > 0 {
				output.WriteString(strconv.Itoa(currentRuneCount))
				output.WriteRune(currentRune)
			}
			currentRuneCount = 1
			currentRune = c
		}
	}
	if currentRuneCount > 0 {
		output.WriteString(strconv.Itoa(currentRuneCount))
		output.WriteRune(currentRune)
	}
	return output.String()
}
