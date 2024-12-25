package main

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

/*
	func TestPart1Example(t *testing.T) {
		operations, registers := loadData("example")
		result := runOperationsWithRegisters(operations, registers)
		if result != 2024 {
			t.Errorf("expected 2024 but got %d", result)
		}
	}
*/
func TestPart2Exploration(t *testing.T) {
	operations, _ := loadData("data")
	registers := map[string]bool{}

	for i := range 45 {
		registers["x"+fmt.Sprintf("%02d", i)] = false
		registers["y"+fmt.Sprintf("%02d", i)] = false
	}
	failures := map[int]bool{}

	findFeeding2(operations, "z"+fmt.Sprintf("%02d", 1), 0)
	findFeeding2(operations, "z"+fmt.Sprintf("%02d", 2), 0)
	for i := range 45 {
		local := maps.Clone(registers)

		local["x"+fmt.Sprintf("%02d", i)] = true
		local["y"+fmt.Sprintf("%02d", i)] = true
		result := runOperationsWithRegisters(slices.Clone(operations), local)
		if result != (1 << (i + 1)) {
			fmt.Println("iteration", i, "both expected", (1 << (i + 1)), "got", result)
			failures[i] = true
			findFeeding2(operations, "z"+fmt.Sprintf("%02d", i+1), 0)

		}

		local = maps.Clone(registers)

		usedRegister := "x" + fmt.Sprintf("%02d", i)
		local[usedRegister] = true
		result = runOperationsWithRegisters(slices.Clone(operations), local)
		if result != (1 << i) {
			failures[i] = true
			fmt.Println("iteration", i, "x only expected", (1 << i), "got", result)
			findOperationsResultingFrom(operations, usedRegister, 0)
		}

		local = maps.Clone(registers)
		usedRegister = "y" + fmt.Sprintf("%02d", i)
		local[usedRegister] = true
		result = runOperationsWithRegisters(slices.Clone(operations), local)
		if result != (1 << i) {
			failures[i] = true
			fmt.Println("iteration", i, "y only expected", (1 << i), "got", result)
			findOperationsResultingFrom(operations, usedRegister, 0)
		}

	}
	// fmt.Println(failures)
}

func findOperationsResultingFrom(ops []operation, register string, depth int) {
	for _, op := range ops {
		if op.left == register || op.right == register {
			tabCrDepth(depth, fmt.Sprint(op))
			findOperationsResultingFrom(ops, op.destination, depth+1)
		}
	}
	fmt.Println()
}

func findFeeding2(ops []operation, register string, depth int) {
	for _, op := range ops {
		if register == op.destination {
			tabCrDepth(depth, fmt.Sprint(op.left," ", op.actionDescription," ", op.right," ",op.destination))

			findFeeding2(ops, op.left, depth+1)

			findFeeding2(ops, op.right, depth+1)
		}
	}
    if depth == 0 {
        fmt.Println()
    }
}
