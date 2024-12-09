package main

import (
	"testing"
)

func Test_calculateChecksum(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		checkSum int
	}{
		{"case01", "2333133121414131402", 1928},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diskMap := condenseDiskMap(createDiskMap(tt.input))
			checkSum := calculateCheckSum(diskMap)
			if tt.checkSum != checkSum {
				t.Errorf("Checksum %v expected %v", checkSum, tt.checkSum)
			}
		})
	}
}

func Test_calculateChecksumBlocks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		checkSum int
	}{
		{"case01", "2333133121414131402", 2858},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diskMap := condenseDiskMapBlocks(createDiskMap(tt.input))
			checkSum := calculateCheckSum(diskMap)
			if tt.checkSum != checkSum {
				t.Errorf("Checksum %v expected %v", checkSum, tt.checkSum)
			}
		})
	}
}
