package main

import (
	"testing"
)

func Test_calcTotalDistance(t *testing.T){
	actualArr1, actualArr2 := splitFile();
	tests := []struct {
		name string
		inputArray1 []int;
		inputArray2 []int;
		secondLocationsMap map[int]int
		totalDistance  int
		similarityScore  int
	}{
		{"actual", actualArr1, actualArr2, generateLocationMap(actualArr2), 1222801, 22545250},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalDistance, similarityScore := calcDistanceAndSimilarity(tt.inputArray1, tt.inputArray2, tt.secondLocationsMap);
			if (tt.totalDistance != totalDistance){
				t.Errorf("Total distance %v expected %v", totalDistance, tt.totalDistance)
			}

			if (tt.similarityScore != similarityScore){
				t.Errorf("Similarity score %v expected %v", similarityScore, tt.similarityScore)
			}
		})
	}
}