package main

import "fmt"

type Container struct {
	size int
}

func main() {

	containers := []*Container{
		{size: 20},
		{size: 15},
		{size: 10},
		{size: 5},
		{size: 5},
	}

	combinations := countCombinations(containers, 25)
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
