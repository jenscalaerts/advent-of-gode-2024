package main

import (
	"advent/2024/parsing"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {
	nodeMap, sets := readDateToMap("data")
	size := 2
	for {
		newSets := calculateNextGroupSize(nodeMap, sets)

		size++
		if len(newSets) == 0 {
			fmt.Println(len(sets[0]))
			return
		}

		if size == 3 {
			count := 0
			for _, set := range newSets {
				for _, computer := range set {
					if strings.Index(computer, "t") == 0 {
						count++
						break
					}
				}

			}
			fmt.Println(count)
		}
		printSets(newSets)
		sets = newSets
	}
}

func calculateNextGroupSize(nodeMap map[string][]string, sets []lan) []lan {

	newSets := []lan{}
	for i := 0; i < len(sets)-3; i++ {
		set := sets[i]
		for _, other := range sets[i+1:] {
			if slices.Equal(set[:len(other)-1], other[:len(other)-1]) {
				if slices.Contains(nodeMap[set.last()], other.last()) {

					newSet := append(slices.Clone(set), other.last())
					newSets = append(newSets, newSet)
				}
			} else {
				break
			}
            
		}
	}

	slices.SortFunc(newSets, compare)
	newSets = slices.CompactFunc(newSets, slices.Equal)
	return newSets

}

func printSets(lans []lan) {
	for _, lan := range lans {
		fmt.Println(lan)
	}
	fmt.Println("===============")
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

func readDateToMap(name string) (map[string][]string, []lan) {
	lines := parsing.ReadLines(name)
	nodes := map[string][]string{}
	groups := make([]lan, len(lines))
	for i, line := range lines {
		split := strings.Split(line, "-")
		nodes[split[0]] = append(nodes[split[0]], split[1])
		nodes[split[1]] = append(nodes[split[1]], split[0])
		slices.Sort(split)
		groups[i] = split

	}
	slices.SortFunc(groups, compare)
	return nodes, groups
}

type lan []string

func compare(l, r lan) int {
	if len(l) != len(r) {
		panic("Different sized compared")
	}
	for i := range l {
		diff := cmp.Compare(l[i], r[i])
		if diff != 0 {
			return diff
		}
	}
	return 0
}

func (l lan) last() string {
	return l[len(l)-1]
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
