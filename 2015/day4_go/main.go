package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func main() {

	input := "ckczppom"
	randoNum := getNumberStr(input)
	fmt.Println("solution is:", randoNum)
}

func getNumberStr(input string) int {
	i := -1
	randoNum := 1
	for i != 0 {
		hash := md5.New()
		inp := fmt.Sprintf("%s%d", input, randoNum)
		fmt.Println("testing:", inp)
		io.WriteString(hash, inp)
		md5str := fmt.Sprintf("%x", hash.Sum(nil))
		i = strings.Index(md5str, "000000")
		if i != 0 {
			randoNum++
		}
	}
	return randoNum
}
