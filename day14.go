package main

import (
	"bufio"
	"fmt"
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

func iterate(r Robot) Robot {
	x := r.startx
	y := r.starty
	for i := 0; i < iterations; i++ {
		x = mod(x+r.vx, tilex)
		y = mod(y+r.vy, tiley)
	}
	x = x % tilex
	y = y % tiley
	r.startx = x
	r.starty = y
	return Robot{r.id, x, y, r.vx, r.vy}
}

func iterateAll(data []Robot) []Robot {
	for i, v := range data {
		data[i] = iterate(v)
		// iterate(&v)
	}
	fmt.Println(data)
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
	fmt.Println(q1, q2, q3, q4)
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
