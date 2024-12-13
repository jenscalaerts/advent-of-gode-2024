package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	machines := parseFile("data")
	fmt.Println(calculateCompleteCost(machines))
	for i := range machines {
		machines[i].end.X += 10000000000000
		machines[i].end.Y += 10000000000000
	}

	fmt.Println(calculateCompleteCost(machines))
}

type clawMachine struct {
	end   pair
	aStep pair
	bStep pair
}

func parseFile(name string) []clawMachine {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	content := make([]string, 0)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	clawMachines := []clawMachine{}
	for i := 0; i < len(content); i += 4 {
		a := parseButtonLine(content[i])
		b := parseButtonLine(content[i+1])
		prize := parseButtonLine(content[i+2])
		clawMachines = append(clawMachines, clawMachine{prize, a, b})

	}
	return clawMachines

}

func parseButtonLine(line string) pair {
	secondPart := line[strings.Index(line, ": ")+2:]

	split := strings.Split(secondPart, ", ")
	x, err := strconv.ParseInt(split[0][2:], 10, 64)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(split[1][2:], 10, 64)
	if err != nil {
		panic(err)
	}
	return pair{X: x, Y: y}
}


func calculateCompleteCost(machines []clawMachine) int64 {
	var sum int64
	for _, machine := range machines {
		cost := calculateCost(machine)
		if cost > 0 {
			sum += cost
		}
	}
	return sum
}

func calculateCost(c clawMachine) int64 {
	bnom := (c.aStep.X * c.end.Y) - (c.aStep.Y * c.end.X)
	denom := (c.bStep.Y * c.aStep.X) - (c.bStep.X * c.aStep.Y)
	if denom == 0 || c.aStep.X == 0 {
		panic("not solvable with equation")
	}
	if bnom%denom != 0 {
		return -1
	}
	B := bnom / denom
	anom := c.end.X - (c.bStep.X * B)
	if anom%c.aStep.X != 0 {
		return -1
	}
	A := anom / c.aStep.X
	return A*3 + B
}

type pair struct {
    X,Y int64
}
