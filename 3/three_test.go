package main

import (
	"testing"
)

func TestPartOneExample(t *testing.T){
	s := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
    result := calculate(s)
    if result != 161{
        t.Errorf("Incorrect result %d expected 161", result)
    }
}


func TestPartTwoExample(t *testing.T){
	s := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
    result := calculateWithConditional(s)
    if result != 48{
        t.Errorf("Incorrect result %d expected 48", result)
    }
}
