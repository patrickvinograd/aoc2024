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

func offMap(c Coord, territory []string) bool {
	if c.x < 0 || c.x >= len(territory[0]) {
		return true
	}
	if c.y < 0 || c.y >= len(territory) {
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
				mul := 0
				for true {
					node := Coord{second.x + (xdiff * mul), second.y + (ydiff * mul)}
					if offMap(node, terr) {
						break
					} else {
						nodes[node] = true
					}
					mul++
				}
				mul = 0
				for true {
					node := Coord{first.x - (xdiff * mul), first.y - (ydiff * mul)}
					if offMap(node, terr) {
						break
					} else {
						nodes[node] = true
					}
					mul++
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
	findNodes(terr)
	fmt.Println(len(nodes))
}
