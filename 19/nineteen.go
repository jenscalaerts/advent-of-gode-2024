package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	patterns, designs := readData("data")
    fmt.Println("part 1:", patterns.countMatching(designs))
    fmt.Println("Part 2:", patterns.totalNumberOfCombinations(designs))
}

type towelPatterns []string

func readData(name string) (towelPatterns, []string) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	patterns := towelPatterns(strings.Split(scanner.Text(), ", "))
	scanner.Scan()
	designs := []string{}
	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return patterns, designs
}

func (pattern towelPatterns) countMatching(designs []string) int {
	memo := map[string]bool{}
	count := 0
	for _, design := range designs {
		if pattern.canMake(design, memo) {
			count++
		}
	}
	return count
}

func (patterns towelPatterns) canMake(s string, memo map[string]bool) bool {
	result, found := memo[s]
	if found {

		return result
	}

	for _, pattern := range patterns {
		if strings.Index(s, pattern) != 0 {
			continue
		}
		if len(pattern) == len(s) {
			memo[s] = true
			return true
		}
		if patterns.canMake(s[len(pattern):], memo) {
			memo[s] = true
			return true
		}
	}
	memo[s] = false
	return false
}

func (pattern towelPatterns) totalNumberOfCombinations(designs []string) int {
	memo := map[string]int{}
	count := 0
	for _, design := range designs {
		count += pattern.possibleCombination(design, memo)
	}
	return count
}

func (patterns towelPatterns) possibleCombination(s string, memo map[string]int) int {
	result, found := memo[s]
	if found {
		return result
	}

	options := 0
	for _, pattern := range patterns {
		if strings.Index(s, pattern) != 0 {
			continue
		}
		if len(pattern) == len(s) {
			memo[s] = 1
			options++
		}
		options += patterns.possibleCombination(s[len(pattern):], memo)
	}
	memo[s] = options
	return options
}
