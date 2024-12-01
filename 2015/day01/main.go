package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var inputFile string;

func main(){
	floor, enteredBasementAt := notQuiteLisp(inputFile)
	fmt.Printf("Exit on floor %#v\n", floor);
	fmt.Printf("Entered the basement at position %#v\n", enteredBasementAt);
}

func notQuiteLisp(input string) (int, int) {
	floor := 0;
	var enteredBasementAt int;
	for pos, r := range input {
		if r == '(' {
			floor++
		} else {
			floor--
		}

		if floor < 0 && enteredBasementAt == 0{
			enteredBasementAt = pos + 1
		}
	}

	return floor, enteredBasementAt
}