package day4

import (
	"bufio"
	"fmt"
	"os"

	"github.com/loissascha/go-assert/assert"
)

func Day4() {
	file, err := os.Open("day4.input")
	assert.Nil(err, "Can't read file!")
	defer file.Close()

	content := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}

	// fmt.Println(content)
	found := search(content)
	fmt.Println("found part 1:", found)
	found = search2(content)
	fmt.Println("found part 2:", found)
}

func masInShape(tL string, tR string, bL string, bR string) int {
	found := 0
	if tL == "M" && tR == "S" {
		found++
		if bL == "M" && bR == "S" {
			found++
		}
	} else if tL == "S" && tR == "M" {
		found++
		if bL == "S" && bR == "M" {
			found++
		}
	} else if tL == "M" && bL == "S" {
		found++
		if tR == "M" && bR == "S" {
			found++
		}
	} else if tL == "S" && bL == "M" {
		found++
		if tR == "S" && bR == "M" {
			found++
		}
	}
	return found
}

func search2(content []string) int {
	found := 0
	for ci, c := range content {
		for i := 0; i < len(c); i++ {
			char := c[i : i+1]

			if char == "A" {
				topLeft := topLeftCharacter(content, ci, i)
				topRight := topRightCharacter(content, ci, i)
				bottomLeft := bottomLeftCharacter(content, ci, i)
				bottomRight := bottomRightCharacter(content, ci, i)

				if topLeft == "" || topRight == "" || bottomLeft == "" || bottomRight == "" {
					continue
				}

				count := masInShape(topLeft, topRight, bottomLeft, bottomRight)
				// fmt.Printf("%v,%v\n,%v,\n%v,%v\n", topLeft, topRight, "A", bottomLeft, bottomRight)
				// fmt.Println("count is", count)
				if count >= 2 {
					found++
				}
			}
		}
	}

	return found
}

func search(content []string) int {
	found := 0
	for ci, c := range content {
		for i := 0; i < len(c); i++ {
			char := c[i : i+1]
			if char == "X" { // try and find match
				if topCharacter(content, ci, i) == "M" {
					if topCharacter(content, ci-1, i) == "A" {
						if topCharacter(content, ci-2, i) == "S" {
							found++
						}
					}
				}
				if bottomCharacter(content, ci, i) == "M" {
					if bottomCharacter(content, ci+1, i) == "A" {
						if bottomCharacter(content, ci+2, i) == "S" {
							found++
						}
					}
				}
				if leftCharacter(content, ci, i) == "M" {
					if leftCharacter(content, ci, i-1) == "A" {
						if leftCharacter(content, ci, i-2) == "S" {
							found++
						}
					}
				}
				if rightCharacter(content, ci, i) == "M" {
					if rightCharacter(content, ci, i+1) == "A" {
						if rightCharacter(content, ci, i+2) == "S" {
							found++
						}
					}
				}
				if topLeftCharacter(content, ci, i) == "M" {
					if topLeftCharacter(content, ci-1, i-1) == "A" {
						if topLeftCharacter(content, ci-2, i-2) == "S" {
							found++
						}
					}
				}
				if topRightCharacter(content, ci, i) == "M" {
					if topRightCharacter(content, ci-1, i+1) == "A" {
						if topRightCharacter(content, ci-2, i+2) == "S" {
							found++
						}
					}
				}
				if bottomLeftCharacter(content, ci, i) == "M" {
					if bottomLeftCharacter(content, ci+1, i-1) == "A" {
						if bottomLeftCharacter(content, ci+2, i-2) == "S" {
							found++
						}
					}
				}
				if bottomRightCharacter(content, ci, i) == "M" {
					if bottomRightCharacter(content, ci+1, i+1) == "A" {
						if bottomRightCharacter(content, ci+2, i+2) == "S" {
							found++
						}
					}
				}
			}
		}
	}
	return found
}

func topRightCharacter(content []string, contentIndex int, lineIndex int) string {
	if contentIndex == 0 {
		return ""
	}
	line := content[contentIndex-1]
	if lineIndex >= len(line)-1 {
		return ""
	}
	return line[lineIndex+1 : lineIndex+2]
}

func topLeftCharacter(content []string, contentIndex int, lineIndex int) string {
	if contentIndex == 0 {
		return ""
	}
	if lineIndex == 0 {
		return ""
	}
	line := content[contentIndex-1]
	return line[lineIndex-1 : lineIndex]
}

func bottomRightCharacter(content []string, contentIndex int, lineIndex int) string {
	if contentIndex >= len(content)-1 {
		return ""
	}
	line := content[contentIndex+1]
	if lineIndex >= len(line)-1 {
		return ""
	}
	return line[lineIndex+1 : lineIndex+2]
}

func bottomLeftCharacter(content []string, contentIndex int, lineIndex int) string {
	if contentIndex >= len(content)-1 {
		return ""
	}
	if lineIndex == 0 {
		return ""
	}
	line := content[contentIndex+1]
	return line[lineIndex-1 : lineIndex]
}

func rightCharacter(content []string, contentIndex int, lineIndex int) string {
	line := content[contentIndex]
	if lineIndex >= len(line)-1 {
		return ""
	}
	return line[lineIndex+1 : lineIndex+2]
}

func leftCharacter(content []string, contentIndex int, lineIndex int) string {
	if lineIndex == 0 {
		return ""
	}
	line := content[contentIndex]
	return line[lineIndex-1 : lineIndex]
}

func bottomCharacter(content []string, contentIndex int, lineIndex int) string {
	if contentIndex >= len(content)-1 {
		return ""
	}
	line := content[contentIndex+1]
	return line[lineIndex : lineIndex+1]
}

func topCharacter(content []string, contentIndex int, lineIndex int) string {
	if contentIndex == 0 {
		return ""
	}
	line := content[contentIndex-1]
	return line[lineIndex : lineIndex+1]
}
