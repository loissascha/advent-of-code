package day9

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	file, err := os.Open("day9.input")
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
	reordered := reorderConvertedLinev2(converted)
	fmt.Println(reordered)
	visualizeElements(reordered)
	checksum := checkSum(reordered)
	fmt.Println("Checksum:", checksum)
}

func part1(line string) {
	converted := convertLine(line)
	fmt.Println(converted)
	reordered := reorderConvertedLine(converted)
	fmt.Println(reordered)
	visualizeElements(reordered)
	checksum := checkSum(reordered)
	fmt.Println("Checksum:", checksum)
}

func visualizeElements(elements []ElemType) {
	for _, v := range elements {
		_, spaceOk := v.(SpaceType)
		if spaceOk {
			fmt.Print(".")
			continue
		}
		fmt.Print(v.getNum())
	}
	fmt.Print("\n")
}

func reorderConvertedLinev2(elements []ElemType) []ElemType {
	workedElementIds := []int{}
	for i := len(elements) - 1; i >= 0; i-- {
		element := elements[i]
		_, ok := element.(NumType)
		if !ok {
			continue
		}
		if slices.Contains(workedElementIds, element.getNum()) {
			continue
		}
		count := getNumElementCount(elements, element.getNum(), i)
		workedElementIds = append(workedElementIds, element.getNum())

		for j := 0; j < len(elements); j++ {
			if j >= i {
				break
			}
			e := elements[j]
			_, ok := e.(SpaceType)
			if !ok {
				continue
			}
			spaces := getSpaceCount(elements, j)

			if spaces >= count {
				// found element
				for y := 0; y < count; y++ {
					elements[j+y] = NumType{num: element.getNum()}
					elements[i-y] = SpaceType{}
				}
				break
			}
		}
	}
	return elements
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
			continue
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
