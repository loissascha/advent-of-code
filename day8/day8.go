package day8

import (
	"bufio"
	"fmt"
	"os"

	"github.com/loissascha/go-assert/assert"
)

type Point struct {
	x, y int
}

func parseAntennnas(grid []string) map[string][]Point {
	antennas := map[string][]Point{}
	for y, line := range grid {
		for x := 0; x < len(line); x++ {
			char := line[x : x+1]
			if char != "." && char != "#" {
				fmt.Println("found antenna for char:", char, "pos:", x, y)
				antennas[char] = append(antennas[char], Point{x, y})
			}
		}
	}
	return antennas
}

func Day8() {
	grid := readFile("day8.input")
	antennas := parseAntennnas(grid)
	fmt.Println("Antennas:", antennas)

	ps := 0

	antinodes := map[Point]struct{}{}
	for ant, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := 0; j < len(points); j++ {
				if i == j {
					continue
				}
				p1, p2 := points[i], points[j]
				fmt.Println("trying to find antinode point for", ant, "with points", p1, "and", p2)

				dx, dy := p2.x-p1.x, p2.y-p1.y
				if dx < 0 {
					dx *= -1
				}
				if dy < 0 {
					dy *= -1
				}
				fmt.Println("distance x:", dx, "distance y:", dy)

				a1x := 0
				if p1.x > p2.x {
					a1x = p1.x + dx
				} else if p1.x < p2.x {
					a1x = p1.x - dx
				}

				a2x := 0
				if p2.x > p1.x {
					a2x = p2.x + dx
				} else if p2.x < p1.x {
					a2x = p2.x - dx
				}

				a1y := 0
				if p1.y > p2.y {
					a1y = p1.y + dy
				} else if p1.y < p2.y {
					a1y = p1.y - dy
				}

				a2y := 0
				if p2.y > p1.y {
					a2y = p2.y + dy
				} else if p2.y < p1.y {
					a2y = p2.y - dy
				}

				a1 := Point{x: a1x, y: a1y}
				a2 := Point{x: a2x, y: a2y}

				if inBounds(a1, grid) {
					fmt.Println("antinote pos:", a1.x, a1.y)
					drawGridWithAntinote(grid, a1.x, a1.y)
					antinodes[a1] = struct{}{}
					ps++
				}

				if inBounds(a2, grid) {
					fmt.Println("antinote pos:", a2.x, a2.y)
					drawGridWithAntinote(grid, a2.x, a2.y)
					antinodes[a2] = struct{}{}
					ps++
				}
			}
		}
	}
	fmt.Println("points:", len(antinodes))
}

func drawGridWithAntinote(grid []string, ax int, ay int) {
	for y, v := range grid {
		for x := 0; x < len(v); x++ {
			char := v[x : x+1]
			if y == ay && x == ax {
				fmt.Print("#")
			} else {
				fmt.Print(char)
			}
		}
		fmt.Print("\n")
	}
}

func inBounds(p Point, grid []string) bool {
	return p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[p.y])
}

func readFile(filepath string) []string {
	file, err := os.Open(filepath)
	assert.Nil(err, "Can't read file")
	defer file.Close()

	// Read the input map from stdin
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		grid = append(grid, line)
	}
	return grid
}
