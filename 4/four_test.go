package main

import "testing"

func TestPart1_example(t *testing.T) {
	data := readGrid("example")
	countXMASStraight(data)

}

func TestPart2_example(t *testing.T) {
	data := readGrid("example")
	result := countX_MAS(data)
	if result != 9 {
		t.Errorf("Unexpected result %d", result)
	}

}
