package main

import (
	_ "embed"
	"image"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	parseInput(inputFile)
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

func atoi(numString string) int {
	value, _ := strconv.Atoi(numString)

	return value
}
