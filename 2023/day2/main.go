package main

import (
	"day2/game"
	"day2/readfile"
	"fmt"
)

func main() {
	filename := "test_input.txt"
	lines, err := readfile.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	for line := range lines {
		g := game.NewGame(line)
		fmt.Println(g)
	}
}
