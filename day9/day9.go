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

type NumType struct {
	num int
}

func (n NumType) getNum() int {
	return n.num
}

type SpaceType struct {
}

func (n SpaceType) getNum() int {
	return 1
}

type ElemType interface {
	getNum() int
}

func Day9() {
	file, err := os.Open("day9.test")
	assert.Nil(err, "Can't open file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		part1(line)
	}
}

func part1(line string) {
	converted := convertLine(line)
	fmt.Println(converted)
	reordered := reorderConvertedLine(converted)
	fmt.Println(reordered)
	checksum := checkSum(reordered)
	fmt.Println("Checksum:", checksum)
}

func checkSum(elements []ElemType) int {
	sum := 0

	for i, v := range elements {
		_, st := v.(SpaceType)
		if st {
			break
		}
		nt := v.(NumType)
		sum += (nt.getNum() * i)
	}

	return sum
}

func reorderConvertedLine(elements []ElemType) []ElemType {
	for true {
		firstSpaceElementIndex := getFirstSpaceElement(elements)
		lastNumElementIndex := getLastNumElement(elements)
		if firstSpaceElementIndex >= lastNumElementIndex {
			break
		}

		elements = replaceElementAtPos(elements, elements[lastNumElementIndex], firstSpaceElementIndex)
		elements = replaceElementAtPos(elements, SpaceType{}, lastNumElementIndex)
	}

	return elements
}

func replaceElementAtPos(elements []ElemType, e ElemType, pos int) []ElemType {
	elements[pos] = e
	return elements
}

func getLastNumElement(e []ElemType) int {
	for i := len(e) - 1; i >= 0; i-- {
		_, ok := e[i].(NumType)
		if ok {
			return i
		}
	}
	return -1
}

func getFirstSpaceElement(e []ElemType) int {
	for i, v := range e {
		_, ok := v.(SpaceType)
		if ok {
			return i
		}
	}
	return -1
}

func convertLine(line string) []ElemType {
	rm := []ElemType{}
	id := 0
	readingType := BLOCKS
	for i := 0; i < len(line); i++ {
		char := line[i : i+1]

		switch readingType {
		case BLOCKS:
			num, err := strconv.Atoi(char)
			assert.Nil(err, "Can't convert num")
			for j := 0; j < num; j++ {
				rm = append(rm, NumType{num: id})
			}
			id++
			readingType = SPACE
			break
		case SPACE:
			num, err := strconv.Atoi(char)
			assert.Nil(err, "Can't convert num")
			for j := 0; j < num; j++ {
				rm = append(rm, SpaceType{})
			}
			readingType = BLOCKS
			break
		default:
			fmt.Println("UNKNOWN readingType")
			break
		}
	}
	return rm
}
