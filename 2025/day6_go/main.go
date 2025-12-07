package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")

	var actions []string
	nums := [][]int{}

	lines := strings.Split(input, "\n")
	for n, line := range lines {
		if line == "" {
			continue
		}
		fmt.Println(line)
		fmt.Println(n, len(lines))
		if n < len(lines)-2 {
			slice := getLineSlice(line)
			fmt.Println(slice)
			nums = append(nums, slice)
		} else {
			// last line
			actions = getActionsLine(line)
			fmt.Println(actions)
		}
	}

	result := 0
	for n, a := range actions {
		fmt.Println("action:", a)
		all := []int{}
		for _, num := range nums {
			all = append(all, num[n])
		}
		fmt.Println("with:", all)

		startNum := all[0]
		for i := 1; i < len(all); i++ {
			switch a {
			case "*":
				startNum *= all[i]
			case "+":
				startNum += all[i]
			}
		}
		result += startNum
	}
	fmt.Println("result:", result)
}

func getActionsLine(line string) []string {
	split := strings.SplitSeq(line, " ")
	items := []string{}
	for l := range split {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		items = append(items, l)
	}
	return items
}

func getLineSlice(line string) []int {
	split := strings.SplitSeq(line, " ")
	items := []int{}
	for l := range split {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		items = append(items, n)
	}
	return items
}

func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}
