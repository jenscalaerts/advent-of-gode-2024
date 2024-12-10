package main

import (
	"advent/2024/grid"
	"testing"
)

func TestExamplePart1(t *testing.T){
    example := grid.ReadGrid("example")
    result := findTrailHeadScores(example)
    if result != 36 {
        t.Fatalf("Unexpected result %d expected 36",result) 
    }
}


func TestExamplePart2(t *testing.T){
    example := grid.ReadGrid("example")
    result := findTrailHeadScoresAllRoutes(example)
    if result != 81 {
        t.Fatalf("Unexpected result %d expected 81",result) 
    }
}
