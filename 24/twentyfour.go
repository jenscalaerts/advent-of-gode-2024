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
}

func main(){
    fmt.Println(part1(loadData("data")))
}

func part1(ops []operation, registers map[string]bool) int {
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
		ops = remaining
	}
	return registers

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
		operations = append(operations, operation{split[0], split[2], split[4], action})

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
