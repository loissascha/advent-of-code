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

func Day2() {
	file, err := os.Open("input2.txt")
	assert.Nil(err, "Can't read file!")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		split := strings.Split(line, " ")
		scanLine(split)
	}

	fmt.Println("Save reports:", safeReports)
}

var safeReports = 0
var unsafeReports = 0

func scanLine(line []string) {
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
			unsafeReports++
			fmt.Println("unsafe because lastNum == num")
			return
		}

		diff := lastNum - num
		if diff < 0 {
			diff *= -1
		}

		// diff > 3 = fail
		if diff > 3 {
			unsafeReports++
			fmt.Println("unsafe because diff > 3")
			return
		}

		if lineType == Unclear {
			if lastNum < num {
				lineType = Increase
			} else {
				lineType = Decrease
			}
		} else if lineType == Increase {
			if lastNum > num {
				unsafeReports++
				fmt.Println("unsafe because lineType increase and no increase. last num:", lastNum, "num:", num)
				return
			}
		} else if lineType == Decrease {
			if lastNum < num {
				unsafeReports++
				fmt.Println("unsafe because lineType decrease and no decrease. last num:", lastNum, "num:", num)
				return
			}
		}

		lastNum = num
	}

	safeReports++
}
