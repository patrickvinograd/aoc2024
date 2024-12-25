package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var links = make(map[string][]string)

func parse(inputs []string) {
	for _, v := range inputs {
		comps := strings.Split(v, "-")
		c1 := comps[0]
		c2 := comps[1]

		links[c1] = append(links[c1], c2)
		links[c2] = append(links[c2], c1)
	}
}

var groups = make(map[string][]string)

func findTriplets() [][]string {
	result := make([][]string, 0)
	for k, v := range links {
		for _, x := range v {
			for _, y := range v {
				if x != y && slices.Contains(links[x], y) {
					t := []string{k, x, y}
					slices.Sort(t)
					tkey := fmt.Sprintf("%s-%s-%s", t[0], t[1], t[2])
					groups[tkey] = t
					// result = append(result, t)
				}
			}
		}
	}
	return result
}

func mkey(group []string) string {
	slices.Sort(group)
	return strings.Join(group, ",")
}

func growGroups() map[string][]string {
	result := make(map[string][]string)
	for k, v := range groups {
		if slices.Contains(v, "ba") && slices.Contains(v, "et") {
			fmt.Println("huh?", k, v)
		}
		first := v[0]
		for _, candidate := range links[first] {
			if !slices.Contains(v, candidate) {
				inall := true
				for _, peer := range v[1:] {
					if !slices.Contains(links[peer], candidate) {
						inall = false
					}
				}
				if inall {
					newgroup := make([]string, len(v))
					copy(newgroup, v)
					newgroup = append(newgroup, candidate)
					slices.Sort(newgroup)
					if slices.Contains(newgroup, "ba") && slices.Contains(newgroup, "et") {
						fmt.Println(first, v, candidate, newgroup)
					}
					newkey := mkey(newgroup)
					result[newkey] = newgroup
				}
			}
		}
	}
	return result
}

func main() {

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	parse(lines)

	findTriplets()
	for i := 0; i < 10; i++ {
		newgroups := growGroups()
		fmt.Println("went from", len(groups), "to", len(newgroups))
		if len(newgroups) == 0 {
			break
		}
		groups = newgroups
	}
	for k, v := range groups {
		fmt.Println(len(v), k)
	}
}
