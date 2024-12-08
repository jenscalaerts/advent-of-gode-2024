package main

import (
	"advent/2024/grid"
	"testing"
)

func TestExample(t *testing.T) {
	g := grid.ReadGrid("example")
	results := calculateUniqueAntiNodes(g)
	if len(results) != 14 {
		t.Fatalf("expecting 14 but was %d", len(results))
	}
}

func TestExample_part2(t *testing.T) {
	g := grid.ReadGrid("example")
	results := calculateUniqueAntiNodesWithResonant(g)
	if len(results) != 34 {
		t.Fatalf("expecting 34 but was %d", len(results))
	}
}

func TestSimpleExample(t *testing.T) {
	g := grid.ReadGrid("simple")
	results := calculateUniqueAntiNodes(g)
	expectedResults := []grid.Coordinate{{X: 0, Y: 0}, {X: 3, Y: 3}}
	for _, expectedResult := range expectedResults {
		if !results[expectedResult] {
			t.Fatalf("expecting %v but was not in resultSet %v", expectedResult, results)
		}
	}
	if len(results) != len(expectedResults) {
		t.Fatalf("expecting %v resultsbut but result set was different %v", len(expectedResults), results)
	}

}

func TestSimplePart2_containsSelf(t *testing.T) {
	g := grid.ReadGrid("simple_part2")
	results := calculateUniqueAntiNodesWithResonant(g)

	if !results[grid.Coordinate{X: 7, Y: 7}] || !results[grid.Coordinate{X: 5, Y: 5}] {
		t.Fatal("Self not found")
	}
}
