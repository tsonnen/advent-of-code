package main

import (
	"testing"
)

func Test_calculateButtonPresses(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		tokenTotal  int
		tokenTotal2 int
	}{
		{"case02", `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`, 480, 875318608908},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clawMachines := parseInput(tt.input)
			tokenTotal := calculateButtonPresses(clawMachines, false)
			tokenTotal2 := calculateButtonPresses(clawMachines, true)

			if tt.tokenTotal != tokenTotal {
				t.Errorf("Token total %v expected %v", tokenTotal, tt.tokenTotal)
			}

			if tt.tokenTotal2 != tokenTotal2 {
				t.Errorf("Token total pt2 %v expected %v", tokenTotal2, tt.tokenTotal2)
			}

		})
	}
}
