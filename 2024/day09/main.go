package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	checkSum := calculateCheckSum(condenseDiskMap(createDiskMap(inputFile)))
	checkSumBlocks := calculateCheckSum(condenseDiskMapBlocks(createDiskMap(inputFile)))

	fmt.Printf("The checksum is %d\n", checkSum)
	fmt.Printf("The checksum  using blocks is %d\n", checkSumBlocks)

}

func createDiskMap(input string) []string {
	var diskMap []string

	digits := strings.Split(input, "")

	pos := 0
	id := 0

	for pos < len(digits) {

		for i := 0; i < atoi(digits[pos]); i++ {
			diskMap = append(diskMap, fmt.Sprintf("%d", id))
		}

		if pos+1 < len(digits) {
			for i := 0; i < atoi(digits[pos+1]); i++ {
				diskMap = append(diskMap, ".")
			}
		}

		id++
		pos += 2
	}

	return diskMap
}

func condenseDiskMapBlocks(diskMap []string) []string {
	blockSize := 0
	var rightPos int
	rightPosScan := len(diskMap) - 1

outer:
	for {
		rightPos = rightPosScan
		for {
			if rightPos <= 0 {
				break outer
			}

			if diskMap[rightPos] == "." {
				rightPos--
			} else {
				break
			}
		}

		rightPosScan = rightPos

		for rightPosScan >= 0 && diskMap[rightPosScan] == diskMap[rightPos] {
			rightPosScan--
		}

		blockSize = rightPos - rightPosScan

		leftPos := 0
		leftPosScan := 0

	leftScan:
		for {
			leftPos = leftPosScan
			for {
				if leftPos >= rightPosScan {
					break leftScan
				}

				if diskMap[leftPos] != "." {
					leftPos++
				} else {
					break
				}
			}

			leftPosScan = leftPos

			for diskMap[leftPosScan] == "." {
				leftPosScan++
			}
			if leftPosScan-leftPos < blockSize {
				continue
			}

			copy(diskMap[leftPos:leftPos+blockSize], diskMap[rightPosScan+1:rightPos+1])
			copy(diskMap[rightPosScan+1:rightPos+1], FillSlice(".", blockSize))
			break
		}
	}

	return diskMap
}

func FillSlice(n string, c int) []string {
	out := make([]string, c)
	for i := 0; i < c; i++ {
		out[i] = n
	}
	return out
}

func condenseDiskMap(diskMap []string) []string {
	pos1 := 0
	pos2 := len(diskMap) - 1

	for diskMap[pos1] != "." {
		pos1++
	}

	for pos1 < pos2 {
		diskMap[pos1] = diskMap[pos2]
		diskMap[pos2] = "."
		for diskMap[pos1] != "." {
			pos1++
		}
		for diskMap[pos2] == "." {
			pos2--
		}
	}

	return diskMap
}

func calculateCheckSum(diskMap []string) int {
	checkSum := 0

	for pos, id := range diskMap {
		if id != "." {
			sum := (atoi(id) * pos)
			checkSum += sum
		}
	}

	return checkSum
}

func atoi(numString string) int {
	value, _ := strconv.Atoi(numString)

	return value
}
