package main

import (
	"fmt"
	"slices"
	"strings"
)

var letters = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var forbiddenLetters = []byte{'i', 'o', 'l'}

func main() {
	password := "cqjxjnds"
	fmt.Println("input:", password)
	for {
		password = increasePassword(password)
		// fmt.Println("input:", password)
		if !hasForbiddenLetter(password) && hasIncreasingStraight(password) && hasPairs(password) >= 2 {
			break
		}
	}

	fmt.Println("next password:", password)
}

func hasPairs(input string) int {
	pairPositions := []int{}
	pairs := 0
	for n, b := range input {
		if n+1 >= len(input) {
			break
		}
		next := input[n+1]
		if next == byte(b) {
			if !slices.Contains(pairPositions, n) && !slices.Contains(pairPositions, n+1) {
				pairPositions = append(pairPositions, n)
				pairPositions = append(pairPositions, n+1)
				pairs++
			}
		}
	}
	return pairs
}

func hasIncreasingStraight(input string) bool {
	for n, l := range letters {
		if n+2 >= len(letters) {
			break
		}
		straight := []byte{l}
		straight = append(straight, letters[n+1])
		straight = append(straight, letters[n+2])
		if strings.Contains(input, string(straight)) {
			return true
		}
	}
	return false
}

func hasForbiddenLetter(input string) bool {
	for _, l := range forbiddenLetters {
		if strings.Contains(input, string(l)) {
			return true
		}
	}
	return false
}

func increasePassword(input string) string {
	newpw := []byte{}

	bytes := []byte(input)
	slices.Reverse(bytes)

	increaseNext := true
	for _, b := range bytes {
		if increaseNext {
			if b == 'z' {
				newpw = append(newpw, 'a')
				increaseNext = true
			} else {
				increaseNext = false
				useL := false
				for _, l := range letters {
					if l == b {
						useL = true
						continue
					}
					if useL {
						newpw = append(newpw, l)
						break
					}
				}
			}
		} else {
			newpw = append(newpw, b)
		}
	}

	slices.Reverse(newpw)

	return string(newpw)
}
