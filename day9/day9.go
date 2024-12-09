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
		part2(line)
	}
}

func part2(line string) {
	converted := convertLine(line)
	fmt.Println(converted)
	reorderConvertedLinev2(converted)
}

func part1(line string) {
	converted := convertLine(line)
	fmt.Println(converted)
	reordered := reorderConvertedLine(converted)
	fmt.Println(reordered)
	checksum := checkSum(reordered)
	fmt.Println("Checksum:", checksum)
}

func reorderConvertedLinev2(elements []ElemType) {
	startOffset := 0
	for true {
		firstSpaceElementIndex := getFirstSpaceElement(elements, startOffset)
		spaceCounts := getSpaceCount(elements, firstSpaceElementIndex)
		fmt.Println("Space counts:", spaceCounts)
		getLastNumElementStartIndex := len(elements)

		putInNum := 0
		putInCount := 0
		foundFittingElement := false
		lastNumElementIndex := 0

		for true {
			lastNumElementIndex = getLastNumElement(elements, getLastNumElementStartIndex)
			numElem := elements[lastNumElementIndex]
			elemCounts := getNumElementCount(elements, numElem.getNum(), lastNumElementIndex)
			fmt.Println("Elem counts for num", numElem.getNum(), ":", elemCounts)

			// fits
			if elemCounts <= spaceCounts {
				putInNum = numElem.getNum()
				putInCount = elemCounts
				foundFittingElement = true
				fmt.Println("fits")
				break
			}

			// if no -> retry
			getLastNumElementStartIndex = lastNumElementIndex - elemCounts
			if getLastNumElementStartIndex < firstSpaceElementIndex {
				fmt.Println("BREAK")
				break
			}
		}

		if lastNumElementIndex < firstSpaceElementIndex {
			fmt.Println("break1")
			break
		}

		if foundFittingElement {
			// replace spaces with num
			// replace num with spaces
			for j := 0; j < putInCount; j++ {
				ni := j + firstSpaceElementIndex
				si := lastNumElementIndex - j
				fmt.Println("ni", ni, "si", si)
				elements[ni] = NumType{num: putInNum}
				elements[si] = SpaceType{}
				startOffset = 0
			}

			fmt.Println(elements)
		} else {
			fmt.Println("not found a fitting element for this space gap!", spaceCounts)
			startOffset += firstSpaceElementIndex
		}
	}
	fmt.Println(elements)
}

func getNumElementCount(e []ElemType, num int, startIndex int) int {
	count := 0
	for i := len(e) - 1; i >= 0; i-- {
		if i > startIndex {
			continue
		}
		v, ok := e[i].(NumType)
		if !ok {
			break
		}
		if v.getNum() != num {
			break
		}
		count++
	}
	return count
}

func getSpaceCount(e []ElemType, startIndex int) int {
	count := 0
	for i, v := range e {
		if i < startIndex {
			continue
		}
		_, ok := v.(SpaceType)
		if !ok {
			break
		}
		count++
	}
	return count
}

func reorderConvertedLine(elements []ElemType) []ElemType {
	for true {
		firstSpaceElementIndex := getFirstSpaceElement(elements, 0)
		lastNumElementIndex := getLastNumElement(elements, len(elements))
		if firstSpaceElementIndex >= lastNumElementIndex {
			break
		}

		elements = replaceElementAtPos(elements, elements[lastNumElementIndex], firstSpaceElementIndex)
		elements = replaceElementAtPos(elements, SpaceType{}, lastNumElementIndex)
	}

	return elements
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

func replaceElementAtPos(elements []ElemType, e ElemType, pos int) []ElemType {
	elements[pos] = e
	return elements
}

func getLastNumElement(e []ElemType, startIndex int) int {
	for i := len(e) - 1; i >= 0; i-- {
		if i > startIndex {
			continue
		}
		_, ok := e[i].(NumType)
		if ok {
			return i
		}
	}
	return -1
}

func getFirstSpaceElement(e []ElemType, startOffset int) int {
	for i, v := range e {
		if i < startOffset {
			continue
		}
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
