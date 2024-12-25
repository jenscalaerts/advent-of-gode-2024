package main

import (
	"advent/2024/parsing"
	"fmt"
	"slices"
	"strings"
)

const XOR = "XOR"
const OR = "OR"
const AND = "AND"

type operation struct {
	left, right, destination string
	act                      action
	actionDescription        string
}

func main() {
	fmt.Println(runOperationsWithRegisters(loadData("data")))
}

func runOperationsWithRegisters(ops []operation, registers map[string]bool) int {
	registers = execute(ops, registers)
	total := 0
	for register, val := range registers {
		if val && register[0] == 'z' {
			total += (1 << parsing.Atoi(register[1:]))
		}

	}
	return total
}

func execute(ops []operation, registers map[string]bool) map[string]bool {
	for slices.ContainsFunc(ops, operation.hasZResult) {
		remaining := ops[:0]
		for _, op := range ops {
			_, leftFilled := registers[op.left]
			_, rightFilled := registers[op.right]
			if leftFilled && rightFilled {
				registers[op.destination] = op.act(registers[op.left], registers[op.right])
			} else {
				remaining = append(remaining, op)
			}
		}
		if len(ops) == len(remaining) {
			for _, op := range ops {
				if op.hasZResult() {
					findFeeding(registers, ops, op.destination, 0)
					fmt.Println("\n===================")
				}
			}
			fmt.Println("Fail")
			break
		}

		ops = remaining

	}
	return registers

}

func findFeeding(registers map[string]bool, ops []operation, lookingFor string, depth int) {
	tabDepth(depth, lookingFor)

	found := false
	for _, op := range ops {
		if op.destination == lookingFor {
			found = true
			tabDepth(depth, fmt.Sprint(op))
			_, foundLeft := registers[op.left]
			if !foundLeft {
				tabCrDepth(depth, "left(\n")
				findFeeding(registers, ops, op.left, depth+1)
				tabCrDepth(depth, ")")
			} else {
				tabCrDepth(depth, fmt.Sprintf("Left in registers %v with value %v", op.left, registers[op.left]))
			}
			_, foundRight := registers[op.right]

			if !foundRight {
				tabCrDepth(depth, "right(\n")
				findFeeding(registers, ops, op.right, depth+1)
				tabCrDepth(depth, ")")
			} else {
				tabCrDepth(depth, fmt.Sprintf("right in registers %v with value %v", op.left, registers[op.left]))
			}
		}

	}
	if !found {
		tabDepth(depth, "Not found")
	}

}
func tabDepth(i int, s string) {
	for range i {
		fmt.Print(" ")
	}
	fmt.Print(s)
}

func tabCrDepth(i int, s string) {
	fmt.Print("\n")
	for range i {
		fmt.Print(" ")
	}
	fmt.Print(s)
}

func loadData(name string) ([]operation, map[string]bool) {
	lines := parsing.ReadLines(name)
	split := slices.Index(lines, "")
	initialValuesRegisters := map[string]bool{}
	for _, line := range lines[:split] {
		split := strings.Split(line, ": ")
		initialValuesRegisters[split[0]] = split[1] == "1"
	}

	operations := []operation{}
	for _, line := range lines[split+1:] {
		split := strings.Split(line, " ")
		var action action
		switch split[1] {
		case AND:
			action = and
		case OR:
			action = or
		case XOR:
			action = xor
		default:
			panic("unexpected operation")
		}
		operations = append(operations, operation{split[0], split[2], split[4], action, split[1]})

	}
	return operations, initialValuesRegisters

}

type action func(bool, bool) bool

func and(l, r bool) bool {
	return l && r
}

func or(l, r bool) bool {
	return l || r
}

func xor(l, r bool) bool {
	return l != r
}

func (o operation) hasZResult() bool {
	return o.destination[0] == 'z'
}
