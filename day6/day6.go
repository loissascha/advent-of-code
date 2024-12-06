package day6

import (
	"bufio"
	"fmt"
	"os"

	"github.com/loissascha/go-assert/assert"
)

var laborMap = [][]string{}

func Day6() {
	readFile("day6.test")
	fmt.Println(laborMap)

	// find first position

	// always find next position
	// check if next position is already found
	// if no -> add 1
}

func readFile(filepath string) {
	file, err := os.Open(filepath)
	assert.Nil(err, "Can't read file!")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		str := []string{}
		for i := 0; i < len(line); i++ {
			char := line[i : i+1]
			str = append(str, char)
		}
		laborMap = append(laborMap, str)
	}
}
