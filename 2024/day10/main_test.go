package main

import (
	"testing"
)

func Test_findTotalHikingScores(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		totalScore int
	}{
		{"case01", `0123
1234
8765
9876`, 1}, {"case02", `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hikingMap, trailHeads := createHikingMap(tt.input)
			totalScore := findHikingScores(hikingMap, trailHeads)
			if tt.totalScore != totalScore {
				t.Errorf("Total score %v expected %v", totalScore, tt.totalScore)
			}
		})
	}
}
