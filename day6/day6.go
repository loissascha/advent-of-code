package day6

import (
	"bufio"
	"fmt"
	"os"

	"github.com/loissascha/go-assert/assert"
)

type PlayerStatus int

const (
	STATUS_UP    PlayerStatus = 1
	STATUS_LEFT  PlayerStatus = 2
	STATUS_RIGHT PlayerStatus = 3
	STATUS_DOWN  PlayerStatus = 4
)

var playerStatus PlayerStatus = STATUS_UP

var laborMap = [][]string{}

func Day6() {
	readFile("day6.test")

	for _, v := range laborMap {
		fmt.Println(v)
	}

	guardPositions := checkMap(laborMap)

	distinctPositions := 0
	for _, v := range guardPositions {
		fmt.Println(v)
		for _, vv := range v {
			if vv > 0 {
				distinctPositions++
			}
		}
	}

	fmt.Println("distinct positions:", distinctPositions)
}

func checkMap(m [][]string) [][]int {
	var guardPositions = [][]int{}

	for _, l := range m {
		iis := []int{}
		for range l {
			iis = append(iis, 0)
		}
		guardPositions = append(guardPositions, iis)
	}

	for li, lm := range m {
		for i, v := range lm {
			if v == "^" {
				fmt.Println("found start position!")
				playerStatus = STATUS_UP
				guardPositions[li][i] = 1
				findNextPosition(li, i, guardPositions)
			}
		}
	}

	return guardPositions
}

func findNextPosition(line int, pos int, guardPositions [][]int) {
	if guardPositions[line][pos] >= 9 {
		fmt.Println("Skipped because guard Position >= 9")
		return
	}
	switch playerStatus {
	case STATUS_UP:
		if line-1 < 0 {
			// out of bounds!
			return
		}
		nextPos := laborMap[line-1][pos]
		if nextPos == "#" {
			// turn right and check if it can go there!
			playerStatus = STATUS_RIGHT
			findNextPosition(line, pos, guardPositions)
			break
		}
		if nextPos == "." || nextPos == "^" {
			// found
			guardPositions[line-1][pos] += 1
			findNextPosition(line-1, pos, guardPositions)
			break
		}
		fmt.Println("no more possible positions up...")
		break
	case STATUS_DOWN:
		if line+1 >= len(laborMap) {
			// out of bounds
			return
		}
		nextPos := laborMap[line+1][pos]
		if nextPos == "#" {
			// turn left and check if it can go there!
			playerStatus = STATUS_LEFT
			findNextPosition(line, pos, guardPositions)
			break
		}
		if nextPos == "." || nextPos == "^" {
			// found
			guardPositions[line+1][pos] += 1
			findNextPosition(line+1, pos, guardPositions)
			break
		}
		fmt.Println("no more possible positions down...")
		break
	case STATUS_RIGHT:
		if pos+1 > len(laborMap[line]) {
			return
		}
		nextPos := laborMap[line][pos+1]
		if nextPos == "#" {
			// turn down and check if it can go there!
			playerStatus = STATUS_DOWN
			findNextPosition(line, pos, guardPositions)
			break
		}
		if nextPos == "." || nextPos == "^" {
			// found
			guardPositions[line][pos+1] += 1
			findNextPosition(line, pos+1, guardPositions)
			break
		}
		fmt.Println("no more possible positions rihgt...")
		break
	case STATUS_LEFT:
		if pos-1 < 0 {
			return
		}
		nextPos := laborMap[line][pos-1]
		if nextPos == "#" {
			// turn up and check if it can go there!
			playerStatus = STATUS_UP
			findNextPosition(line, pos, guardPositions)
			break
		}
		if nextPos == "." || nextPos == "^" {
			// found
			guardPositions[line][pos-1] += 1
			findNextPosition(line, pos-1, guardPositions)
			break
		}
		fmt.Println("no more possible positions left...")
		break
	}
}

func readFile(filepath string) {
	file, err := os.Open(filepath)
	assert.Nil(err, "Can't read file!")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		str := []string{}
		for i := 0; i < len(line); i++ {
			char := line[i : i+1]
			str = append(str, char)
		}
		laborMap = append(laborMap, str)
	}
}
