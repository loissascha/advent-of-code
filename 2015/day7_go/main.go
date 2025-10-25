package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Variable struct {
	Name  string
	Value uint8
}

var variables map[string]Variable = map[string]Variable{}

func main() {
	content, err := os.ReadFile("test_input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.SplitSeq(string(content), "\n")
	for line := range lines {
		fmt.Println(line)
		processLine(line)
	}
}

func processLine(line string) {
	if strings.TrimSpace(line) == "" {
		return
	}
	split := strings.Split(line, " -> ")
	action := split[0]
	varname := split[1]

	actionResult := getActionResult(action)

	variables[varname] = Variable{
		Name:  varname,
		Value: actionResult,
	}

	fmt.Println("Variable value: ", varname, actionResult)
}

func getActionResult(input string) uint8 {
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
		return Lshift(vara.Value, uint8(shift))
	} else if strings.Contains(input, "RSHIFT") {
		split := strings.Split(input, " RSHIFT ")
		vara := variables[split[0]]
		shift, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		return Rshift(vara.Value, uint8(shift))
	} else if strings.Contains(input, "NOT") {
		varStr := strings.TrimLeft(input, "NOT ")
		vara := variables[varStr]
		return Not(vara.Value)
	} else {
		v, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		return uint8(v)
	}

	return 0
}

func And(a, b uint8) uint8 {
	return a & b
}

func Or(a, b uint8) uint8 {
	return a | b
}

func Not(a uint8) uint8 {
	return ^a
}

func Lshift(num, shift uint8) uint8 {
	return num << shift
}

func Rshift(num, shift uint8) uint8 {
	return num >> shift
}
