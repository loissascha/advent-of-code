package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/loissascha/go-assert/assert"
)

type LineInput struct {
	result  int
	numbers []int
}

func Day7() {
	inputs := readFile("day7.test")
	fmt.Println(inputs)
}

func readFile(filepath string) []LineInput {
	result := []LineInput{}
	file, err := os.Open(filepath)
	assert.Nil(err, "Can't read file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		before, after, found := strings.Cut(line, ": ")
		if !found {
			continue
		}
		r, err := strconv.Atoi(before)
		assert.Nil(err, "Can't conv before")

		numbers := []int{}
		as := strings.Split(after, " ")
		for _, a := range as {
			n, err := strconv.Atoi(a)
			assert.Nil(err, "Can't conv after a "+a)
			numbers = append(numbers, n)
		}

		li := LineInput{
			result:  r,
			numbers: numbers,
		}
		result = append(result, li)
	}
	return result
}
