package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/loissascha/go-assert/assert"
)

var firstSplit = []int{}
var secondSplit = []int{}

var diff = 0
var simScore = 0

func Day1() {
	readFile("input1.txt")
	fmt.Println(len(firstSplit), len(secondSplit))
	l := len(firstSplit)
	ogFirstSplit := make([]int, len(firstSplit))
	ogSecondSplit := make([]int, len(secondSplit))
	copy(ogFirstSplit, firstSplit)
	copy(ogSecondSplit, secondSplit)

	for i := 0; i < l; i++ {
		compareLowest()
	}
	fmt.Println("diff is", diff)

	similarityScore(ogFirstSplit, ogSecondSplit)
	fmt.Println("sim score", simScore)
}

func similarityScore(split []int, controlSplit []int) {
	for _, v := range split {
		// find how often in the right list
		similars := 0
		for _, x := range controlSplit {
			if x == v {
				similars++
			}
		}

		sc := v * similars
		simScore += sc
	}
}

func compareLowest() {
	fi := getLowestIndex(firstSplit)
	si := getLowestIndex(secondSplit)

	d := firstSplit[fi] - secondSplit[si]
	if d < 0 {
		d = d * -1
	}

	diff += d

	firstSplit = removeIndex(firstSplit, fi)
	secondSplit = removeIndex(secondSplit, si)
}

func getLowestIndex(s []int) int {
	lowest := 0
	lowestI := 0
	set := false
	for i, v := range s {
		if !set {
			lowest = v
			lowestI = i
			set = true
			continue
		}
		if v < lowest {
			lowest = v
			lowestI = i
		}
	}
	return lowestI
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func readFile(path string) {
	file, err := os.Open(path)
	assert.Nil(err, "Can't read file!")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		split := strings.Split(line, " ")
		firstDone := false
		for _, v := range split {
			if v == "" || v == " " {
				continue
			}
			i, err := strconv.Atoi(v)
			assert.Nil(err, "Strconv Atoi failed")
			if !firstDone {
				firstSplit = append(firstSplit, i)
				firstDone = true
				continue
			}
			secondSplit = append(secondSplit, i)
		}
	}
}
