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
	safetyFactor := moveRobotsAndCalcSafetyFactor(robots, height, width)

	fmt.Printf("The safety factor is %d", safetyFactor)
}

type SecurityRobot struct {
	pos image.Point
	vel image.Point
}

var numSeconds = 100

func moveRobotsAndCalcSafetyFactor(robots []SecurityRobot, height int, width int) int {
	moveRobots(robots, height, width)
	return calculateSafetyFactor(robots, height, width)
}

func moveRobots(robots []SecurityRobot, height int, width int) {
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
	}
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
