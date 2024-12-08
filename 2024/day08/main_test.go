package main

import (
	"testing"
)

func Test_calculateAntinodeCount(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		antinodeCount int
	}{
		{"case01", `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			locations, mapSize := getAntennaLocations(tt.input)
			antinodeCount := calculateAntinodeResonanceCount(locations, mapSize)
			if tt.antinodeCount != antinodeCount {
				t.Errorf("Antinode Count %v expected %v", antinodeCount, tt.antinodeCount)
			}
		})
	}
}

func Test_calculateAntinodeResonanceCount(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		antinodeCount int
	}{
		{"case01", `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`, 34},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			locations, mapSize := getAntennaLocations(tt.input)
			antinodeCount := calculateAntinodeResonanceCount(locations, mapSize)
			if tt.antinodeCount != antinodeCount {
				t.Errorf("Antinode Count %v expected %v", antinodeCount, tt.antinodeCount)
			}
		})
	}
}