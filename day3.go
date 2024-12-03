package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse(data []byte) int {
	total := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllSubmatch(data, -1)
	for _, v := range matches {
		x, _ := strconv.Atoi(string(v[1]))
		y, _ := strconv.Atoi(string(v[2]))
		total += x * y
	}
	return total
}

func main() {

	data, err := os.ReadFile(os.Args[1])
	check(err)

	fmt.Println(parse(data))
}
