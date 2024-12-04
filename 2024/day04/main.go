package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	totalInstances := searchAll(inputFile)
	totalXMas := searchXMas(inputFile)

	fmt.Printf("The total count is %d\n", totalInstances)
	fmt.Printf("The total X-Mas count is %d\n", totalXMas)
}

func searchAll(input string) int {
	inputLines := strings.Split(input, "\n")
	totalInstances := 0
	graph := make([][]string, len(inputLines))
	for i, line := range inputLines {
		graph[i] = strings.Split(line, "")
	}

	for y, line := range graph {
		for x, letter := range line {
			if letter != "X" && letter != "S" {
				continue
			}

			totalInstances += searchHorizontal(x, y, graph)
			totalInstances += searchVertical(x, y, graph)
			totalInstances += searchDiagRight(x, y, graph)
			totalInstances += searchDiagLeft(x, y, graph)
		}
	}
	return totalInstances
}

func searchHorizontal(x int, y int, graph [][]string) int {
	line := graph[y]
	if len(line) < x+4 {
		return 0
	}

	word := strings.Join(line[x:x+4], "")
	if word == "XMAS" || word == "SAMX" {
		return 1
	}

	return 0
}

func searchVertical(x int, y int, graph [][]string) int {
	if len(graph) < y+4 {
		return 0
	}
	letters := make([]string, 4)
	for i := 0; i < 4; i++ {
		letters[i] = graph[y+i][x]
	}

	foundWord := strings.Join(letters, "")

	if foundWord == "XMAS" || foundWord == "SAMX" {
		return 1
	}

	return 0
}

func searchDiagRight(x int, y int, graph [][]string) int {
	if len(graph) >= y+4 && len(graph[y]) >= x+4 {
		letters := make([]string, 4)
		for i := 0; i < 4; i++ {
			letters[i] = graph[y+i][x+i]
		}

		foundWord := strings.Join(letters, "")

		if foundWord == "XMAS" || foundWord == "SAMX" {
			return 1
		}
	}

	return 0
}

func searchDiagLeft(x int, y int, graph [][]string) int {
	if len(graph) >= y+4 && x >= 3 {
		letters := make([]string, 4)

		for i := 0; i < 4; i++ {
			letters[i] = graph[y+i][x-i]
		}

		foundWord := strings.Join(letters, "")

		if foundWord == "XMAS" || foundWord == "SAMX" {
			return 1
		}
	}

	return 0
}

func searchXMas(input string) int {
	inputLines := strings.Split(input, "\n")
	totalInstances := 0
	graph := make([][]string, len(inputLines))
	for i, line := range inputLines {
		graph[i] = strings.Split(line, "")
	}

	for y, line := range graph {
		for x, letter := range line {
			if letter != "M" && letter != "S" {
				continue
			}

			if len(line) < x+3 || len(graph) < y+3 {
				continue
			}

			letters1 := make([]string, 3)
			letters2 := make([]string, 3)
			for i := 0; i < 3; i++ {
				letters1[i] = graph[y+i][x+i]
				letters2[i] = graph[y+i][x+(2-i)]
			}

			word1 := strings.Join(letters1, "")
			word2 := strings.Join(letters2, "")

			word1Correct := word1 == "SAM" || word1 == "MAS"
			word2Correct := word2 == "SAM" || word2 == "MAS"

			if word1Correct && word2Correct {
				totalInstances++
			}
		}
	}

	return totalInstances
}
