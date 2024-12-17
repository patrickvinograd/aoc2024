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

func isBox(pos Coord, dir Dir) bool {
	return grid[mc[dir](pos)] == "[" || grid[mc[dir](pos)] == "]"
}

func isBlocked(boxpos Coord, dir Dir) bool {
	bc1 := boxpos
	bc2 := Coord{boxpos.x + 1, boxpos.y}
	return grid[mc[dir](bc1)] == "#" || grid[mc[dir](bc2)] == "#"
}

func anyBlocked(convoy *list.List, dir Dir) bool {
	blocked := false
	for e := convoy.Front(); e != nil; e = e.Next() {
		bc1 := e.Value.(Coord)
		if isBlocked(bc1, dir) {
			blocked = true
		}
	}
	return blocked
}

func shiftBoxes(convoy *list.List, dir Dir) {
	for e := convoy.Back(); e != nil; e = e.Prev() {
		bc1 := e.Value.(Coord)
		bc2 := Coord{bc1.x + 1, bc1.y}
		grid[mc[dir](bc1)] = "["
		grid[mc[dir](bc2)] = "]"
		grid[bc1] = "."
		grid[bc2] = "."
	}
}

func convoy(pos Coord, dir Dir) *list.List {
	result := list.New()
	searchList := list.New()
	firstBox := mc[dir](pos)
	if grid[firstBox] == "]" {
		firstBox = Coord{firstBox.x - 1, firstBox.y}
	}
	searchList.PushBack(firstBox)
	for e := searchList.Front(); e != nil; e = searchList.Front() {
		searchList.Remove(e)
		box := e.Value.(Coord)
		result.PushBack(box)
		search1 := mc[dir](box)
		if grid[search1] == "[" {
			searchList.PushBack(search1)
		} else if grid[search1] == "]" {
			nbox := Coord{search1.x - 1, search1.y}
			searchList.PushBack(nbox)
		}
		search2 := mc[dir](box)
		search2 = Coord{search2.x + 1, search2.y}
		if grid[search2] == "[" {
			searchList.PushBack(search2)
		}
	}
	return result
}

func shiftUp(pos Coord) (Coord, bool) {
	if grid[mc[Up](pos)] == "#" {
		return pos, false
	} else if grid[mc[Up](pos)] == "." {
		newBot := mc[Up](pos)
		grid[newBot] = "@"
		return newBot, true
	} else if isBox(pos, Up) {
		boxes := convoy(pos, Up)
		if !anyBlocked(boxes, Up) {
			shiftBoxes(boxes, Up)
			newBot := mc[Up](pos)
			grid[newBot] = "@"
			return newBot, true
		}
	}
	return pos, false
}

func shiftDown(pos Coord) (Coord, bool) {
	if grid[mc[Down](pos)] == "#" {
		return pos, false
	} else if grid[mc[Down](pos)] == "." {
		newBot := mc[Down](pos)
		grid[newBot] = "@"
		return newBot, true
	} else if isBox(pos, Down) {
		boxes := convoy(pos, Down)
		if !anyBlocked(boxes, Down) {
			shiftBoxes(boxes, Down)
			newBot := mc[Down](pos)
			grid[newBot] = "@"
			return newBot, true
		}
	}
	return pos, false
}

func shift(pos Coord, cmd Dir) Coord {
	// fmt.Println(pos, cmd, space, canMove)
	if cmd == Up {
		_, shifted := shiftUp(pos)
		if shifted {
			grid[Coord{pos.x, pos.y}] = "."
			return Coord{pos.x, pos.y - 1}
		} else {
			return pos
		}
	} else if cmd == Down {
		_, shifted := shiftDown(pos)
		if shifted {
			grid[Coord{pos.x, pos.y}] = "."
			return Coord{pos.x, pos.y + 1}
		} else {
			return pos
		}
	} else if cmd == Right {
		space, canMove := firstSpace(pos, cmd)
		if canMove {
			for x := space.x; x > pos.x; x-- {
				grid[Coord{x, space.y}] = grid[Coord{x - 1, space.y}]
			}
			grid[Coord{pos.x, pos.y}] = "."
			return Coord{pos.x + 1, pos.y}
		} else {
			return pos
		}
	} else if cmd == Left {
		space, canMove := firstSpace(pos, cmd)
		if canMove {
			for x := space.x; x < pos.x; x++ {
				grid[Coord{x, space.y}] = grid[Coord{x + 1, space.y}]
			}
			grid[Coord{pos.x, pos.y}] = "."
			return Coord{pos.x - 1, pos.y}
		} else {
			return pos
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
		if v == "[" {
			result += (100 * k.y) + k.x
		}
	}
	return result
}

func findStart() Coord {
	for k, v := range grid {
		if v == "@" {
			return k
		}
	}
	return Coord{-1, -1}
}

func makeGrid(lines []string) {
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if string(lines[y][x]) == "#" {
				grid[Coord{x * 2, y}] = "#"
				grid[Coord{x*2 + 1, y}] = "#"
			} else if string(lines[y][x]) == "O" {
				grid[Coord{x * 2, y}] = "["
				grid[Coord{x*2 + 1, y}] = "]"
			} else if string(lines[y][x]) == "." {
				grid[Coord{x * 2, y}] = "."
				grid[Coord{x*2 + 1, y}] = "."
			} else if string(lines[y][x]) == "@" {
				grid[Coord{x * 2, y}] = "@"
				grid[Coord{x*2 + 1, y}] = "."
			}
		}
	}
	ymax = len(lines)
	xmax = len(lines[0]) * 2
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
	makeGrid(gridlines)
	start := findStart()
	fmt.Println(start)
	displayGrid()
	process(start, cmdlines)
	displayGrid()
	fmt.Println(score())
}
