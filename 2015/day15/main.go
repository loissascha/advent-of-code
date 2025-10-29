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

	sprinkles := &Ingredient{
		Name:       "Sprinkles",
		Capacity:   5,
		Durability: -1,
		Flavor:     0,
		Texture:    0,
		Calories:   5,
	}

	peanutButter := &Ingredient{
		Name:       "Peanut Butter",
		Capacity:   -1,
		Durability: 3,
		Flavor:     0,
		Texture:    0,
		Calories:   1,
	}

	frosting := &Ingredient{
		Name:       "Frosting",
		Capacity:   0,
		Durability: -1,
		Flavor:     4,
		Texture:    0,
		Calories:   6,
	}

	sugar := &Ingredient{
		Name:       "Sugar",
		Capacity:   -1,
		Durability: 0,
		Flavor:     0,
		Texture:    2,
		Calories:   8,
	}

	// butterscotch := &Ingredient{
	// 	Name:       "Butterscotch",
	// 	Capacity:   -1,
	// 	Durability: -2,
	// 	Flavor:     6,
	// 	Texture:    3,
	// 	Calories:   8,
	// }
	//
	// cinnamon := &Ingredient{
	// 	Name:       "Cinnamon",
	// 	Capacity:   2,
	// 	Durability: 3,
	// 	Flavor:     -2,
	// 	Texture:    -1,
	// 	Calories:   3,
	// }

	list := []*Ingredient{
		sprinkles, peanutButter, frosting, sugar,
	}

	highestScore := 0
	for a := range 100 {
		for b := range 100 {
			for c := range 100 {
				for d := range 100 {
					if a+b+c+d == 100 {
						sprinkles.Spoons = a
						peanutButter.Spoons = b
						frosting.Spoons = c
						sugar.Spoons = d
						score := getScore(list)
						if score > highestScore {
							highestScore = score
						}
					}
				}
			}
			// if a+b == 100 {
			// 	butterscotch.Spoons = a
			// 	cinnamon.Spoons = b
			// 	score := getScore(list)
			// 	if score > highestScore {
			// 		highestScore = score
			// 	}
			// }
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
	if overallCal != 500 {
		return 0
	}
	return overallCap * overallDura * overallFla * overallText
}
