package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/loissascha/go-assert/assert"
)

type MulResult struct {
	raw string
	res int
}

func Day3() {
	file, err := os.Open("day3.input")
	assert.Nil(err, "Open File failed!")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		muls := getMulsFromInput(line)
		fmt.Println(muls)

		for _, v := range muls {
			sum += v.res
		}
	}
	fmt.Println("sum:", sum)
}

func getMulsFromInput(input string) []MulResult {
	res := []MulResult{}

	currentRead := ""
	readingFirstNum := true
	firstNum := ""
	secondNum := ""
	for i := 0; i < len(input); i++ {
		char := input[i : i+1]
		// fmt.Println(char)

		fmt.Println("currentRead:", currentRead, "next char:", char)
		switch currentRead {
		case "":
			if char == "m" {
				currentRead += char
			}
			break
		case "m":
			if char == "u" {
				currentRead += char
			} else {
				currentRead = ""
			}
			break
		case "mu":
			if char == "l" {
				currentRead += char
			} else {
				currentRead = ""
			}
			break
		case "mul":
			if char == "(" {
				currentRead += char
			} else {
				currentRead = ""
			}
			break
		case "mul(":
			// reading nums
			switch char {
			case ",":
				if !readingFirstNum {
					// reset read
					currentRead = ""
					readingFirstNum = true
					firstNum = ""
					secondNum = ""
					fmt.Println("skip because not reading first num and ,", currentRead)
					continue
				} else {
					readingFirstNum = false
				}
				break
			case ")": // end of num readings
				if firstNum == "" || secondNum == "" {
					// reset read
					currentRead = ""
					readingFirstNum = true
					firstNum = ""
					secondNum = ""
					fmt.Println("First or Second num is empty", firstNum, secondNum, currentRead)
					continue
				}
				fN, err := strconv.Atoi(firstNum)
				if err != nil {
					// reset read
					currentRead = ""
					readingFirstNum = true
					firstNum = ""
					secondNum = ""
					fmt.Println(err, currentRead)
					continue
				}
				sN, err := strconv.Atoi(secondNum)
				if err != nil {
					// reset read
					currentRead = ""
					readingFirstNum = true
					firstNum = ""
					secondNum = ""
					fmt.Println(err, currentRead)
					continue
				}
				mul := fN * sN
				currentRead += firstNum + "," + secondNum + ")"
				fmt.Println("success with", currentRead)
				mulRes := MulResult{
					raw: currentRead,
					res: mul,
				}
				// add to result and reset
				res = append(res, mulRes)
				currentRead = ""
				readingFirstNum = true
				firstNum = ""
				secondNum = ""
				continue
			default:
				// check if its a number -> if no reset
				_, err := strconv.Atoi(char)
				if err != nil {
					currentRead = ""
					readingFirstNum = true
					firstNum = ""
					secondNum = ""
					continue
				}
				if readingFirstNum {
					firstNum += char
				} else {
					secondNum += char
				}
				break
			}
			break
		default:
			currentRead = ""
			break
		}
	}

	return res
}
