package main

import (
	"fmt"
	"testing"
)

func TestExamplePart1(t *testing.T) {
    patterns, designs := readData("example") 
    result := patterns.countMatching(designs)
    if result != 6 {
        t.Fatalf("Expected 6 but got %d", result)
    }

}


func TestExampleSingle(t *testing.T) {
    patterns, _ := readData("example") 
    result := patterns.canMake("brgr", map[string]bool{})
    if !result  {
        fmt.Println("should be able to make")
    }

}

func TestSimple(t *testing.T) {
    patterns := towelPatterns{"s", "t"}
    if !patterns.canMake("st", map[string]bool{}){
        fmt.Println("should be able to make")
    }
    
}

func TestExamplePart2(t *testing.T) {
    patterns, designs := readData("example") 
    result := patterns.totalNumberOfCombinations(designs)
    if result != 16 {
        t.Fatalf("Expected 16 but got %d", result)
    }

}
