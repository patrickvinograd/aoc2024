package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse7(line string) (int, []int) {
	parts := strings.Split(line, ":")
	total, _ := strconv.Atoi(parts[0])
	vstr, _ := strings.CutPrefix(parts[1], " ")
	vstrs := strings.Split(vstr, " ")
	vals := make([]int, len(vstrs))
	for i, s := range vstrs {
		vals[i], _ = strconv.Atoi(s)
	}
	return total, vals
}

var mul = func(a int, b int) int {
	return a * b
}

var add = func(a int, b int) int {
	return a + b
}

func try(target int, accum int, oper func(int, int) int, vals []int) bool {
	// fmt.Println(target, accum, oper, vals)
	if len(vals) == 0 {
		return accum == target
	}
	if try(target, oper(accum, vals[0]), add, vals[1:]) {
		return true
	}
	if try(target, oper(accum, vals[0]), mul, vals[1:]) {
		return true
	}
	return false
}

func sum(vals []int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}

func canBeTrue(total int, vals []int) (int, bool) {
	if try(total, 0, add, vals) {
		return total, true
	}
	if try(total, 1, mul, vals) {
		return total, true

	}
	return 0, false

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		total, vals := parse7(line)
		fmt.Println(total, vals)
		// fmt.Println(vals[1:])
		sum, valid := canBeTrue(total, vals)
		if valid {
			result += sum
		}
	}
	fmt.Println(result)
}
