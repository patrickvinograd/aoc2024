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

var triplets = make(map[string]bool)

func findTriplets() [][]string {
	result := make([][]string, 0)
	for k, v := range links {
		for _, x := range v {
			for _, y := range v {
				if x != y && slices.Contains(links[x], y) {
					t := []string{k, x, y}
					slices.Sort(t)
					tkey := fmt.Sprintf("%s-%s-%s", t[0], t[1], t[2])
					triplets[tkey] = true
					// result = append(result, t)
				}
			}
		}
	}
	return result
}

func countHistorians() int {
	result := 0
	for k, _ := range triplets {
		if string(k[0]) == "t" || string(k[3]) == "t" || string(k[6]) == "t" {
			result++
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
	fmt.Println(links)
	findTriplets()
	// fmt.Println(len(findTriplets()))
	fmt.Println(triplets)
	fmt.Println(countHistorians())
}
