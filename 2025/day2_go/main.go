package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func main() {
	input, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}

	ranges := []Range{}

	rangesSplit := strings.SplitSeq(input, ",")
	for r := range rangesSplit {
		r = strings.TrimRight(r, "\n")
		nums := strings.Split(r, "-")
		if len(nums) != 2 {
			panic("nums len wrong for r: " + r)
		}
		start, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, Range{
			Start: start,
			End:   end,
		})
	}

	fmt.Println("ranges:", ranges)

	invalidSum := 0
	for _, r := range ranges {
		for n := r.Start; n <= r.End; n++ {
			nStr := strconv.Itoa(n)
			if len(nStr)%2 == 0 {
				splitLen := len(nStr) / 2
				firstHalft := nStr[:splitLen]
				secondHalf := nStr[splitLen:]
				if firstHalft == secondHalf {
					invalidSum += n
				}
			}
		}
	}
	fmt.Println("invalid sum:", invalidSum)
}

func readFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
