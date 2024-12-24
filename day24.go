package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type State struct {
	node  string
	value bool
}

type Rule struct {
	a      string
	b      string
	op     string
	output string
}

var state = make(map[string]bool)

func operate(a string, b string, op string) bool {
	if op == "AND" {
		return state[a] && state[b]
	} else if op == "OR" {
		return state[a] || state[b]
	} else if op == "XOR" {
		return state[a] != state[b]
	} else {
		panic("Unknown operator " + op)
	}
}

func execute(initState []State, rules []Rule) {
	for _, s := range initState {
		state[s.node] = s.value
	}
	runlist := list.New()
	for _, r := range rules {
		runlist.PushBack(r)
	}
	for e := runlist.Front(); e != nil; e = runlist.Front() {
		rule := runlist.Remove(e).(Rule)
		_, af := state[rule.a]
		_, bf := state[rule.b]
		if af == true && bf == true {
			state[rule.output] = operate(rule.a, rule.b, rule.op)
		} else {
			runlist.PushBack(rule)
		}
	}
}

func zvals() int {
	result := 0
	for k, v := range state {
		suffix, isz := strings.CutPrefix(k, "z")
		if !isz {
			continue
		}
		num, _ := strconv.Atoi(suffix)
		if v == true {
			result = result | (1 << num)
		}
	}
	return result
}

func main() {

	initStates := make([]State, 0)
	rules := make([]Rule, 0)
	scanner := bufio.NewScanner(os.Stdin)

	re_a := regexp.MustCompile(`(.*): (\d)`)
	re_b := regexp.MustCompile(`(.*) (.*) (.*) -> (.*)`)

	instates := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			instates = false
		} else if instates {
			matches := re_a.FindStringSubmatch(line)
			// fmt.Println(matches)
			var val bool
			if matches[2] == "1" {
				val = true
			} else if matches[2] == "0" {
				val = false
			}
			initStates = append(initStates, State{matches[1], val})
		} else if !instates {
			matches := re_b.FindStringSubmatch(line)
			// fmt.Println(matches)
			rules = append(rules, Rule{matches[1], matches[3], matches[2], matches[4]})
		}
	}
	fmt.Println(initStates, rules)
	execute(initStates, rules)
	fmt.Println(state)
	fmt.Println(zvals())
	// fmt.Println(strings.Join(output, ","))

}
