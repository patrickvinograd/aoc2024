package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type XY struct {
	x int
	y int
}

var terr []string
var seen [][]bool
var idCounter int = 0

var regionMap map[int]*list.List = make(map[int]*list.List)

func offMap(x int, y int) bool {
	if x < 0 || x >= len(terr[0]) {
		return true
	}
	if y < 0 || y >= len(terr) {
		return true
	}
	return false
}

func initSeen(terr []string) {
	seen = make([][]bool, len(terr))
	for i := 0; i < len(terr); i++ {
		seen[i] = make([]bool, len(terr[i]))
	}
}

func addAdjacent(x int, y int, val byte) bool {
	if offMap(x, y) {
		return false
	}
	if seen[y][x] == true {
		return false
	}
	return (terr[y][x] == val)
}

func bfs(searchList *list.List, val byte, id int) {
	e := searchList.Front()
	for e != nil {
		x := e.Value.(XY).x
		y := e.Value.(XY).y
		regionMap[id].PushBack(XY{x, y})
		seen[y][x] = true
		searchList.Remove(e)
		if addAdjacent(x-1, y, val) {
			searchList.PushBack(XY{x - 1, y})
			seen[y][x-1] = true
		}
		if addAdjacent(x, y-1, val) {
			searchList.PushBack(XY{x, y - 1})
			seen[y-1][x] = true
		}
		if addAdjacent(x+1, y, val) {
			searchList.PushBack(XY{x + 1, y})
			seen[y][x+1] = true
		}
		if addAdjacent(x, y+1, val) {
			searchList.PushBack(XY{x, y + 1})
			seen[y+1][x] = true
		}
		e = searchList.Front()
	}
}

func startBfs(x int, y int) {
	// fmt.Println(seen)
	id := idCounter
	regionMap[id] = list.New()
	val := terr[y][x]
	searchList := list.New()
	searchList.PushBack(XY{x, y})
	bfs(searchList, val, id)
	idCounter++
}

func scan(terr []string) {
	for y := 0; y < len(terr); y++ {
		for x := 0; x < len(terr[0]); x++ {
			if seen[y][x] == true {
				continue
			}
			startBfs(x, y)
			// printRegions()
		}
	}
}

func borders(x int, y int) int {
	result := 0
	val := terr[y][x]
	if offMap(x-1, y) || terr[y][x-1] != val {
		result++
	}
	if offMap(x, y-1) || terr[y-1][x] != val {
		result++
	}
	if offMap(x+1, y) || terr[y][x+1] != val {
		result++
	}
	if offMap(x, y+1) || terr[y+1][x] != val {
		result++
	}
	return result
}

func scoreRegions() int {
	total := 0
	for k, v := range regionMap {
		area := v.Len()
		perim := 0
		for e := v.Front(); e != nil; e = e.Next() {
			perim += borders(e.Value.(XY).x, e.Value.(XY).y)
		}
		score := area * perim
		fmt.Println(k, score)
		total += score
	}
	return total
}

func printRegions() {
	for k, v := range regionMap {
		fmt.Print(k, " ")
		printList(v)
	}
}

func printList(l *list.List) {
	current := l.Front()
	for current != nil {
		fmt.Printf("%d,%d ", current.Value.(XY).x, current.Value.(XY).y)
		current = current.Next()
	}
	fmt.Println("")
}

func main() {

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	terr = lines
	initSeen(lines)
	scan(lines)
	fmt.Println(scoreRegions())
}
