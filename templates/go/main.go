package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	inputValues := splitFile()

	fmt.Println(inputValues)
}

func splitFile() []string {
	lines := strings.Split(inputFile, "\n")

	return lines
}
