package main

import (
	"testing"
)

func TestSumBetween(t *testing.T){
    result := sumOfIndicesBetweenInclusive(1,4)
    if result != 10 {
        t.Fatalf("expected 10 but got %d", result)
    }
}


func TestMain(t *testing.T){
    part2()
}
