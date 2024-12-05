package main

import "testing"

func TestExample_part1(t *testing.T){
    queue := load("example")
    result := sumValidMiddles(queue)
    if result !=  143{
        t.Fatalf("incorrect result %d exected 143", result)
    }
}

func TestExample_part2(t *testing.T){
    queue := load("example")
    result := correctAndSumInvalidMiddles(queue)
    if result !=  123{
        t.Fatalf("incorrect result %d exected 123", result)
    }
}
