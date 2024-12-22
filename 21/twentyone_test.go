package main

import (
	"testing"
)

func lksja(t *testing.T) {
	result := part1("example")
	if result != 126384 {
		t.Errorf("expected %d but got %d", 126384, result)
	}

}

func TestPart1Simple_1Level(t *testing.T) {
	result := findBestForRoute("A", map[memo]int{}, 3)
	if result != 1 {
		t.Errorf("expected %d but got %d", 1, result)
	}
}

func TestPart1Example2(t *testing.T) {
	result := findBestForRoute("vA", map[memo]int{}, 2)
	if result != 6 {
		t.Errorf("expected %d but got %d", 6, result)
	}

}

func TestPart1Example3(t *testing.T) {
	result := findBestForRoute("vA", map[memo]int{}, 1)
	if result != 16 {
		t.Errorf("expected %d but got %d", 16, result)
	}

}


func TestPart1SimpleExample(t *testing.T) {
	result := part1ForLines([]string{"029A"} )
	if result != 1972 {
		t.Errorf("expected %d but got %d", 1972, result)
	}
}

func TestPart1Example(t *testing.T) {
	result := part1("example")
	if result != 126384 {
		t.Errorf("expected %d but got %d", 126384, result)
	}
}



const dlkjs = `
179A
14 <<^A^^A>>AvvvA
28 <<vAA>^A>A<AA>AvAA^A<vAAA>^A
64 <<vA A >A >^A A vA <^A >A vA ^A <<vA >>^A A vA ^A <vA >^A A <A >A <<vA >A >^A A A vA <^A >A

<<vAA>A>^AAvA<^A>AvA^A<<vA>>^AAvA^A<vA>^AA<A>A<<vA>A>^AAAvA<^A>A
<<vA >>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>AAvAA
<    A
   <    < v   A  A  >  ^  A  >  A   <   A    A >  A  v   A   A  ^ A   <   v  A   A A >  ^   A
<<vAA>^A>A<AA>AvAA^A<vAAA>^A   
<<vAA>^A>A<AA>AvAA^A<vAAA>^A

179A 179 64 11456
`
