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
	var height = 103
	var width = 101

	robots := parseInput(inputFile)
	// safetyFactor := moveRobotsAndCalcSafetyFactor(robots, height, width)

	// fmt.Printf("The safety factor is %d\n", safetyFactor)

	easterEgg := moveRobots(robots, height, width, true)

	fmt.Printf("The easter egg happens at %d seconds\n", easterEgg)
}

type SecurityRobot struct {
	pos image.Point
	vel image.Point
}

var numSeconds = 100

func moveRobotsAndCalcSafetyFactor(robots []SecurityRobot, height int, width int) int {
	moveRobots(robots, height, width, false)
	return calculateSafetyFactor(robots, height, width)
}

func checkRobotsEasterEgg(robots []SecurityRobot, height int, width int) bool {
	var image [][]string
	for y := 0; y < height; y++ {
		var row []string
		for x := 0; x < width; x++ {
			row = append(row, ".")
		}
		image = append(image, row)
	}
	for _, robot := range robots {
		image[robot.pos.Y][robot.pos.X] = "#"
	}

	robCount := 0
	for y := 0; y < height; y++ {
		robCount = 0
		for x := 0; x < width; x++ {
			if image[y][x] == "#" {
				robCount++
			}
			if robCount > 10 {
				return true
			}
			if image[y][x] == "." {
				robCount = 0
			}
		}
	}

	return false
}

func moveRobots(robots []SecurityRobot, height int, width int, checkEasterEgg bool) int {
	for i := 0; i < numSeconds; i++ {
		for rI := range robots {
			robot := &robots[rI]
			robot.pos = robot.pos.Add(robot.vel)

			if robot.pos.X < 0 {
				robot.pos.X = width + robot.pos.X
			} else if robot.pos.X >= width {
				robot.pos.X = robot.pos.X - width
			}

			if robot.pos.Y < 0 {
				robot.pos.Y = height + robot.pos.Y
			} else if robot.pos.Y >= height {
				robot.pos.Y = robot.pos.Y - height
			}
		}

		if checkEasterEgg && checkRobotsEasterEgg(robots, height, width) {
			return i + 1
		} else if checkEasterEgg {
			numSeconds++
		}
	}

	return -1
}

func calculateSafetyFactor(robots []SecurityRobot, height int, width int) int {
	vertMiddle := (height - 1) / 2
	horzMiddle := (width - 1) / 2

	var quad1, quad2, quad3, quad4 int

	for _, robot := range robots {
		if robot.pos.X < horzMiddle {
			if robot.pos.Y < vertMiddle {
				quad1++
			} else if robot.pos.Y > vertMiddle {
				quad3++
			}
		} else if robot.pos.X > horzMiddle {
			if robot.pos.Y < vertMiddle {
				quad2++
			} else if robot.pos.Y > vertMiddle {
				quad4++
			}
		}
	}

	return quad1 * quad2 * quad3 * quad4
}

func parseInput(input string) []SecurityRobot {
	var robots []SecurityRobot
	for _, line := range strings.Split(input, "\n") {
		var posX, posY, dx, dy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &posX, &posY, &dx, &dy)
		robots = append(robots, SecurityRobot{image.Pt(posX, posY), image.Pt(dx, dy)})
	}

	return robots
}
