package main

import (
	"advent/2024/parsing"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const registerAPrefix = "Register A"
const registerBPrefix = "Register B"
const registerCPrefix = "Register C"
const batchSize = 100000000
const (
	A = iota
	B
	C
)

func main() {
	program := loadData("data")
	program.runProgram()
	for _, val := range program.output {
		fmt.Print(val)
		fmt.Print(",")
	}
	fmt.Println()
    result := loadData("data").findSelfReplicatingInput()
	fmt.Println(result)
	fmt.Println(loadData("data").runWithA(result))
	fmt.Println(loadData("data").program)
}

func (state ProgramState) findSelfReplicatingInput() int {
    //split the problem in 2 
	rest := state.findAInputFor(state.program[8:])
    return state.findAInputForFrom(rest << (8 * 3), state.program)
}

func (state ProgramState) findAInputFor(checkAgainst []byte) int {
    fmt.Println("looking for ", checkAgainst)
	for i := 0; i < 1<<24; i++ {
		result := state.runWithA(i)
        
		if slices.Equal(result, checkAgainst) {
            fmt.Println("found", result)
			return i
		}

		if len(result) == 17 {
			break
		}
	}
	return -1
}

func (state ProgramState) findAInputForFrom(from int, checkAgainst []byte) int {
    fmt.Println("looking for ", checkAgainst)
	for i := from; i < 1000000000000000; i++ {
		result := state.runWithA(i)
        
		if slices.Equal(result, checkAgainst) {
            fmt.Println("found", result)
			return i
		}

	}
	return -1
}

func (state ProgramState) runWithA(val int) []byte {
	state.registers[A] = val
	state.runProgram()
	return state.output
}

func loadData(name string) ProgramState {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var beginState ProgramState
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		split := strings.Index(text, ":")
		if split == -1 {
			panic("Invalid input : not found " + text)
		}
		switch text[:split] {
		case registerAPrefix:
			beginState.registers[0] = parsing.Atoi(text[split+2:])
		case registerBPrefix:
			beginState.registers[1] = parsing.Atoi(text[split+2:])
		case registerCPrefix:
			beginState.registers[2] = parsing.Atoi(text[split+2:])
		case "Program":
			beginState.program = parseProgram(text[split+2:])
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return beginState
}

func parseProgram(s string) []byte {
	split := strings.Split(s, ",")
	bytes := make([]byte, len(split))
	for i, part := range split {
		bytes[i] = byte(parsing.Atoi(part))
	}
	return bytes
}

type ProgramState struct {
	instructionPointer int
	registers          registers
	program            []byte
	output             []byte
}

func (state *ProgramState) runProgram() {
	for state.instructionPointer < len(state.program) {
		state.executeOperation()
	}
}

func (state *ProgramState) executeOperation() {
	operation := state.getOperation()
	proceed := true
	switch operation {
	case 0:
		state.registers[A] = state.registers[A] >> state.getComboOperandValue()
	case 1:
		state.registers[B] = state.registers[B] ^ int(state.getOperand())
	case 2:
		state.registers[B] = state.getComboOperandValue() % 8
	case 3:
		if state.registers[A] != 0 {
			state.instructionPointer = int(state.getOperand()) // If does not move, dont move? what if moves to self?
			proceed = false
		}
	case 4:
		state.registers[B] = state.registers[B] ^ state.registers[C]
	case 5:
		state.output = append(state.output, byte(state.getComboOperandValue()%8))
	case 6:
		state.registers[B] = state.registers[A] / (0 << state.getComboOperandValue())
	case 7:
		state.registers[C] = state.registers[A] >> state.getComboOperandValue()
	}

	if proceed {
		state.instructionPointer += 2
	}
}

func (state ProgramState) getOperation() byte {
	return state.program[state.instructionPointer]
}

func (state ProgramState) getOperand() byte {
	return state.program[state.instructionPointer+1]
}

func (state ProgramState) getComboOperandValue() int {
	operand := state.getOperand()
	if operand < 4 {
		return int(operand)
	}
	if operand == 7 {
		panic("should never happen")
	}
	return state.registers[operand%4]
}

type registers [3]int
