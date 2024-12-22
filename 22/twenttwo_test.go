package main

import (
	"fmt"
	"testing"
)

func TestMix(t *testing.T) {
	secret := 42
	value := 15
	result := mix(value, secret)
	assertEquals(37, result, t)
}

func TestPrune(t *testing.T) {
	secret := 100000000
	result := prune(secret)
	assertEquals(16113920, result, t)
}

func assertEquals(expected, result int, t *testing.T) {
	if expected != result {
		t.Errorf("Expected %d but got %d, a difference of %d",
			expected, result, expected-result)
	}
}

func TestAdvance(t *testing.T) {
	result := advance(123)
	assertEquals(15887950, result, t)
}

func TestAdvanceMultipleTimes(t *testing.T) {
	expectedResults := []int{
		15887950,
		16495136,
		527345,
		704524,
		1553684,
		12683156,
		11100544,
		12249484,
		7753432,
		5908254,
	}
    secret := 123
    for _, expected := range expectedResults{
        secret = advance(secret)
        assertEquals(expected,secret, t)
    }
}

func TestExample(t *testing.T){
    result := part1("example")
    assertEquals(37327623, result, t)
}


func TestExamplePart2(*testing.T){

    _, m := part2("example2")
    fmt.Println(m[fmt.Sprint([]int{-2, 1, -1,3})])
}

func TestPart2Simple(*testing.T){
    best, _ := part2WithData([]int{123})
    fmt.Println(best)
}
