package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type City struct {
	Name string
}

type Route struct {
	Start    *City
	Target   *City
	Distance int
}

var cities = []*City{}
var routes = []*Route{}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.SplitSeq(string(content), "\n")
	for line := range lines {
		if line == "" {
			continue
		}
		fmt.Println(line)
		city1Name := getCity1(line)
		city2Name := getCity2(line)
		distance := getDistance(line)

		city1 := getCity(city1Name)
		if city1 == nil {
			city1 = addCity(city1Name)
		}

		city2 := getCity(city2Name)
		if city2 == nil {
			city2 = addCity(city2Name)
		}

		addRoute(city1, city2, distance)
	}

	shortestDistance := -1
	longestDistance := -1
	for _, city := range cities {

		node := buildNodesFromStartCity(city, nil, 0)

		printNodeTree(node, 1, len(cities))

		fastest := findFastestRoute(node, 1, len(cities))
		longest := findLongestRoute(node, 1, len(cities))
		if shortestDistance == -1 || fastest < shortestDistance {
			shortestDistance = fastest
		}
		if longestDistance == -1 || longestDistance < longest {
			longestDistance = longest
		}

	}
	fmt.Println("shortest distance:", shortestDistance)
	fmt.Println("longest distance:", longestDistance)
}

func getRoutesForCity(city *City) []*Route {
	res := []*Route{}
	for _, route := range routes {
		if route.Start == city || route.Target == city {
			res = append(res, route)
		}
	}
	return res
}

func addRoute(startCity *City, endCity *City, distance int) {
	r := &Route{
		Start:    startCity,
		Target:   endCity,
		Distance: distance,
	}
	routes = append(routes, r)
}

func addCity(name string) *City {
	c := &City{
		Name: name,
	}
	cities = append(cities, c)
	return c
}

func getCity(name string) *City {
	for _, c := range cities {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func getDistance(input string) int {
	split := strings.Split(input, "=")
	distStr := strings.TrimSpace(split[1])
	dist, err := strconv.Atoi(distStr)
	if err != nil {
		panic(err)
	}
	return dist
}

func getCity1(input string) string {
	split := strings.Split(input, "to")
	cityStr := strings.TrimSpace(split[0])
	return cityStr
}

func getCity2(input string) string {
	split := strings.Split(input, "to")
	cityStr := strings.TrimSpace(strings.Split(strings.TrimSpace(split[1]), "=")[0])
	return cityStr
}
