package main

import (
	"fmt"
	"testing"
)

func TestTripletsPart(t *testing.T) {
    data := readData("example")    
    fmt.Println(data.findSets())
}
