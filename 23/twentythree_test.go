package main

import (
	"fmt"
	"testing"
)

func TestTripletsPart(t *testing.T) {
    m, groups := readDateToMap("example")
    groups = calculateNextGroupSize(m, groups)
    fmt.Println(groups)
    groups = calculateNextGroupSize(m,groups)
    fmt.Println(groups)
    groups = calculateNextGroupSize(m,groups)
    fmt.Println(groups)
}
