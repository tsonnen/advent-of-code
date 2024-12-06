package main

import (
	"testing"
)

func Test_simulateGuardPath(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		uniquePositions int
	}{
		{"case01", `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`, 41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid, guard := getGridAndGuard(tt.input)
			uniquePositions := simulateGuardPath(grid, guard)
			if tt.uniquePositions != uniquePositions {
				t.Errorf("Unique positions %v expected %v", uniquePositions, tt.uniquePositions)
			}
		})
	}
}
