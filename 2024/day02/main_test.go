package main

import (
	"testing"
)

func Test_calculateSafeReports(t *testing.T) {
	tests := []struct {
		name          string
		reportStrings []string
		safeReports   int
	}{
		{"case01", []string{"7 6 4 2 1", "1 2 7 8 9", "9 7 6 2 1", "1 3 2 4 5", "8 6 4 4 1", "1 3 6 7 9"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			safeReports := calculateSafeReports(tt.reportStrings)
			if tt.safeReports != safeReports {
				t.Errorf("Safe reports %v expected %v", safeReports, tt.safeReports)
			}
		})
	}
}

func Test_calculateSafeReportsTolerateError(t *testing.T) {
	tests := []struct {
		name          string
		reportStrings []string
		safeReports   int
	}{
		{"case01", []string{"7 6 4 2 1", "1 2 7 8 9", "9 7 6 2 1", "1 3 2 4 5", "8 6 4 4 1", "1 3 6 7 9"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			safeReports := calculateSafeReportsTolerateError(tt.reportStrings)
			if tt.safeReports != safeReports {
				t.Errorf("Safe reports tolerate error %v expected %v", safeReports, tt.safeReports)
			}
		})
	}
}
