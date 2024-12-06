package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

var north = coordinate{0, -1}

func main() {
	g := readGrid("data")
	route := determineRoute(g)
	fmt.Println(len(route))
	fmt.Println(findPositionsResultingInCircle(g))
}

type guardSnapshot struct {
	loc coordinate
	dir coordinate
}

func findPositionsResultingInCircle(g grid) int{
    count := 0
    for _, loc := range g.findAll('.'){
        cp := g.copy()
        cp.set(loc, '#')
        if determineRouteIsCircular(cp) {
            count ++
        }
    }
    return count
}

func determineRouteIsCircular(g grid) bool {
	starts := g.findAll('^')
	if len(starts) > 1 {
		log.Fatal("invalid input")
	}
	currentLocation := starts[0]
	direction := north
	passedLocations := map[guardSnapshot]bool{
		{currentLocation, direction}: true,
	}
	for {
		nextLocation := currentLocation.plus(direction)
		nextValue := g.get(nextLocation)
		if nextValue == 0 {
			return false
		}

		if passedLocations[guardSnapshot{nextLocation, direction}] {
			return true
		}
		if nextValue == '#' {
			direction = direction.rotateRight()
		}
		if nextValue == '.' || nextValue == '^' {
			currentLocation = nextLocation
		}
		passedLocations[guardSnapshot{currentLocation, direction}] = true
	}
}

func determineRoute(g grid) map[coordinate]bool {
	starts := g.findAll('^')
	if len(starts) > 1 {
		log.Fatal("invalid input")
	}
	currentLocation := starts[0]
	direction := north
	passedLocations := map[coordinate]bool{
		currentLocation: true,
	}
	for {
		nextLocation := currentLocation.plus(direction)
		nextValue := g.get(nextLocation)
		if nextValue == 0 {
			break
		}
		if nextValue == '#' {
			direction = direction.rotateRight()
		}
		if nextValue == '.' || nextValue == '^' {
			currentLocation = nextLocation
			passedLocations[currentLocation] = true
		}
	}
	return passedLocations
}

type coordinate struct {
	x, y int
}

func (c coordinate) rotateRight() coordinate {
	return coordinate{-c.y, c.x}
}

type grid []string

func (g grid) get(c coordinate) byte {
	if c.x < 0 || c.x >= len(g[0]) {
		return 0
	}
	if c.y < 0 || c.y >= len(g) {
		return 0
	}
	val := g[c.y][c.x]
	return val
}

func (g grid) findAll(b byte) []coordinate {
	var coords []coordinate
	for y, row := range g {
		for x, el := range []byte(row) {
			if el == b {
				coords = append(coords, coordinate{x, y})
			}
		}
	}
	return coords
}

func (l coordinate) plus(r coordinate) coordinate {
	return coordinate{l.x + r.x, l.y + r.y}
}
func readGrid(name string) grid {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}
func (g grid) copy() grid {
	return slices.Clone(g)
}

func (g *grid) set(c coordinate, val byte) {
	gr := *g
	gr[c.y] = gr[c.y][:c.x] + string(val) + gr[c.y][c.x+1:]
}
func (g grid) print() {
	for _, row := range g {
		fmt.Println(row)
	}
	fmt.Println()
}
