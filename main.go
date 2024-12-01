package main

import (
	"fmt"
)

var firstSplit = []int{}
var secondSplit = []int{}

var diff = 0
var simScore = 0

func main() {
	readFile("input1.txt")
	fmt.Println(len(firstSplit), len(secondSplit))
	l := len(firstSplit)
	ogFirstSplit := make([]int, len(firstSplit))
	ogSecondSplit := make([]int, len(secondSplit))
	copy(ogFirstSplit, firstSplit)
	copy(ogSecondSplit, secondSplit)

	for i := 0; i < l; i++ {
		compareLowest()
	}
	fmt.Println("diff is", diff)

	similarityScore(ogFirstSplit, ogSecondSplit)
	fmt.Println("sim score", simScore)
}
