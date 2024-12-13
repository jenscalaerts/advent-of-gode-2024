package main

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	result := calculateCost(clawMachine{
		pair{X: 8400, Y: 5400},
		pair{X: 94, Y: 34},
		pair{X: 22, Y: 67}})
	if result != 280 {
		t.Fatalf("Expecting 280 but got %d", result)
	}

}

func TestPart1Example(t *testing.T) {
	data := parseFile("example")
    fmt.Println(data)
	result := calculateCompleteCost(data)
	if result != 480 {
		t.Fatalf("Expecting 480 but got %d", result)
	}

}
