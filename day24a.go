package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"slices"
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

func printRegister(prefix string) int {
	result := 0
	for k, v := range state {
		suffix, isz := strings.CutPrefix(k, prefix)
		if !isz {
			continue
		}
		num, _ := strconv.Atoi(suffix)
		if v == true {
			result = result | (1 << num)
		}
	}
	fmt.Println(result)
	fmt.Printf("%b\n", result)
	return result
}

var noded = make(map[string]bool)

func printFlowchartMermaid(states []State, rules []Rule) {
	fmt.Println("flowchart LR")
	for _, r := range rules {
		if !noded[r.a] {
			fmt.Printf("    %s(%s)\n", r.a, r.a)
			noded[r.a] = true
		}
		if !noded[r.b] {
			fmt.Printf("    %s(%s)\n", r.b, r.b)
			noded[r.b] = true
		}
		if !noded[r.output] {
			fmt.Printf("    %s(%s)\n", r.output, r.output)
			noded[r.output] = true
		}
		opnode := fmt.Sprintf("%s.%s.%s", r.a, r.b, r.output)
		fmt.Printf("    %s((%s))\n", opnode, r.op)
		fmt.Printf("    %s --> %s\n", r.a, opnode)
		fmt.Printf("    %s --> %s\n", r.b, opnode)
		fmt.Printf("    %s --> %s\n", opnode, r.output)
	}

}

func printFlowchart(states []State, rules []Rule) {
	fmt.Println("digraph LR {")
	for _, r := range rules {
		if !noded[r.a] {
			fmt.Printf("    %s;\n", r.a)
			noded[r.a] = true
		}
		if !noded[r.b] {
			fmt.Printf("    %s;\n", r.b)
			noded[r.b] = true
		}
		if !noded[r.output] {
			fmt.Printf("    %s;\n", r.output)
			noded[r.output] = true
		}
		opnode := fmt.Sprintf("%s_%s_%s", r.a, r.b, r.output)
		fmt.Printf("    %s [ label = \"%s\" ];\n", opnode, r.op)
		fmt.Printf("    %s -> %s;\n", r.a, opnode)
		fmt.Printf("    %s -> %s;\n", r.b, opnode)
		fmt.Printf("    %s -> %s;\n", opnode, r.output)
	}
	fmt.Println("}")
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
	// fmt.Println(initStates, rules)
	// printFlowchart(initStates, rules)
	execute(initStates, rules)
	x := printRegister("x")
	y := printRegister("y")
	z := printRegister("z")

	fmt.Printf("Expected x+y = %d, (%b)\n", x+y, x+y)
	fmt.Printf("Got z        = %d, (%b)\n", z, z)

	swaps := []string{"kth", "z12", "z26", "gsd", "tbt", "z32", "qnf", "vpm"}
	slices.Sort(swaps)
	fmt.Println(strings.Join(swaps, ","))
}
