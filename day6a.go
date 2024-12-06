package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Dir int

const (
	None Dir = iota
	Up
	Down
	Left
	Right
)

var m = map[Dir]func(int, int) (int, int){
	Up:    func(x int, y int) (int, int) { return x, y - 1 },
	Down:  func(x int, y int) (int, int) { return x, y + 1 },
	Left:  func(x int, y int) (int, int) { return x - 1, y },
	Right: func(x int, y int) (int, int) { return x + 1, y },
}

func next(x int, y int, d Dir) (int, int) {
	return m[d](x, y)
}

func turn(x Dir) Dir {
	switch x {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		panic("Unknown dir")
	}
}

func offMap(x int, y int, territory []string) bool {
	if x < 0 || x >= len(territory[0]) {
		return true
	}
	if y < 0 || y >= len(territory) {
		return true
	}
	return false
}

type Coord struct {
	x int
	y int
}

// # = 35
// ^ = 94
func findStart(lines []string) Coord {
	for y, v := range lines {
		x := strings.IndexByte(v, 94)
		if x != -1 {
			return Coord{x, y}
		}
	}
	return Coord{-1, -1}
}

func tryLoop(start Coord, lines []string, tx int, ty int) bool {
	var visited = make(map[Coord]Dir)
	var d Dir = Up
	visited[start] = d
	// fmt.Println(visited[Coord{0, 0}])
	x := start.x
	y := start.y
	for true {
		nx, ny := next(x, y, d)
		// fmt.Println(x, y, nx, ny)
		if offMap(nx, ny, lines) {
			// fmt.Println(visited)
			return false
		} else if lines[ny][nx] == 35 {
			d = turn(d)
		} else if nx == tx && ny == ty {
			d = turn(d)
		} else if visited[Coord{nx, ny}] == d {
			return true
		} else {
			x = nx
			y = ny
			visited[Coord{x, y}] = d
		}
	}
	return false
}

func checkLoops(start Coord, lines []string) int {
	total := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if tryLoop(start, lines, x, y) {
				// fmt.Println("loop", x, y)
				total++
			}
		}
	}
	return total
}

func main() {

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	c := findStart(lines)
	//visited[c] = true
	// walk(c, lines)
	fmt.Println(checkLoops(c, lines))
	//fmt.Println(len(visited))
}
