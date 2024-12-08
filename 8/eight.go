package main

import (
	"advent/2024/grid"
	"fmt"
	"slices"
)

func main() {
	g := grid.ReadGrid("data")
	fmt.Println(len(calculateUniqueAntiNodes(g)))
	fmt.Println(len(calculateUniqueAntiNodesWithResonant(g)))
}

func calculateUniqueAntiNodes(g grid.Grid) map[grid.Coordinate]bool {
	coordinatesByValue := g.AsValueCoordinateMap()
	results := map[grid.Coordinate]bool{}
	for value, coordinates := range coordinatesByValue {
		if value == '.' {
			continue
		}

		for i, left := range coordinates[:len(coordinates)-1] {
			for _, right := range coordinates[i+1:] {

				diff := left.Minus(right)
				antiNode1 := left.Plus(diff)
				if g.IsInGrid(antiNode1) {
					results[antiNode1] = true
				}
				antiNode2 := right.Minus(diff)
				if g.IsInGrid(antiNode2) {
					results[antiNode2] = true
				}
			}
		}
	}
	return results
}

func calculateUniqueAntiNodesWithResonant(g grid.Grid) map[grid.Coordinate]bool {
	coordinatesByValue := g.AsValueCoordinateMap()
	results := map[grid.Coordinate]bool{}
	for value, coordinates := range coordinatesByValue {
		if value == '.' {
			continue
		}
		antiNodes := findResonantAntiNodesInRange(g, coordinates)
		for _, antiNode := range antiNodes {
			results[antiNode] = true
		}
	}
	return results
}

func findResonantAntiNodesInRange(g grid.Grid, coordinates []grid.Coordinate) []grid.Coordinate {
	results := slices.Clone([]grid.Coordinate{})
	for i, left := range coordinates[:len(coordinates)-1] {
		for _, right := range coordinates[i+1:] {

			diff := left.Minus(right)
			for l := left; g.IsInGrid(l); l = l.Minus(diff) {
				fmt.Println(l)
				results = append(results, l)
			}
			for l := right; g.IsInGrid(l); l = l.Plus(diff) {
				results = append(results, l)
			}
		}
	}
	return results
}
