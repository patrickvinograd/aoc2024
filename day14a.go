package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
	id     int
	startx int
	starty int
	vx     int
	vy     int
}

var tilex = 101
var tiley = 103
var iterations = 100

func parse(line string, num int) Robot {
	re := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	matches := re.FindStringSubmatch(line)
	startx, _ := strconv.Atoi(matches[1])
	starty, _ := strconv.Atoi(matches[2])
	vx, _ := strconv.Atoi(matches[3])
	vy, _ := strconv.Atoi(matches[4])

	return Robot{num, startx, starty, vx, vy}
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func iterate(r Robot, iters int) Robot {
	x := r.startx
	y := r.starty
	for i := 0; i < iters; i++ {
		x = mod(x+r.vx, tilex)
		y = mod(y+r.vy, tiley)
	}
	x = x % tilex
	y = y % tiley
	r.startx = x
	r.starty = y
	return Robot{r.id, x, y, r.vx, r.vy}
}

func count(x int, y int, data []Robot) int {
	result := 0
	for _, v := range data {
		if v.startx == x && v.starty == y {
			result++
		}
	}
	return result
}

func display(data []Robot) {
	for y := 0; y < tiley; y++ {
		for x := 0; x < tilex; x++ {
			if count(x, y, data) > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func iterateAll(data []Robot) []Robot {
	minScore := math.MaxInt
	minIter := 0
	for c := 0; c < 9000; c++ {
		for i, v := range data {
			data[i] = iterate(v, 1)
			score := score(data)
			if score < minScore {
				minScore = score
				minIter = c
			}
		}
		if c == 8269 || c == 5442 {
			display(data)
		}
	}
	fmt.Println(minIter, minScore)
	return data
}

func score(data []Robot) int {
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	for _, v := range data {
		if v.startx < tilex/2 {
			if v.starty < tiley/2 {
				q1++
			} else if v.starty > tiley/2 {
				q3++
			}
		} else if v.startx > tilex/2 {
			if v.starty < tiley/2 {
				q2++
			} else if v.starty > tiley/2 {
				q4++
			}
		}
	}
	// fmt.Println(q1, q2, q3, q4)
	return q1 * q2 * q3 * q4
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	data := make([]Robot, 0)

	// result := 0
	num := 0
	for scanner.Scan() {
		line := scanner.Text()
		robot := parse(line, num)
		data = append(data, robot)
		num++
	}

	fmt.Println(data)
	data = iterateAll(data)
	fmt.Println(score(data))
	// fmt.Println(result)
}
