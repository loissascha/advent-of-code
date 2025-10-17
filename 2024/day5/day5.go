package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/loissascha/go-assert/assert"
)

type Rule struct {
	left  int
	right int
}

type Update struct {
	elements []int
}

type SliceRule struct {
	number    int
	needLeft  []int
	needRight []int
}

var rules = []Rule{}
var updates = []Update{}
var overallSum = 0
var correctedSum = 0

func Day5() {
	readFile("day5.input")

	fmt.Println("found", len(rules), "rules")
	fmt.Println("found", len(updates), "updates")

	for _, u := range updates {
		ok := u.testUpdate()

		// Part2 reorder based on rules!
		if !ok {
			u.correctlyOrder()
		}
	}

	fmt.Println("Part 1 Sum:", overallSum)
	fmt.Println("Part 2 Sum:", correctedSum)
}

func (u *Update) correctlyOrder() {
	fmt.Println("trying to reorder update with elements:", u.elements)

	// get all relevant rules
	rrules := []Rule{}
	for _, v := range u.elements {
		rs := findRulesForNumber(v)
		for _, r := range rs {
			if u.ruleValidForUpdate(r) {
				if !slices.Contains(rrules, r) {
					rrules = append(rrules, r)
				}
			}
		}
	}

	fmt.Println("this are all relevant rules for this update:", rrules)

	// test to make sure all elements are included
	for _, v := range u.elements {
		exists := false
		for _, r := range rrules {
			if r.left == v || r.right == v {
				exists = true
			}
		}
		assert.True(exists, fmt.Sprintf("No rule exists for element: %v", v))
	}

	createSliceForRules(u.elements, rrules)
}

func createSliceForRules(eles []int, rs []Rule) {
	sliceRules := []SliceRule{}
	for _, e := range eles {
		lefts := []int{}
		rights := []int{}
		for _, r := range rs {
			if r.left == e {
				rights = append(rights, r.right)
			}
			if r.right == e {
				lefts = append(lefts, r.left)
			}
		}
		sliceRules = append(sliceRules, SliceRule{
			number:    e,
			needLeft:  lefts,
			needRight: rights,
		})
	}
	fmt.Println(sliceRules)

	addedRules := 0
	maxRules := len(sliceRules)
	res := []int{}

	for addedRules < maxRules {
		for _, v := range sliceRules {
			if slices.Contains(res, v.number) {
				continue
			}
			leftsFound := true
			for _, l := range v.needLeft {
				if !slices.Contains(res, l) {
					leftsFound = false
				}
			}
			if leftsFound {
				fmt.Println("lefts found for", v.number)
				res = append(res, v.number)
				addedRules++
			}
		}
	}
	fmt.Println("finished slice:", res)

	middle := getMiddleElement(res)
	correctedSum += middle
}

func (u *Update) testUpdate() bool  {
	fmt.Println("testing update with:", u.elements)
	ruleFailed := false
	for i, element := range u.elements {
		rs := findRulesForNumber(element)
		for _, r := range rs {
			if !u.ruleValidForUpdate(r) {
				continue
			}
			ff := u.ruleFulfilled(i, r)
			if !ff {
				ruleFailed = true
			}
		}
	}
	if !ruleFailed {
		middle := getMiddleElement(u.elements)
		overallSum += middle
		fmt.Println("overallSum ++", middle)
		return true
	}
	return false
}

func (u *Update) ruleValidForUpdate(rule Rule) bool {
	hasLeft := false
	hasRight := false
	for _, v := range u.elements {
		if rule.left == v {
			hasLeft = true
		}
		if rule.right == v {
			hasRight = true
		}
	}
	return hasLeft == true && hasRight == true
}

func getMiddleElement[T any](a []T) T {
	i := len(a) / 2
	return a[i]
}

func (u *Update) ruleFulfilled(index int, rule Rule) bool {
	num := u.elements[index]
	// number on left side of rule -> look at all the elements after
	if rule.left == num {
		for i, v := range u.elements {
			if i <= index {
				continue
			}
			if v == rule.right {
				return true
			}
		}
	}
	// number on right side of rule -> look at all the elements before
	if rule.right == num {
		for i, v := range u.elements {
			if i >= index {
				continue
			}
			if v == rule.left {
				return true
			}
		}
	}
	return false
}

func findRulesForNumber(number int) []Rule {
	res := []Rule{}
	for _, v := range rules {
		if v.left == number || v.right == number {
			res = append(res, v)
		}
	}
	return res
}

func readFile(filepath string) {
	file, err := os.Open(filepath)
	assert.Nil(err, "Can't open file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// rule
		if strings.Contains(line, "|") {
			sp := strings.Split(line, "|")
			assert.True(len(sp) == 2, "Rule length is wrong!")
			r0, err := strconv.Atoi(sp[0])
			assert.Nil(err, "StrConv sp0 failed")
			r1, err := strconv.Atoi(sp[1])
			assert.Nil(err, "StrConv sp1 failed")
			rules = append(rules, Rule{
				left:  r0,
				right: r1,
			})
		}
		// updates
		if strings.Contains(line, ",") {
			sp := strings.Split(line, ",")
			e := []int{}
			for _, v := range sp {
				r, err := strconv.Atoi(v)
				assert.Nil(err, "Strconv r failed")
				e = append(e, r)
			}
			updates = append(updates, Update{
				elements: e,
			})
		}
	}
}
