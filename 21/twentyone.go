package main

import (
	"advent/2024/grid"
	"advent/2024/parsing"
	"fmt"
	"slices"
)

type keypad map[rune]grid.Coordinate
type directionalPad map[rune]grid.Coordinate

var numericPad = keypad{
	'7': {X: 0, Y: 0},
	'8': {X: 1, Y: 0},
	'9': {X: 2, Y: 0},
	'4': {X: 0, Y: 1},
	'5': {X: 1, Y: 1},
	'6': {X: 2, Y: 1},
	'1': {X: 0, Y: 2},
	'2': {X: 1, Y: 2},
	'3': {X: 2, Y: 2},
	'0': {X: 1, Y: 3},
	'A': {X: 2, Y: 3},
}

var directional = directionalPad{
	'^': {X: 1, Y: 0},
	'A': {X: 2, Y: 0},
	'<': {X: 0, Y: 1},
	'v': {X: 1, Y: 1},
	'>': {X: 2, Y: 1},
}
var forbidden = grid.CreateCoordinate(0, 0)

func main() {
	fmt.Println(part1("data"))
	fmt.Println(part2("data"))
}

func part2(name string) int {
	codes := parsing.ReadLines(name)
	return part2ForLines(codes)
}
func part2ForLines(lines []string) int {
	memo := map[memo]int{}

	totalComplexity := 0
	for _, code := range lines {
		start := grid.CreateCoordinate(2, 3)
		forbidden := grid.CreateCoordinate(0, 3)
		routes := generateAllPossibleRoutes(start, forbidden, code, numericPad)
		codeNumber := parsing.Atoi(code[:len(code)-1])

		best := findBestForMultipeRoute(routes, memo, 25)
		complexity := best * codeNumber
		totalComplexity += complexity

	}
	return totalComplexity
}


func part1(name string) int {
	codes := parsing.ReadLines(name)
	return part1ForLines(codes)
}
func part1ForLines(lines []string) int {
	memo := map[memo]int{}

	totalComplexity := 0
	for _, code := range lines {
		start := grid.CreateCoordinate(2, 3)
		forbidden := grid.CreateCoordinate(0, 3)
		routes := generateAllPossibleRoutes(start, forbidden, code, numericPad)
		codeNumber := parsing.Atoi(code[:len(code)-1])

		best := findBestForMultipeRoute(routes, memo, 2)
		complexity := best * codeNumber
		totalComplexity += complexity

	}
	return totalComplexity
}

func findBestForRoute(route string, memos map[memo]int, depth int) int {

	s := 'A'
	best := 0

	memoKey := memo{route, depth}
	m, found := memos[memoKey]
	if found {
		return m
	}

	if depth == 0 {
		memos[memoKey] = len(route)
		return len(route)
	}

	for _, step := range route {
		e := directional[step]
		var m move = move{directional[s], e, depth - 1}
		possibles := generatePossibleRoutes(m.start, m.end, forbidden)
		b := findBestForMultipeRoute(possibles, memos, m.depth)
		best += b
		s = step
	}
	memos[memoKey] = best
	return best
}

func findBestForMultipeRoute(possibles []string, memo map[memo]int, depth int) int {
	best := findBestForRoute(possibles[0], memo, depth)
	for _, possible := range possibles[1:] {
		posBest := findBestForRoute(possible, memo, depth)
		if best > posBest {
			best = posBest
		}
	}
	return best
}

func findBestRoutes(m move, memos map[memo]int) int {

	if m.depth == 3 {
		dif := m.start.Minus(m.end)
		return abs(dif.X) + abs(dif.Y)
	}
	possibles := generatePossibleRoutes(m.start, m.end, forbidden)
	best := findBestForMultipeRoute(possibles, memos, m.depth)
	return best
}

type move struct {
	start, end grid.Coordinate
	depth      int
}

type memo struct {
	route string
	depth int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i

}

func repeat(r rune, times int) string {
	runes := make([]rune, times)
	for i := range times {
		runes[i] = r
	}
	return string(runes)
}

func (k keypad) calculateVector(from, to rune) grid.Coordinate {
	return k.find(to).Minus(k.find(from))
}

func (k keypad) find(r rune) grid.Coordinate {
	coord, found := k[r]
	if !found {
		panic("rune not found" + string(r))
	}
	return coord
}

func (k keypad) vectorize(input string) []grid.Coordinate {
	vectors := make([]grid.Coordinate, len(input))
	for i, val := range input {
		vectors[i] = k[val]
	}
	return vectors
}

func (k directionalPad) vectorize(input string) []grid.Coordinate {
	return keypad(k).vectorize(input)
}

func generateAllPossibleRoutes(start, forbidden grid.Coordinate, route string, k keypad) []string {
	routes := generatePossibleRoutes(start, k[rune(route[0])], forbidden)
	s := rune(route[0])
	for _, v := range route[1:] {
		newRoutes := []string{}
		for _, newRoute := range generatePossibleRoutes(k[s], k[v], forbidden) {
			for _, route := range routes {
				newRoutes = append(newRoutes, route+newRoute)
			}
		}
		routes = newRoutes
		s = v
	}

	return routes
}

func generatePossibleRoutes(start, end, forbidden grid.Coordinate) []string {
	possibilities := []string{}
	dif := end.Minus(start)
	if start.Plus(grid.Coordinate{X: dif.X}) != forbidden {
		buttons := ""
		if dif.X > 0 {
			buttons += repeat('>', abs(dif.X))
		} else {
			buttons += repeat('<', abs(dif.X))
		}
		if dif.Y < 0 {
			buttons += repeat('^', abs(dif.Y))
		} else {
			buttons += repeat('v', abs(dif.Y))
		}
		buttons += "A"
		possibilities = append(possibilities, buttons)
	}
	if start.Plus(grid.Coordinate{Y: dif.Y}) != forbidden {
		buttons := ""
		if dif.Y < 0 {
			buttons += repeat('^', abs(dif.Y))
		} else {
			buttons += repeat('v', abs(dif.Y))
		}
		if dif.X > 0 {
			buttons += repeat('>', abs(dif.X))
		} else {
			buttons += repeat('<', abs(dif.X))
		}
		buttons += "A"
		possibilities = append(possibilities, buttons)
	}

	return slices.Compact(possibilities)
}
