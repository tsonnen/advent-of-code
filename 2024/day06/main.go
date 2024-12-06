package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	grid, guard := getGridAndGuard(inputFile)
	uniquePositions := simulateGuardPath(grid, guard)
	loopPositions := simLoop(grid, guard)

	fmt.Printf("State visits %d unique positions\n", uniquePositions)
	fmt.Printf("%d loop positions\n", loopPositions)

}

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

var (
	UP    = Direction{0, -1}
	DOWN  = Direction{0, 1}
	LEFT  = Direction{-1, 0}
	RIGHT = Direction{1, 0}
)

type State struct {
	pos Point
	dir Direction
}

func (g *State) turnRight() {
	// turn right 90 degrees
	switch g.dir {
	case UP:
		g.dir = RIGHT
	case RIGHT:
		g.dir = DOWN
	case LEFT:
		g.dir = UP
	case DOWN:
		g.dir = LEFT
	}
}

func getGridAndGuard(input string) ([][]byte, State) {
	var grid [][]byte
	var guard State

	for y, line := range strings.Split(input, "\n") {
		row := make([]byte, len(line))
		for x, cell := range line {
			if cell == '^' {
				guard = State{Point{x, y}, UP}
				row[x] = '.' // replace the guard with an empty space
			} else {
				row[x] = byte(cell)
			}
		}
		grid = append(grid, row)
	}

	return grid, guard
}

// helper func to check if the position is in bounds
func isInBounds(p Point, grid [][]byte) bool {
	return p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[0])
}

func simulateGuardPathLoop(grid [][]byte, guard State) int {
	states := make(map[State]struct{})
	states[guard] = struct{}{}

	for {
		nextPosition := Point{
			x: guard.pos.x + guard.dir.dx,
			y: guard.pos.y + guard.dir.dy,
		}

		if !isInBounds(nextPosition, grid) {
			return 0
		}

		if grid[nextPosition.y][nextPosition.x] == '#' {
			guard.turnRight()
		} else {
			guard.pos = nextPosition
		}

		state := guard
		if _, exists := states[state]; exists {
			return 1
		}
		states[state] = struct{}{}
	}
}

func simulateGuardPath(grid [][]byte, guard State) int {
	visited := make(map[Point]bool)
	visited[guard.pos] = true

	for {
		// calculate nextPosition position
		nextPosition := Point{
			x: guard.pos.x + guard.dir.dx,
			y: guard.pos.y + guard.dir.dy,
		}

		// check if guard would leave the area
		if !isInBounds(nextPosition, grid) {
			break
		}

		// check if there's an obstacle
		if grid[nextPosition.y][nextPosition.x] == '#' {
			guard.turnRight()
			continue
		}

		// move forward
		guard.pos = nextPosition
		visited[guard.pos] = true
	}

	return len(visited)
}

func simLoop(grid [][]byte, guard State) int {
	start := guard.pos
	count := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != '.' || (Point{x, y} == start) {
				continue
			}

			newGrid := make([][]byte, len(grid))
			for i := range grid {
				newGrid[i] = make([]byte, len(grid[i]))
				copy(newGrid[i], grid[i])
			}
			newGrid[y][x] = '#'

			if simulateGuardPathLoop(newGrid, guard) > 0 {
				count++
			}
		}
	}
	return count
}
