package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	allTotal := 0
	allClean := 0

	lines := strings.SplitSeq(string(content), "\n")
	for line := range lines {
		if line == "" {
			continue
		}

		if strings.Index(line, "\"") != 0 {
			fmt.Println(line)
			panic("line not starting with \"")
		}

		if strings.LastIndex(line, "\"") != len(line)-1 {
			fmt.Println(line)
			panic("line not ending with \"")
		}

		totalLen := len(line)

		line = strings.TrimPrefix(line, "\"")
		line = strings.TrimSuffix(line, "\"")

		// fmt.Println("cleanline:", line)
		line = strings.ReplaceAll(line, "\\\"", "\"")
		line = strings.ReplaceAll(line, "\\\\", "\\")

		for strings.Contains(line, "\\x") {
			i := strings.Index(line, "\\x")
			hexa := line[i : i+4]
			h, err := hex.DecodeString(hexa)
			if err != nil {
				panic(err)
			}
			fmt.Println("found hexa:", hexa, h)
			lineFirst := line[0:i]
			lineSecond := line[i+4:]
			// fmt.Println("line before:", lineFirst, lineSecond)
			line = lineFirst + "'" + lineSecond
		}

		cleanLen := len(line)

		allTotal += totalLen
		allClean += cleanLen

		// fmt.Println("workedLine: ", line)
		// fmt.Println("total len:", totalLen, "clean len:", cleanLen)
	}

	finalCount := allTotal - allClean
	fmt.Println("solution:", finalCount)
}
