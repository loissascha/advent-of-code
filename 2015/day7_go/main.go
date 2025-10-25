package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Variable struct {
	Name       string
	Value      uint16
	Operations []string
	Done       bool
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

	fmt.Println(variables)
	fmt.Println("var a value:", variables["a"].Value)
	fmt.Println("var lx value:", variables["lx"].Value)
}

// TODO: a gate provides no signal until all of it's inputs have a signal

func processLine(line string) {
	if strings.TrimSpace(line) == "" {
		return
	}
	split := strings.Split(line, " -> ")
	action := split[0]
	varname := split[1]

	actionResult, found := getActionResult(action)

	variables[varname] = Variable{
		Name:       varname,
		Value:      actionResult,
		Operations: []string{line},
		Done:       found,
	}

	fmt.Println("Variable value: ", varname, actionResult)
}

func getActionResult(input string) (uint16, bool) {
	if strings.Contains(input, "AND") {
		split := strings.Split(input, " AND ")
		vara, found := variables[split[0]]
		if !found {
			return 0, false
		}
		varb, found := variables[split[1]]
		if !found {
			return 0, false
		}
		res := And(uint16(vara.Value), uint16(varb.Value))
		fmt.Println("and res:", res)
		return res, true
	} else if strings.Contains(input, "OR") {
		split := strings.Split(input, " OR ")
		vara, found := variables[split[0]]
		if !found {
			return 0, false
		}
		varb, found := variables[split[1]]
		if !found {
			return 0, false
		}
		res := Or(uint16(vara.Value), uint16(varb.Value))
		fmt.Println("or res:", res)
		return res, true
	} else if strings.Contains(input, "LSHIFT") {
		split := strings.Split(input, " LSHIFT ")
		vara, found := variables[split[0]]
		if !found {
			return 0, false
		}
		shift, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		res := Lshift(uint16(vara.Value), uint16(shift))
		fmt.Println("lshfit res:", res)
		return res, true
	} else if strings.Contains(input, "RSHIFT") {
		split := strings.Split(input, " RSHIFT ")
		vara, found := variables[split[0]]
		if !found {
			return 0, false
		}
		shift, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		res := Rshift(uint16(vara.Value), uint16(shift))
		fmt.Println("rshfit res:", res)
		return res, true
	} else if strings.Contains(input, "NOT") {
		varStr := strings.TrimLeft(input, "NOT ")
		vara, found := variables[varStr]
		if !found {
			return 0, false
		}
		res := Not(uint16(vara.Value))
		fmt.Println("not res:", res)
		return res, true
	} else {
		v, err := strconv.Atoi(input)
		if err != nil {
			// its a variable name
			v, found := variables[input]
			if !found {
				return 0, false
			}
			return v.Value, true
		}
		return uint16(v), true
	}
	panic("operation wrong")
}

func And(a, b uint16) uint16 {
	return a & b
}

func Or(a, b uint16) uint16 {
	return a | b
}

func Not(a uint16) uint16 {
	return ^a
}

func Lshift(num, shift uint16) uint16 {
	return num << shift
}

func Rshift(num, shift uint16) uint16 {
	return num >> shift
}
