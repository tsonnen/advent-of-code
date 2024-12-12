package main

import (
	"testing"
)

func Test_simulateBlinks(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		numberBlinks int
		stoneCount   int
	}{
		{"case01", "0 1 10 99 999", 1, 7}, {"case02", "125 17", 6, 22}, {"case03", "125 17", 25, 55312},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stones := parseInput(tt.input)
			stoneCount := simulateBlinks(stones, tt.numberBlinks)

			if tt.stoneCount != stoneCount {
				t.Errorf("Number of stones %v expected %v", stoneCount, tt.stoneCount)
			}

		})
	}
}
