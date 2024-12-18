package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
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

var bytes []Coord
var grid = make(map[Coord]string)

var xmax int = 70
var ymax int = 70
var bytecount int = 1024

// var xmax int = 6
// var ymax int = 6
// var bytecount int = 12

// reset per meta-search
var bestCost = make(map[Coord]int)
var searchList = list.New()

func offMap(c Coord) bool {
	if c.x < 0 || c.x > xmax {
		return true
	}
	if c.y < 0 || c.y > ymax {
		return true
	}
	return false
}

type Coord struct {
	x int
	y int
}

func checkRoute(currentPos Coord, neighbor Coord) {
	delta := 1
	newCost := bestCost[currentPos] + delta
	moveCost, ok := bestCost[neighbor]
	if !ok {
		bestCost[neighbor] = newCost
		searchList.PushBack(neighbor)
	} else {
		if newCost < moveCost {
			bestCost[neighbor] = newCost
			searchList.PushBack(neighbor)
		}
	}
}

func trySearch(current Coord, direction Dir) {
	neighbor := mf[direction](current)
	if grid[neighbor] != "#" && !offMap(neighbor) {
		// fmt.Println("searching -> ", current, neighbor)
		checkRoute(current, neighbor)
	}
}

func search(start Coord, end Coord) {
	bestCost[start] = 0

	searchList.PushBack(start)

	for e := searchList.Front(); e != nil; e = searchList.Front() {
		current := searchList.Remove(e).(Coord)
		// fmt.Println("searching from", current)
		trySearch(current, Up)
		trySearch(current, Left)
		trySearch(current, Down)
		trySearch(current, Right)
	}

}

func metaSearch(start Coord, end Coord) {
	for i := bytecount; i < len(bytes); i++ {
		grid[bytes[i]] = "#"
		bestCost = make(map[Coord]int)
		searchList = list.New()
		search(start, end)

		_, ok := bestCost[end]
		if !ok {
			// displayGrid()
			fmt.Printf("No route at byte %d: %d,%d\n", i, bytes[i].x, bytes[i].y)
			return
		}
	}
}

func makeGrid(lines []string) {
	bytes = make([]Coord, 0)
	for _, l := range lines {
		vstrs := strings.Split(l, ",")
		xs := vstrs[0]
		ys := vstrs[1]
		x, _ := strconv.Atoi(xs)
		y, _ := strconv.Atoi(ys)
		bytes = append(bytes, Coord{x, y})
	}
	for i := 0; i < bytecount; i++ {
		grid[bytes[i]] = "#"
	}
}

func displayGrid() {
	for y := 0; y <= ymax; y++ {
		for x := 0; x <= xmax; x++ {
			b, ok := grid[Coord{x, y}]
			if ok {
				fmt.Print(b)
			} else {
				fmt.Print(".")
			}
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
	start := Coord{0, 0}
	end := Coord{xmax, ymax}

	fmt.Println(start, end)
	makeGrid(lines)
	displayGrid()
	metaSearch(start, end)
}
