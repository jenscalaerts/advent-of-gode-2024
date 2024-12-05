package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)
func main(){
    data := load("data")
    result1 := sumValidMiddles(data)
    fmt.Printf("Result 1: %d", result1)
}


func sumValidMiddles(queue printQueue) int{
    sum := 0
    for _, up := range(queue.updates){
        if up.satisfies(queue.rules){
            sum += up.getMiddle()

            fmt.Println(up.getMiddle())
        }
    }
    return sum;
}

func load(name string) printQueue {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rules := map[int][]int{}

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		rule, err := parseRule(text)
        fmt.Println(rule)
		if err != nil {
			log.Fatal(err)
		}
		rules[rule.before] = append(rules[rule.before], rule.after)
	}
	var updates []update
	for scanner.Scan() {
		text := scanner.Text()
		numStrings := strings.Split(text, ",")
		var u update
		for _, numString := range numStrings {
			num, err := strconv.ParseInt(numString, 0, 0)
			if err != nil {
				log.Fatal(err)
			}
			u = append(u, int(num))
		}
		updates = append(updates, u)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return printQueue{rules, updates}
}

func (up update) getMiddle() int {
	return up[(len(up)/2)]
}

func (up update) satisfies(rules map[int][]int) bool {
    fmt.Println(up)
	for index, val := range up {
		mustBeBefore := rules[val]
        fmt.Printf("must be before %v",mustBeBefore)
		for _, val := range mustBeBefore {
			if slices.Contains(up[:index], val) {
                fmt.Printf("failure contains %d before index %d", val, index)
				return false
			}
		}
	}
    fmt.Println("+++++++++++++++++++++") 
	return true

}

func parseRule(s string) (rule, error) {

	split := strings.Index(s, "|")
	before, err := strconv.ParseInt(s[:split], 0, 0)
	if err != nil {
		return rule{}, fmt.Errorf("Fail to parse %v, %w", s, err)
	}

	after, err := strconv.ParseInt(s[split+1:], 0, 0)
	if err != nil {
		return rule{}, fmt.Errorf("Fail to parse %v, %w", s, err)
	}
	return rule{int(before), int(after)}, nil
}

type printQueue struct {
	rules   map[int][]int
	updates []update
}

type rule struct {
	before, after int
}

type update []int
