package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/loissascha/go-assert/assert"
)

var stoneCount = 0

func Day11() {
	input := readFile("day11.input")
	fmt.Println("Input:", input)
	stones := inputToStones(input)
	fmt.Println("stones:", stones)
	// newStones := []int{}
	for si, stone := range stones {
		// stoneRes := []int{stone}
		blink(stone, 75)
		// for i := 0; i < 75; i++ {
		// 	stoneRes = blink(stoneRes)
		// 	fmt.Println("stone", si, "run", i)
		// }
		fmt.Println("finished with stone", si)
		// newStones = append(newStones, stoneRes...)
	}
	// fmt.Println("stones:", newStones)
	fmt.Println("stones count:", stoneCount)
}

func blink(stone int, calculationsLeft int) {
	// res := []int{}
	// for _, stone := range stones {
	if calculationsLeft <= 0 {
		stoneCount++
		fmt.Println("reached end! stoneCount increased")
		return
	}
	// fmt.Println("blink for stone", stone, "with calculationsLeft:", calculationsLeft, "stone count:", stoneCount)
	calculationsLeft--
	stoneStr := fmt.Sprintf("%v", stone)
	digits := len(stoneStr)
	if stone == 0 {
		// rule 1
		// fmt.Println("rule 1 for stone:", stone)
		stone = 1
		// stoneCount++
		blink(1, calculationsLeft)
		// res = append(res, stone)
	} else if digits%2 == 0 {
		// rule 2
		// fmt.Println("rule 2 for stone:", stone)
		firstStone := stoneStr[0 : digits/2]
		secondStone := stoneStr[digits/2:]
		// fmt.Println("creating 2 stones, first with:", firstStone, "second with:", secondStone)
		fn, err := strconv.Atoi(firstStone)
		assert.Nil(err, "strconv first")
		sn, err := strconv.Atoi(secondStone)
		assert.Nil(err, "strconv second")
		s1 := fn
		s2 := sn
		// stoneCount += 2
		blink(s1, calculationsLeft)
		blink(s2, calculationsLeft)
		// res = append(res, s1)
		// res = append(res, s2)
	} else {
		// rule 3
		// fmt.Println("rule 3 for stone:", stone)
		stone *= 2024
		// stoneCount++
		blink(stone, calculationsLeft)
		// res = append(res, stone)
	}
	// fmt.Println("after stonecount:", stoneCount)
	// }
	// return res
}

func inputToStones(input string) []int {
	res := []int{}
	split := strings.Split(input, " ")
	for _, v := range split {
		num, err := strconv.Atoi(v)
		assert.Nil(err, "strconv fail")
		res = append(res, num)
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
