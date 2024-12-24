package main

import (
	"bufio"
	"fmt"
	"os"
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

func iterate(secret int, iterations int) int {
	x := secret
	for i := 0; i < iterations; i++ {
		x = evolve(x)
	}
	return x
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

	sum := 0
	for _, v := range inputs {
		sum += iterate(v, 2000)
	}
	fmt.Println(sum)
}
