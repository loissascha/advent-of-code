package main

import (
	"day4/readfile"
	"fmt"
)

func main() {
	lines, err := readfile.ReadLines("test_input.txt")
	if err != nil {
		panic(err)
	}

	for line := range lines {
		fmt.Println(line)
	}
}
