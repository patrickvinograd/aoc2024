package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func mix(a int, b int) int {
	return a ^ b
}

func prune(a int) int {
	return mod(a, 16777216)
}

func evolve(secret int) int {
	x := prune(mix(secret, secret*64))
	y := prune(mix(x, x/32))
	z := prune(mix(y, y*2048))
	return z
}

var priceMap = make(map[int][]int)
var changeMap = make(map[int][]int)

func priceFor(sequence []int, index int) int {
	changes := changeMap[index]
	prices := priceMap[index]
	for i := 0; i < len(changes)-4+1; i++ {
		// fmt.Println("checking ", sequence, " vs. ", changes[i:i+4])
		if slices.Equal(sequence, changes[i:i+4]) {
			return prices[i+4-1]
			// fmt.Println(index, "match", sequence, "at", i, ": ", changes[i:i+4], "price: ", prices[i+4-1])
		}
	}
	return 0
}

func totalPriceFor(sequence []int) int {
	result := 0
	for i := 0; i < len(priceMap); i++ {
		result += priceFor(sequence, i)
	}
	return result
}

func tryAllSequences() int {
	best := 0
	count := 0
	for a := -9; a < 10; a++ {
		for b := -9; b < 10; b++ {
			for c := -9; c < 10; c++ {
				for d := -9; d < 10; d++ {
					sequence := []int{a, b, c, d}
					if count%10000 == 0 {
						fmt.Println("checking", sequence)
					}
					thisprice := totalPriceFor(sequence)
					if thisprice > best {
						best = thisprice
						fmt.Println("new best", best)
					}
					count++
				}
			}
		}
	}
	return best
}

func iterate(secret int, iterations int, index int) {
	prices := make([]int, iterations)
	changes := make([]int, iterations)

	curr := secret
	for i := 0; i < iterations; i++ {
		next := evolve(curr)
		prices[i] = next % 10
		changes[i] = (next % 10) - (curr % 10)
		curr = next
	}
	priceMap[index] = prices
	changeMap[index] = changes
}

func main() {

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var inputs []int
	for _, s := range lines {
		seed, _ := strconv.Atoi(s)
		inputs = append(inputs, seed)
	}

	// seq := []int{-1, -1, 0, 2}
	// iterate(123, 10, 0)
	// fmt.Println(priceMap)
	// fmt.Println(changeMap)
	// priceFor(seq, 0)

	for i, v := range inputs {
		iterate(v, 2000, i)
		// seq := []int{-2, 1, -1, 3}
		// priceFor(seq, i)
	}
	fmt.Println(tryAllSequences())

}
