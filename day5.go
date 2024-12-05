package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseRules(r []string) map[int][]int {
	result := make(map[int][]int)
	for _, v := range r {
		vals := strings.Split(v, "|")
		a, _ := strconv.Atoi(vals[0])
		b, _ := strconv.Atoi(vals[1])
		result[a] = append(result[a], b)
	}
	return result
}

func parseUpdates(u []string) [][]int {
	result := make([][]int, 0)
	for _, v := range u {
		vals := strings.Split(v, ",")
		update := make([]int, len(vals))
		for i, u := range vals {
			update[i], _ = strconv.Atoi(u)
		}
		result = append(result, update)
	}
	return result
}

func valid(update []int, rules map[int][]int) bool {
	for i := 0; i < len(update); i++ {
		for j := i; j < len(update); j++ {
			if i == j {
				continue
			}
			if slices.Contains(rules[update[j]], update[i]) {
				return false
			}
		}
	}
	return true
}

func processUpdates(updates [][]int, rules map[int][]int) int {
	total := 0
	for _, update := range updates {
		if valid(update, rules) {
			//fmt.Println(i, update)
			total += update[len(update)/2]
		}
	}
	return total
}

func main() {

	var r []string
	var u []string
	scanner := bufio.NewScanner(os.Stdin)

	var gap = false
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			gap = true
			continue
		}
		if gap == false {
			r = append(r, line)
		} else {
			u = append(u, line)
		}
	}
	rules := parseRules(r)
	updates := parseUpdates(u)
	fmt.Println(processUpdates(updates, rules))
}
