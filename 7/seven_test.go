package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	result := isPossibleFormula(10, 0, []int{10})
	if !result {
		t.Fatal("expected true")
	}
}

func TestSimple_2Elements(t *testing.T) {
	result := isPossibleFormula(10, 0, []int{2, 5})
	if !result {
		t.Fatal("expected true")
	}
}

func TestSimple_2Elements_plus(t *testing.T) {
	result := isPossibleFormula(10, 0, []int{5, 5})
	if !result {
		t.Fatal("expected true")
	}
}

func TestSimple_whenNotPossible_returnsFalse(t *testing.T) {
	result := isPossibleFormula(10, 0, []int{3, 3})
	if result {
		t.Fatal("expected true")
	}
}

func TestSimple_3whenNotPossible_returnsFalse(t *testing.T) {
	result := isPossibleFormula(10, 0, []int{3, 3, 3})
	if result {
		t.Fatal("expected true")
	}
}

func TestPart1Example(t *testing.T) {
	calibrations := readCalibrations("example")
	result := sumValidCalibrations(calibrations)
	if result != 3749 {
		t.Fatalf("Expected 3749 but got %d", result)
	}
}

func TestPart2Example(t *testing.T) {
	calibrations := readCalibrations("example")
	result := sumValidCalibrationsWithConcat(calibrations)
	if result != 11387 {
		t.Fatalf("Expected 11387 but got %d", result)
	}
}

func TestPart2_possibleWithConcat(t *testing.T) {
	result := isPossibleFormulaWithConcat(192,0, []int{17, 8, 14})
    if !result {
        t.Fatal("expecte true")
    }
}
