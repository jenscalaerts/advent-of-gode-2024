package grid

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Coordinate struct {
	X, Y int
}

func (c Coordinate) RotateRight() Coordinate {
	return Coordinate{-c.Y, c.X}
}

type Grid []string

func (g Grid) Get(c Coordinate) byte {
	if !g.IsInGrid(c) {
		return 0
	}
	val := g[c.Y][c.X]
	return val
}

func (g Grid) FindAll(b byte) []Coordinate {
	var coords []Coordinate
	for y, row := range g {
		for x, el := range []byte(row) {
			if el == b {
				coords = append(coords, Coordinate{x, y})
			}
		}
	}
	return coords
}

func (l Coordinate) Plus(r Coordinate) Coordinate {
	return Coordinate{l.X + r.X, l.Y + r.Y}
}
func ReadGrid(name string) Grid {
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

func (g Grid) Copy() Grid {
	return slices.Clone(g)
}

func (g *Grid) Set(c Coordinate, val byte) {
	gr := *g
	gr[c.Y] = gr[c.Y][:c.X] + string(val) + gr[c.Y][c.X+1:]
}
func (g Grid) Print() {
	for _, row := range g {
		fmt.Println([]byte(row))
	}
	fmt.Println()
}
func (g Grid) AsValueCoordinateMap() map[byte][]Coordinate {
	result := make(map[byte][]Coordinate, 10)
	for y, row := range g {
		for x, val := range []byte(row) {
			result[val] = append(result[val], Coordinate{x, y})
		}
	}
	return result
}

func (g Grid) IsInGrid(c Coordinate) bool {
	return c.X >= 0 && c.X < len(g[0]) && c.Y >= 0 && c.Y < len(g)
}

func (l Coordinate) Minus(r Coordinate) Coordinate {
	return Coordinate{l.X - r.X, l.Y - r.Y}
}

func (g Grid) GetCardinalAdjecents(c Coordinate) []AdjecentResult {
	result := []AdjecentResult{}
	for _, coord := range c.adjecents() {
		symbol := g.Get(coord)
		if symbol != 0 {
			result = append(result, AdjecentResult{coord, symbol})
		}
	}
	return result
}

type AdjecentResult struct {
	Loc   Coordinate
	Value byte
}

func (c Coordinate) adjecents() []Coordinate {
	var coords = []Coordinate{
		{c.X - 1, c.Y},
		{c.X + 1, c.Y},
		{c.X, c.Y - 1},
		{c.X, c.Y + 1},
	}
	return coords
}
func (g Grid) Coordinates() []Coordinate {
	coordinates := make([]Coordinate, len(g)*len(g[0]))
	var i int
	for y := range len(g) {
		for x := range len(g[0]) {
			coordinates[i] = Coordinate{x, y}
			i++
		}
	}
	return coordinates
}

func (g Grid) EmptyCopy() Grid {
    data := []string{}
	for range len(g) {
        data = append(data, string(make([]byte, len(g[0]))))
	}
    return Grid(data)
}

func CreateCoordinate(x,y int) Coordinate{
    return Coordinate{x,y}
}

func (left Coordinate) DotProduct(right Coordinate)int{
    return left.X * right.X + left.Y * right.Y
}

