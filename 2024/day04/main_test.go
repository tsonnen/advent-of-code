package main

import (
	"testing"
)

func Test_searchAll(t *testing.T) {
	tests := []struct {
		name  string
		input string
		total int
	}{
		{"case01",
			`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total := searchAll(tt.input)
			if tt.total != total {
				t.Errorf("Total %v expected %v", total, tt.total)
			}
		})
	}
}

func Test_searchXMas(t *testing.T) {
	tests := []struct {
		name  string
		input string
		total int
	}{
		{"case01",
			`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total := searchXMas(tt.input)
			if tt.total != total {
				t.Errorf("Total %v expected %v", total, tt.total)
			}
		})
	}
}
