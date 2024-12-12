package main

import (
	"testing"
)

func Test_findRegions(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		fencePrice      int
		discountedPrice int
	}{
		{"case02", `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`, 772, 436},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gardenPlan := parseInput(tt.input)
			fencePrice, discountedPrice := calcFencePrice(gardenPlan)

			if tt.fencePrice != fencePrice {
				t.Errorf("Fence Price %v expected %v", fencePrice, tt.fencePrice)
			}

			if tt.discountedPrice != discountedPrice {
				t.Errorf("Discounted fence Price %v expected %v", discountedPrice, tt.discountedPrice)
			}

		})
	}
}
