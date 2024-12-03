package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	lines := strings.Split(inputFile, "\n")

	total := parseCommands(lines)
	totalExpanded := parseCommandsExpanded(lines)

	fmt.Printf("The total is %d\n", total)
	fmt.Printf("The total using expanded is %d\n", totalExpanded)
}

func parseCommands(lines []string) int {
	total := 0
	r := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			valuesToMultiply := r.FindStringSubmatch(match)
			total += atoi(valuesToMultiply[1]) * atoi(valuesToMultiply[2])
		}
	}

	return total
}

func parseCommandsExpanded(lines []string) int {
	total := 0
	shouldMultiply := true
	r := regexp.MustCompile(`(?:mul\(([0-9]{1,3}),([0-9]{1,3})\))|(?:do\(\))|(?:don't\(\))`)
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			if match == "do()" {
				shouldMultiply = true
				continue
			}
			if match == "don't()" {
				shouldMultiply = false
				continue
			}
			if shouldMultiply {
				valuesToMultiply := r.FindStringSubmatch(match)
				total += atoi(valuesToMultiply[1]) * atoi(valuesToMultiply[2])
			}
		}
	}

	return total
}

func atoi(numString string) int {
	value, _ := strconv.Atoi(numString)

	return value
}
