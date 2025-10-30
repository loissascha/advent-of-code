package main

import "fmt"

type Container struct {
	size int
}

func main() {

	containers := []*Container{
		{size: 33},
		{size: 14},
		{size: 18},
		{size: 20},
		{size: 45},
		{size: 35},
		{size: 16},
		{size: 35},
		{size: 1},
		{size: 13},
		{size: 18},
		{size: 13},
		{size: 50},
		{size: 44},
		{size: 48},
		{size: 6},
		{size: 24},
		{size: 41},
		{size: 30},
		{size: 42},
	}

	combinations := countCombinations(containers, 150)
	fmt.Println("there are", combinations, "combinations")

}

func countCombinations(containers []*Container, target int) int {
	var dfs func(start, remaining int) int
	dfs = func(start, remaining int) int {
		if remaining == 0 {
			return 1
		}
		if remaining < 0 {
			return 0
		}

		count := 0
		for i := start; i < len(containers); i++ {
			count += dfs(i+1, remaining-containers[i].size)
		}
		return count
	}

	return dfs(0, target)
}
