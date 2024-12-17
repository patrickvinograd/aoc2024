package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"slices"
	"strings"
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

func turns(from Dir, to Dir) int {

	if from == Down {
		if to == Down {
			return 0
		} else if to == Left || to == Right {
			return 1
		} else {
			return 2
		}
	} else if from == Up {
		if to == Up {
			return 0
		} else if to == Left || to == Right {
			return 1
		} else {
			return 2
		}
	} else if from == Left {
		if to == Left {
			return 0
		} else if to == Up || to == Down {
			return 1
		} else {
			return 2
		}
	} else if from == Right {
		if to == Right {
			return 0
		} else if to == Up || to == Down {
			return 1
		} else {
			return 2
		}
	}
	panic("Unknown turns")
}

var grid = make(map[Coord]string)
var xmax int
var ymax int

var bestCost = make(map[Pos]int)
var bestPath = make(map[Pos][]Pos)
var searchList = list.New()

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

type Pos struct {
	x       int
	y       int
	heading Dir
}

type PathElem struct {
	x       int
	y       int
	heading Dir
	prev    *Pos
}

func costPlus(move Dir, currHeading Dir, endHeading Dir) int {
	return turns(currHeading, move)*1000 + 1 + turns(move, endHeading)*1000
}

func checkRoute(currentPos Pos, neighbor Coord, move Dir, desiredDir Dir) {
	delta := costPlus(move, currentPos.heading, desiredDir)
	newCost := bestCost[currentPos] + delta
	newPos := Pos{neighbor.x, neighbor.y, desiredDir}
	// fmt.Println("cost for new pos: ", newPos, newCost)
	moveCost, ok := bestCost[newPos]
	if !ok {
		bestCost[newPos] = newCost
		searchList.PushBack(newPos)
		bestPath[newPos] = append(bestPath[newPos], currentPos)
	} else {
		if newCost == moveCost && !slices.Contains(bestPath[newPos], currentPos) {
			bestPath[newPos] = append(bestPath[newPos], currentPos)
			// fmt.Println("equal path", newPos, bestPath[newPos])
		} else if newCost < moveCost {
			bestCost[newPos] = newCost
			bestPath[newPos] = make([]Pos, 0)
			bestPath[newPos] = append(bestPath[newPos], currentPos)
			// fmt.Println("new best path", newPos, bestPath[newPos])
			searchList.PushBack(newPos)
		}
	}
}

func trySearch(current Pos, direction Dir) {
	neighbor := mf[direction](Coord{current.x, current.y})
	if grid[neighbor] == "." || grid[neighbor] == "E" {
		// fmt.Println("searching -> ", current, neighbor)
		checkRoute(current, neighbor, direction, Up)
		checkRoute(current, neighbor, direction, Left)
		checkRoute(current, neighbor, direction, Down)
		checkRoute(current, neighbor, direction, Right)
	}
}

func search(start Coord, end Coord) {
	bestCost[Pos{start.x, start.y, Right}] = 0
	bestCost[Pos{start.x, start.y, Up}] = 1000
	bestCost[Pos{start.x, start.y, Down}] = 1000
	bestCost[Pos{start.x, start.y, Left}] = 2000

	searchList.PushBack(Pos{start.x, start.y, Right})

	e := searchList.Front()
	for e != nil {
		current := searchList.Remove(e).(Pos)
		// fmt.Println("searching from", current)
		trySearch(current, Up)
		trySearch(current, Left)
		trySearch(current, Down)
		trySearch(current, Right)

		e = searchList.Front()
	}
	// fmt.Println(bestCost)
	// fmt.Println(bestPath)
	// for k, v := range bestPath {
	// 	if len(v) > 1 {
	// 		fmt.Println(k, bestCost[k])
	// 		for _, pe := range v {
	// 			fmt.Println("    ", pe)
	// 		}
	// 	}
	// }
	backtrack(end)

}

func backtrack(end Coord) {
	best := Up
	cost := bestCost[Pos{end.x, end.y, Up}]
	if x := bestCost[Pos{end.x, end.y, Down}]; x < cost {
		best = Down
		cost = x
	}
	if x := bestCost[Pos{end.x, end.y, Left}]; x < cost {
		best = Left
		cost = x
	}
	if x := bestCost[Pos{end.x, end.y, Right}]; x < cost {
		best = Right
		cost = x
	}
	fmt.Println(best, cost)

	btList := list.New()
	btList.PushBack(Pos{end.x, end.y, best})
	e := btList.Front()
	for e != nil {
		btList.Remove(e)
		current := e.Value.(Pos)
		// fmt.Println("current: ", current)
		bestRoutes[Coord{current.x, current.y}] = true
		prevs := bestPath[current]
		for _, v := range prevs {
			// fmt.Println(v)
			if seen[v] != true {
				seen[v] = true
				btList.PushBack(v)
			}
		}
		seen[current] = true
		e = btList.Front()
	}
	displayGrid()
	// fmt.Println(bestRoutes)
	fmt.Println(len(bestRoutes))
}

var seen = make(map[Pos]bool)
var bestRoutes = make(map[Coord]bool)

// S = 83
func findStart(lines []string) Coord {
	for y, v := range lines {
		x := strings.IndexByte(v, 83)
		if x != -1 {
			return Coord{x, y}
		}
	}
	return Coord{-1, -1}
}

// E = 69
func findEnd(lines []string) Coord {
	for y, v := range lines {
		x := strings.IndexByte(v, 69)
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
			if bestRoutes[Coord{x, y}] == true {
				fmt.Print("O")
			} else {
				fmt.Print(grid[Coord{x, y}])
			}
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
	end := findEnd(gridlines)
	fmt.Println(start, end)
	makeGrid(gridlines)
	// displayGrid()
	search(start, end)
	// process(start, cmdlines)
	// displayGrid()
	// fmt.Println(score())
}
