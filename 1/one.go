package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	result, err := readFromFile("data")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("The part 1 result is '%d'", result.calculateDistance())
	fmt.Printf("The part 2 result is '%d'", result.calculateSimilarity())

}


func readFromFile(filename string) (*locationEntries, error) {

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return readFromString(string(content))

}

func readFromString(s string) (*locationEntries, error) {

	lineSplit := strings.Split(s, "\n")
	locationEntries := locationEntries{}

	for _, line := range lineSplit {
		if line == "" {
			continue
		}
		split := strings.Split(line, "   ")
		if len(split) != 2 {
			return nil, errors.New(fmt.Sprintf("Invalid input line='%v' %v", line, split))
		}
		left, err := strconv.ParseInt(split[0], 0, 0)
		if err != nil {
			return nil, fmt.Errorf("Can not parse int '%s'  %w", split[0], err)
		}

		right, err := strconv.ParseInt(split[1], 0, 0)
		if err != nil {
			return nil, fmt.Errorf("Can not parse int '%s'  %w", split[1], err)
		}
		locationEntries.add(int(left), int(right))

	}
	return  &locationEntries, nil;
}

type locationEntries struct {
	left  []int
	right []int
}

func (l *locationEntries) add(left, right int) {
	l.left = append(l.left, left)
	l.right = append(l.right, right)
}

func (l locationEntries) calculateDistance() int {
	slices.Sort(l.left)
	slices.Sort(l.right)
	sum := 0
	for i := range len(l.left) {
		sum = sum + abs(l.left[i]-l.right[i])
	}
	return sum
}

func (l locationEntries) calculateSimilarity() int {
	leftMap := countOccurences(l.left)
	rightMap := countOccurences(l.right)

	sum := 0
	for key := range leftMap {
		sum = sum + leftMap[key]*key*rightMap[key]

        fmt.Println(sum)
	}
	return sum
}

func countOccurences(numbers []int) map[int]int {
	m := make(map[int]int)
	for _, val := range numbers {
		m[val]++
	}
	return m
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
