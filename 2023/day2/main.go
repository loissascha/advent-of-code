package main

import (
	"day2/game"
	"day2/readfile"
	"fmt"
)

func main() {
	filename := "input.txt"
	lines, err := readfile.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	possibleGames := []game.Game{}

	sum := 0
	sum2 := 0

	checkFor := map[string]int{}
	checkFor["red"] = 12
	checkFor["green"] = 13
	checkFor["blue"] = 14

	for line := range lines {
		g := game.NewGame(line)
		if g.PossibleWithInput(checkFor) {
			possibleGames = append(possibleGames, *g)
			fmt.Println(g)
			sum += g.Number
		}
		setPower := g.GetPowerOfMinSets()
		fmt.Println("set", g.Number, "power:", setPower)
		sum2 += setPower
	}

	fmt.Println("sum:", sum)
	fmt.Println("sum2:", sum2)
}
