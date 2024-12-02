package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	unknown  = 0
	negative = -1
	positive = 1
	maxDif   = 3
)

func main() {
	data := readFile("data")
	result1 := calculateNumberOfSave(data, isReportSafeNoDampner)
	result2 := calculateNumberOfSave(data, isReportSafeWithDampner)
	fmt.Printf("Part 1: %v", result1)
	fmt.Printf("Part 2: %v", result2)
}

func readFile(name string) [][]int {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open file %v", err)
	}
	scanner := bufio.NewScanner(file)
	var data [][]int
	for scanner.Scan() {

		text := scanner.Text()
		split := strings.Split(text, " ")
		var report []int
		for _, num := range split {
			el, err := strconv.ParseInt(num, 0, 0)
			if err != nil {
				log.Fatalf("failed to open file %v", err)
			}
			report = append(report, int(el))

		}
		data = append(data, report)
	}

	return data
}

type safeCheck func([]int) bool

func calculateNumberOfSave(data [][]int, s safeCheck) int {
	total := 0
	for _, report := range data {
		safe := s(report)
		fmt.Printf("report %t\n", safe)
		if safe {
			total++
		}
	}
	return total
}

func isReportSafeNoDampner(report []int) bool {
	sign := toSign(report[0] - report[1])
	return isReportSafe(report, sign)
}

func isReportSafe(report []int, sign int) bool {
	fmt.Printf("testing %v %d \n", report, sign)
	previous := report[0]
	for _, el := range report[1:] {
		if !isSafe(previous, el, sign) {
			return false
		}
		previous = el
	}

	return true
}

func isReportSafeWithDampner(report []int) bool {
	previous := report[0]
	sign := toSign(report[0] - report[1])
	fmt.Printf("starting %v\n", report)
	for index, el := range report[1:] {
		if !isSafe(previous, el, sign) {
			if index == 0 {
                fmt.Println("start")
				return isReportSafeNoDampner(report[1:]) || isReportSafeNoDampner(append([]int{report[0]}, report[2:]...))
			}
			if index == len(report)-2 {
				return true
			}

            fmt.Println("middle")
            fmt.Println(report)
            return isReportSafeNoDampner(append(slices.Clone(report[:index-1]), report[index:]...)) ||
            isReportSafeNoDampner(append(slices.Clone(report[:index]), report[index+1:]...)) || isReportSafeNoDampner(append(report[:index+1], report[index+2:]...))
		}
		previous = el
	}

	return true
}

func isSafe(previous int, el int, sign int) bool {

	diff := previous - el
	if diff > maxDif || diff < -maxDif {
		fmt.Printf("TOO BIG diff %d prev %d cur %d ", diff, previous, el)
		return false
	}

	newSign := toSign(diff)

	if newSign == unknown {
		fmt.Println("NO CHANGE")
		return false
	}
	if newSign != sign {
		fmt.Println("SIGN CHANGE")
		return false
	}
	return true
}

func isSafeWithDampner(report []int) int {
	previous := report[0]
	sign := toSign(report[0] - report[1])
	for _, el := range report[1:] {
		diff := previous - el
		if diff > maxDif || diff < -maxDif {
			fmt.Printf("TOO BIG diff %d prev %d cur %d ", diff, previous, el)
			return 0
		}

		fmt.Printf("diff[%d]", diff)
		newSign := toSign(diff)

		fmt.Printf("sign %d newSign %d; ", sign, newSign)
		if newSign == unknown || newSign != sign {
			fmt.Println("SIGN CHANGE")
			return 0
		}
		previous = el
	}

	return 1
}

func toSign(diff int) int {
	if diff > 0 {
		return positive
	}
	if diff < 0 {
		return negative
	}
	return unknown
}
