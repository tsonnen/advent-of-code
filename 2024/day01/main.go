package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	histPos1, histPos2 := splitFile()
	secondLocationsMap := generateLocationMap(histPos2)

	totalDistance, similarityScore := calcDistanceAndSimilarity(histPos1, histPos2, secondLocationsMap)
	fmt.Printf("Total distance: %v\n", totalDistance)
	fmt.Printf("Similarity score: %v\n", similarityScore)
}

func splitFile() ([]int, []int) {
	var arr1 []int
	var arr2 []int
	lines := strings.Split(inputFile, "\n")
	for _, line := range lines {
		var n1, n2 int
		fmt.Sscanf(line, "%d   %d", &n1, &n2)
		arr1 = append(arr1, n1)
		arr2 = append(arr2, n2)
	}

	return arr1, arr2
}

func generateLocationMap(locationArray []int) map[int]int {
	locationMap := make(map[int]int)

	for _, location := range locationArray {
		locationMap[location]++
	}

	return locationMap
}

func calcDistanceAndSimilarity(histPos1 []int, histPos2 []int, secondLocationsMap map[int]int) (int, int) {
	totalDistance := 0
	similarityScore := 0

	slices.Sort(histPos1)
	slices.Sort(histPos2)

	for i := 0; i < len(histPos1); i++ {
		totalDistance += int(math.Abs(float64(histPos1[i] - histPos2[i])))
		similarityScore += secondLocationsMap[histPos1[i]] * histPos1[i]
	}

	return totalDistance, similarityScore
}
