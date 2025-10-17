package game

import (
	"maps"
	"strconv"
	"strings"
)

type Game struct {
	Number int
	Sets   []Set
}

func (g *Game) GetPowerOfMinSets() int {
	bag := g.GetMinNecessaryBagToPlayAllSets()
	sum := 0
	for _, v := range bag {
		if sum == 0 {
			sum = v
			continue
		}
		sum *= v
	}
	return sum
}

func (g *Game) GetMinNecessaryBagToPlayAllSets() map[string]int {
	bag := map[string]int{}
	for _, set := range g.Sets {
		for k, v := range set.Content {
			b, ok := bag[k]
			if !ok {
				bag[k] = v
				continue
			}
			if v > b {
				bag[k] = v
			}
		}
	}
	return bag
}

func (g *Game) PossibleWithInput(input map[string]int) bool {
	for _, set := range g.Sets {
		bag := maps.Clone(input)
		for k, v := range set.Content {
			b, ok := bag[k]
			if !ok {
				return false
			}
			if b < v {
				return false
			}
			bag[k] = bag[k] - v
		}
	}
	return true
}

type Set struct {
	Content map[string]int
}

func NewGame(line string) *Game {
	splits := strings.SplitN(line, ":", 2)
	if len(splits) != 2 {
		panic("error split")
	}
	return &Game{
		Number: getGameId(splits[0]),
		Sets:   getGameSets(splits[1]),
	}
}

// input: Game 5
func getGameId(gameName string) int {
	gameName = strings.TrimLeft(gameName, "Game ")
	n, err := strconv.Atoi(gameName)
	if err != nil {
		panic(err)
	}
	return n
}

// input:  3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func getGameSets(input string) []Set {
	result := []Set{}
	splits := strings.SplitSeq(input, ";")
	for item := range splits {
		s := Set{
			Content: getSetContent(item),
		}
		result = append(result, s)
	}
	return result
}

// input:  3 blue, 4 red
func getSetContent(input string) map[string]int {
	result := map[string]int{}
	splits := strings.SplitSeq(input, ",")
	for item := range splits {
		item = strings.TrimSpace(item)
		isplit := strings.SplitN(item, " ", 2)
		if len(isplit) != 2 {
			panic("isplit len not 2")
		}
		iamount, err := strconv.Atoi(isplit[0])
		if err != nil {
			panic(err)
		}
		result[isplit[1]] = iamount
	}
	return result
}
