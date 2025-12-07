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

	robotLocPt2, warehouseMapPt2, _ := parseInputPart2(inputFile)
	gpsSumPt2 := moveRobotPt2(robotLocPt2, warehouseMapPt2, moves)

	fmt.Printf("The GPS sum is %d\n", gpsSum)
	fmt.Printf("The GPS sum for part 2 is %d\n", gpsSumPt2)
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

func parseInputPart2(input string) (image.Point, map[image.Point]string, []string) {
	readingMap := true
	warehouseMap := make(map[image.Point]string)
	var moves []string
	var robotLoc image.Point

	for y, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			readingMap = false
		}
		for lineIndex, char := range strings.Split(line, "") {
			if readingMap {
				x := lineIndex * 2
				switch char {
				case ".", "#":
					warehouseMap[image.Point{x, y}] = char
					warehouseMap[image.Point{x + 1, y}] = char
				case "@":
					robotLoc = image.Pt(x, y)
					warehouseMap[image.Point{x, y}] = "@"
					warehouseMap[image.Point{x + 1, y}] = "."
				case "O":
					warehouseMap[image.Point{x, y}] = "["
					warehouseMap[image.Point{x + 1, y}] = "]"
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

func moveRobotPt2(robotLoc image.Point, warehouseMap map[image.Point]string, moves []string) int {
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

		validMove, numBoxes := canMakeMovePt2(robotLoc, warehouseMap, movementDirection)

		if validMove {
			if movementDirection == image.Pt(1, 0) {
				for n := range numBoxes {
					rSide := robotLoc.Add(image.Pt((numBoxes-n)*2, 0))
					lSide := rSide.Sub(image.Pt(1, 0))

					warehouseMap[rSide] = "]"
					warehouseMap[lSide] = "["
				}
			} else if movementDirection == image.Pt(-1, 0) {
				for n := range numBoxes {
					lSide := robotLoc.Sub(image.Pt((numBoxes-n)*2, 0))
					rSide := lSide.Add(image.Pt(1, 0))

					warehouseMap[rSide] = "]"
					warehouseMap[lSide] = "["
				}
			} else {

			}

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

func canMakeMovePt2(robotLoc image.Point, warehouseMap map[image.Point]string, movementDirection image.Point) (bool, int) {
	numBoxes := 0
	nextPos := robotLoc.Add(movementDirection)

	for {
		switch warehouseMap[nextPos] {
		case ".":
			return true, numBoxes
		case "[", "]":
			numBoxes++
			if movementDirection == image.Pt(1, 0) || movementDirection == image.Pt(-1, 0) {
				// If moving horizontally, move over to the next box
				nextPos = nextPos.Add(movementDirection.Mul(2))
			} else {
				nextPos = nextPos.Add(movementDirection)
			}
		case "#":
			return false, -1
		}
	}
}
