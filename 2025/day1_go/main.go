package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	dialPosition := 50
	zeros := 0

	lines := strings.SplitSeq(string(content), "\n")
	for line := range lines {
		if line == "" {
			continue
		}

		fmt.Println("line:", line)

		if turnStr, ok := strings.CutPrefix(line, "L"); ok {
			turn, err := strconv.Atoi(turnStr)
			if err != nil {
				panic(err)
			}
			dialPosition -= turn
			for dialPosition < 0 {
				dialPosition = 99 - (dialPosition * -1) + 1
			}
			fmt.Println("new dial position:", dialPosition)
		} else if turnStr, ok := strings.CutPrefix(line, "R"); ok {
			turn, err := strconv.Atoi(turnStr)
			if err != nil {
				panic(err)
			}
			dialPosition += turn
			for dialPosition > 99 {
				dialPosition = 0 + (dialPosition - 100)
			}
			fmt.Println("new dial position:", dialPosition)
		} else {
			panic("unknown prefix for line:" + line)
		}

		if dialPosition == 0 {
			zeros++
		}
	}

	fmt.Println("zeros: ", zeros)
}
