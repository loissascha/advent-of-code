package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var letterMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	inputFile := "input.txt"
	lines, err := readLines(inputFile)
	if err != nil {
		panic(err)
	}
	sum := 0
	for line := range lines {
		res := getLineValue([]byte(line))
		fmt.Println(res)
		sum += res
	}
	fmt.Println("Sum:", sum)
}

func replaceHighestNumber(line string) string {
	foundHigh := false
	highestIdx := -1
	highestKey := ""
	highestValue := ""
	for k, v := range letterMap {
		idx := strings.Index(line, k)
		if idx > -1 && idx > highestIdx {
			foundHigh = true
			highestIdx = idx
			highestKey = k
			highestValue = v
		}
	}
	if foundHigh {
		line = line[:highestIdx] + highestValue + line[highestIdx+len(highestKey):]
	}
	return line
}

func replaceLowestNumber(line string) string {
	foundLow := false
	lowestIdx := len(line) + 1
	lowestKey := ""
	lowestValue := ""
	for k, v := range letterMap {
		idx := strings.Index(line, k)
		if idx > -1 && idx < lowestIdx {
			lowestIdx = idx
			lowestKey = k
			lowestValue = v
			foundLow = true
		}
	}
	if foundLow {
		line = strings.Replace(line, lowestKey, lowestValue, 1)
	}
	return line
}

func getLineValue(line []byte) int {
	// fmt.Println(string(line))
	resStr := ""

	lineStr := string(line)
	fmt.Println("line:", lineStr)
	lineStr = replaceLowestNumber(lineStr)
	lineStr = replaceHighestNumber(lineStr)
	fmt.Println("update line:", lineStr)
	line = []byte(lineStr)

	for _, ch := range line {
		if ch >= byte('0') && ch <= byte('9') {
			resStr += string(ch)
		}
	}

	if len(resStr) == 1 {
		resStr = resStr + resStr
	}

	if len(resStr) > 2 {
		resStr = resStr[:1] + resStr[len(resStr)-1:]
	}

	// fmt.Println("resStr", resStr)

	res, err := strconv.Atoi(resStr)
	if err != nil {
		panic(err)
	}
	return res
}

func readLines(inputFile string) (chan string, error) {
	out := make(chan string, 1)
	file, err := os.Open(inputFile)
	if err != nil {
		return out, err
	}

	go func() {
		defer file.Close()

		line := ""
		for {
			var data = make([]byte, 1)
			read, err := file.Read(data)
			if err != nil {
				break
			}
			chunk := data[:read]
			split := bytes.SplitN(chunk, []byte("\n"), 2)
			if len(split) == 2 {
				line += string(split[0])
				out <- line
				line = ""
				line += string(split[1])
				continue
			}
			line += string(chunk)
		}

		if line != "" {
			out <- line
		}
		close(out)
	}()

	return out, nil

}
