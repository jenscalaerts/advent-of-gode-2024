package main

import (
	"advent/2024/grid"
	"fmt"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	g := grid.ReadGrid("example")
	base := calculateBase(g)
	cheats := findCheats(base)

	cheatCounts := map[int]int{}
	for _, cheat := range cheats {
		cheatCounts[cheat.score]++
	}
	fmt.Println(cheatCounts)

}

func TestExamplePart2(t *testing.T) {
	g := grid.ReadGrid("example")
	base := calculateBase(g)
	cheats := findLongCheats(base)

	cheatCounts := map[int]int{}
	for _, cheat := range cheats {
		if cheat.score >= 50 {
			cheatCounts[cheat.score]++
		}
	}
	fmt.Println(cheatCounts)
}

func TestGeneration(t *testing.T) {
	fmt.Println(calculatePossibleMoves(2))
}
