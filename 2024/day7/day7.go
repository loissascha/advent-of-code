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
	inputs := readFile("day7.input")

	sum := 0
	for _, input := range inputs {
		sum += input.trySolve()
	}
	fmt.Println("sum:", sum)
}

func (input *LineInput) trySolve() int {
	spaces := len(input.numbers) - 1

	operators := generateCombinations(spaces)

	for i := 0; i < len(operators); i++ {
		op := operators[i]

		sum := 0
		for j, v := range input.numbers {
			if j == 0 {
				sum += v
				continue
			}
			pop := op[j-1 : j]
			switch pop {
			case "+":
				sum += v
				break
			case "*":
				sum *= v
				break
			case "|":
				sumstr := fmt.Sprintf("%v%v", sum, v)
				newsum, err := strconv.Atoi(sumstr)
				assert.Nil(err, "newsum not possible")
				sum = newsum
				break
			default:
				panic("INVALID POP")
			}
		}
		if sum == input.result {
			return sum
		}
	}

	return 0
}

func generateCombinations(n int) []string {
	if n == 0 {
		return []string{""}
	}
	smallerCombinations := generateCombinations(n - 1)
	combinations := []string{}
	for _, comb := range smallerCombinations {
		combinations = append(combinations, comb+"+")
		combinations = append(combinations, comb+"*")
		combinations = append(combinations, comb+"|")
	}
	return combinations
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
