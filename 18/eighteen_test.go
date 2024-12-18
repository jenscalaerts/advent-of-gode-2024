package main

import (
	"advent/2024/grid"
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	data := readData("example")
	result, res := findRoute(data, 7, 7, 12)
	for y := range 7 {
		for x := range 7 {
            fmt.Print("\t")
            fmt.Print(res[grid.CreateCoordinate(x,y)])
		}
        fmt.Println()
	}
	if result != 22 {
		t.Fatalf("expected %d but got %d", 22, result)
	}

}


func TestExamplePart2(t *testing.T) {
	data := readData("example")
	result:= part2(data, 7,7)
    expected := grid.CreateCoordinate(6,1)
	if result != expected {
		t.Fatalf("expected %v but got %v", expected, result.X)
	}

}
