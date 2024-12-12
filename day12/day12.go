package day12

import (
	"bufio"
	"fmt"
	"os"

	"github.com/loissascha/go-assert/assert"
)

type PlotLine struct {
	char   string
	fields []int
	y      int
}

type CombinedPlotLine struct {
	char string
	rows []PlotLine
}

var maxX = 0
var maxY = 0

func Day12() {
	m := readFile("day12.input")
	maxY = len(m)
	for _, v := range m {
		maxX = len(v)
		fmt.Println(v)
	}
	combinedPlotLines := []CombinedPlotLine{}
	for y, line := range m {
		plotLines := mapLine(line, y)
		combinedPlotLines = combinePlotLines(combinedPlotLines, plotLines)
	}
	sum := 0
	for _, v := range combinedPlotLines {
		fmt.Println(v)
		per := printCombinedPlotLine(v)
		perRaw := calculatePerimeterString(per)
		perimeter := countPer(perRaw)
		fields := countFields(perRaw)
		fmt.Println("perimeter:", perimeter)
		fmt.Println("fields:", fields)
		sum += (perimeter * fields)
	}
	fmt.Println("Sum:", sum)
}

func countFields(input [][]string) int {
	fields := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			char := input[y][x]
			if char != "+" && char != " " {
				fields++
			}
		}
	}
	return fields
}

func countPer(input [][]string) int {
	per := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			char := input[y][x]
			if char == "+" {
				per++
			}
		}
	}
	return per
}

func calculatePerimeterString(input [][]string) [][]string {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			char := input[y][x]
			if char == " " {
				continue
			}
			if char == "+" {
				continue
			}
			fmt.Println("no skipo", char, "at pos", x, y)
			leftX := x - 2
			rightX := x + 2
			topY := y - 2
			bottomY := y + 2
			writeLeft := false
			writeRight := false
			writeTop := false
			writeBottom := false
			if leftX > 0 {
				left := input[y][leftX]
				if left == " " {
					writeLeft = true
				}
			} else {
				writeLeft = true
			}

			if rightX < len(input[y]) {
				right := input[y][rightX]
				if right == " " {
					writeRight = true
				}
			} else {
				writeRight = true
			}

			if topY > 0 {
				top := input[topY][x]
				if top == " " {
					writeTop = true
				}
			} else {
				writeTop = true
			}

			if bottomY < len(input) {
				bottom := input[bottomY][x]
				if bottom == " " {
					writeBottom = true
				}
			} else {
				writeBottom = true
			}

			if writeLeft {
				fmt.Println("+ 1", x, y)
				input[y-1][x-1] = "+"
				input[y+1][x-1] = "+"
			}

			if writeRight {
				fmt.Println("+ 2", x, y)
				input[y-1][x+1] = "+"
				input[y+1][x+1] = "+"
			}

			if writeTop {
				fmt.Println("+ 3", x, y)
				input[y-1][x+1] = "+"
				input[y-1][x-1] = "+"
			}

			if writeBottom {
				fmt.Println("+ 4", x, y)
				input[y+1][x+1] = "+"
				input[y+1][x-1] = "+"
			}
		}
	}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			char := input[y][x]
			fmt.Print(char)
		}
		fmt.Print("\n")
	}

	return input
}

func printCombinedPlotLine(cpl CombinedPlotLine) [][]string {
	res := [][]string{}
	for yy := -1; yy < maxY; yy++ {
		resline := []string{}
		emptyline := []string{}
		for xx := -1; xx < maxX; xx++ {
			foundX := false
			for _, pl := range cpl.rows {
				y := pl.y
				if y != yy {
					continue
				}
				for _, x := range pl.fields {
					if x != xx {
						continue
					}
					resline = append(resline, pl.char+" ")
					resline = append(resline, " ")
					foundX = true
				}
			}
			if !foundX {
				resline = append(resline, " ")
				resline = append(resline, " ")
			}
			emptyline = append(emptyline, " ")
			emptyline = append(emptyline, " ")
		}
		res = append(res, resline)
		res = append(res, emptyline)
	}
	return res
}

func addToPerimeterMap(perimeterMap map[uint]map[uint]int, x int, y int) map[uint]map[uint]int {
	foundY := false
	for yy := range perimeterMap {
		if yy != uint(y) {
			continue
		}
		foundY = true
		perimeterMap[uint(y)][uint(x)] = 1
	}

	if !foundY {
		// create
		perimeterMap[uint(y)] = make(map[uint]int)
		perimeterMap[uint(y)][uint(x)] = 1
	}
	return perimeterMap
}

func hasPos(cpl CombinedPlotLine, x int, y int) bool {
	for _, pl := range cpl.rows {
		yy := pl.y
		if yy != y {
			continue
		}
		for _, xx := range pl.fields {
			if xx == x {
				return true
			}
		}
	}

	return false
}

func combinePlotLines(combinedPlotLines []CombinedPlotLine, plotLines []PlotLine) []CombinedPlotLine {
	workedPlotLinesIndexes := []int{}
	for i, cpl := range combinedPlotLines {
		for ii, pl := range plotLines {
			if pl.char != cpl.char {
				continue
			}
			canCombine := false
			for _, plfield := range pl.fields {
				for _, row := range cpl.rows {
					for _, field := range row.fields {
						if field == plfield && (pl.y == row.y-1 || pl.y == row.y+1) {
							// should be connected because fields match
							canCombine = true
						}
					}
				}
			}
			if canCombine {
				combinedPlotLines[i].rows = append(combinedPlotLines[i].rows, pl)
				workedPlotLinesIndexes = append(workedPlotLinesIndexes, ii)
			}
		}
	}

	for i, pl := range plotLines {
		found := false
		for _, wp := range workedPlotLinesIndexes {
			if wp == i {
				found = true
			}
		}
		if found {
			continue
		}
		combinedPlotLines = append(combinedPlotLines, CombinedPlotLine{char: pl.char, rows: []PlotLine{pl}})
	}

	return combinedPlotLines
}

func mapLine(line []string, y int) []PlotLine {
	plotLines := []PlotLine{}
	prevChar := ""
	currentFields := []int{}
	for x, char := range line {
		if char == prevChar {
			currentFields = append(currentFields, x)
			continue
		}

		if len(currentFields) > 0 {
			plotLines = append(plotLines, PlotLine{char: prevChar, fields: currentFields, y: y})
		}
		prevChar = char
		currentFields = []int{x}
	}
	if len(currentFields) > 0 {
		plotLines = append(plotLines, PlotLine{char: prevChar, fields: currentFields, y: y})
	}
	// fmt.Println(plotLines)
	return plotLines
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
