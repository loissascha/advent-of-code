package main

import (
	"day4/readfile"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	Id             int
	WinningNumbers []int
	TipNumbers     []int
}

func (c *Card) getPoints() int {
	points := 0
	for _, tipNum := range c.TipNumbers {
		if slices.Contains(c.WinningNumbers, tipNum) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func newCardFromLine(line string) *Card {
	splitLine := strings.SplitN(line, ":", 2)

	// card number (id)
	cardNumStr := strings.TrimLeft(splitLine[0], "Card ")
	cardNum, err := strconv.Atoi(cardNumStr)
	if err != nil {
		panic(err)
	}

	card := &Card{
		Id:             cardNum,
		WinningNumbers: []int{},
		TipNumbers:     []int{},
	}

	numsSplit := strings.Split(splitLine[1], "|")
	winningLine := strings.TrimSpace(numsSplit[0])
	tipLine := strings.TrimSpace(numsSplit[1])

	winningSplit := strings.SplitSeq(winningLine, " ")
	for v := range winningSplit {
		if v == "" {
			continue
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		card.WinningNumbers = append(card.WinningNumbers, n)
	}

	tipSpilt := strings.SplitSeq(tipLine, " ")
	for v := range tipSpilt {
		if v == "" {
			continue
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		card.TipNumbers = append(card.TipNumbers, n)
	}

	return card
}

func main() {
	lines, err := readfile.ReadLines("test_input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for line := range lines {
		card := newCardFromLine(line)
		cardPoints := card.getPoints()
		sum += cardPoints
		fmt.Println(card, "worth points:", cardPoints)
	}

	fmt.Println("sum:", sum)
}
