package main

import "fmt"

type ReindeerState int

const (
	StateFly  ReindeerState = 0
	StateRest ReindeerState = 1
)

type Reindeer struct {
	FlySpeed int
	FlyTime  int
	RestTime int
	State    ReindeerState
}

func (r *Reindeer) simulateSeconds(second int) int {
	distance := 0
	currentTime := 0
	for range second {
		switch r.State {
		case StateFly:
			distance += r.FlySpeed
			currentTime++
			if currentTime >= r.FlyTime {
				currentTime = 0
				r.State = StateRest
			}
		case StateRest:
			currentTime++
			if currentTime >= r.RestTime {
				currentTime = 0
				r.State = StateFly
			}
		}
	}
	return distance
}

func main() {

	comet := &Reindeer{
		FlySpeed: 14,
		FlyTime:  10,
		RestTime: 127,
		State:    StateFly,
	}

	comatDistance := comet.simulateSeconds(1000)
	fmt.Println("comet distance:", comatDistance)

}
