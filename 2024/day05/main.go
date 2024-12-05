package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string
var GLOBAL_RULES map[int][]int

func main() {
	rules, pageUpdates := getRulesAndPageUpdates(inputFile)
	totalValid := sumMiddleValidUpdates(rules, pageUpdates, false)
	totalResolveInvalid := sumMiddleValidUpdates(rules, pageUpdates, true)

	fmt.Printf("Sum of valid updates is %d\n", totalValid)
	fmt.Printf("Sum of resolved invalid updates is %d\n", totalResolveInvalid)

}

func getRulesAndPageUpdates(input string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	var pageUpdates [][]int

	lines := strings.Split(input, "\n")
	readRules := true

	for _, line := range lines {
		if line == "" {
			readRules = false
			continue
		}

		if readRules {
			var n1, n2 int
			fmt.Sscanf(line, "%d|%d", &n1, &n2)
			rules[n1] = append(rules[n1], n2)
			continue
		} else {
			pageNumStrings := strings.Split(line, ",")

			pageNums := []int{}

			for _, page := range pageNumStrings {
				pageNums = append(pageNums, atoi(page))
			}
			pageUpdates = append(pageUpdates, pageNums)
		}
	}

	GLOBAL_RULES = rules

	return rules, pageUpdates
}

func sortPages(page1, page2 int) int {
	if slices.Contains(GLOBAL_RULES[page2], page1) {
		return 1
	}

	return -1
}

func sumMiddleValidUpdates(rules map[int][]int, pageUpdates [][]int, resolveInvalid bool) int {
	total := 0

	for _, pageUpdate := range pageUpdates {
		valid := true

		for position, pageNum := range pageUpdate {
			if position == 0 {
				continue
			}

			for i := 0; i < position; i++ {
				if slices.Contains(rules[pageNum], pageUpdate[i]) {
					valid = false
					break
				}
			}

			if !valid {
				break
			}
		}

		if valid && !resolveInvalid {
			total += pageUpdate[int(len(pageUpdate)/2)]
		} else if !valid && resolveInvalid {
			slices.SortFunc(pageUpdate, sortPages)
			total += pageUpdate[int(len(pageUpdate)/2)]
		}
	}

	return total
}

func atoi(numString string) int {
	value, _ := strconv.Atoi(numString)

	return value
}
