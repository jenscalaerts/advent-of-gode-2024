package main

import (
	"advent/2024/grid"
	"advent/2024/parsing"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	data := readData("data")
	length := findRoute(data, 71, 1024)
	fmt.Println("Part 1:", length)
	index := part2(data, 71)
	fmt.Println(index)
}

func part2(data map[grid.Coordinate]int, bound int) grid.Coordinate {
	indices := make([]int, len(data))
	for i := range indices {
		indices[i] = i
	}
	result, _ := slices.BinarySearchFunc(indices, true, func(index int, _ bool) int {
		result := findRoute(data, bound, index)
		if result != -1 {
			return -1
		}
		if findRoute(data, bound, index-1) != -1 {
			return 0
		}
		return 1
	})

	for key, val := range data {
		if val == result {
			return key
		}
	}
	panic("not found")
}

func readData(name string) map[grid.Coordinate]int {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	droppingBricks := map[grid.Coordinate]int{}

	var i int
	for scanner.Scan() {
		i++
		split := strings.Split(scanner.Text(), ",")
		droppingBricks[grid.CreateCoordinate(parsing.Atoi(split[0]), parsing.Atoi(split[1]))] = i
	}

	return droppingBricks
}

func findRoute(bricks map[grid.Coordinate]int, bound, timestamp int) int {
	tails := []coordinateAndScore{{coord: grid.CreateCoordinate(0, 0), length: 0}}
	bestScores := map[grid.Coordinate]int{}
	for len(tails) > 0 {
		tail := tails[0]
		adjecents := tail.coord.GetAdjectents(bound, bound)
		pathLength := tail.length + 1
		for _, adjecent := range adjecents {
			brickTime, found := bricks[adjecent]
			if (!found || brickTime > timestamp) && (bestScores[adjecent] == 0 || bestScores[adjecent] > pathLength) {
				if adjecent == grid.CreateCoordinate(bound-1, bound-1) {
					return pathLength
				}
				bestScores[adjecent] = pathLength
				tails = append(tails, coordinateAndScore{adjecent, pathLength})
			}
		}
		tails = tails[1:]
	}
	return -1
}

type coordinateAndScore struct {
	coord  grid.Coordinate
	length int
}
