package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	equations := parseEquations(inputFile)
	total := sumValidEquations(equations, false)
	totalConcat := sumValidEquations(equations, true)

	fmt.Printf("The sum of valid equations is %d\n", total)
	fmt.Printf("The sum of valid equations allowing concat is %d\n", totalConcat)

}

type Equation struct {
	total  int
	values []int
}

func sumValidEquations(equations []Equation, allowConcat bool) int {
	total := 0
	for _, equation := range equations {
		base := 2

		if allowConcat {
			base = 3
		}
		if equation.canMake(base) {
			total += equation.total
		}
	}

	return total
}

var operators = [3]string{"+", "*", "||"}

func (e *Equation) canMake(base int) bool {
	ops := int(math.Pow(float64(base), float64(len(e.values)-1)))
outer:
	for p := range ops {
		permNum, total := p, e.values[0]

		for pos, value := range e.values[1:] {
			switch operators[permNum%base] {
			case "+":
				{
					total += value
					break
				}
			case "*":
				{
					total *= value
					break
				}
			case "||":
				{
					total = concat(total, value)
					break
				}

			}

			if pos == len(e.values[1:])-1 && total == e.total {
				return true
			}
			if total > e.total {
				continue outer
			}

			permNum /= base
		}
	}

	return false
}

func concat(i1, i2 int) int {
	combo := fmt.Sprintf("%d%d", i1, i2)
	return atoi(combo)
}

func parseEquations(input string) []Equation {
	var equations []Equation
	for _, line := range strings.Split(input, "\n") {
		splitLine := strings.Split(line, ": ")
		total := atoi(splitLine[0])
		var values []int

		for _, value := range strings.Split(splitLine[1], " ") {
			values = append(values, atoi(value))
		}

		equations = append(equations, Equation{total, values})
	}

	return equations
}

func atoi(numString string) int {
	value, _ := strconv.Atoi(numString)

	return value
}
