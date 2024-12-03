package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var total float64
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
	slices.Sort(left)
	slices.Sort(right)
	// fmt.Println(left)
	// fmt.Println(right)

	for i, _ := range left {
		total += math.Abs((float64)(left[i] - right[i]))
	}
	fmt.Println((int)(total))
}
