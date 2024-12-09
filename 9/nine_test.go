package main

import (
	"fmt"
	"testing"
)

func TestSumBetween(t *testing.T){
    result := sumOfIndicesBetweenInclusive(1,4)
    if result != 10 {
        t.Fatalf("expected 10 but got %d", result)
    }
}


func TestMain(t *testing.T){
    //main()
}

func TestRead(t *testing.T){
    data := readData("data")
    var sum int64
    for _,da := range data {
        sum += int64(da)
    }
    fmt.Println(sum)
}

