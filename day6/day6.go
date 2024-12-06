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
}

func readFile(filepath string) {
	file, err := os.Open(filepath)
	assert.Nil(err, "Can't read file!")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
