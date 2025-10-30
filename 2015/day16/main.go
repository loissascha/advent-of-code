package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AuntSue struct {
	id         int
	attributes map[string]int
}

func (a *AuntSue) matchesOriginal(original *AuntSue) bool {
	matches := true
	for k, v := range a.attributes {
		ov, found := original.attributes[k]
		if found {
			if k == "cats" || k == "trees" {
				if v <= ov {
					matches = false
				}
			} else if k == "pomeranians" || k == "goldfish" {
				if v >= ov {
					matches = false
				}
			} else if ov != v {
				matches = false
			}
		}
	}
	return matches
}

func main() {
	original := &AuntSue{
		id: 0,
		attributes: map[string]int{
			"children":    3,
			"cats":        7,
			"samoyeds":    2,
			"pomeranians": 3,
			"akitas":      0,
			"vizslas":     0,
			"goldfish":    5,
			"trees":       3,
			"cars":        2,
			"perfumes":    1,
		},
	}

	corr := &AuntSue{
		id: 0,
		attributes: map[string]int{
			"children": 3,
			"cats":     7,
			"goldfish": 5,
			"perfumes": 1,
		},
	}

	wrong := &AuntSue{
		id: 0,
		attributes: map[string]int{
			"children": 3,
			"goldfish": 4,
			"trees":    3,
			"cars":     1,
			"perfumes": 1,
		},
	}

	corrTrue := corr.matchesOriginal(original)
	wrongTrue := wrong.matchesOriginal(original)

	fmt.Println(corrTrue, wrongTrue)

	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.SplitSeq(string(content), "\n")

	for line := range lines {
		if line == "" {
			continue
		}
		sue := newSueFromLine(line)
		if sue.matchesOriginal(original) {
			fmt.Println("found matching sue:", sue.id)
		}
	}
}

func newSueFromLine(line string) *AuntSue {
	split := strings.SplitN(line, ":", 2)
	idStr := strings.TrimLeft(split[0], "Sue ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	res := &AuntSue{
		id:         id,
		attributes: map[string]int{},
	}

	attributesStr := split[1]
	attributesSplit := strings.SplitSeq(attributesStr, ",")
	for v := range attributesSplit {
		split := strings.Split(v, ":")
		attrName := strings.TrimSpace(split[0])
		attrValueStr := strings.TrimSpace(split[1])
		attrValue, err := strconv.Atoi(attrValueStr)
		if err != nil {
			panic(err)
		}
		res.attributes[attrName] = attrValue
	}

	return res
}
