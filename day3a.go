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
	re := regexp.MustCompile(`(mul)\((\d+),(\d+)\)|(do)\(\)|(don)'t\(\)`)
	matches := re.FindAllSubmatch(data, -1)
	var active bool = true
	for _, v := range matches {
		if string(v[4]) == "do" {
			active = true
		} else if string(v[5]) == "don" {
			active = false
		} else if string(v[1]) == "mul" && active {
			x, _ := strconv.Atoi(string(v[2]))
			y, _ := strconv.Atoi(string(v[3]))
			total += x * y
		}
	}
	return total
}

func main() {

	data, err := os.ReadFile(os.Args[1])
	check(err)

	fmt.Println(parse(data))
}
