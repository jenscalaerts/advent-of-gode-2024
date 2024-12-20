package main

import (
	"advent/2024/grid"
	"fmt"
)

func main() {
	g := grid.ReadGrid("data")
	base := calculateBase(g)
	cheats := findCheats(base)
	fmt.Println(calcScore(cheats))
	fmt.Println(calcScore(findLongCheats(base)))
}

func calcScore(cheats []cheatScore)int{
	count := 0
	for _, cheat := range cheats {
		if cheat.score >= 100 {
			count++
		}
	}
    return count
}

func calculateBase(g grid.Grid) precalculations {
	startLocation := g.FindAll('S')[0]
	distancesFromStart := calculateDistances(startLocation, g)
	endLocation := g.FindAll('E')[0]
	distancesFromEnd := calculateDistances(endLocation, g)
	return precalculations{
		begin:     startLocation,
		end:       endLocation,
		fromStart: distancesFromStart,
		fromEnd:   distancesFromEnd,
	}
}

func findCheats(calcs precalculations) []cheatScore {
	baseDistance := calcs.fromStart[calcs.end]
	foundCheats := []cheatScore{}

	for coordinate, fromStart := range calcs.fromStart {
		for _, cheat := range calculatePossibleMoves(2) {
			fromEnd, found := calcs.fromEnd[coordinate.Plus(cheat)]
			if !found {
				continue
			}
			length := fromStart + fromEnd + 2
			if baseDistance > length {
				cs := cheatScore{coordinate, coordinate.Plus(cheat), baseDistance - length}
				foundCheats = append(foundCheats, cs)
			}
		}
	}
	return foundCheats
}

func findLongCheats(calcs precalculations) []cheatScore {
	baseDistance := calcs.fromStart[calcs.end]
	foundCheats := []cheatScore{}
	possibleMoves := calculatePossibleMoves(20)
	for coordinate, fromStart := range calcs.fromStart {
		for _, cheat := range possibleMoves {
			fromEnd, found := calcs.fromEnd[coordinate.Plus(cheat)]
			if !found {
				continue
			}
			length := fromStart + fromEnd + abs(cheat.X) +abs(cheat.Y)
			if baseDistance > length {
				cs := cheatScore{coordinate, coordinate.Plus(cheat), baseDistance - length}
				foundCheats = append(foundCheats, cs)
			}
		}
	}

	return foundCheats
}

func calculatePossibleMoves(length int)[]grid.Coordinate{
    possibleMoves:= []grid.Coordinate{}
	for cheatLength := 2; cheatLength <= length; cheatLength++ {
		for x := -cheatLength + 1; x < cheatLength; x++ {
			plus := grid.CreateCoordinate(x, cheatLength-abs(x))
			minus := grid.CreateCoordinate(x, -(cheatLength - abs(x)))
			possibleMoves = append(possibleMoves, plus, minus)
		}

		plus := grid.CreateCoordinate(cheatLength, 0)
		minus := grid.CreateCoordinate(-cheatLength, 0)
		possibleMoves = append(possibleMoves, plus, minus)
	}
    return possibleMoves
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type cheatScore struct {
	begin, end grid.Coordinate
	score      int
}

func calculateDistances(from grid.Coordinate, g grid.Grid) map[grid.Coordinate]int {

	distances := map[grid.Coordinate]int{from: 0}

	tails := []grid.Coordinate{from}
	for len(tails) > 0 {
		tail := tails[0]
		tails = append(tails, nextSteps(tail, g, distances)...)
		tails = tails[1:]
	}

	return distances
}
func nextSteps(c grid.Coordinate, g grid.Grid, bests map[grid.Coordinate]int) []grid.Coordinate {
	nextSteps := []grid.Coordinate{}
	nextScore := bests[c] + 1
	for _, adj := range g.GetCardinalAdjecents(c) {
        dist, found := bests[adj.Loc]
		if adj.Value != '#' && (!found || dist > nextScore) {

			bests[adj.Loc] = nextScore
			nextSteps = append(nextSteps, adj.Loc)
		}
	}
	return nextSteps
}

type precalculations struct {
	begin, end         grid.Coordinate
	fromStart, fromEnd map[grid.Coordinate]int
}
