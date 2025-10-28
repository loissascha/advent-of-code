package main

import (
	"fmt"
	"slices"
)

var letters = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var forbiddenLetters = []byte{'i', 'o', 'l'}

func main() {
	input := "cqjxjnds"
	fmt.Println("input:", input)
	for range 15 {
		input = increasePassword(input)
		fmt.Println("input:", input)
	}
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
