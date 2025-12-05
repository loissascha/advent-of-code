package main

import (
	"fmt"
	"iter"
	"os"
	"slices"
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

	part2(lines)
}

func part2(lines iter.Seq[string]) {
	freshIdRange := []Range{}
	for line := range lines {
		if strings.Contains(line, "-") {
			fmt.Println("line:", line)
			sp := strings.SplitN(line, "-", 2)

			freshIdRange = append(freshIdRange, Range{
				Start: strToInt(sp[0]),
				End:   strToInt(sp[1]),
			})

			continue
		}
	}

	slices.SortFunc(freshIdRange, func(a, b Range) int {
		if a.Start < b.Start {
			return -1
		}
		if a.Start > b.Start {
			return 1
		}
		return 0
	})

	found := true
	for found {
		found = false
		for n, item := range freshIdRange {
			if n+1 == len(freshIdRange) {
				break
			}
			next := freshIdRange[n+1]

			if isInRange(item.Start, item.End, next.Start) || isInRange(item.Start, item.End, next.End) {
				newStart := min(item.Start, next.Start)
				newEnd := max(next.End, item.End)
				freshIdRange[n].Start = newStart
				freshIdRange[n].End = newEnd
				freshIdRange = slices.Delete(freshIdRange, n+1, n+2)
				found = true
				break
			}
		}
	}

	sum := 0
	for _, item := range freshIdRange {
		sum += item.End - (item.Start - 1)
	}

	fmt.Println("sum:", sum)

}

func isInRange(start, end, pos int) bool {
	if start <= pos && end >= pos {
		return true
	}
	return false
}

func part1(lines iter.Seq[string]) {
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
