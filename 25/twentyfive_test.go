package main

import (
	"fmt"
	"testing"
)

func TestPart1Example(t *testing.T) {
	locks, keys := readData("example")
    fmt.Println(locks, keys)
	result := getNumberOfOverlaps(locks, keys)
	if result != 3 {
		t.Error("Expected 3 but got ", result)
	}

}
