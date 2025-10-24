package main

import "fmt"

const GRID_SIZE int = 1000

type Light struct {
	status bool
}

type Grid struct {
	items [GRID_SIZE][GRID_SIZE]Light
}

func main() {

	grid := Grid{
		items: [GRID_SIZE][GRID_SIZE]Light{},
	}

	for y := range GRID_SIZE {
		for x := range GRID_SIZE {
			grid.items[x][y].status = false
		}
	}

	fmt.Println(grid)

	grid.turnOn(0, 0, 999, 999)
	// grid.toggle(0, 0, 999, 0)
	grid.turnOff(499, 499, 500, 500)

	grid.printLit()
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

func (g *Grid) turnOff(startX, startY, stopX, stopY int) {
	for y := startY; y <= stopY; y++ {
		for x := startX; x <= stopX; x++ {
			g.items[x][y].status = false
		}
	}
}

func (g *Grid) toggle(startX, startY, stopX, stopY int) {
	for y := startY; y <= stopY; y++ {
		for x := startX; x <= stopX; x++ {
			g.items[x][y].status = !g.items[x][y].status
		}
	}
}

func (g *Grid) turnOn(startX, startY, stopX, stopY int) {
	for y := startY; y <= stopY; y++ {
		for x := startX; x <= stopX; x++ {
			g.items[x][y].status = true
		}
	}
}
