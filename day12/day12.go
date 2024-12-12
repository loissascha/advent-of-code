package day12

import (
	"bufio"
	"fmt"
	"os"

	"github.com/loissascha/go-assert/assert"
)

type Plot struct {
	char   string
	fields map[int][]int
}

func Day12() {
	m := readFile("day12.test")
	for _, v := range m {
		fmt.Println(v)
	}
	plots := findPlots(m)
	fmt.Println(plots)
	fmt.Println("amount plots:", len(plots))
	for _, v := range plots {
		amount := 0
		for _, p := range v.fields {
			amount += len(p)
		}
		fmt.Println(v.char, "with:", amount)
	}
}

func findPlots(input [][]string) []Plot {
	plots := []Plot{}
	findPlot(input, 0, 0)
	return plots
}

func findPlot(input [][]string, x int, y int) {
	foundPositions := make(map[int][]int)
	foundPositions[0] = []int{0}
	foundPositions = findDir(input, x, y, 1, 0, foundPositions)
	foundPositions = findDir(input, x, y, -1, 0, foundPositions)
	foundPositions = findDir(input, x, y, 0, -1, foundPositions)
	foundPositions = findDir(input, x, y, 0, 1, foundPositions)
	fmt.Println("find plot result for x", x, "y", y)
	fmt.Println(foundPositions)
}

func hasFoundPosition(input map[int][]int, x int, y int) bool {
	for yy, v := range input {
		if y != yy {
			continue
		}
		for _, n := range v {
			if n == x {
				return true
			}
		}
	}
	return false
}

func findDir(input [][]string, x int, y int, dirX int, dirY int, foundPositions map[int][]int) map[int][]int {
	char := input[y][x]
	if !hasFoundPosition(foundPositions, x, y) {
		next := input[y+dirY][x+dirX]
		if next == char && hasFoundPosition(foundPositions, x+dirX, y+dirY) {
			fmt.Println("found pos: ", x, y)
			ll := foundPositions[y]
			ll = append(ll, x)
			foundPositions[y] = ll
			foundPositions = findDir(input, x+dirX, y+dirY, 1, 0, foundPositions)
			foundPositions = findDir(input, x+dirX, y+dirY, -1, 0, foundPositions)
			foundPositions = findDir(input, x+dirX, y+dirY, 0, 1, foundPositions)
			foundPositions = findDir(input, x+dirX, y+dirY, 0, -1, foundPositions)
		}
	}
	return foundPositions
}

func readFile(filepath string) [][]string {
	file, err := os.Open(filepath)
	assert.Nil(err, "Can't open file")
	defer file.Close()

	res := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ls := []string{}
		for i := 0; i < len(line); i++ {
			char := line[i : i+1]
			ls = append(ls, char)
		}
		res = append(res, ls)
	}
	return res
}
