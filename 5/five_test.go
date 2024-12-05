package main

import "testing"

func TestExample(t *testing.T){
    queue := load("example")
    result := sumValidMiddles(queue)
    if result !=  143{
        t.Fatalf("incorrect result %d exected 143", result)
    }
}
