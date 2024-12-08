package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	locations, mapSize := getAntennaLocations(inputFile)
	antinodeCount := calculateAntinodeCount(locations, mapSize)
	antinodeCountResonance := calculateAntinodeResonanceCount(locations, mapSize)

	fmt.Printf("There are %d antinode locations\n", antinodeCount)
	fmt.Printf("There are %d antinode locations accounting for resonance\n", antinodeCountResonance)
}

func getAntennaLocations(input string) (map[string][]image.Point, map[image.Point]bool) {
	locations := make(map[string][]image.Point)
	mapCoords := make(map[image.Point]bool)
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			mapCoords[image.Point{x, y}] = true
			if char != "." {
				locations[char] = append(locations[char], image.Point{x, y})
			}
		}
	}

	return locations, mapCoords
}

func calculateAntinodeCount(locations map[string][]image.Point, mapCoords map[image.Point]bool) int {
	antinodeLocations := make(map[image.Point]bool)

	for _, frequency := range locations {
		for i, pos1 := range frequency {
			for _, pos2 := range frequency[i+1:] {
				delta := pos2.Sub(pos1)

				if mapCoords[pos2.Add(delta)] {
					antinodeLocations[pos2.Add(delta)] = true
				}

				if mapCoords[pos1.Sub(delta)] {
					antinodeLocations[pos1.Sub(delta)] = true
				}
			}
		}
	}

	return len(antinodeLocations)
}

func calculateAntinodeResonanceCount(locations map[string][]image.Point, mapCoords map[image.Point]bool) int {
	antinodeLocations := make(map[image.Point]bool)

	for _, frequency := range locations {
		for i, pos1 := range frequency {
			for _, pos2 := range frequency[i+1:] {
				delta := pos2.Sub(pos1)
				tmpPos1 := pos1
				tmpPos2 := pos2
				for mapCoords[tmpPos2] || mapCoords[tmpPos1] {

					if mapCoords[tmpPos2] {
						antinodeLocations[tmpPos2] = true
					}

					if mapCoords[tmpPos1] {
						antinodeLocations[tmpPos1] = true
					}

					tmpPos1 = tmpPos1.Sub(delta)
					tmpPos2 = tmpPos2.Add(delta)
				}
			}
		}
	}

	return len(antinodeLocations)
}
