package main

import (
	"advent/2024/grid"
	"fmt"
	"testing"
)

func TestPart1Simple(t *testing.T) {
	st := readData("simple")
	executeCommands(st)
	result := st.locats[grid.CreateCoordinate(1, 2)].symbol
	if result != 'O' {
		fmt.Println(string(st.locats[grid.CreateCoordinate(1, 2)].symbol))
		t.Fatal("not there")
	}
}

func TestPart1Example(t *testing.T) {
	st := readData("example")
	executeCommands(st)
	result := calculateGps(st.locats)
	if result != 10092 {
		t.Fatalf("Expected 10092 but got %d", result)
	}
}

func TestPart2Example(t *testing.T) {
	st := readDataPart2("example")
	executeCommandsBoxes(st)
	result := calcSumPart2(st.locats)
    st.print()
	if result != 9021 {
		t.Fatalf("Expected 9021 but got %d", result)
	}
}
func TestPart2Simple(t *testing.T) {
	st := readDataPart2("simple")
	executeCommandsBoxes(st)
	result := st.locats[grid.CreateCoordinate(1, 2)].symbol
	resultright := st.locats[grid.CreateCoordinate(2, 2)].symbol
	if result != '[' && resultright == ']' {
		fmt.Println(string(st.locats[grid.CreateCoordinate(1, 2)].symbol))
		t.Fatal("not there")
	}
}
func TestPart2SimpleWithPush(t *testing.T) {
	st := readDataPart2("example2")
	executeCommandsBoxes(st)
}
