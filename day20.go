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

func scanShortcuts() {
	for x := 0; x < xmax; x++ {
		for y := 0; y < ymax; y++ {
			curr := Coord{x, y}
			if grid[curr] != "#" {
				continue
			}
			for _, rel1 := range alldirs {
				for _, rel2 := range alldirs {
					p1 := mf[rel1](curr)
					p2 := mf[rel2](curr)
					pc1, on1 := pathCost[p1]
					pc2, on2 := pathCost[p2]
					if on1 && on2 && pc2-pc1-2 > 0 {
						shortcuts[Shortcut{curr, p2}] = pc2 - pc1 - 2
						// fmt.Println("shortcut", curr, p2, pc2-pc1-2)
					}
				}
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
	displayGrid()
	findPath(start, end)
	fmt.Println(pathCost)
	fmt.Println(end, pathCost[end])
	scanShortcuts()
	fmt.Println(summarize())
}
