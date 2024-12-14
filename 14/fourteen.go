package main

import (
	"advent/2024/grid"
	"advent/2024/parsing"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	robots := readFile("data")
	room := room{101, 103}
	locs := locationsAfter(100, room, robots)
	quadrants := countQuadrants(locs, room)
	fmt.Println(calculateSafetyFactor(quadrants))
    
    for i :=99;i < 100000; i+=101 {
        
        grid := robotGrid{}
	    locs := locationsAfter(i, room, robots)
        for _, loc := range locs {
            grid[loc.Y][loc.X] = 1
        }
         
        fmt.Println()
        fmt.Println("-------", i, "-----")
        grid.print()
    }

}


type robotGrid [103][101]byte
func (r *robotGrid)print(){
    for _,row := range r{
        for _, s := range row{
            if s == 0 {
                fmt.Print(" ")
            }else{
                fmt.Print("O")
            }
        }
        fmt.Println()
        
    }
}

type robot struct {
	start    grid.Coordinate
	velocity grid.Coordinate
}
type room struct {
	width, length int
}

type robotLocation grid.Coordinate

func calculateSafetyFactor(counts [5]int) int {
	factor := 1
	for _, val := range counts[1:] {
		factor *= val
	}
	return factor
}

func countQuadrants(locations []robotLocation, room room) [5]int {
	var counts [5]int
	for _, loc := range locations {
		quad := loc.getQuadrant(room)
		counts[quad]++
	}
	return counts
}

func locationsAfter(time int, room room, robots []robot) []robotLocation {
	locations := make([]robotLocation, len(robots))
	for i, robot := range robots {
		locations[i] = robot.calculatePositionAfter(time, room)
	}
	return locations
}

func readFile(name string) []robot {
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewScanner(f)
	robots := make([]robot, 0)
	for reader.Scan() {
		row := reader.Text()
		delim := strings.Index(row, " ")
		robots = append(robots, robot{readTuple(row[2:delim]), readTuple(row[delim+3:])})
	}
	return robots
}

func readTuple(s string) grid.Coordinate {
	comma := strings.Index(s, ",")
	x := parsing.Atoi(s[:comma])
	y := parsing.Atoi(s[comma+1:])
	return grid.Coordinate{X: x, Y: y}
}

func (rob robot) calculatePositionAfter(time int, room room) robotLocation {
	x := (time*rob.velocity.X + rob.start.X) % room.width
	if x < 0 {
		x = room.width + x
	}
	y := (time*rob.velocity.Y + rob.start.Y) % room.length
	if y < 0 {
		y = room.length + y
	}
	return robotLocation{x, y}

}

func (loc robotLocation) getQuadrant(r room) int {
	if loc.X == r.middleX() || loc.Y == r.middleY() {
		return 0
	}
	if loc.X < r.middleX() && loc.Y < r.middleY() {
		return 1
	}
	if loc.X > r.middleX() && loc.Y < r.middleY() {
		return 2
	}

	if loc.X < r.middleX() && loc.Y > r.middleY() {
		return 3
	}

	return 4
}

func (r room) middleX() int {
	return r.width / 2
}

func (r room) middleY() int {
	return r.length / 2
}
