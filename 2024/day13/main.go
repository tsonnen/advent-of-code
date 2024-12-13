package main

import (
	_ "embed"
	"fmt"
	"image"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	clawMachines := parseInput(inputFile)
	tokenTotal := calculateButtonPresses(clawMachines, false)
	tokenTotal2 := calculateButtonPresses(clawMachines, true)

	fmt.Printf("Total token cost %d\n", tokenTotal)
	fmt.Printf("Total token cost pt2 %d\n", tokenTotal2)

}

type ClawMachine struct {
	buttonA image.Point
	buttonB image.Point

	prizeLoc image.Point
}

func parseInput(input string) []ClawMachine {
	var clawMachines []ClawMachine
	lines := strings.Split(input, "\n")

	buttonRegex := regexp.MustCompile(`Button (?:A|B): X\+([0-9]+), Y\+([0-9]+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)

	for i := 0; i < len(lines); {
		buttonAVals := buttonRegex.FindStringSubmatch(lines[i])
		buttonBVals := buttonRegex.FindStringSubmatch(lines[i+1])
		prizeVals := prizeRegex.FindStringSubmatch(lines[i+2])

		clawMachines = append(clawMachines,
			ClawMachine{
				buttonA:  image.Pt(atoi(buttonAVals[1]), atoi(buttonAVals[2])),
				buttonB:  image.Pt(atoi(buttonBVals[1]), atoi(buttonBVals[2])),
				prizeLoc: image.Pt(atoi(prizeVals[1]), atoi(prizeVals[2])),
			})
		i += 4
	}

	return clawMachines
}

func calculateButtonPresses(clawMachines []ClawMachine, isPart2 bool) int {
	total := 0

	for _, clawMachine := range clawMachines {
		prizeLoc := clawMachine.prizeLoc

		if isPart2 {
			prizeLoc = prizeLoc.Add(image.Point{10000000000000, 10000000000000})
		}
		aPresses :=
			(clawMachine.buttonB.Y*prizeLoc.X - clawMachine.buttonB.X*prizeLoc.Y) /
				(clawMachine.buttonA.X*clawMachine.buttonB.Y - clawMachine.buttonA.Y*clawMachine.buttonB.X)
		bPresses :=
			(clawMachine.buttonA.Y*prizeLoc.X - clawMachine.buttonA.X*prizeLoc.Y) /
				(clawMachine.buttonA.Y*clawMachine.buttonB.X - clawMachine.buttonA.X*clawMachine.buttonB.Y)
		if clawMachine.buttonA.Mul(aPresses).Add(clawMachine.buttonB.Mul(bPresses)) == prizeLoc {
			total += aPresses*3 + bPresses
		}
	}

	return total
}

func atoi(numString string) int {
	value, _ := strconv.Atoi(numString)

	return value
}
