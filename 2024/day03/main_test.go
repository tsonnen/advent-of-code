package main

import (
	"testing"
)

func Test_parseCommands(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		total int
	}{
		{"case01", []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}, 161},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total := parseCommands(tt.lines)
			if tt.total != total {
				t.Errorf("Total %v expected %v", total, tt.total)
			}
		})
	}
}

func Test_parseCommandsExpanded(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		total int
	}{
		{"case01", []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}, 48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total := parseCommandsExpanded(tt.lines)
			if tt.total != total {
				t.Errorf("Total %v expected %v", total, tt.total)
			}
		})
	}
}
