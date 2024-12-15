package main

import (
	"advent/2024/grid"
	"bufio"
	"fmt"
	"os"
	"slices"
)

var up = grid.CreateCoordinate(0, -1)
var down = grid.CreateCoordinate(0, 1)
var left = grid.CreateCoordinate(-1, 0)
var right = grid.CreateCoordinate(1, 0)

type locations map[grid.Coordinate]location

type location struct {
	symbol  rune
	canMove bool
}

type state struct {
	locats   locations
	robot    grid.Coordinate
	commands []grid.Coordinate
}
type box [2]grid.Coordinate

func main() {
	s := readData("data")
	executeCommands(s)
	fmt.Println(calculateGps(s.locats))
	s = readDataPart2("data")
    executeCommandsBoxes(s)
    fmt.Println(calcSumPart2(s.locats))
}

func calculateGps(s locations) int {
	var total int
	for coord, val := range s {
		if val.symbol == 'O' {
			total += 100*coord.Y + coord.X
		}
	}
	return total
}

func executeCommands(s state) {
	for _, command := range s.commands {
		movingLocations := []grid.Coordinate{s.robot}

		nextCoordinate := s.robot
		for {
			nextCoordinate = nextCoordinate.Plus(command)
			nextLocation, locationFound := s.locats[nextCoordinate]
			if !locationFound {
				s.locats.moveAll(movingLocations, command)
				s.robot = s.robot.Plus(command)
				break
			}
			if nextLocation.canMove {
				movingLocations = append(movingLocations, nextCoordinate)
			}
			if !nextLocation.canMove {
				break
			}
		}
	}
}

func executeCommandsBoxes(s state) {
	for _, command := range s.commands {
		nextCoordinate := s.robot.Plus(command)
		nextLocation, locationFound := s.locats[nextCoordinate]
		if !locationFound {
			s.robot = s.robot.Plus(command)
			continue
		}
		if !nextLocation.canMove {
			continue
		}
		touchingBlock := s.locats.findBlockAt(nextCoordinate)
		movingBlocks, canMove := touchingBlock.touchesEdge(s.locats, command)
		if canMove {
			s.locats.movingBlocks(movingBlocks, command)
			s.robot = s.robot.Plus(command)
		}
	}
}

func (l *locations) movingBlocks(boxes []box, direction grid.Coordinate) {
	locs := *l

	for _, b := range boxes {
		delete(locs, b[0])
		delete(locs, b[1])
	}
	for _, b := range boxes {
		locs[b[0].Plus(direction)] = location{'[', true}
		locs[b[1].Plus(direction)] = location{']', true}
	}

}

func (l locations) findBlockAt(c grid.Coordinate) box {
	val := l[c]
	if val.symbol == ']' {
		return box{c.Plus(left), c}
	} else {
		return box{c, c.Plus(right)}
	}

}

func (b box) touchesEdge(s locations, direction grid.Coordinate) ([]box, bool) {
	touchedBoxes := []box{b}
	for _, loc := range b.getNextLocations(direction) {
		val, filled := s[loc]
		if !filled {
			continue
		}
		if !val.canMove {
			return []box{}, false
		}
		var touchedBox box
		if val.symbol == '[' {
			touchedBox = box{loc, loc.Plus(right)}
		}
		if val.symbol == ']' {

			touchedBox = box{loc.Plus(left), loc}
		}
		ap, canMove := touchedBox.touchesEdge(s, direction)
		if !canMove {
			return touchedBoxes, false
		}
		touchedBoxes = append(touchedBoxes, ap...)
	}
	return touchedBoxes, true
}

func (b box) getNextLocations(direction grid.Coordinate) []grid.Coordinate {
	if direction.Y == 0 {
		if direction.X == 1 {
			return []grid.Coordinate{b[1].Plus(direction)}
		} else {
			return []grid.Coordinate{b[0].Plus(direction)}
		}
	} else {
		return []grid.Coordinate{b[0].Plus(direction), b[1].Plus(direction)}
	}
}

func (l *locations) moveAll(moving []grid.Coordinate, direction grid.Coordinate) {
	locs := *l
	for i := len(moving) - 1; i > 0; i-- {
		locs[moving[i].Plus(direction)] = locs[moving[i]]
	}
	if len(moving) > 1 {
		delete(locs, moving[1])
	}

}

func readData(name string) state {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	locations := map[grid.Coordinate]location{}
	row := 0
	var start grid.Coordinate
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		for col, val := range line {
			if val == '.' {
				continue
			}
			if val == '@' {
				start = grid.CreateCoordinate(col, row)
				continue
			}
			locations[grid.CreateCoordinate(col, row)] = createLocation(val)
		}
		row++
	}
	commands := []grid.Coordinate{}
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			var command grid.Coordinate
			switch r {
			case '<':
				command = left
			case '^':
				command = up
			case '>':
				command = right
			case 'v':
				command = down
			}
			commands = append(commands, command)
		}
	}

	return state{locations, start, commands}
}

func createLocation(symbol rune) location {
	if symbol == 'O' {
		return location{symbol, true}
	}
	if symbol == '#' {
		return location{symbol, false}
	}
	fmt.Println(string(symbol))
	panic("unexpected rune")
}

func readDataPart2(name string) state {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	locations := map[grid.Coordinate]location{}
	row := 0
	var start grid.Coordinate
	for scanner.Scan() {
		col := 0
		line := scanner.Text()
		if line == "" {
			break
		}
		for _, val := range line {
			if val == '.' {
				col += 2
				continue
			}
			if val == '@' {
				start = grid.CreateCoordinate(col, row)
				col += 2
				continue
			}
			if val == 'O' {
				locations[grid.CreateCoordinate(col, row)] = location{'[', true}
				locations[grid.CreateCoordinate(col+1, row)] = location{']', true}
			}
			if val == '#' {
				locations[grid.CreateCoordinate(col, row)] = location{'#', false}
				locations[grid.CreateCoordinate(col+1, row)] = location{'#', false}
			}
			col += 2

		}
		row++
	}
	commands := []grid.Coordinate{}
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			var command grid.Coordinate
			switch r {
			case '<':
				command = left
			case '^':
				command = up
			case '>':
				command = right
			case 'v':
				command = down
			}
			commands = append(commands, command)
		}
	}

	return state{locations, start, commands}
}

func calcSumPart2(locs locations) int {
	sum := 0
	for c, loc := range locs {
		if loc.symbol == '[' {
			sum += 100*c.Y + c.X
		}
	}
	return sum
}
