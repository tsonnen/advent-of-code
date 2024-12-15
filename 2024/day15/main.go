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
	robotLoc, warehouseMap, moves := parseInput(inputFile)
	gpsSum := moveRobot(robotLoc, warehouseMap, moves)

	fmt.Printf("The GPS sum is %d\n", gpsSum)
}

func parseInput(input string) (image.Point, map[image.Point]string, []string) {
	readingMap := true
	warehouseMap := make(map[image.Point]string)
	var moves []string
	var robotLoc image.Point

	for y, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			readingMap = false
		}
		for x, char := range strings.Split(line, "") {
			if readingMap {
				warehouseMap[image.Point{x, y}] = char
				if char == "@" {
					robotLoc = image.Pt(x, y)
				}
			} else {
				moves = append(moves, char)
			}
		}
	}

	return robotLoc, warehouseMap, moves
}

func moveRobot(robotLoc image.Point, warehouseMap map[image.Point]string, moves []string) int {
	for _, move := range moves {
		var movementDirection image.Point
		switch move {
		case "^":
			movementDirection = image.Pt(0, -1)
		case "<":
			movementDirection = image.Pt(-1, 0)
		case ">":
			movementDirection = image.Pt(1, 0)
		case "v":
			movementDirection = image.Pt(0, 1)
		}

		validMove, numBoxes := canMakeMove(robotLoc, warehouseMap, movementDirection)

		if validMove {
			warehouseMap[robotLoc] = "."
			robotLoc = robotLoc.Add(movementDirection)
			warehouseMap[robotLoc] = "@"

			for i := range numBoxes {
				boxLoc := robotLoc.Add(movementDirection.Mul(i + 1))
				warehouseMap[boxLoc] = "O"
			}
		}
	}

	gpsSum := 0

	for pos := range warehouseMap {
		if warehouseMap[pos] == "O" {
			gpsSum += (100 * pos.Y) + pos.X
		}
	}

	return gpsSum
}

func canMakeMove(robotLoc image.Point, warehouseMap map[image.Point]string, movementDirection image.Point) (bool, int) {
	numBoxes := 0
	nextPos := robotLoc.Add(movementDirection)

	for {
		switch warehouseMap[nextPos] {
		case ".":
			return true, numBoxes
		case "O":
			numBoxes++
			nextPos = nextPos.Add(movementDirection)
		case "#":
			return false, -1
		}
	}
}
