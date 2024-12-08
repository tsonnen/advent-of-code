package main

import (
	"testing"
)

func Test_sumValidEquations(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		allowConcat bool
		total       int
	}{
		{"case01", `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`, false, 3749},
		{"case02", `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`, true, 11387},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equations := parseEquations(tt.input)
			total := sumValidEquations(equations, tt.allowConcat)
			if tt.total != total {
				t.Errorf("Total %v expected %v", total, tt.total)
			}
		})
	}
}
