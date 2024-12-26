package main

import (
	"advent/2024/parsing"
	"fmt"
)

func main(){
    result := getNumberOfOverlaps(readData("data"))
    fmt.Println(result)
}

func getNumberOfOverlaps(locks, keys [][5]int) int {
	possibleCombinations := 0
	for _, lock := range locks {
		for _, key := range keys {
			if !isOverlapping(lock, key) {
				possibleCombinations++
			}
		}
	}
	return possibleCombinations

}

func readData(name string) (locks, keys [][5]int) {
	lines := parsing.ReadLines(name)
	locks = [][5]int{}
	keys = [][5]int{}
	for i := 0; i < len(lines); i += 8 {
		if lines[i] == "" {
			continue
		}
		if lines[i] == "#####" {
			locks = append(locks, parseLock(lines[i:i+7]))
		} else {
			keys = append(keys, parseKey(lines[i:i+7]))
		}
	}
	return
}

func isOverlapping(lock, key [5]int) bool {
	for i := range lock {
		if lock[i]+key[i] > 5 {
			return true
		}
	}
	return false

}

func parseLock(lockLines []string) [5]int {
	return parseLength(lockLines[1:], '#')
}

func parseKey(lockLines []string) [5]int {
	lengths := parseLength(lockLines[:7], '.')
	for i, val := range lengths {
		lengths[i] = 6 - val

	}
	return lengths
}

func parseLength(lockLines []string, character rune) [5]int {

	var lengths [5]int
	for _, line := range lockLines {
		for i, char := range line {
			if char == character {
				lengths[i]++
			}
		}
	}
	return lengths
}
