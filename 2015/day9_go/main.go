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

type Path struct {
	Start    *City
	Distance int
	Parent   *Path
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
	// var shortestPath *Path
	for _, city := range cities {

		node := buildNodesFromStartCity(city, nil, 0)
		// fmt.Println("node for startcity:", city.Name, node)
		// printNodeTree(node, 1, len(cities))

		fastest := findFastestRoute(node, 1, len(cities))
		// fmt.Println("fastest:", fastest)
		if shortestDistance == -1 || fastest < shortestDistance {
			shortestDistance = fastest
		}

		// fmt.Println("start point:", city.Name)
		// paths := tryWithStartPoint(city, 0, nil)
		// for _, path := range paths {
		// 	count := countPathParents(path)
		// 	if count < len(cities) {
		// 		continue
		// 	}
		// 	fmt.Println("count:", count)
		// 	if path.Distance < shortestDistance || shortestDistance == -1 {
		// 		shortestDistance = path.Distance
		// 		shortestPath = path
		// 	}
		// 	// fmt.Println("to:", path.Start.Name, "Distance:", path.Distance)
		// 	if path.Parent != nil {
		// 		// fmt.Println("parent: ", path.Parent.Start.Name)
		// 	}
		// }
	}
	fmt.Println("shortest distance:", shortestDistance)
	// printPath(shortestPath)
}

func printPath(path *Path) {
	if path == nil {
		return
	}
	fmt.Printf("%s", path.Start.Name)
	for path.Parent != nil {
		fmt.Printf(" -> %s", path.Parent.Start.Name)
		path = path.Parent
	}
	fmt.Printf("\n")
}

func countPathParents(path *Path) int {
	count := 1
	for path.Parent != nil {
		count++
		path = path.Parent
	}
	return count
}

func tryWithStartPoint(city *City, distance int, parent *Path) []*Path {
	paths := []*Path{}
	if parent != nil {
		paths = append(paths, parent)
	}

	rs := getRoutesForCity(city)
	for _, r := range rs {
		// start point
		path := &Path{}
		path.Start = city
		path.Distance = distance
		path.Parent = parent
		// fmt.Println("start pos:", city.Name)

		// for each target -> create another path that has the parent and the distance
		// fmt.Println(r.Start, r.Target, r.Distance)

		c := r.Start
		if city == r.Start {
			c = r.Target
		}
		// fmt.Println("this route:", c.Name)

		// fmt.Println("checking if city", c.Name, "is already in path")
		if parentHasCityInPath(path, c) {
			// fmt.Println("parent has city")
			continue
		}
		// fmt.Println("parent does not have city")

		paths = append(paths, tryWithStartPoint(c, path.Distance+r.Distance, path)...)
	}

	return paths
}

func parentHasCityInPath(path *Path, city *City) bool {
	for {
		if path.Parent != nil && path.Parent.Start == city {
			return true
		}
		if path.Parent == nil {
			break
		}
		path = path.Parent
	}
	return false
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
