package main

import "fmt"

type Node struct {
	City     *City
	Nodes    []*Node
	Parent   *Node
	Distance int
}

func printNodeTree(node *Node, count int, targetCount int) {
	if count == targetCount {
		fmt.Printf("%s -> %d\n", node.City.Name, node.Distance)
	}
	for _, n := range node.Nodes {
		printNodeTree(n, count+1, targetCount)
		// fmt.Printf("%s, ", n.City.Name)
	}
	fmt.Printf("\n")
}

func buildNodesFromStartCity(city *City, startNode *Node, distance int) *Node {
	node := &Node{
		City:     city,
		Nodes:    []*Node{},
		Parent:   startNode,
		Distance: distance,
	}

	rs := getRoutesForCity(city)
	for _, r := range rs {
		c := r.Start
		if r.Start == city {
			c = r.Target
		}

		if nodeHasCity(c, node) {
			continue
		}

		node.Nodes = append(node.Nodes, buildNodesFromStartCity(c, node, distance+r.Distance))
	}

	return node
}

func nodeHasCity(city *City, node *Node) bool {
	if node.City == city {
		return true
	}
	if node.Parent != nil {
		if node.Parent.City == city {
			return true
		}
		return nodeHasCity(city, node.Parent)
	}
	return false
}
