package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/loissascha/go-assert/assert"
)

type LineType int

const (
	Unclear  LineType = 0
	Increase LineType = 1
	Decrease LineType = 2
)

var safeReports = 0
var unsafeReports = 0

func Day2() {
	file, err := os.Open("input2.txt")
	assert.Nil(err, "Can't read file!")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		split := strings.Split(line, " ")
		safe := scanLine(split)
		if safe {
			safeReports++
			continue
		}

		// check if it's safe if one number is removed from the list
		hasSafe := false
		for i := 0; i < len(split); i++ {
			var newSplit = make([]string, len(split))
			copy(newSplit, split)
			newSplit = removeIndex(newSplit, i)
			safe := scanLine(newSplit)
			if safe {
				hasSafe = true
				break
			}
		}
		if hasSafe {
			safeReports++
			continue
		}
		unsafeReports++
	}

	fmt.Println("Safe reports:", safeReports)
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func scanLine(line []string) bool {
	fmt.Println(line)

	lineType := Unclear
	firstRun := true
	lastNum := 0

	for _, v := range line {
		num, err := strconv.Atoi(v)
		assert.Nil(err, "Strconv Atoi failed for "+v)

		if firstRun {
			lastNum = num
			firstRun = false
			continue
		}

		// no diff = fail
		if lastNum == num {
			fmt.Println("unsafe because lastNum == num")
			return false
		}

		diff := lastNum - num
		if diff < 0 {
			diff *= -1
		}

		// diff > 3 = fail
		if diff > 3 {
			fmt.Println("unsafe because diff > 3")
			return false
		}

		if lineType == Unclear { // first time -> define which lineType
			if lastNum < num {
				lineType = Increase
			} else {
				lineType = Decrease
			}
		} else if lineType == Increase { // if number doesn't increase -> failed
			if lastNum > num {
				fmt.Println("unsafe because lineType increase and no increase. last num:", lastNum, "num:", num)
				return false
			}
		} else if lineType == Decrease { // if number doesn't decrease -> failed
			if lastNum < num {
				fmt.Println("unsafe because lineType decrease and no decrease. last num:", lastNum, "num:", num)
				return false
			}
		}

		lastNum = num
	}

	return true
}
