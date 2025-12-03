package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.SplitSeq(content, "\n")

	result := 0
	for line := range lines {
		if line == "" {
			continue
		}
		r, err := checkLine(line)
		if err != nil {
			panic(err)
		}
		fmt.Println("largest:", r)
		result += r
	}

	fmt.Println("result:", result)
}

func checkLine(line string) (int, error) {
	fmt.Println("line:", line)
	largestCombination := 0

	for n := 0; n < len(line); n++ {
		if n+1 == len(line) {
			break
		}
		this := line[n]

		for i := n + 1; i < len(line); i++ {
			next := line[i]
			comb := fmt.Sprintf("%s%s", string(this), string(next))
			c, err := strconv.Atoi(comb)
			if err != nil {
				return 0, err
			}
			if c > largestCombination {
				largestCombination = c
			}
		}
	}

	return largestCombination, nil
}

func readFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
