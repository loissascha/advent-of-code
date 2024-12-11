package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/loissascha/go-assert/assert"
)

type Stone struct {
	num int
}

func Day11() {
	input := readFile("day11.input")
	fmt.Println("Input:", input)
	stones := inputToStones(input)
	fmt.Println("stones:", stones)
	for i := 0; i < 75; i++ {
		stones = blink(stones)
	}
	fmt.Println("stones:", stones)
	fmt.Println("stones count:", len(stones))
}

func blink(stones []Stone) []Stone {
	res := []Stone{}
	for _, stone := range stones {
		stoneStr := fmt.Sprintf("%v", stone.num)
		digits := len(stoneStr)
		if stone.num == 0 {
			// rule 1
			fmt.Println("rule 1 for stone:", stone)
			stone.num = 1
			res = append(res, stone)
		} else if digits%2 == 0 {
			// rule 2
			fmt.Println("rule 2 for stone:", stone)
			firstStone := stoneStr[0 : digits/2]
			secondStone := stoneStr[digits/2:]
			fmt.Println("creating 2 stones, first with:", firstStone, "second with:", secondStone)
			fn, err := strconv.Atoi(firstStone)
			assert.Nil(err, "strconv first")
			sn, err := strconv.Atoi(secondStone)
			assert.Nil(err, "strconv second")
			s1 := Stone{num: fn}
			s2 := Stone{num: sn}
			res = append(res, s1)
			res = append(res, s2)
		} else {
			// rule 3
			fmt.Println("rule 3 for stone:", stone)
			stone.num *= 2024
			res = append(res, stone)
		}
	}
	return res
}

func inputToStones(input string) []Stone {
	res := []Stone{}
	split := strings.Split(input, " ")
	for _, v := range split {
		num, err := strconv.Atoi(v)
		assert.Nil(err, "strconv fail")
		res = append(res, Stone{num: num})
	}
	return res
}

func readFile(filepath string) string {
	file, err := os.Open(filepath)
	assert.Nil(err, "Can't read file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		return line
	}
	return ""
}
