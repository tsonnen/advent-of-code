package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputFile string

func main() {
	stones := parseInput(inputFile)
	total := simulateBlinks(stones, 75)

	fmt.Printf("There are %d stones\n", total)
}

func parseInput(input string) []int {
	var stones []int
	for _, number := range strings.Split(input, " ") {
		stones = append(stones, atoi(number))
	}

	return stones
}

func simulateBlinks(stones []int, blinks int) int {
	start := time.Now()
	total := 0
	for _, stone := range stones {
		total += blinkStoneCache(stone, blinks)
	}

	passed := time.Since(start)
	fmt.Printf("%d seconds passed\n", passed/time.Duration(math.Pow(10, 9)))

	return total
}

type StoneBlinkResult struct {
	val       int
	numBlinks int
}

var cache = make(map[StoneBlinkResult]int)

func blinkStoneCache(stone int, blinks int) int {
	if cachedVal, ok := cache[StoneBlinkResult{val: stone, numBlinks: blinks}]; ok {
		return cachedVal
	}

	if blinks == 0 {
		return 1
	}

	if stone == 0 {
		value := blinkStoneCache(1, blinks-1)
		cache[StoneBlinkResult{val: stone, numBlinks: blinks}] = value
		return value
	} else if len(strconv.Itoa(stone))%2 == 0 {
		stoneString := strings.Split(strconv.Itoa(stone), "")
		midPoint := len(stoneString) / 2
		val1 := atoi(strings.Join(stoneString[:midPoint], ""))
		val2 := atoi(strings.Join(stoneString[midPoint:], ""))
		value := blinkStoneCache(val1, blinks-1) + blinkStoneCache(val2, blinks-1)
		cache[StoneBlinkResult{val: stone, numBlinks: blinks}] = value
		return value
	}
	value := blinkStoneCache(stone*2024, blinks-1)
	cache[StoneBlinkResult{val: stone, numBlinks: blinks}] = value
	return value
}

func atoi(numString string) int {
	value, _ := strconv.Atoi(numString)

	return value
}
