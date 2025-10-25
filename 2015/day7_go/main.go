package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Variable struct {
	Name  string
	Value int
}

var variables map[string]Variable = map[string]Variable{}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.SplitSeq(string(content), "\n")
	for line := range lines {
		fmt.Println(line)
		processLine(line)
	}

	fmt.Println("var a value:", variables["a"].Value)
}

func processLine(line string) {
	if strings.TrimSpace(line) == "" {
		return
	}
	split := strings.Split(line, " -> ")
	action := split[0]
	varname := split[1]

	actionResult := getActionResult(action)

	if actionResult < 0 {
		actionResult = 65535 + actionResult + 1
	}

	variables[varname] = Variable{
		Name:  varname,
		Value: actionResult,
	}

	fmt.Println("Variable value: ", varname, actionResult)
}

func getActionResult(input string) int {
	if strings.Contains(input, "AND") {
		split := strings.Split(input, " AND ")
		vara := variables[split[0]]
		varb := variables[split[1]]
		return And(vara.Value, varb.Value)
	} else if strings.Contains(input, "OR") {
		split := strings.Split(input, " OR ")
		vara := variables[split[0]]
		varb := variables[split[1]]
		return Or(vara.Value, varb.Value)
	} else if strings.Contains(input, "LSHIFT") {
		split := strings.Split(input, " LSHIFT ")
		vara := variables[split[0]]
		shift, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		return Lshift(vara.Value, int(shift))
	} else if strings.Contains(input, "RSHIFT") {
		split := strings.Split(input, " RSHIFT ")
		vara := variables[split[0]]
		shift, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		return Rshift(vara.Value, int(shift))
	} else if strings.Contains(input, "NOT") {
		varStr := strings.TrimLeft(input, "NOT ")
		vara := variables[varStr]
		return Not(vara.Value)
	} else {
		v, err := strconv.Atoi(input)
		if err != nil {
			// its a variable name
			return variables[input].Value
		}
		return int(v)
	}

	return 0
}

func And(a, b int) int {
	return a & b
}

func Or(a, b int) int {
	return a | b
}

func Not(a int) int {
	return ^a
}

func Lshift(num, shift int) int {
	return num << shift
}

func Rshift(num, shift int) int {
	return num >> shift
}
