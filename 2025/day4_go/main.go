package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x      int
	y      int
	symbol string
}

func main() {
	input := readFile("input.txt")
	lines := strings.Split(input, "\n")

	grid := map[string]Pos{}

	for y, line := range lines {
		if line == "" {
			continue
		}
		for x := 0; x < len(line); x++ {
			grid[fmt.Sprintf("%d-%d", x, y)] = Pos{
				x:      x,
				y:      y,
				symbol: string(line[x]),
			}
		}
	}

	fmt.Println(grid)

	removedOverall := 0
	accessible := 1
	for accessible > 0 {
		accessible = 0
		for k, v := range grid {
			rollsAround := 0
			topLeft, ok := grid[fmt.Sprintf("%d-%d", v.x-1, v.y-1)]
			if ok {
				if topLeft.symbol == "@" {
					rollsAround++
				}
			}
			top, ok := grid[fmt.Sprintf("%d-%d", v.x, v.y-1)]
			if ok {
				if top.symbol == "@" {
					rollsAround++
				}
			}
			topRight, ok := grid[fmt.Sprintf("%d-%d", v.x+1, v.y-1)]
			if ok {
				if topRight.symbol == "@" {
					rollsAround++
				}
			}
			right, ok := grid[fmt.Sprintf("%d-%d", v.x+1, v.y)]
			if ok {
				if right.symbol == "@" {
					rollsAround++
				}
			}
			bottomRight, ok := grid[fmt.Sprintf("%d-%d", v.x+1, v.y+1)]
			if ok {
				if bottomRight.symbol == "@" {
					rollsAround++
				}
			}
			bottom, ok := grid[fmt.Sprintf("%d-%d", v.x, v.y+1)]
			if ok {
				if bottom.symbol == "@" {
					rollsAround++
				}
			}
			bottomLeft, ok := grid[fmt.Sprintf("%d-%d", v.x-1, v.y+1)]
			if ok {
				if bottomLeft.symbol == "@" {
					rollsAround++
				}
			}
			left, ok := grid[fmt.Sprintf("%d-%d", v.x-1, v.y)]
			if ok {
				if left.symbol == "@" {
					rollsAround++
				}
			}
			if v.symbol == "@" {
				if rollsAround < 4 {
					grid[k] = Pos{
						x:      v.x,
						y:      v.y,
						symbol: ".",
					}
					accessible++
				}
			}
		}
		removedOverall += accessible
	}

	fmt.Println("there are", removedOverall, "accessible")
}

func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}
