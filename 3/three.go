package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const beginToken = "mul("
const endToken = ")"
const comma = ","

func main() {
	file, err := os.ReadFile("data")
	if err != nil {
		log.Fatal(err)
	}
    part1 := calculate(string(file))
    part2 :=  calculateWithConditional(string(file))

	fmt.Printf("The part 1 total is '%d'\n", part1 )
	fmt.Printf("The part 2 total is '%d'\n", part2)
}

func calculateWithConditional(s string) int {
	total := 0
	for {
		index := strings.Index(s, "don't()")
		if index == -1 {
			return total + calculate(s)
		}
        total += calculate(s[:index])
        s = s[index+7:]
        index = strings.Index(s, "do()") 
        if index == -1 {
            return total
        }
        s = s[index+4:]
	}
}

func calculate(s string) int {
    //Don't want to regex today
	total := 0
	for {
		index := strings.Index(s, beginToken)
		if index == -1 {
			break
		}
		s = s[index+4:]
		comma := strings.Index(s, comma)
		if comma == -1 {
			break
		}
		if comma > 3 {
			continue
		}
		firstNumber, err := strconv.ParseInt(s[:comma], 0, 0)
		if err != nil {
			continue
		}
		s = s[comma+1:]
		end := strings.Index(s, endToken)
		if end > 3 {
			continue
		}
		if end == -1 {
			break
		}
		secondNumber, err := strconv.ParseInt(s[:end], 0, 0)
		if err != nil {
			continue
		}
		s = s[end+1:]

		total += int(firstNumber) * int(secondNumber)
	}
	return total
}
