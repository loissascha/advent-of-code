package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/loissascha/go-assert/assert"
)

type Game struct {
	baX int
	baY int
	bbX int
	bbY int
	pX  int
	pY  int
}

func Day13() {
	games := readFile("day13.test")

	tokens := 0
	for _, game := range games {
		fmt.Println("game:", game)
		tokens += calcButtonPresses(game.baX, game.baY, game.bbX, game.bbY, game.pX, game.pY)
	}
	fmt.Println("final tokens:", tokens)
}

func calcButtonPresses(aX int, aY int, bX int, bY int, priceX int, priceY int) int {
	if priceX > priceY {
		if aX > bX {
			maxMultisA, rX, rY := getPossibleMultis(aX, aY, bX, bY, priceX, priceY)
			maxMultisB, leftOverX, leftOverY := getPossibleMultis(bX, bY, 0, 0, rX, rY)
			if leftOverX > 0 || leftOverY > 0 || (maxMultisA == 0 && maxMultisB == 0) {
				fmt.Println("NOT POSSIBLE")
			} else {
				fmt.Println("A Button multis:", maxMultisA)
				fmt.Println("B Button multis:", maxMultisB)
				tokensA := maxMultisA * 3
				tokensB := maxMultisB * 1
				fmt.Println("tokens:", tokensA+tokensB)
				return tokensA + tokensB
			}
		} else {
			maxMultisB, rX, rY := getPossibleMultis(bX, bY, aX, aY, priceX, priceY)
			maxMultisA, leftOverX, leftOverY := getPossibleMultis(aX, aY, 0, 0, rX, rY)
			if leftOverX > 0 || leftOverY > 0 || (maxMultisA == 0 && maxMultisB == 0) {
				fmt.Println("NOT POSSIBLE")
			} else {
				fmt.Println("B Button multis:", maxMultisB)
				fmt.Println("A Button multis:", maxMultisA)
				tokensA := maxMultisA * 3
				tokensB := maxMultisB * 1
				fmt.Println("tokens:", tokensA+tokensB)
				return tokensA + tokensB
			}
		}
	} else {
		if aY > bY {
			maxMultisA, rX, rY := getPossibleMultis(aY, aX, bY, bX, priceY, priceX)
			maxMultisB, leftOverX, leftOverY := getPossibleMultis(bY, bX, 0, 0, rX, rY)
			if leftOverX > 0 || leftOverY > 0 || (maxMultisA == 0 && maxMultisB == 0) {
				fmt.Println("NOT POSSIBLE")
			} else {
				fmt.Println("A Button multis:", maxMultisA)
				fmt.Println("B Button multis:", maxMultisB)
				tokensA := maxMultisA * 3
				tokensB := maxMultisB * 1
				fmt.Println("tokens:", tokensA+tokensB)
				return tokensA + tokensB
			}
		} else {
			maxMultisB, rX, rY := getPossibleMultis(bY, bX, aY, aX, priceY, priceX)
			maxMultisA, leftOverX, leftOverY := getPossibleMultis(aY, aX, 0, 0, rX, rY)
			if leftOverX > 0 || leftOverY > 0 || (maxMultisA == 0 && maxMultisB == 0) {
				fmt.Println("NOT POSSIBLE")
			} else {
				fmt.Println("B Button multis:", maxMultisB)
				fmt.Println("A Button multis:", maxMultisA)
				tokensA := maxMultisA * 3
				tokensB := maxMultisB * 1
				fmt.Println("tokens:", tokensA+tokensB)
				return tokensA + tokensB
			}
		}
	}
	return 0
}

func getPossibleMultis(aX int, aY int, bX int, bY int, priceX int, priceY int) (maxMultis int, remainingX int, remainingY int) {
	testX := 0
	testY := 0
	times := 0
	lastPossibleTimes := 0
	rX := 0
	rY := 0
	for true {
		testX += aX
		testY += aY
		times++
		remainingX := priceX - testX
		remainingY := priceY - testY
		if remainingX < 0 || remainingY < 0 {
			break
		}
		dividableX := false
		if bX == 0 {
			dividableX = true
		} else {
			if remainingX%bX == 0 {
				dividableX = true
			}
		}
		dividableY := false
		if bY == 0 {
			dividableY = true
		} else {
			if remainingY%bY == 0 {
				dividableY = true
			}
		}
		if dividableX && dividableY {
			lastPossibleTimes = times
			rX = remainingX
			rY = remainingY
		}
	}
	fmt.Println("last possible times X:", lastPossibleTimes, "remaining X:", rX, "remaining Y:", rY)
	return lastPossibleTimes, rX, rY
}

func readFile(filepath string) []Game {
	file, err := os.Open(filepath)
	assert.Nil(err, "Can't open file")
	defer file.Close()

	res := []Game{}

	scanner := bufio.NewScanner(file)
	currentGame := Game{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			res = append(res, currentGame)
			currentGame = Game{}
			continue
		}
		afterButtonA, found := strings.CutPrefix(line, "Button A:")
		if found {
			split := strings.Split(afterButtonA, ",")
			xStr := strings.TrimPrefix(strings.TrimSpace(split[0]), "X+")
			yStr := strings.TrimPrefix(strings.TrimSpace(split[1]), "Y+")
			numX, err := strconv.Atoi(xStr)
			assert.Nil(err, "Strconv fail")
			numY, err := strconv.Atoi(yStr)
			assert.Nil(err, "Strconv fail")
			currentGame.baX = numX
			currentGame.baY = numY
			continue
		}
		afterButtonB, found := strings.CutPrefix(line, "Button B:")
		if found {
			split := strings.Split(afterButtonB, ",")
			xStr := strings.TrimPrefix(strings.TrimSpace(split[0]), "X+")
			yStr := strings.TrimPrefix(strings.TrimSpace(split[1]), "Y+")
			numX, err := strconv.Atoi(xStr)
			assert.Nil(err, "Strconv fail")
			numY, err := strconv.Atoi(yStr)
			assert.Nil(err, "Strconv fail")
			currentGame.bbX = numX
			currentGame.bbY = numY
			continue
		}
		afterPrize, found := strings.CutPrefix(line, "Prize:")
		if found {
			split := strings.Split(afterPrize, ",")
			xStr := strings.TrimPrefix(strings.TrimSpace(split[0]), "X=")
			yStr := strings.TrimPrefix(strings.TrimSpace(split[1]), "Y=")
			numX, err := strconv.Atoi(xStr)
			assert.Nil(err, "Strconv fail")
			numY, err := strconv.Atoi(yStr)
			assert.Nil(err, "Strconv fail")
			currentGame.pX = numX
			currentGame.pY = numY
			continue
		}
	}
	return res
}
