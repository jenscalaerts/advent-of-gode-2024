package main

import (
	"advent/2024/grid"
	"fmt"
	"testing"
)

func TestExample(t *testing.T){
    g := grid.ReadGrid("example")
    part1, part2 := calculateMinScore(g)
    fmt.Println(part1)
    fmt.Println(part2)
}


