package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type Dir int

const (
	Up Dir = iota
	Left
	Down
	Right
)

var mf = map[Dir]func(Coord) Coord{
	Up:    func(pt Coord) Coord { return Coord{pt.x, pt.y - 1} },
	Down:  func(pt Coord) Coord { return Coord{pt.x, pt.y + 1} },
	Left:  func(pt Coord) Coord { return Coord{pt.x - 1, pt.y} },
	Right: func(pt Coord) Coord { return Coord{pt.x + 1, pt.y} },
}

type Coord struct {
	x int
	y int
}

type Shortcut struct {
	one Coord
	two Coord
}

var grid = make(map[Coord]string)
var xmax int
var ymax int

var pathCost = make(map[Coord]int)
var shortcuts = make(map[Shortcut]int)
var searchList = list.New()

func offMap(c Coord) bool {
	if c.x < 0 || c.x >= xmax {
		return true
	}
	if c.y < 0 || c.y >= ymax {
		return true
	}
	return false
}

var alldirs = []Dir{Up, Left, Down, Right}

func tryNext(current Coord) Coord {
	for _, dir := range alldirs {
		neighbor := mf[dir](current)
		if grid[neighbor] != "#" && !offMap(neighbor) {
			_, ok := pathCost[neighbor]
			if !ok {
				return neighbor
			}
		}
	}
	panic("No path found")
}

func findPath(start Coord, end Coord) {
	pathCost[start] = 0
	current := start
	for current != end {
		next := tryNext(current)
		pathCost[next] = pathCost[current] + 1
		current = next
	}
}

func mdist(p1 Coord, p2 Coord) int {
	xdiff := p1.x - p2.x
	if xdiff < 0 {
		xdiff = xdiff * -1
	}
	ydiff := p1.y - p2.y
	if ydiff < 0 {
		ydiff = ydiff * -1
	}
	return xdiff + ydiff
}

func scanShortcuts() {
	for spoint, scost := range pathCost {
		for epoint, ecost := range pathCost {
			slen := mdist(spoint, epoint)
			savings := ecost - scost - slen
			if slen <= 20 && savings > 0 {
				shortcuts[Shortcut{spoint, epoint}] = savings
			}
		}
	}
}

func summarize() int {
	count := 0
	counts := make(map[int]int)
	for _, v := range shortcuts {
		counts[v]++
		if v >= 100 {
			count++
		}
	}
	fmt.Println(counts)
	return count
}

func makeGrid(lines []string) (Coord, Coord) {
	var start, end Coord
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			val := string(lines[y][x])
			grid[Coord{x, y}] = val
			if val == "S" {
				start = Coord{x, y}
			}
			if val == "E" {
				end = Coord{x, y}
			}
		}
	}
	ymax = len(lines)
	xmax = len(lines[0])
	return start, end
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

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	start, end := makeGrid(lines)
	fmt.Println(start, end)
	// displayGrid()
	findPath(start, end)
	// fmt.Println(pathCost)
	fmt.Println(end, pathCost[end])
	scanShortcuts()
	fmt.Println(summarize())
}
