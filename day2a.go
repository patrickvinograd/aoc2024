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

func parse(line string) []int {
	vstrs := strings.Split(line, " ")
	levels := make([]int, len(vstrs))
	for i, s := range vstrs {
		levels[i], _ = strconv.Atoi(s)
	}
	return levels
}

func dampSafe(reports []int) bool {
	if isSafe(reports) {
		return true
	}
	fmt.Println(reports)
	for i := 0; i < len(reports); i++ {
		sub := make([]int, len(reports))
		_ = copy(sub, reports)
		slices.Delete(sub, i, i+1)
		// fmt.Println(reports)
		// fmt.Println(sub[0 : len(sub)-1])
		if isSafe(sub[0 : len(sub)-1]) {
			return true
		}
	}
	return false
}

func isSafe(reports []int) bool {
	// fmt.Println(reports)
	var min, max int
	if reports[1] > reports[0] {
		max = 3
		min = 1
	} else {
		max = -1
		min = -3
	}
	for i := 1; i < len(reports); i++ {
		diff := reports[i] - reports[i-1]
		if diff > max || diff < min {
			return false
		}
	}
	return true
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	safe := 0
	for scanner.Scan() {
		line := scanner.Text()
		reports := parse(line)
		if dampSafe(reports) {
			safe++
		}

	}

	fmt.Println(safe)
}
