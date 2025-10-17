package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/loissascha/go-assert/assert"
)

type MulResult struct {
	raw string
	res int
}

var doEnabled = true

func Day3() {
	file, err := os.Open("day3.input")
	assert.Nil(err, "Open File failed!")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		muls := readMem(line)
		fmt.Println(muls)

		for _, v := range muls {
			sum += v.res
		}
	}
	fmt.Println("sum:", sum)
}

func Next(s string, currentIndex int) string {
	i := currentIndex + 1
	return s[i : i+1]
}

func Until(input string, startIndex int, until string) (string, bool) {
	str := input[startIndex+1:]
	before, _, ok := strings.Cut(str, until)
	if !ok {
		return "", false
	}
	return before, true
}

func readMem(input string) []MulResult {
	res := []MulResult{}

	for i := 0; i < len(input); i++ {
		char := input[i : i+1]

		// do don't
		if char == "d" {
			customI := i
			if Next(input, customI) == "o" {
				customI++
				if Next(input, customI) == "(" {
					customI++
					if Next(input, customI) == ")" {
						doEnabled = true
						fmt.Println("found do")
					}
				} else if Next(input, customI) == "n" {
					customI++
					if Next(input, customI) == "'" {
						customI++
						if Next(input, customI) == "t" {
							customI++
							if Next(input, customI) == "(" {
								customI++
								if Next(input, customI) == ")" {
									doEnabled = false
									fmt.Println("found dont")
								}
							}
						}
					}
				}
			}
		}

		// mul(x,y)
		if char == "m" {
			customI := i
			if Next(input, customI) == "u" {
				customI++
				if Next(input, customI) == "l" {
					customI++
					if Next(input, customI) == "(" {
						customI++
						fmt.Println("found mul(")
						firstUntil, found := Until(input, customI, ",")
						if !found {
							continue
						}
						fmt.Println("firstUntil:", firstUntil)

						customI += len(firstUntil) + 1
						secondUntil, found := Until(input, customI, ")")
						if !found {
							continue
						}
						fmt.Println("secondUntil", secondUntil)

						// first until and second until are valid numbers = success!

						fN, err := strconv.Atoi(firstUntil)
						if err != nil {
							fmt.Println("firstUntil not valid")
							continue
						}
						sN, err := strconv.Atoi(secondUntil)
						if err != nil {
							fmt.Println("secondUntil not valid")
							continue
						}

						if doEnabled {
							m := fN * sN
							r := MulResult{
								raw: fmt.Sprintf("mul(%v,%v)", fN, sN),
								res: m,
							}
							res = append(res, r)
							fmt.Println("added!")
						} else {
							fmt.Println("doEnabled false")
						}
					}
				}
			}
		}
	}

	return res
}
