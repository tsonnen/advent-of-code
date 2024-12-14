package main

import (
	"testing"
)

func Test_moveRobotsAndCalcSafetyFactor(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		height       int
		width        int
		safetyFactor int
	}{
		{"case01", `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`, 7, 11, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			securityRobots := parseInput(tt.input)
			safetyFactor := moveRobotsAndCalcSafetyFactor(securityRobots, tt.height, tt.width)

			if tt.safetyFactor != safetyFactor {
				t.Errorf("Safety factor %v expected %v", safetyFactor, tt.safetyFactor)
			}

		})
	}
}
