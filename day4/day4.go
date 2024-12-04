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

	fmt.Println(content)
	found := search(content)
	fmt.Println("found:", found)
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
							fmt.Println("found top XMAS")
							found++
						}
					}
				}
				if bottomCharacter(content, ci, i) == "M" {
					if bottomCharacter(content, ci+1, i) == "A" {
						if bottomCharacter(content, ci+2, i) == "S" {
							fmt.Println("found bottom XMAS")
							found++
						}
					}
				}
				if leftCharacter(content, ci, i) == "M" {
					if leftCharacter(content, ci, i-1) == "A" {
						if leftCharacter(content, ci, i-2) == "S" {
							fmt.Println("found left XMAS")
							found++
						}
					}
				}
				if rightCharacter(content, ci, i) == "M" {
					if rightCharacter(content, ci, i+1) == "A" {
						if rightCharacter(content, ci, i+2) == "S" {
							fmt.Println("found right XMAS")
							found++
						}
					}
				}
				if topLeftCharacter(content, ci, i) == "M" {
					if topLeftCharacter(content, ci-1, i-1) == "A" {
						if topLeftCharacter(content, ci-2, i-2) == "S" {
							fmt.Println("found top left XMAS")
							found++
						}
					}
				}
				if topRightCharacter(content, ci, i) == "M" {
					if topRightCharacter(content, ci-1, i+1) == "A" {
						if topRightCharacter(content, ci-2, i+2) == "S" {
							fmt.Println("found top right XMAS")
							found++
						}
					}
				}
				if bottomLeftCharacter(content, ci, i) == "M" {
					if bottomLeftCharacter(content, ci+1, i-1) == "A" {
						if bottomLeftCharacter(content, ci+2, i-2) == "S" {
							fmt.Println("found bottom left XMAS")
							found++
						}
					}
				}
				if bottomRightCharacter(content, ci, i) == "M" {
					if bottomRightCharacter(content, ci+1, i+1) == "A" {
						if bottomRightCharacter(content, ci+2, i+2) == "S" {
							fmt.Println("found bottom right XMAS")
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
