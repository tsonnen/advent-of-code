package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string;

func main(){
	histPos1, histPos2:= splitFile();
	
	totalDistance, similarityScore := calcTotalDistance(histPos1, histPos2);
	fmt.Printf("Total distance: %v\n", totalDistance);
	fmt.Printf("Similarity score: %v\n", similarityScore);
}

func splitFile()([]int, []int) {
	var arr1 []int;
	var arr2 []int;
	lines := strings.Split(inputFile, "\n")
	for _, line := range lines{
		positions := strings.Split(line, "   ");
		if(len(positions) == 2){
			pos1, _ := strconv.Atoi(positions[0]);
			pos2, _ := strconv.Atoi(positions[1]);
			arr1 = append(arr1, pos1);
			arr2 = append(arr2, pos2);
		}
	}

	return arr1, arr2;
}

func calcTotalDistance(histPos1 []int, histPos2 []int) (int, int){
	totalDistance := 0;
	similarityScore := 0;

	if(!slices.IsSorted(histPos1)){
		slices.Sort(histPos1);
	}

	if(!slices.IsSorted((histPos2))){
		slices.Sort(histPos2);
	}

	for i :=0; i< len(histPos1); i++{
		totalDistance += int(math.Abs(float64(histPos1[i] - histPos2[i])));
		n, found := slices.BinarySearch(histPos2, histPos1[i]);
		if found{
			count := 1;
			for j:= n + 1; j < len(histPos2); j++{
				if(histPos2[j] != histPos2[n]){
					break;
				}

				count++;
			}

			similarityScore += histPos1[i] * count;
		}
	}

	return totalDistance, similarityScore;
}