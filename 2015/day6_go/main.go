package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const GRID_SIZE int = 1000

type Light struct {
	status     bool
	brightness int
}

type Grid struct {
	items [GRID_SIZE][GRID_SIZE]Light
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(contents), "\n")

	grid := Grid{
		items: [GRID_SIZE][GRID_SIZE]Light{},
	}

	for y := range GRID_SIZE {
		for x := range GRID_SIZE {
			grid.items[x][y].status = false
		}
	}

	for _, line := range split {
		grid.processLine(line)
	}

	grid.printLit()
	grid.printBrightness()
}

func (g *Grid) processLine(line string) {
	if strings.Contains(line, "turn on") {
		line = strings.TrimLeft(line, "turn on ")
		startX, startY, stopX, stopY := getCoords(line)
		g.turnOn(startX, startY, stopX, stopY)
	} else if strings.Contains(line, "turn off") {
		line = strings.TrimLeft(line, "turn off ")
		startX, startY, stopX, stopY := getCoords(line)
		g.turnOff(startX, startY, stopX, stopY)
	} else if strings.Contains(line, "toggle") {
		line = strings.TrimLeft(line, "toggle ")
		startX, startY, stopX, stopY := getCoords(line)
		g.toggle(startX, startY, stopX, stopY)
	}
}

func getCoords(line string) (startX int, startY int, stopX int, stopY int) {
	split := strings.Split(line, " through ")
	startCoords := split[0]
	stopCoords := split[1]
	startX, startY = toCoords(startCoords)
	stopX, stopY = toCoords(stopCoords)
	return startX, startY, stopX, stopY
}

func toCoords(input string) (int, int) {
	split := strings.Split(input, ",")
	x, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return x, y
}

func (g *Grid) printLit() {
	lit := 0
	for y := range GRID_SIZE {
		for x := range GRID_SIZE {
			if g.items[x][y].status {
				lit++
			}
		}
	}
	fmt.Println("lit:", lit)
}

func (g *Grid) printBrightness() {
	lit := 0
	for y := range GRID_SIZE {
		for x := range GRID_SIZE {
			lit += g.items[x][y].brightness
		}
	}
	fmt.Println("brightness:", lit)
}

func (g *Grid) turnOff(startX, startY, stopX, stopY int) {
	for y := startY; y <= stopY; y++ {
		for x := startX; x <= stopX; x++ {
			g.items[x][y].status = false
			g.items[x][y].brightness = max(g.items[x][y].brightness-1, 0)
			// g.items[x][y].brightness -= 1
			// if g.items[x][y].brightness < 0 {
			// 	g.items[x][y].brightness = 0
			// }
		}
	}
}

func (g *Grid) toggle(startX, startY, stopX, stopY int) {
	for y := startY; y <= stopY; y++ {
		for x := startX; x <= stopX; x++ {
			g.items[x][y].status = !g.items[x][y].status
			g.items[x][y].brightness += 2
		}
	}
}

func (g *Grid) turnOn(startX, startY, stopX, stopY int) {
	for y := startY; y <= stopY; y++ {
		for x := startX; x <= stopX; x++ {
			g.items[x][y].status = true
			g.items[x][y].brightness += 1
		}
	}
}
