package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/loissascha/go-assert/assert"
)

type TrailHead struct {
	x int
	y int
}

func Day10() {
	m := [][]int{}

	file, err := os.Open("day10.input")
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

	startTrails(m)
}

func startTrails(m [][]int) {
	sumTrails := 0
	sumTrails2 := 0
	for y, line := range m {
		for x, num := range line {
			if num == 0 {
				// part 1
				possibleNines := findPossibleNines(m, num, x, y)
				trails := []TrailHead{}
				for _, v := range possibleNines {
					trails = addTrailHead(trails, v)
				}
				fmt.Println("Possible Nines:", len(trails))
				fmt.Println(trails)
				sumTrails += len(trails)

				// part 2
				possibleTrails := findNextNum(m, num, x, y)
				sumTrails2 += possibleTrails
			}
		}
	}
	fmt.Println("sumtrails:", sumTrails)
	fmt.Println("sumtrails2:", sumTrails2)
}

func addTrailHead(e []TrailHead, i TrailHead) []TrailHead {
	found := false
	for _, v := range e {
		if v.x == i.x && v.y == i.y {
			found = true
		}
	}
	if !found {
		e = append(e, i)
	}
	return e
}

func findPossibleNines(m [][]int, currentNum int, x int, y int) []TrailHead {
	trails := []TrailHead{}
	if y-1 >= 0 {
		top := m[y-1][x]
		if top == currentNum+1 {
			// fmt.Println("top field:", top)
			if top == 9 {
				trails = addTrailHead(trails, TrailHead{x: x, y: y - 1})
			}
			trails = append(trails, findPossibleNines(m, top, x, y-1)...)
		}
	}

	if y+1 < len(m) {
		bottom := m[y+1][x]
		if bottom == currentNum+1 {
			// fmt.Println("bottom field:", bottom)
			if bottom == 9 {
				trails = addTrailHead(trails, TrailHead{x: x, y: y + 1})
			}
			trails = append(trails, findPossibleNines(m, bottom, x, y+1)...)
		}
	}

	if x-1 >= 0 {
		left := m[y][x-1]
		if left == currentNum+1 {
			// fmt.Println("left field:", left)
			if left == 9 {
				trails = addTrailHead(trails, TrailHead{x: x - 1, y: y})
			}
			trails = append(trails, findPossibleNines(m, left, x-1, y)...)
		}
	}

	if x+1 < len(m[y]) {
		right := m[y][x+1]
		if right == currentNum+1 {
			// fmt.Println("right field:", right)
			if right == 9 {
				trails = addTrailHead(trails, TrailHead{x: x + 1, y: y})
			}
			trails = append(trails, findPossibleNines(m, right, x+1, y)...)
		}
	}

	return trails
}

func findNextNum(m [][]int, currentNum int, x int, y int) (possibleTrails int) {
	trails := 0
	if y-1 >= 0 {
		top := m[y-1][x]
		if top == currentNum+1 {
			// fmt.Println("top field:", top)
			if top == 9 {
				trails++
			}
			trails += findNextNum(m, top, x, y-1)
		}
	}

	if y+1 < len(m) {
		bottom := m[y+1][x]
		if bottom == currentNum+1 {
			// fmt.Println("bottom field:", bottom)
			if bottom == 9 {
				trails++
			}
			trails += findNextNum(m, bottom, x, y+1)
		}
	}

	if x-1 >= 0 {
		left := m[y][x-1]
		if left == currentNum+1 {
			// fmt.Println("left field:", left)
			if left == 9 {
				trails++
			}
			trails += findNextNum(m, left, x-1, y)
		}
	}

	if x+1 < len(m[y]) {
		right := m[y][x+1]
		if right == currentNum+1 {
			// fmt.Println("right field:", right)
			if right == 9 {
				trails++
			}
			trails += findNextNum(m, right, x+1, y)
		}
	}

	// if currentNum == 8 && trails > 1 {
	// 	return 1
	// }

	return trails
}
