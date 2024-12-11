package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(brutelong(readData("data"), 25))
	fmt.Println(brutelong(readData("data"), 75))
}

func readData(filename string) map[int]int {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	line := string(file)
	split := strings.Split(line[:len(line)-1], " ")
	values := make(map[int]int, len(split))
	for _, val := range split {
		val, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		values[val]++
	}
	return values
}

func brutelong(data map[int]int, iterations int) int {
	for range iterations {
		data = blinkOnce(data)
	}
	total := 0
	for _, amount := range data {
		total += amount
	}

	return total
}

func blinkOnce(data map[int]int) map[int]int {
	nextDate := make(map[int]int, 1000)
	for val, amount := range data {
		for _, result := range blink(val) {
			nextDate[result] += amount
		}
	}
	return nextDate
}

func blink(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	numberOfDigits := math.Floor(math.Log10(float64(stone)))
	if int(numberOfDigits)%2 == 1 {
		shed := int(math.Pow(10, math.Floor(numberOfDigits/2)+1))
		left := stone / shed
		right := stone % shed
		return []int{left, right}
	}
	return []int{2024 * stone}
}
