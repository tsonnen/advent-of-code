package main

import (
	_ "embed"
	"fmt"
	"image"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	hikingMap, trailHeads := createHikingMap(inputFile)
	totalScore, totalRating := findHikingScoresAndRatings(hikingMap, trailHeads)

	fmt.Printf("Total Score is %d\n", totalScore)
	fmt.Printf("Total Rating is %d\n", totalRating)

}

func createHikingMap(input string) (map[image.Point]int, []image.Point) {
	var trailHeads []image.Point
	hikingMap := make(map[image.Point]int)

	for y, line := range strings.Split(input, "\n") {
		for x, char := range strings.Split(line, "") {
			elevation := atoi(char)

			hikingMap[image.Point{x, y}] = elevation

			if elevation == 0 {
				trailHeads = append(trailHeads, image.Point{x, y})
			}
		}
	}

	return hikingMap, trailHeads
}

func findHikingScoresAndRatings(hikingMap map[image.Point]int, trailHeads []image.Point) (int, int) {
	totalScores := 0
	totalRatings := 0
	for _, trailHead := range trailHeads {
		endpoints, rating := findHikingEndpoints(trailHead, hikingMap)
		totalScores += len(endpoints)
		totalRatings += rating
	}

	return totalScores, totalRatings
}

func findHikingEndpoints(curPos image.Point, hikingMap map[image.Point]int) (map[image.Point]int, int) {
	curElevation := hikingMap[curPos]
	endpoints := make(map[image.Point]int)
	possibleNextSteps := []image.Point{curPos.Sub(image.Point{1, 0}), curPos.Add(image.Point{1, 0}), curPos.Sub(image.Point{0, 1}), curPos.Add(image.Point{0, 1})}
	rating := 0

	for _, nextStep := range possibleNextSteps {
		elevation, ok := hikingMap[nextStep]

		if !ok {
			continue
		}

		if elevation != curElevation+1 {
			continue
		}

		if elevation == 9 {
			endpoints[nextStep] = 1
			rating++
			continue
		}

		childEndpoints, childRating := findHikingEndpoints(nextStep, hikingMap)
		rating += childRating

		for pos := range childEndpoints {
			endpoints[pos] = 1
		}
	}
	return endpoints, rating
}

func atoi(numString string) int {
	value, _ := strconv.Atoi(numString)

	return value
}
