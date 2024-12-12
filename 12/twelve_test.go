package main

import (
	"advent/2024/grid"
	"fmt"
	"testing"
)

func TestPart1Example(t *testing.T){
    g :=grid.ReadGrid("example") 
    result := part1(g)
    fmt.Println(result)
}


func TestPart2Example(t *testing.T){
    g :=grid.ReadGrid("example") 
    result := part2(g)
    if result != 1206{
    t.Fatal(result)
    }
}

func TestPartSimple(t *testing.T){
    g := grid.Grid{"A"}
    result := part2(g)
    if result != 4 {
        t.Fatalf("expected 4 got %d", result)
    }
}

func TestPartSimple2(t *testing.T){
    g := grid.Grid{"AA","BB"}
    result := part2(g)
    if result != 16 {
        t.Fatalf("expected 16 got %d", result)
    }
}

func TestPartSimple3(t *testing.T){
    g := grid.Grid{"AB","BB"}
    result := part2(g)
    if result != 22 {
        t.Fatalf("expected 22 got %d", result)
    }
}
