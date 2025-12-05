package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func main() {
	input := readFile("input.txt")
	lines := strings.SplitSeq(input, "\n")

	// freshIdMap := map[int]bool{}
	// freshIdSlice := []int{}
	freshIdRange := []Range{}
	availableIngredients := []int{}

	for line := range lines {
		if line == "" {
			continue
		}

		// fresh id ranges
		if strings.Contains(line, "-") {
			sp := strings.SplitN(line, "-", 2)
			freshIdRange = append(freshIdRange, Range{
				Start: strToInt(sp[0]),
				End:   strToInt(sp[1]),
			})

			// for n := strToInt(sp[0]); n <= strToInt(sp[1]); n++ {
			// freshIdMap[n] = true
			// freshIdSlice = append(freshIdSlice, n)
			// }
			continue
		}

		// available ingredients
		availableIngredients = append(availableIngredients, strToInt(line))
	}

	// fmt.Println(freshIdMap)
	fmt.Println(availableIngredients)

	freshs := 0
	fmt.Println("starting getting freshs")
	for _, ing := range availableIngredients {
		// _, ok := freshIdMap[ing]
		// if ok {
		// 	freshs++
		// }
		found := false
		for _, r := range freshIdRange {
			if r.Start <= ing && r.End >= ing {
				found = true
				break
			}
		}
		if found {
			freshs++
		}
	}

	fmt.Println("fresh:", freshs)
}

func strToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}
