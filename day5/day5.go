package day5

import (
	"bufio"
	"fmt"
	"os"
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

var rules = []Rule{}
var updates = []Update{}

func Day5() {
	// read rules
	// read updates
	readFile("day5.test")

	fmt.Println("found", len(rules), "rules")
	fmt.Println("found", len(updates), "updates")

	for _, u := range updates {
		u.testUpdate()
	}

	// go number by number
	// find rules that matter for this number
	// check if rules are fullfilled for this number (by checking all the elements before and next of it)

	// if correctly ordered -> get middle number and add together
}

func (u *Update) testUpdate() {
	fmt.Println("testing update with:", u.elements)
	ruleFailed := false
	for i, element := range u.elements {
		// fmt.Println("finding rules for number", element)
		rs := findRulesForNumber(element)
		// fmt.Println(rs)
		// go through each rule and test if it fullfills it?
		for _, r := range rs {
			// test if elements from rule are even included
			elementsInRule := 0
			for _, v := range u.elements {
				if r.left == v || r.right == v {
					elementsInRule++
				}
			}
			if elementsInRule < 2 {
				// fmt.Println("skipped rule because not all elements in rule!")
				continue
			}
			ff := u.ruleFulfilled(i, r)
			if ff {
				// fmt.Println("rule fulfilled")
				// fmt.Println(r)
			} else {
				// fmt.Println("rule not fulfilled")
				// fmt.Println(r)
				ruleFailed = true
			}
		}
	}
	if !ruleFailed {
		fmt.Println("rule success!")
	}
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
