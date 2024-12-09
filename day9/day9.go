package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/loissascha/go-assert/assert"
)

const (
	BLOCKS = iota
	SPACE
)

func Day9() {
	file, err := os.Open("day9.input")
	assert.Nil(err, "Can't open file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		converted := convertLine(line)
		fmt.Println(converted)
		reordered := reorderConvertedLine(converted)
		checksum := checkSum(reordered)
		fmt.Println("Checksum:", checksum)
	}
}

func checkSum(line string) int {
	id := 0
	checksum := 0
	for i := 0; i < len(line); i++ {
		char := line[i : i+1]
		if char == "." {
			break
		}
		num, err := strconv.Atoi(char)
		assert.Nil(err, "Can't convert num checksum")
		checksum += (num * id)
		id++
	}
	return checksum
}

func reorderConvertedLine(line string) string {
	result := line

	for true {
		firstDot := getFirstDot(result)
		lastNonDot := getLastNonDot(result)
		if firstDot >= lastNonDot {
			break
		}
		lastChar := getCharPos(result, lastNonDot)
		result = replaceLineCharAtPos(result, lastChar, firstDot)
		result = replaceLineCharAtPos(result, ".", lastNonDot)
		fmt.Println(result)
	}

	return result
}

func replaceLineCharAtPos(line string, char string, pos int) string {
	result := line[0:pos]
	result += char
	result += line[pos+1:]
	return result
}

func getCharPos(str string, i int) string {
	return str[i : i+1]
}

func getFirstDot(str string) int {
	for i := 0; i < len(str); i++ {
		char := str[i : i+1]
		if char == "." {
			return i
		}
	}
	return -1
}

func getLastNonDot(str string) int {
	for i := len(str) - 1; i >= 0; i-- {
		char := str[i : i+1]
		if char != "." {
			return i
		}
	}
	return -1
}

func convertLine(line string) string {
	id := 0
	readingType := BLOCKS
	result := ""
	for i := 0; i < len(line); i++ {
		char := line[i : i+1]

		switch readingType {
		case BLOCKS:
			num, err := strconv.Atoi(char)
			assert.Nil(err, "Can't convert num")
			for j := 0; j < num; j++ {
				result = fmt.Sprintf("%v%v", result, id)
			}
			id++
			readingType = SPACE
			break
		case SPACE:
			num, err := strconv.Atoi(char)
			assert.Nil(err, "Can't convert num")
			for j := 0; j < num; j++ {
				result = fmt.Sprintf("%v.", result)
			}
			readingType = BLOCKS
			break
		default:
			fmt.Println("UNKNOWN readingType")
			break
		}
	}
	return result
}
