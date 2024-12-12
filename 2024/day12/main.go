package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	gardenPlan := parseInput(inputFile)
	fencePrice, discountedPrice := calcFencePrice(gardenPlan)

	fmt.Printf("The total price of fence is %d\n", fencePrice)
	fmt.Printf("The total price of discounted fence is %d\n", discountedPrice)
}

func parseInput(input string) map[image.Point]string {
	gardenPlan := make(map[image.Point]string)
	for y, line := range strings.Split(input, "\n") {
		for x, char := range strings.Split(line, "") {
			gardenPlan[image.Point{x, y}] = char
		}
	}

	return gardenPlan
}

func calcFencePrice(gardenPlan map[image.Point]string) (int, int) {
	visited := make(map[image.Point]bool)
	price := 0
	discountedPrice := 0

	for pos := range gardenPlan {
		if visited[pos] {
			continue
		}
		visited[pos] = true

		curArea := 1
		perimeter, sides := 0, 0
		queue := []image.Point{pos}

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			for _, dir := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				nextPos := p.Add(dir)
				if gardenPlan[nextPos] != gardenPlan[p] {
					perimeter++
					r := p.Add(image.Point{-dir.Y, dir.X})
					if gardenPlan[r] != gardenPlan[p] || gardenPlan[r.Add(dir)] == gardenPlan[p] {
						sides++
					}
				} else if !visited[nextPos] {
					visited[nextPos] = true
					queue = append(queue, nextPos)
					curArea++
				}
			}
		}

		price += curArea * perimeter
		discountedPrice += curArea * sides
	}

	return price, discountedPrice
}
