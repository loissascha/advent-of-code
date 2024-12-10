package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/loissascha/go-assert/assert"
)

var hasPath = [][]int{}

func Day10() {
	m := [][]int{}

	file, err := os.Open("day10.test")
	assert.Nil(err, "Can't open file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l := []int{}
		for i := 0; i < len(line); i++ {
			char := line[i : i+1]
			num, err := strconv.Atoi(char)
			assert.Nil(err, "Can't strconv")
			l = append(l, num)
		}
		fmt.Println(l)
		m = append(m, l)
	}

	for i := 0; i < len(m); i++ {
		rl := []int{}
		for j := 0; j < len(m[i]); j++ {
			rl = append(rl, 0)
		}
		hasPath = append(hasPath, rl)
	}

	startTrails(m)

	sum := 0
	for _, line := range hasPath {
		for _, num := range line {
			sum += num
		}
	}
	fmt.Println("sum:", sum)
}

func startTrails(m [][]int) {
	sumTrails := 0
	for y, line := range m {
		for x, num := range line {
			if num == 0 {
				// if m[y][x] == 0 {
				// 	fmt.Println("correct")
				// }
				possibleTrails := findNextNum(m, num, x, y)
				sumTrails += possibleTrails
				fmt.Println("Possible Trails:", possibleTrails)
			}
		}
	}
	fmt.Println("sumtrails:", sumTrails)
}

func findNextNum(m [][]int, currentNum int, x int, y int) (possibleTrails int) {
	trails := 0
	if y-1 >= 0 {
		top := m[y-1][x]
		if top == currentNum+1 {
			if hasPath[y-1][x] != 1 {
				fmt.Println("top field:", top)
				if top == 9 {
					trails++
				}
				trails += findNextNum(m, top, x, y-1)
				hasPath[y-1][x] = 1
			}
		}
	}

	if y+1 < len(m) {
		bottom := m[y+1][x]
		if bottom == currentNum+1 {
			if hasPath[y+1][x] != 1 {
				fmt.Println("bottom field:", bottom)
				if bottom == 9 {
					trails++
				}
				trails += findNextNum(m, bottom, x, y+1)
				hasPath[y+1][x] = 1
			}
		}
	}

	if x-1 >= 0 {
		left := m[y][x-1]
		if left == currentNum+1 {
			if hasPath[y][x-1] != 1 {
				fmt.Println("left field:", left)
				if left == 9 {
					trails++
				}
				trails += findNextNum(m, left, x-1, y)
				hasPath[y][x-1] = 1
			}
		}
	}

	if x+1 < len(m[y]) {
		right := m[y][x+1]
		if right == currentNum+1 {
			if hasPath[y][x+1] != 1 {
				fmt.Println("right field:", right)
				if right == 9 {
					trails++
				}
				trails += findNextNum(m, right, x+1, y)
				hasPath[y][x+1] = 1
			}
		}
	}

	return trails
}
