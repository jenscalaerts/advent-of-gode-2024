package main

import (
	"slices"
	"testing"
)

func TestSimple_bst(t *testing.T) {
	state := ProgramState{program: []byte{2, 6},
		registers: registers{0, 0, 9}}
	state.runProgram()
	if state.registers[B] != 1 {
		t.Fatal("bad")
	}
}

func TestSimple_simpleExample(t *testing.T) {
	state := ProgramState{program: []byte{5, 0, 5, 1, 5, 4},
		registers: registers{10, 0, 0}}
	state.runProgram()
	expected := []byte{0, 1, 2}
	if !slices.Equal(state.output, expected) {
		t.Fatalf("expected %v but got %v", expected, state.output)
	}
}

func TestSimple_simpleExample2(t *testing.T) {
	state := ProgramState{program: []byte{0, 1, 5, 4, 3, 0},
		registers: registers{2024, 0, 0}}
	state.runProgram()

	if state.registers[A] != 0 {
		t.Fatalf("expected %v but got %v", 0, state.registers[A])
	}
	expected := []byte{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}
	if !slices.Equal(state.output, expected) {
		t.Fatalf("expected %v but got %v", expected, state.output)
	}
}

func TestSimple_simpleExample3(t *testing.T) {
	state := ProgramState{program: []byte{1, 7},
		registers: registers{0, 29, 0}}
	state.runProgram()

	if state.registers[B] != 26 {
		t.Fatalf("expected %v but got %v", 26, state.registers[A])
	}
}

func TestSimple_simpleExample4(t *testing.T) {
	state := ProgramState{program: []byte{4, 0},
		registers: registers{0, 2024, 43690}}
	state.runProgram()

	if state.registers[B] != 44354 {
		t.Fatalf("expected %v but got %v", 44354, state.registers[A])
	}
}

func TestExample(t *testing.T) {
	program := loadData("example")
	program.runProgram()
	expected := []byte{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}
	if !slices.Equal(expected, program.output) {
		t.Fatalf("expected %v but got %v", expected, program.output)
	}

}

func TestExample_part2(t *testing.T) {
	program := loadData("example_2")
    result := program.findSelfReplicatingInput()
    expected := 117440
	if expected != result{
		t.Fatalf("expected %v but got %v", expected, result)
	}
}
