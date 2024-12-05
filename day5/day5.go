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
var overallSum = 0

func Day5() {
	readFile("day5.test")

	fmt.Println("found", len(rules), "rules")
	fmt.Println("found", len(updates), "updates")

	for _, u := range updates {
		ok, failedRules := u.testUpdate()
		// reorder based on rules!
		if !ok {
			fmt.Println("failed by rules:")
			fmt.Println(failedRules)
		}

	}

	fmt.Println("Overall Sum:", overallSum)
}

func (u *Update) testUpdate() (bool, []Rule) {
	fmt.Println("testing update with:", u.elements)
	ruleFailed := false
	failedRules := []Rule{}
	for i, element := range u.elements {
		rs := findRulesForNumber(element)
		for _, r := range rs {
			if !u.ruleValidForUpdate(r) {
				continue
			}
			ff := u.ruleFulfilled(i, r)
			if !ff {
				ruleFailed = true
				failedRules = append(failedRules, r)
			}
		}
	}
	if !ruleFailed {
		middle := getMiddleElement(u.elements)
		overallSum += middle
		return true, []Rule{}
	}
	return false, failedRules
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
