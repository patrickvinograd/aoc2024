package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func score(left []int, right []int) int {
	total := 0
	var counts = make(map[int]int)
	for _, v := range right {
		counts[v] = counts[v] + 1
	}
	for _, v := range left {
		total += v * counts[v]
	}
	return total
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	left := make([]int, 0)
	right := make([]int, 0)

	for scanner.Scan() {
		var l int
		var r int
		line := scanner.Text()
		fmt.Sscanf(line, "%d    %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}

	fmt.Println(score(left, right))
}
