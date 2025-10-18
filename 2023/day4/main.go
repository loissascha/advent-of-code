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
	Instances      int
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

func (c *Card) getMatches() int {
	matches := 0
	for _, tipNum := range c.TipNumbers {
		if slices.Contains(c.WinningNumbers, tipNum) {
			matches += 1
		}
	}
	return matches

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
		Instances:      1,
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
	lines, err := readfile.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	cards := []*Card{}
	for line := range lines {
		card := newCardFromLine(line)
		cards = append(cards, card)
		cardPoints := card.getPoints()
		sum += cardPoints
		// fmt.Println(card, "worth points:", cardPoints)
	}

	// fmt.Println("sum:", sum)

	// part 2
	for n, card := range cards {
		cardMatches := card.getMatches()
		fmt.Println("card matches:", cardMatches)
		for i := range cardMatches {
			instancePlusCard := cards[n+i+1]
			if instancePlusCard == nil {
				fmt.Println("nil card")
			} else {
				fmt.Println("instancepluscard:", instancePlusCard)
				instancePlusCard.Instances += card.Instances
			}
			// fmt.Println(instancePlusCard)
		}

	}

	cardCount := 0
	for _, card := range cards {
		fmt.Println(card)
		cardCount += card.Instances
	}

	fmt.Println("card count:", cardCount)
}
