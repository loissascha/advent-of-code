package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ReindeerState int

const (
	StateFly  ReindeerState = 0
	StateRest ReindeerState = 1
)

type Reindeer struct {
	Name     string
	FlySpeed int
	FlyTime  int
	RestTime int
	State    ReindeerState
	Points   int
}

func (r *Reindeer) simulateSeconds(second int) int {
	fmt.Println("simulate seconds", second)
	distance := 0
	currentTime := 0
	r.State = StateFly
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
		fmt.Println("distance", distance)
	}
	return distance
}

func main() {

	// comet := &Reindeer{
	// 	Name:     "Comet",
	// 	FlySpeed: 14,
	// 	FlyTime:  10,
	// 	RestTime: 127,
	// 	State:    StateFly,
	// }
	//
	// dancer := &Reindeer{
	// 	Name:     "Dancer",
	// 	FlySpeed: 16,
	// 	FlyTime:  11,
	// 	RestTime: 162,
	// 	State:    StateFly,
	// }

	// comatDistance := comet.simulateSeconds(1000)
	// dancerDistance := dancer.simulateSeconds(1000)
	// fmt.Println("comet distance:", comatDistance)
	// fmt.Println("dancer distance:", dancerDistance)

	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	reindeers := []*Reindeer{}

	lines := strings.SplitSeq(string(content), "\n")
	for line := range lines {
		if line == "" {
			continue
		}
		r := lineToReindeer(line)
		reindeers = append(reindeers, r)
	}

	// fmt.Println("there are", len(reindeers), "reindeers")
	//
	// simulationTime := 2503
	// longestDistance := 0
	// for _, r := range reindeers {
	// 	d := r.simulateSeconds(simulationTime)
	// 	if d > longestDistance {
	// 		longestDistance = d
	// 	}
	// }
	//
	// fmt.Println("longest distance:", longestDistance)

	// tr := []*Reindeer{comet, dancer}
	// simulateOverTime(tr, 1000)

	simulateOverTime(reindeers, 2503)
}

func simulateOverTime(reindeers []*Reindeer, simulationTime int) {
	for n := range simulationTime {
		var winningReindeer *Reindeer
		winningDistance := 0
		for _, r := range reindeers {
			d := r.simulateSeconds(n + 1)
			if d > winningDistance {
				winningDistance = d
				winningReindeer = r
			}
		}
		if winningReindeer != nil {
			winningReindeer.Points++
		}
	}

	for _, r := range reindeers {
		fmt.Println(r.Name, "with points:", r.Points)
	}
}

func lineToReindeer(line string) *Reindeer {
	split := strings.SplitN(line, "can fly", 2)
	name := strings.TrimSpace(split[0])
	line = split[1]

	split = strings.SplitN(line, "km/s for", 2)
	speedStr := strings.TrimSpace(split[0])
	speed, err := strconv.Atoi(speedStr)
	if err != nil {
		panic(err)
	}
	line = split[1]

	split = strings.SplitN(line, "seconds", 2)
	timeStr := strings.TrimSpace(split[0])
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		panic(err)
	}
	line = split[1]

	restTimeStr := strings.TrimLeft(line, ", but then must rest for ")
	restTimeStr = strings.TrimRight(restTimeStr, " seconds.")
	restTime, err := strconv.Atoi(restTimeStr)
	if err != nil {
		panic(err)
	}

	return &Reindeer{
		Name:     name,
		FlySpeed: speed,
		FlyTime:  time,
		RestTime: restTime,
		State:    StateFly,
	}
}
