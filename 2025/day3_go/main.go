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
		r, err := checkLinePart2(line)
		if err != nil {
			panic(err)
		}
		fmt.Println("largest:", r)
		result += r
	}

	fmt.Println("result:", result)
}

const NUM_LEN = 12

func findLargestNumForIndex(nums []int, index int, startIndex int) (int, int) {
	largestNumber := 0
	largestIndex := 0
	for n := startIndex + 1; n < len(nums); n++ {
		minK := NUM_LEN - (index + 1)
		k := len(nums) - (n + 1)
		if k < minK {
			continue
		}

		if nums[n] > largestNumber {
			largestNumber = nums[n]
			largestIndex = n
		}
	}
	return largestNumber, largestIndex
}

func checkLinePart2(line string) (int, error) {
	fmt.Println("line:", line)

	nums := []int{}
	for n := range len(line) {
		this := string(line[n])
		num, _ := strconv.Atoi(this)
		nums = append(nums, num)
	}

	stack := make([]int, NUM_LEN)
	lastIndex := -1

	for i := range NUM_LEN {
		stack[i], lastIndex = findLargestNumForIndex(nums, i, lastIndex)
	}

	fmt.Println("stack:", stack)

	lc := ""
	for n := range len(stack) {
		lc = fmt.Sprintf("%s%d", lc, stack[n])
	}
	fmt.Println("lc:", lc)

	lcn, _ := strconv.Atoi(lc)

	return lcn, nil
}

func checkLine(line string) (int, error) {
	fmt.Println("line:", line)
	largestCombination := 0

	for n := 0; n < len(line); n++ {
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
