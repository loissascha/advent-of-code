package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/loissascha/go-logger/logger"
)

type Variable struct {
	Name      string
	Value     uint16
	Operation string
	Done      bool
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

	for hasUndone() {
		retry()
	}

	fmt.Println(variables)
	fmt.Println("var a value:", variables["a"].Value)
	fmt.Println("var lx value:", variables["lx"].Value)
}

func hasUndone() bool {
	for _, v := range variables {
		if !v.Done {
			return true
		}
	}
	return false
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
		Name:      varname,
		Value:     actionResult,
		Operation: line,
		Done:      found,
	}

	if found {
		fmt.Println("Variable value: ", varname, actionResult)
	}
}

func retry() {
	for _, v := range variables {
		if !v.Done {
			processLine(v.Operation)
		}
	}
}

func getValueForPosition(input string) (uint16, bool) {
	n, err := strconv.Atoi(input)
	if err != nil {
		v, found := variables[input]
		if !found {
			return 0, false
		}
		return v.Value, v.Done
	}
	return uint16(n), true
}

func getActionResult(input string) (uint16, bool) {
	if strings.Contains(input, "AND") {
		split := strings.Split(input, " AND ")

		posa, found := getValueForPosition(split[0])
		if !found {
			logger.Error(nil, "AND: Can't find var: {VarName}", split[0])
			return 0, false
		}

		posb, found := getValueForPosition(split[1])
		if !found {
			logger.Error(nil, "AND: Can't find var: {VarName}", split[1])
			return 0, false
		}

		res := And(posa, posb)
		fmt.Println("and res:", res)
		return res, true
	} else if strings.Contains(input, "OR") {
		split := strings.Split(input, " OR ")

		posa, found := getValueForPosition(split[0])
		if !found {
			logger.Error(nil, "OR: Can't find var: {VarName}", split[0])
			return 0, false
		}

		posb, found := getValueForPosition(split[1])
		if !found {
			logger.Error(nil, "OR: Can't find var: {VarName}", split[1])
			return 0, false
		}

		res := Or(posa, posb)
		fmt.Println("or res:", res)
		return res, true
	} else if strings.Contains(input, "LSHIFT") {
		split := strings.Split(input, " LSHIFT ")
		posa, found := getValueForPosition(split[0])
		if !found {
			logger.Error(nil, "LSHIFT: Can't find var: {VarName}", split[0])
			return 0, false
		}
		posb, found := getValueForPosition(split[1])
		if !found {
			logger.Error(nil, "LSHIFT: Can't find var: {VarName}", split[1])
			return 0, false
		}
		res := Lshift(posa, posb)
		fmt.Println("lshfit res:", res)
		return res, true
	} else if strings.Contains(input, "RSHIFT") {
		split := strings.Split(input, " RSHIFT ")
		posa, found := getValueForPosition(split[0])
		if !found {
			logger.Error(nil, "RSHIFT: Can't find var: {VarName}", split[0])
			return 0, false
		}
		posb, found := getValueForPosition(split[1])
		if !found {
			logger.Error(nil, "RSHIFT: Can't find var: {VarName}", split[1])
			return 0, false
		}
		res := Rshift(posa, posb)
		fmt.Println("rshfit res:", res)
		return res, true
	} else if strings.Contains(input, "NOT") {
		varStr := strings.TrimLeft(input, "NOT ")
		posa, found := getValueForPosition(varStr)
		if !found {
			logger.Error(nil, "NOT: Can't find var: {VarName}", varStr)
			return 0, false
		}
		res := Not(posa)
		fmt.Println("not res:", res)
		return res, true
	} else {
		v, err := strconv.Atoi(input)
		if err != nil {
			// its a variable name
			v, found := variables[input]
			if !found {
				logger.Error(nil, "Parse: Can't find var: {VarName}", input)
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
