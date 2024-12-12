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

func Day12() {
	m := readFile("day12.test")
	for _, v := range m {
		fmt.Println(v)
	}
	combinedPlotLines := []CombinedPlotLine{}
	for y, line := range m {
		plotLines := mapLine(line, y)
		combinedPlotLines = combinePlotLines(combinedPlotLines, plotLines)
	}
	for _, v := range combinedPlotLines {
		fmt.Println(v)
		calculatePerimeter(v)
	}
	fmt.Println(len(combinedPlotLines))
}

func calculatePerimeter(cpl CombinedPlotLine) {
	perimeter := 0
	for _, pl := range cpl.rows {
		y := pl.y
		for _, x := range pl.fields {
			// check if there exists an item on the right and the bottom and the bottom right
			// check if there exists an item on the right and the top and the top right
			// check if there exists an item on the left and the top and the top right
			// check if there exists an item on the left and the bottom and the bottom left
		}
	}
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
