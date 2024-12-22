package main

import (
	"advent/2024/parsing"
	"fmt"
)

func main() {
	fmt.Println(part1("data"))
    result, _ :=part2("data")
	fmt.Println(result)
}

func mix(value, secret int) int {
	return value ^ secret
}

func prune(secret int) int {
	return secret % 16777216
}

func advance(secret int) int {
	secret = prune(mix(secret*64, secret))
	secret = prune(mix(secret/32, secret))
	secret = prune(mix(secret*2048, secret))
	return secret
}

func readData(name string) []int {
	lines := parsing.ReadLines(name)
	data := make([]int, len(lines))
	for i, line := range lines {
		data[i] = parsing.Atoi(line)
	}
	return data
}

func part1(name string) int {
	secrets := readData(name)
	for range 2000 {
		for i, secret := range secrets {
			secrets[i] = advance(secret)
		}
	}
	sum := 0
	for _, secret := range secrets {
		sum += secret
	}
	return sum
}

func part2(name string) (int, map[string][]int) {
	s := readData(name)
    return part2WithData(s)
}

func part2WithData(s []int) (int, map[string][]int) {
	prices, diffs := calculatePrices(s)
	priceByChanges := map[string][]int{}
	for i, row := range diffs {
		for j := 3; j < len(row); j++ {
			d := row[j-3 : j+1]
			key := fmt.Sprint(d)
			_, found := priceByChanges[key]
			if !found {
				priceByChanges[key] = initializeAtMinus1(len(s))
			}
			if priceByChanges[key][i] == -1 {
				priceByChanges[key][i] = prices[i][j+1]
			}
		}
	}

	return calculateMaxFullPrice(priceByChanges), priceByChanges
}

func calculateMaxFullPrice(prices map[string][]int) int {
	max := 0
	for _, val := range prices {
		sum := 0
		for _, price := range val {
			if price != -1 {
				sum += price
			}
		}
		if max < sum {
			max = sum
		}
	}
	return max

}

func initializeAtMinus1(length int) []int {
	slice := make([]int, length)
	for i := range slice {
		slice[i] = -1
	}
	return slice

}

func calculatePrices(secrets []int) ([][2001]int, [][2000]int) {
	prices := make([][2001]int, len(secrets))
	diffs := make([][2000]int, len(secrets))
	for i, secret := range secrets {
		sec := secret
		prices[i][0] = sec % 10
		for j := range 2000 {
			sec = advance(sec)
			prices[i][j+1] = sec % 10
			diffs[i][j] = prices[i][j+1] - prices[i][j]
		}
	}
	return prices, diffs
}
