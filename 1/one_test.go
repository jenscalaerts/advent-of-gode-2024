package main

import "testing"

func TestPart1_single(t *testing.T) {
	input := "3   4"
	entries, err := readFromString(input)
	if err != nil {
		t.Errorf("%v", err)
	}
    result := entries.calculateDistance()
	if result != 1 {
		t.Errorf("Expected result 1 but was %d", result)
	}

}

func TestPart1_multiLine(t *testing.T) {
	input := "3   4\n1   7"
	entries, err := readFromString(input)
	if err != nil {
		t.Errorf("%v", err)
	}
    result := entries.calculateDistance()
	if result != 7 {
		t.Errorf("Expected result 7 but was %d", result)
	}

}

func TestPart1_example(t *testing.T) {
	entries, err := readFromFile("example_1_1")

	if err != nil {
		t.Errorf("%v", err)
	}
    result := entries.calculateDistance()
	if result != 11 {
		t.Errorf("Expected result 11 but was %d", result)
	}
}

func TestPart2_multiLine(t *testing.T) {
    entries := locationEntries{[]int{3,3,4,5}, []int{4,4,3,1}}
	result := entries.calculateSimilarity()

	if result != 14 {
		t.Errorf("Expected result 14 but was %d", result)
	}
}

func TestPart2_example(t *testing.T) {
	entries, err := readFromFile("example_1_1")

	if err != nil {
		t.Errorf("%v", err)
	}
    result := entries.calculateSimilarity()
	if result != 31 {
		t.Errorf("Expected result 31 but was %d", result)
	}
}
