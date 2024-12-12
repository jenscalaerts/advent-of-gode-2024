package main

import (
	"advent/2024/grid"
	"fmt"
)

var floodDirections = []grid.Coordinate{{X: 0, Y: -1}, {X: -1, Y: 0}}

func main() {
	g := grid.ReadGrid("data")
	fmt.Println(part1(g))
	fmt.Println(part2(g))
}

func part1(g grid.Grid) int {
	areas, _ := floodAndCalcAll(g)
	return calcSum(areas)
}

func calcSum(areas []areaDescription) int {

	var sum int
	for _, area := range areas {
		if area.size == 0 {
			continue
		}
		sum += area.circumference * area.size
	}
	return sum
}

func part2(g grid.Grid) int {
	amounts, floodGrid := floodAndCalcAll(g)
	for i, val := range amounts {
		val.circumference = 0
		amounts[i] = val
	}
	amounts = calcRunningEdges(floodGrid, amounts)
	for _, am := range amounts {
		if am.size != 0 {
			fmt.Printf("%v %d %d\n", string(am.symbol), am.size, am.circumference)
		}
	}
	return calcSum(amounts)
}

func calcRunningEdges(floodGrid intGrid, amounts []areaDescription) []areaDescription {

	for _, dir := range []grid.Coordinate{{X: 0, Y: -1}, {X: 0, Y: 1}} {
		for y := range len(floodGrid) {
			cur := 0
			for x := range len(floodGrid[0]) {
				coord := grid.Coordinate{X: x, Y: y}
				edgeValue := floodGrid.Get(coord)
				other := floodGrid.Get(coord.Plus(dir))

				if other == edgeValue { //not an edge
					edgeValue = 0
				}
				if cur != edgeValue {
					cur = edgeValue
					amounts[cur].circumference++
				}
			}
		}
	}

	for _, dir := range []grid.Coordinate{{X: 1, Y: 0}, {X: -1, Y: 0}} {
		for x := range len(floodGrid[0]) {
			cur := 0
			for y := range len(floodGrid) {
				coord := grid.Coordinate{X: x, Y: y}
				edgeValue := floodGrid.Get(coord)
				other := floodGrid.Get(coord.Plus(dir))

				if other == edgeValue { //not an edge
					edgeValue = 0
				}
				if cur != edgeValue {
					cur = edgeValue
					amounts[cur].circumference++
					fmt.Println("vert")
					fmt.Println(cur)
				}
			}
		}
	}
	return amounts
}

func floodAndCalcAll(g grid.Grid) ([]areaDescription, intGrid) {

	var nextArea int
	nextArea++
	areas := make([]areaDescription, 1000000)
	floodGrid := make([][]int, len(g))
	for i := range len(g) {
		floodGrid[i] = make([]int, len(g[0]))
	}
	for _, coord := range g.Coordinates() {
		val := g.Get(coord)
		nextArea = floodAndCalc(floodGrid, coord, areas, nextArea, val)
	}
	return areas, floodGrid
}

func floodAndCalc(g intGrid, coord grid.Coordinate, areas []areaDescription, end int, val byte) int {

	value := val
	for _, dir := range floodDirections {
		flood := g.Get(coord.Plus(dir))
		area := areas[flood]
		if flood == 0 || area.symbol != value {
			continue
		}
		edges := 2 // remove 1 from touching and 1 from self
		edgeValue := g.Get(coord.Plus(grid.Coordinate{X: dir.Y, Y: dir.X}))
		if edgeValue == flood {
			edges -= 2
		} else if areas[edgeValue].symbol == value {
			edges -= 2
			for _, oldCoord := range g.FindAll(edgeValue) {
				g.Set(oldCoord, flood)
			}
			area = area.add(areas[edgeValue].size, areas[edgeValue].circumference)
			areas[edgeValue] = areaDescription{}
		}
		areas[flood] = area.add(1, edges)
		g.Set(coord, flood)
		return end
	}
	//circumference 4 nothing top left or right, the rest is subtracted later
	areas[end] = areaDescription{size: 1, circumference: 4, symbol: value}
	g.Set(coord, end)
	end++
	return end
}

type areaDescription struct {
	size, circumference int
	symbol              byte
}

func (a areaDescription) add(size, circumference int) areaDescription {
	return areaDescription{a.size + size, a.circumference + circumference, a.symbol}
}

type intGrid [][]int

func (g intGrid) Get(c grid.Coordinate) int {
	if !g.IsInGrid(c) {
		return 0
	}
	return g[c.Y][c.X]
}

func (g intGrid) Set(c grid.Coordinate, val int) {
	g[c.Y][c.X] = val
}

func (g intGrid) FindAll(b int) []grid.Coordinate {
	var coords []grid.Coordinate
	for y, row := range g {
		for x, el := range row {
			if el == b {
				coords = append(coords, grid.Coordinate{X: x, Y: y})
			}
		}
	}
	return coords
}
func (g intGrid) IsInGrid(c grid.Coordinate) bool {
	return c.X >= 0 && c.X < len(g[0]) && c.Y >= 0 && c.Y < len(g)
}
