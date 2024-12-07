package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	calibrations := readCalibrations("data")
	fmt.Println(sumValidCalibrations(calibrations))
	fmt.Println(sumValidCalibrationsWithConcat(calibrations))
}

type calibration struct {
	total   int
	numbers []int
}

func isPossibleFormula(expectedTotal, total int, numbers []int) bool {
	if len(numbers) == 0 {
		return expectedTotal == total
	}
	if total > expectedTotal {
		return false
	}

	return isPossibleFormula(expectedTotal, total*numbers[0], numbers[1:]) ||
		isPossibleFormula(expectedTotal, total+numbers[0], numbers[1:])
}

func isPossibleFormulaWithConcat(expectedTotal, total int, numbers []int) bool {
	if len(numbers) == 0 {
		return expectedTotal == total
	}
	if total > expectedTotal {
		return false
	}

	return isPossibleFormulaWithConcat(expectedTotal, total*numbers[0], numbers[1:]) ||
		isPossibleFormulaWithConcat(expectedTotal, total+numbers[0], numbers[1:]) ||
		isPossibleFormulaWithConcat(expectedTotal, convertString(fmt.Sprint(total)+fmt.Sprint(numbers[0])), numbers[1:])
}

func readCalibrations(name string) []calibration {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []calibration
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Index(text, ":")
		data = append(data, calibration{
			total:   convertString(text[:split]),
			numbers: parseNumbers(text[split+2:]),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func parseNumbers(s string) []int {
	split := strings.Split(s, " ")
	data := make([]int, len(split))
	for i, num := range split {
		data[i] = convertString(num)
	}
	return data
}

func convertString(s string) int {

	val, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	return int(val)
}

func sumValidCalibrations(calibrations []calibration) int {
	total := 0
	for _, calibration := range calibrations {
		if isPossibleFormula(calibration.total, 0, calibration.numbers) {
			total += calibration.total
		}
	}
	return total
}

func sumValidCalibrationsWithConcat(calibrations []calibration) int {
	total := 0
	for _, calibration := range calibrations {
		if isPossibleFormulaWithConcat(calibration.total, 0, calibration.numbers) {
			total += calibration.total
		}
	}
	return total
}
