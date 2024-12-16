package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type Dir int

const (
	Up Dir = iota
	Left
	Down
	Right
)

var m = map[Dir]func(int, int) (int, int){
	Up:    func(x int, y int) (int, int) { return x, y - 1 },
	Down:  func(x int, y int) (int, int) { return x, y + 1 },
	Left:  func(x int, y int) (int, int) { return x - 1, y },
	Right: func(x int, y int) (int, int) { return x + 1, y },
}

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
	prev    *Pos
}

func costPlus(move Dir, currHeading Dir, endHeading Dir) int {
	return turns(currHeading, move)*1000 + 1 + turns(move, endHeading)*1000
}

func checkRoute(currentPos Pos, neighbor Coord, move Dir, desiredDir Dir) {
	delta := costPlus(move, currentPos.heading, desiredDir)
	newCost := bestCost[currentPos] + delta
	newPos := Pos{neighbor.x, neighbor.y, desiredDir, nil}
	// fmt.Println("cost for new pos: ", newPos, newCost)
	moveCost, ok := bestCost[newPos]
	if !ok {
		bestCost[newPos] = newCost
		searchList.PushBack(newPos)
	} else {
		if newCost < moveCost {
			bestCost[newPos] = newCost
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
	bestCost[Pos{start.x, start.y, Right, nil}] = 0
	bestCost[Pos{start.x, start.y, Up, nil}] = 1000
	bestCost[Pos{start.x, start.y, Down, nil}] = 1000
	bestCost[Pos{start.x, start.y, Left, nil}] = 2000

	searchList.PushBack(Pos{start.x, start.y, Right, nil})

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
	fmt.Println(bestCost[Pos{end.x, end.y, Up, nil}])
	fmt.Println(bestCost[Pos{end.x, end.y, Left, nil}])
	fmt.Println(bestCost[Pos{end.x, end.y, Down, nil}])
	fmt.Println(bestCost[Pos{end.x, end.y, Right, nil}])

}

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
	end := findEnd(gridlines)
	fmt.Println(start, end)
	makeGrid(gridlines)
	// displayGrid()
	search(start, end)
	// process(start, cmdlines)
	// displayGrid()
	// fmt.Println(score())
}
