package main

import (
	"bufio"
	"fmt"
	"os"
)

type Lock struct {
	pins []int
}

type Key struct {
	pins []int
}

var locks = make([]Lock, 0)
var keys = make([]Key, 0)

func makeLock(input []string) Lock {
	pins := make([]int, 0)
	for x := 0; x < len(input[0]); x++ {
		for y := 1; y < len(input); y++ {
			if string(input[y][x]) == "." {
				pins = append(pins, y-1)
				break
			}
		}
	}
	return Lock{pins}
}

func makeKey(input []string) Key {
	pins := make([]int, 0)
	for x := 0; x < len(input[0]); x++ {
		for y := len(input) - 2; y >= 0; y-- {
			if string(input[y][x]) == "." {
				pins = append(pins, (len(input) - 2 - y))
				break
			}
		}
	}
	return Key{pins}
}

func parse(inputs [][]string) {
	for _, input := range inputs {
		if input[0] == "#####" {
			locks = append(locks, makeLock(input))
		} else {
			keys = append(keys, makeKey(input))
		}
	}
	// fmt.Println("locks", len(locks), locks)
	// fmt.Println("keys", len(keys), keys)
}

func test(lock Lock, key Key) bool {
	max := 5
	for i := 0; i < len(lock.pins); i++ {
		if lock.pins[i]+key.pins[i] > max {
			return false
		}
	}
	return true
}

func testAll() int {
	result := 0
	for _, lock := range locks {
		for _, key := range keys {
			if test(lock, key) {
				result++
			}
		}
	}
	return result
}

func main() {

	var inputs [][]string
	scanner := bufio.NewScanner(os.Stdin)

	current := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			inputs = append(inputs, current)
			current = make([]string, 0)
		} else {
			current = append(current, line)
		}
	}
	inputs = append(inputs, current)
	parse(inputs)
	fmt.Println(testAll())
}
