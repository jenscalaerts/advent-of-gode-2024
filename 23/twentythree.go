package main

import (
	"advent/2024/parsing"
	"fmt"
	"slices"
	"strings"
)

func main(){
    dat := readData("data")
    fmt.Println(len(dat.findSets()))
}

func (nm nodeMap) findSets() map[string]bool {
    all := map[string]bool{}
	for _, node := range nm {
        if node.name[0] == 't' {
            cycles := node.hasNodeAtRange(3, node)
            for _, cycle := range cycles {
                cycleName := make([]string, len(cycle))
                for i, n := range cycle {
                    cycleName[i] = n.name
                }
                slices.Sort(cycleName)
                key := strings.Join(cycleName,"")
                all[key] = true
            }
        }
	}
	return all
}

func (n *nodeType) hasNodeAtRange(distance int, node *nodeType) [][]*nodeType {
	if distance == 0 {
		if n == node{
            return [][]*nodeType{{}}
        }
        return[][]*nodeType{}

	}
    cycles := [][]*nodeType{}
	for _, edge := range n.edges {
        paths := edge.hasNodeAtRange(distance - 1, node)
        for i, path := range paths {
           paths[i] = append(path, n) 
        }
        cycles = append(cycles, paths...)
    }
    return cycles
}

func readData(name string) nodeMap {
	lines := parsing.ReadLines(name)
	nodes := nodeMap{}
	for _, line := range lines {
		split := strings.Split(line, "-")

		left := nodes.getNode(split[0])
		right := nodes.getNode(split[1])
		left.edges = append(left.edges, right)
		right.edges = append(right.edges, left)
	}
	return nodes
}

type nodeMap map[string]*nodeType

func (nm *nodeMap) getNode(name string) *nodeType {
	n := *nm
	node := n[name]
	if node == nil {
		node = &nodeType{name: name}
		n[name] = node
	}
	return node

}

type nodeType struct {
	name  string
	edges []*nodeType
}
