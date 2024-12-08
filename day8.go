package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	x int
	y int
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

var antennas = make(map[byte][]Coord)
var nodes = make(map[Coord]bool)

func scanMap(lines []string) {
	for y, v := range lines {
		for x, b := range []byte(v) {
			if b == 46 {
				continue
			} else {
				antennas[b] = append(antennas[b], Coord{x, y})
			}
		}
	}
}

func findNodes(terr []string) {
	for _, ants := range antennas {
		for i := 0; i < len(ants); i++ {
			for j := i + 1; j < len(ants); j++ {
				first := ants[i]
				second := ants[j]
				xdiff := second.x - first.x
				ydiff := second.y - first.y
				node1 := Coord{second.x + xdiff, second.y + ydiff}
				node2 := Coord{first.x - xdiff, first.y - ydiff}
				if !offMap(node1.x, node1.y, terr) {
					// fmt.Println(freq, node1)
					nodes[node1] = true
				}
				if !offMap(node2.x, node2.y, terr) {
					// fmt.Println(freq, node2)
					nodes[node2] = true
				}
			}
		}
	}
}

func main() {

	var terr []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		terr = append(terr, line)
	}

	scanMap(terr)
	// fmt.Println(antennas)
	findNodes(terr)
	fmt.Println(len(nodes))
}
