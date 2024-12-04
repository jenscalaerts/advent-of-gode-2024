package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

var nw = coordinate{-1, -1}
var ne = coordinate{1, -1}
var sw = coordinate{-1, 1}
var se = coordinate{1, 1}

func main() {
	g := readGrid("data")
	fmt.Printf("Part 1: %d\n", countXMASStraight(g))
	fmt.Printf("Part 2: %d\n", countX_MAS(g))
}

func countXMASStraight(g grid) int {
	locations := g.findAll('X')
	directions := coordinate{0, 0}.adjecents()
	count := 0
	for _, loc := range locations {
		for _, dir := range directions {
			if g.get(loc.plus(dir.times(3))) != nil {
				wordBytes := make([]byte, 4)
				for i := range 4 {
					wordBytes[i] = *g.get(loc.plus(dir.times(i)))
				}
				word := string(wordBytes)
				if word == "XMAS" {
					count++
				}
			}

		}
	}
	return count
}

func countX_MAS(g grid) int {
	locations := g.findAll('A')
	masCount := 0
	for _, loc := range locations {
		if isMS(g, loc, []coordinate{ne, sw}) && isMS(g, loc, []coordinate{nw, se}) {
			masCount++
		}
	}
	return masCount
}

func isMS(g grid, loc coordinate, dirs []coordinate) bool {
	var others []byte
	for _, dir := range dirs {
		c := loc.plus(dir)
		val := g.get(c)
		if val != nil {
			others = append(others, *val)
		}
	}
	return slices.Contains(others, 'M') && slices.Contains(others, 'S')
}


//I misunderstood the part to allow snaking. Keeping this around for later
func countXMASComplicado(g grid) {
	locations := g.findAll('X')[:1]

	for _, nextItem := range []byte("MAS") {
		locations = adjecentLocationsMatching(g, locations, nextItem)
		fmt.Println(locations)
	}
	fmt.Printf("Part 1: %d\n", locations)

}

func adjecentLocationsMatching(g grid, coords []coordinate, value byte) []coordinate {
	var allAdjecents []coordinate
	for _, c := range coords {
		for _, adj := range g.getAdjecents(c) {
			if adj.value == value {
				allAdjecents = append(allAdjecents, adj.loc)
			}
		}
	}
	return allAdjecents
}

type grid []string
type coordinate struct {
	x, y int
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

func (g grid) get(c coordinate) *byte {
	if c.x < 0 || c.x >= len(g[0]) {
		return nil
	}
	if c.y < 0 || c.y >= len(g) {
		return nil
	}
	val := g[c.y][c.x]
	return &val
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

func (g grid) getAdjecents(c coordinate) []adjecentResult {
	var result []adjecentResult
	for _, a := range c.adjecents() {
		adj := g.get(a)
		if adj != nil {
			result = append(result, adjecentResult{a, *adj})
		}
	}
	return result
}

type adjecentResult struct {
	loc   coordinate
	value byte
}

func (c coordinate) adjecents() []coordinate {
	var coords = []coordinate{
		{c.x - 1, c.y - 1},
		{c.x - 1, c.y},
		{c.x - 1, c.y + 1},
		{c.x, c.y - 1},
		{c.x, c.y + 1},
		{c.x + 1, c.y - 1},
		{c.x + 1, c.y},
		{c.x + 1, c.y + 1},
	}
	return coords
}

func (l coordinate) plus(r coordinate) coordinate {
	return coordinate{l.x + r.x, l.y + r.y}
}
func (l coordinate) times(i int) coordinate {
	return coordinate{l.x * i, l.y * i}
}
