package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

	var data map[string]any

	err = json.Unmarshal(content, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

	sum := getSum(data)
	fmt.Println("sum:", sum)
}

func workValue(d any) int {
	sum := 0
	n, ok := d.(int)
	if ok {
		fmt.Println("found int", n)
		sum += n
		return sum
	}
	ss, ok := d.(string)
	if ok {
		i, err := strconv.Atoi(strings.TrimSpace(ss))
		if err == nil {
			fmt.Println("found int", i)
			sum += i
			return sum
		}
		// fmt.Println("ERROR: found unclear string")
		return sum
	}
	m, ok := d.(map[string]any)
	if ok {
		// fmt.Println("found map")
		sum += getSum(m)
		return sum
	}
	s, ok := d.([]any)
	if ok {
		// fmt.Println("found slice")
		sum += getSliceSum(s)
		return sum
	}
	f, ok := d.(float64)
	if ok {
		sum += int(f)
		return sum
	}
	fmt.Println("unknown value:", d)
	return sum
}

func getSum(data map[string]any) int {
	sum := 0

	for _, d := range data {
		// fmt.Println("KEY:", m)
		// fmt.Println("VALUE:", d)
		sum += workValue(d)
	}

	return sum
}

func getSliceSum(data []any) int {
	sum := 0
	for _, d := range data {
		// fmt.Println("VALUE:", d)
		sum += workValue(d)
	}
	return sum
}
