package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	reportStrings := strings.Split(inputFile, "\n")

	safeReports := calculateSafeReports(reportStrings)
	safeReportsTolerate := calculateSafeReportsTolerateError(reportStrings)

	fmt.Printf("There are %v safe reports\n", safeReports)
	fmt.Printf("There are %v safe reports tolerating an error \n", safeReportsTolerate)
}

func isReportSafe(reportValues []string) bool {
	increasing := atoi(reportValues[0]) < atoi(reportValues[1])

	for i := range len(reportValues) - 1 {
		diff := math.Abs(float64(atoi(reportValues[i]) - atoi(reportValues[i+1])))
		if diff < 1 || diff > 3 {
			return false
		}
		currInc := atoi(reportValues[i]) < atoi(reportValues[i+1])
		if currInc != increasing {
			return false
		}
	}
	return true
}

func calculateSafeReports(reportStrings []string) int {
	safeReports := 0

	for _, report := range reportStrings {
		reportValues := strings.Split(report, " ")

		if isReportSafe(reportValues) {
			safeReports++
		}
	}

	return safeReports
}

func calculateSafeReportsTolerateError(reportStrings []string) int {
	safeReports := 0

outer:
	for _, report := range reportStrings {
		reportValues := strings.Split(report, " ")

		for i := range len(reportValues) {
			newReportRow := slices.Clone(reportValues)

			newReportRow = append(newReportRow[:i], newReportRow[i+1:]...)

			if isReportSafe(newReportRow) {
				safeReports++
				continue outer
			}

		}
	}

	return safeReports
}

func atoi(numString string) int {
	value, _ := strconv.Atoi(numString)

	return value
}
