package main

import (
	"testing"
)

func Test_notQuiteLisp(t *testing.T){
	tests := []struct {
		name string
		input string
		floor  int
		enteredBasementAt  int
	}{
		{"actual", inputFile, 280, 1797},
		{"case1", "(())", 0,0},
		{"case2", "()()", 0,0},
		{"case3", ")())())", -3, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			floor, enteredBasementAt := notQuiteLisp(tt.input);
			if (tt.floor != floor){
				t.Errorf("Exited on floor %v expected to exit on floor %v", floor, tt.floor)
			}

			if (tt.enteredBasementAt != enteredBasementAt){
				t.Errorf("Entered basement at position %v expected to enter at position %v", enteredBasementAt, tt.enteredBasementAt)
			}
		})
	}
}