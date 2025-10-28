package main

import "fmt"

func main() {
	input := "3113322113"
	// input := "1"

	for range 50 {
		fmt.Println("Input:", len(input))
		input = runCycle(input)
		fmt.Println("Output:", len(input))
	}

	fmt.Println("result len:", len(input))
}

func runCycle(input string) string {
	output := ""

	var currentRune rune
	currentRuneCount := 0
	for _, c := range input {
		if currentRune == c {
			currentRuneCount++
		} else {
			if currentRuneCount > 0 {
				output += fmt.Sprint(currentRuneCount) + string(currentRune)
			}
			currentRuneCount = 1
			currentRune = c
		}
	}
	if currentRuneCount > 0 {
		output += fmt.Sprint(currentRuneCount) + string(currentRune)
	}
	return output
}
