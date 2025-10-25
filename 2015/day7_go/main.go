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
	content, err := os.ReadFile("test_input.txt")
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

	// if actionResult < 0 {
	// 	actionResult = 65535 + actionResult + 1
	// }

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
		res := And(uint8(vara.Value), uint8(varb.Value))
		fmt.Println("and res:", res)
		return int(res)
	} else if strings.Contains(input, "OR") {
		split := strings.Split(input, " OR ")
		vara := variables[split[0]]
		varb := variables[split[1]]
		res := Or(uint8(vara.Value), uint8(varb.Value))
		fmt.Println("or res:", res)
		return int(res)
	} else if strings.Contains(input, "LSHIFT") {
		split := strings.Split(input, " LSHIFT ")
		vara := variables[split[0]]
		shift, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		res := Lshift(uint(vara.Value), uint(shift))
		fmt.Println("lshfit res:", res)
		return int(res)
	} else if strings.Contains(input, "RSHIFT") {
		split := strings.Split(input, " RSHIFT ")
		vara := variables[split[0]]
		shift, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		res := Rshift(uint(vara.Value), uint(shift))
		fmt.Println("rshfit res:", res)
		return int(res)
	} else if strings.Contains(input, "NOT") {
		varStr := strings.TrimLeft(input, "NOT ")
		vara := variables[varStr]
		res := Not(uint8(vara.Value))
		fmt.Println("not res:", res)
		return int(res)
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

func And(a, b uint8) uint8 {
	return a & b
}

func Or(a, b uint8) uint8 {
	return a | b
}

func Not(a uint8) uint8 {
	return ^a
}

func Lshift(num, shift uint) uint {
	return num << shift
}

func Rshift(num, shift uint) uint {
	return num >> shift
}
