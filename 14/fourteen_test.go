package main

import (
	"advent/2024/grid"
	"testing"
    "fmt"
)

func TestPart1Simple(t *testing.T) {
	rob := robot{start: grid.Coordinate{X: 2, Y: 4}, velocity: grid.Coordinate{X: 2, Y: -3}}
	room := room{11, 7}
	position := rob.calculatePositionAfter(1, room)
	expected := robotLocation{X: 4, Y: 1}
	if position != expected {
		t.Fatalf("Expected %v but wat %v", expected, position)
	}

}

func TestPart1WithTeleport(t *testing.T) {
	rob := robot{start: grid.Coordinate{X: 2, Y: 4}, velocity: grid.Coordinate{X: 2, Y: -3}}
	room := room{11, 7}
	position := rob.calculatePositionAfter(2, room)
    expected := robotLocation{X: 6, Y: 5}
    if position != expected {
		t.Fatalf("Expected %v but wat %v", expected, position)
	}

}


func TestPart1Example(t *testing.T) {
	robots := readFile("example")
	room := room{11, 7}
    locs:=locationsAfter(100, room, robots)
    fmt.Println(locs)
    result := countQuadrants(locs, room)
    expect := [5]int{3,1,3,4,1}
    if expect != result {
        t.Fatal("Expected", expect, "but got", result)
    }
    factor := calculateSafetyFactor(result)
    if factor != 12 {
        t.Fatal("Expected", 12, "but got", factor)
    }
    
}

