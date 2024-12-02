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

func Day2() {
	file, err := os.Open("input2.txt")
	assert.Nil(err, "Can't read file!")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")
		safe := lineSafe(split)
		if safe {
			safeReports++
			continue
		}

		// Problem Dampener
		hasSafe := hasSafeVariation(split)
		if hasSafe {
			safeReports++
			continue
		}
	}

	fmt.Println("Safe reports:", safeReports)
}

// check if it has a variant where one element is removed that is safe
func hasSafeVariation(split []string) bool {
	hasSafe := false
	for i := 0; i < len(split); i++ {
		var newSplit = make([]string, len(split))
		copy(newSplit, split)
		newSplit = removeIndex(newSplit, i)
		safe := lineSafe(newSplit)
		if safe {
			hasSafe = true
			break
		}
	}
	return hasSafe
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func getDiff(num1 int, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		diff *= -1
	}
	return diff
}

func lineSafe(line []string) bool {
	fmt.Println(line)

	lineType := Unclear
	firstRun := true
	lastNum := 0

	for _, v := range line {
		num, err := strconv.Atoi(v)
		assert.Nil(err, "Strconv Atoi failed for "+v)

		// first run, just set lastNum and disable first run
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

		diff := getDiff(lastNum, num)

		// diff > 3 = fail
		if diff > 3 {
			fmt.Println("unsafe because diff > 3")
			return false
		}

		switch lineType {
		case Unclear: // first time -> define which lineType
			if lastNum < num {
				lineType = Increase
			} else {
				lineType = Decrease
			}
			break
		case Increase: // if number doesn't increase -> failed
			if lastNum > num {
				fmt.Println("unsafe because lineType increase and no increase. last num:", lastNum, "num:", num)
				return false
			}
			break
		case Decrease: // if number doesn't decrease -> failed
			if lastNum < num {
				fmt.Println("unsafe because lineType decrease and no decrease. last num:", lastNum, "num:", num)
				return false
			}
			break
		}
		lastNum = num
	}

	return true
}
