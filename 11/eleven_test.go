package main

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T){
    data := readData("example") 
    fmt.Println(data)
    result := brutelong(data,25)
    if result != 55312 {
        t.Fatalf("No workie %d", result)
    }
    fmt.Println(data)
}


func TestExample_2(t *testing.T){
    data := readData("example") 
    result := brutelong(data, 75)
    fmt.Println(result)
}
func TestBlink(t *testing.T){
    result := blink(3219)
    if 32 != result[0] && 19 != result[1]{
        t.Fatalf("result is not 32 19 %d %d\n", result[0], result[1])
    }
}

func TestBlinkZero(t *testing.T){
    result := blink(0)
    if 1 != result[0] {
        t.Fatalf("result is not 1 %v\n", result)
    }
}

func TestBlinkNotEvenLength(t *testing.T){
    result := blink(111)
    if 224664 != result[0] {
        t.Fatalf("result is not 224664 %v\n", result)
    }
}
