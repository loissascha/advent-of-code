package main

import "fmt"

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
	Spoons     int
}

func (i *Ingredient) getScore(spoons int) (c, d, f, t, ca int) {
	c = i.Capacity * spoons
	d = i.Durability * spoons
	f = i.Flavor * spoons
	t = i.Texture * spoons
	ca = i.Calories * spoons
	return c, d, f, t, ca
}

func main() {

	butterscotch := &Ingredient{
		Name:       "Butterscotch",
		Capacity:   -1,
		Durability: -2,
		Flavor:     6,
		Texture:    3,
		Calories:   8,
	}

	cinnamon := &Ingredient{
		Name:       "Cinnamon",
		Capacity:   2,
		Durability: 3,
		Flavor:     -2,
		Texture:    -1,
		Calories:   3,
	}

	list := []*Ingredient{
		butterscotch, cinnamon,
	}

	highestScore := 0
	for a := range 100 {
		for b := range 100 {
			if a+b == 100 {
				butterscotch.Spoons = a
				cinnamon.Spoons = b
				score := getScore(list)
				if score > highestScore {
					highestScore = score
				}
			}
		}
	}

	fmt.Println("highest score:", highestScore)

	// score := getScore(list)
	// fmt.Println("score:", score)

}

func getScore(ings []*Ingredient) int {
	overallCap := 0
	overallDura := 0
	overallFla := 0
	overallText := 0
	overallCal := 0

	for _, i := range ings {
		ca, dur, fla, text, cal := i.getScore(i.Spoons)
		overallCap += ca
		overallDura += dur
		overallFla += fla
		overallText += text
		overallCal += cal
	}

	overallCap = max(overallCap, 0)
	overallDura = max(overallDura, 0)
	overallFla = max(overallFla, 0)
	overallText = max(overallText, 0)
	overallCal = max(overallCal, 0)
	// overallCal = 1 // TODO: remove
	return overallCap * overallDura * overallFla * overallText * overallCal
}
