package main

import (
	"fmt"
	"testing"
)

func TestPartOneExample(t *testing.T){
    data := readFile("example_2_1")
    result := calculateNumberOfSave(data, isReportSafeNoDampner); 
    if result != 2 {
        t.Errorf("Expected 2 but result is %d", result)
    }
}

func TestPartTwoExample(t *testing.T){
    fmt.Println("slkdjflskjdfklskfljdlsfja;sdkjfa;slkfj")
    data := readFile("example_2_1")
    result := calculateNumberOfSave(data, isReportSafeWithDampner); 
    if result != 4 {
        t.Errorf("Expected 4 but result is %d", result)
    }
}


func TestPartTwo_startWrong_equal(t *testing.T){
    result := isReportSafeWithDampner([]int{1,1,4,5})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}

func TestPartTwo_startWrong_sign(t *testing.T){
    result := isReportSafeWithDampner([]int{1,-1,2,4})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}

func TestPartTwo_startWrong_diff(t *testing.T){
    result := isReportSafeWithDampner([]int{1,-5,3,5})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}

func TestPartTwo_startWrong_diff2(t *testing.T){
    result := isReportSafeWithDampner([]int{1,5,3,1})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}


func TestPartTwo_startWrong_diff3(t *testing.T){
    result := isReportSafeWithDampner([]int{9,5,3,1})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}

func TestPartTwo_wrong2_equal(t *testing.T){

    result := isReportSafeWithDampner([]int{4,5,5,7})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}

func TestPartTwo_wrongMiddle_diff(t *testing.T){
    result := isReportSafeWithDampner([]int{1,3,8,5})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}


func TestPartTwo_wrongMiddle_equal(t *testing.T){
    result := isReportSafeWithDampner([]int{1,3,5,5,7})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}

func TestPartTwo_wrongMiddle_sign(t *testing.T){
    result := isReportSafeWithDampner([]int{1,3,5,4,7})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}


func TestPartTwo_wrong4_diff(t *testing.T){
    result := isReportSafeWithDampner([]int{1,3,5,9,7})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}

func TestPartTwo_wrong4Equal(t *testing.T){
    result := isReportSafeWithDampner([]int{1,3,5,5,7})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}

func TestPartTwo_wrongEnd_diff(t *testing.T){
    result := isReportSafeWithDampner([]int{1,3,5,7,70})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}

func TestPartTwo_wrongEnd_Equal(t *testing.T){
    result := isReportSafeWithDampner([]int{1,3,5,7,7})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}

func TestPartTwo_wrongEnd_sign(t *testing.T){
    result := isReportSafeWithDampner([]int{1,3,5,7,7})
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}


func TestPartTwo_begin_sign(t *testing.T){
    result := isReportSafeWithDampner([]int{83,81,82,83,85,87,90,92 })
    if !result {
        t.Errorf("Expected result but result is %t", result)
    }
}
