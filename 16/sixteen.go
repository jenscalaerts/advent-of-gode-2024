package main

import (
	"advent/2024/grid"
	"fmt"
)

type maze grid.Grid

const rotationCost = 1000

func main() {
	g := grid.ReadGrid("data")
	fmt.Println(calculateMinScore(g))
}

func calculateMinScore(g grid.Grid) (int, int) {
	startLocation := g.FindAll('S')
	if len(startLocation) != 1 {
		panic("invalid input")
	}
	g.Set(startLocation[0], '.')
	start := reindeer{startLocation[0], grid.Coordinate{X: 1, Y: 0}}
	bestStates := map[reindeer]scoreAndPrevious{
		start: {score: 1, previous: []reindeer{}},
	}

	for tails := []reindeerWithScore{{start, 1}}; len(tails) != 0; tails = tails[1:] {
		tail := tails[0]
		nextSteps := nextSteps(g, tail.rein)
		for nextStep, scoreAdd := range nextSteps {
			nextScore := tail.score + scoreAdd
			if bestStates[nextStep].score > nextScore || bestStates[nextStep].score == 0 {
				bestStates[nextStep] = scoreAndPrevious{[]reindeer{tail.rein}, nextScore}
				tails = append(tails, reindeerWithScore{nextStep, tail.score + scoreAdd})
			} else if nextScore == bestStates[nextStep].score {
				current := bestStates[nextStep]
				current.previous = append(current.previous, tail.rein)
				bestStates[nextStep] = current
			}
		}
	}
	bestScore, bestReindeer := findBestEnd(bestStates, g)

	return bestScore - 1, calculateAlwaysVisited(bestReindeer, bestStates)
}

func findBestEnd(bestStates map[reindeer]scoreAndPrevious, g grid.Grid) (int, reindeer) {
	end := g.FindAll('E')[0]
	best := scoreAndPrevious{}
	var bestReindeer reindeer
	for state, score := range bestStates {
		if state.coord == end {
			if best.score == 0 || best.score > score.score {
				best = score
				bestReindeer = state
			}
		}
	}
	return best.score, bestReindeer
}

func calculateAlwaysVisited(bestReindeer reindeer, bestStates map[reindeer]scoreAndPrevious) int {
	routeCoordinates := map[grid.Coordinate]bool{}
	alreadyVisited := map[reindeer]bool{}
	for ends := []reindeer{bestReindeer}; len(ends) > 0; ends = ends[1:] {
		next := ends[0]
		_, found := alreadyVisited[next]
		if !found {
			alreadyVisited[next] = true
			routeCoordinates[next.coord] = true
			ends = append(ends, bestStates[next].previous...)
		}

	}
	return len(routeCoordinates)
}

type reindeer struct {
	coord     grid.Coordinate
	direction grid.Coordinate
}

type reindeerWithScore struct {
	rein  reindeer
	score int
}

type scoreAndPrevious struct {
	previous []reindeer
	score    int
}

func (r reindeer) nextCoord() grid.Coordinate {
	return r.coord.Plus(r.direction)
}

func nextSteps(m grid.Grid, r reindeer) map[reindeer]int {
	allowedSteps := map[reindeer]int{
		{r.coord, r.direction.RotateRight()}: rotationCost,
		{r.coord, r.direction.RotateLeft()}:  rotationCost,
	}
	move := r.nextCoord()
	if grid.Grid.Get(m, move) != '#' {
		allowedSteps[reindeer{r.nextCoord(), r.direction}] = 1
	}
	return allowedSteps
}
