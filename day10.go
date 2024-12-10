package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func offMap(x int, y int, territory [][]int) bool {
	if x < 0 || x >= len(territory[0]) {
		return true
	}
	if y < 0 || y >= len(territory) {
		return true
	}
	return false
}

type XY struct {
	x int
	y int
}

type Coord struct {
	x      int
	y      int
	elev   int
	origin *Coord
}

var searchList *list.List = list.New()
var trailheads = make(map[XY][]Coord)

func toTerr(lines []string) [][]int {
	result := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, b := range line {
			row[j], _ = strconv.Atoi(string(b))
		}
		result[i] = row
	}
	return result
}

func findNines(terr [][]int) {
	for y := 0; y < len(terr); y++ {
		for x := 0; x < len(terr[0]); x++ {
			if terr[y][x] == 9 {
				nine := Coord{x, y, 9, nil}
				nine.origin = &nine
				searchList.PushBack(nine)
			}
		}
	}
}

func checkN(x int, y int, neighborElev int, terr [][]int, origin *Coord) bool {
	if offMap(x, y, terr) {
		return false
	}
	elev := terr[y][x]
	if elev == neighborElev-1 {
		push := Coord{x, y, elev, origin}
		searchList.PushBack(push)
		return true
	}
	return false
}

func checkNeighbors(current Coord, terr [][]int) {
	checkN(current.x-1, current.y, current.elev, terr, current.origin)
	checkN(current.x, current.y-1, current.elev, terr, current.origin)
	checkN(current.x+1, current.y, current.elev, terr, current.origin)
	checkN(current.x, current.y+1, current.elev, terr, current.origin)
}

func pushTrailhead(current Coord) {
	key := XY{current.x, current.y}
	trailheads[key] = append(trailheads[key], current)
}

func search(terr [][]int) {
	findNines(terr)
	for true {
		if searchList.Len() == 0 {
			break
		}
		current := searchList.Remove(searchList.Front()).(Coord)
		if current.elev == 0 {
			pushTrailhead(current)
		} else {
			checkNeighbors(current, terr)
		}
	}
}

func scoreTrailheads() int {
	total := 0
	for k, v := range trailheads {
		uniques := make(map[XY]bool)
		for _, route := range v {
			uniques[XY{route.origin.x, route.origin.y}] = true
		}
		fmt.Println(k, uniques)
		total += len(uniques)
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
	terr := toTerr(lines)
	// fmt.Println(terr)
	search(terr)
	// fmt.Println(trailheads)
	fmt.Println(scoreTrailheads())
}
