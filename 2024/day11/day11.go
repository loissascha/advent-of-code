package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/loissascha/go-assert/assert"
)

func Day11() {
	input := readFile("day11.input")
	fmt.Println("Input:", input)
	stones := inputToStones(input)
	fmt.Println("stones:", stones)
	for i := 0; i < 75; i++ {
		stones = blink(stones)
	}

	sum := 0
	for _, count := range stones {
		sum += count
	}
	fmt.Println("stones:", sum)
}

func blink(stones map[int]int) map[int]int {
	newStones := make(map[int]int)
	for stone, count := range stones {
		fmt.Println("stone", stone, "with count", count)
		stoneStr := fmt.Sprintf("%v", stone)
		digits := len(stoneStr)
		if stone == 0 {
			// rule 1
			newStones[1] += count
		} else if digits%2 == 0 {
			// rule 2
			firstStone := stoneStr[0 : digits/2]
			secondStone := stoneStr[digits/2:]
			fn, err := strconv.Atoi(firstStone)
			assert.Nil(err, "strconv first")
			sn, err := strconv.Atoi(secondStone)
			assert.Nil(err, "strconv second")
			s1 := fn
			s2 := sn
			newStones[s1] += count
			newStones[s2] += count
		} else {
			n := stone * 2024
			newStones[n] += count
		}
	}
	return newStones
}

func inputToStones(input string) map[int]int {
	res := make(map[int]int)
	split := strings.Split(input, " ")

	for _, v := range split {
		num, err := strconv.Atoi(v)
		assert.Nil(err, "strconv fail")
		res[num] += 1
		// res = append(res, num)
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
