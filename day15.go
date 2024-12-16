package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Dir int

const (
	Up Dir = iota
	Down
	Left
	Right
)

var mc = map[Dir]func(Coord) Coord{
	Up:    func(pt Coord) Coord { return Coord{pt.x, pt.y - 1} },
	Down:  func(pt Coord) Coord { return Coord{pt.x, pt.y + 1} },
	Left:  func(pt Coord) Coord { return Coord{pt.x - 1, pt.y} },
	Right: func(pt Coord) Coord { return Coord{pt.x + 1, pt.y} },
}

func todir(cmd string) Dir {
	switch cmd {
	case "^":
		return Up
	case ">":
		return Right
	case "<":
		return Left
	case "v":
		return Down
	default:
		panic("Unknown dir")
	}
}

var grid = make(map[Coord]string)
var xmax int
var ymax int

func offMap(x int, y int) bool {
	if x < 0 || x >= xmax {
		return true
	}
	if y < 0 || y >= ymax {
		return true
	}
	return false
}

type Coord struct {
	x int
	y int
}

func firstSpace(pos Coord, cmd Dir) (Coord, bool) {
	curr := pos
	for true {
		curr = mc[cmd](curr)
		val := grid[curr]
		if val == "." {
			return curr, true
		} else if val == "#" {
			return Coord{-1, -1}, false
		}
	}
	return Coord{-1, -1}, false
}

func shift(pos Coord, cmd Dir) Coord {
	space, canMove := firstSpace(pos, cmd)
	fmt.Println(pos, cmd, space, canMove)
	if canMove {
		if cmd == Up {
			for y := space.y; y < pos.y; y++ {
				grid[Coord{space.x, y}] = grid[Coord{space.x, y + 1}]
			}
			grid[Coord{pos.x, pos.y}] = "."
			return Coord{pos.x, pos.y - 1}
		} else if cmd == Down {
			for y := space.y; y > pos.y; y-- {
				grid[Coord{space.x, y}] = grid[Coord{space.x, y - 1}]
			}
			grid[Coord{pos.x, pos.y}] = "."
			return Coord{pos.x, pos.y + 1}

		} else if cmd == Right {
			for x := space.x; x > pos.x; x-- {
				grid[Coord{x, space.y}] = grid[Coord{x - 1, space.y}]
			}
			grid[Coord{pos.x, pos.y}] = "."
			return Coord{pos.x + 1, pos.y}

		} else if cmd == Left {
			for x := space.x; x < pos.x; x++ {
				grid[Coord{x, space.y}] = grid[Coord{x + 1, space.y}]
			}
			grid[Coord{pos.x, pos.y}] = "."
			return Coord{pos.x - 1, pos.y}
		}
	}
	return pos
}

func process(start Coord, cmds []string) {
	loc := start
	for _, line := range cmds {
		for _, cmdbyte := range line {
			cmd := todir(string(cmdbyte))
			loc = shift(loc, cmd)
		}
	}
}

func score() int {
	result := 0
	for k, v := range grid {
		if v == "O" {
			result += (100 * k.y) + k.x
		}
	}
	return result
}

// # = 35
// @ = 64
// O = 79
// ^ = 94
// v = 118
// < = 60
// > = 62
func findStart(lines []string) Coord {
	for y, v := range lines {
		x := strings.IndexByte(v, 64)
		if x != -1 {
			return Coord{x, y}
		}
	}
	return Coord{-1, -1}
}

func makeGrid(lines []string) {
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			grid[Coord{x, y}] = string(lines[y][x])
		}
	}
	ymax = len(lines)
	xmax = len(lines[0])
}

func displayGrid() {
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			fmt.Print(grid[Coord{x, y}])
		}
		fmt.Println("")
	}
}

func main() {

	var gridlines []string
	var cmdlines []string
	scanner := bufio.NewScanner(os.Stdin)

	ingrid := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			ingrid = false
			continue
		}
		if ingrid {
			gridlines = append(gridlines, line)
		} else {
			cmdlines = append(cmdlines, line)
		}
	}
	start := findStart(gridlines)
	fmt.Println(start)
	makeGrid(gridlines)
	displayGrid()
	process(start, cmdlines)
	displayGrid()
	fmt.Println(score())
}
