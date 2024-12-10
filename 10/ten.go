package main

import (
	"advent/2024/grid"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	data := grid.ReadGrid("data")
	fmt.Printf("Part 1 %d\n", findTrailHeadScores(data))
	fmt.Printf("Part 2 %d\n", findTrailHeadScoresAllRoutes(data))
}

func findTrailHeadScoresAllRoutes(g grid.Grid) int {
	trailHeads := g.FindAll('0')
	trailEnds := map[grid.Coordinate][]grid.Coordinate{}
	for _, trailHead := range trailHeads {
		trailEnds[trailHead] = []grid.Coordinate{trailHead}
	}

	for nextLevelInt := 1; nextLevelInt < 10; nextLevelInt++ {
		nextLevel := strconv.Itoa(nextLevelInt)[0]
		trailEnds = findNextSteps(nextLevel, trailEnds, g)
	}
	sum := 0
	for _, starts := range trailEnds {
		sum += len(starts)
	}
	return sum
}

func findNextSteps(expectedValue byte, trailEnds map[grid.Coordinate][]grid.Coordinate, g grid.Grid) map[grid.Coordinate][]grid.Coordinate {
	newTrailEnds := map[grid.Coordinate][]grid.Coordinate{}
	for trailEnd, fromStrats := range trailEnds {
		for _, adjecent := range findMatchingAdjecents(expectedValue, trailEnd, g) {
			newTrailEnds[adjecent.Loc] = append(newTrailEnds[adjecent.Loc], fromStrats...)
		}
	}
	return newTrailEnds
}

func findMatchingAdjecents(expectedValue byte, c grid.Coordinate, g grid.Grid) []grid.AdjecentResult {
	adj := g.GetCardinalAdjecents(c)
	adj = slices.DeleteFunc(adj, func(a grid.AdjecentResult) bool {
		return a.Value != expectedValue
	})
	return adj
}
func findTrailHeadScores(g grid.Grid) int {
	trailHeads := g.FindAll('0')
	trailEnds := map[grid.Coordinate]map[grid.Coordinate]bool{}
	for _, trailHead := range trailHeads {
		trailEnds[trailHead] = map[grid.Coordinate]bool{trailHead: true}
	}

	for nextLevelInt := 1; nextLevelInt < 10; nextLevelInt++ {
		nextLevel := strconv.Itoa(nextLevelInt)[0]
		newTrailEnds := map[grid.Coordinate]map[grid.Coordinate]bool{}
		for trailEnd, fromStrats := range trailEnds {
			adjecents := g.GetCardinalAdjecents(trailEnd)
			for _, adjecent := range adjecents {
				if adjecent.Value == nextLevel {
					val, exists := newTrailEnds[adjecent.Loc]
					if !exists {
						val = make(map[grid.Coordinate]bool, 0)
					}
					for coord := range fromStrats {
						val[coord] = true
					}
					newTrailEnds[adjecent.Loc] = val
				}
			}
		}
		trailEnds = newTrailEnds
	}
	sum := 0
	for _, starts := range trailEnds {
		sum += len(starts)
	}
	return sum
}
