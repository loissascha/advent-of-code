package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dialPosition = 50
var zeros = 0

func clickLeft() {
	dialPosition--
	if dialPosition == 0 {
		zeros++
	}
	if dialPosition < 0 {
		dialPosition = 99
	}
}

func clickRight() {
	dialPosition++
	if dialPosition > 99 {
		dialPosition = 0
	}
	if dialPosition == 0 {
		zeros++
	}
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

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
			for range turn {
				clickLeft()
			}
		} else if turnStr, ok := strings.CutPrefix(line, "R"); ok {
			turn, err := strconv.Atoi(turnStr)
			if err != nil {
				panic(err)
			}
			for range turn {
				clickRight()
			}
		} else {
			panic("unknown prefix for line:" + line)
		}
	}

	fmt.Println("zeros: ", zeros)
}
