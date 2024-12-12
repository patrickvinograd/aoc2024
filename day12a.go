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

var cMap = make(map[XY]int)

func valAt(x int, y int) byte {
	if offMap(x, y) {
		return 0
	} else {
		return terr[y][x]
	}
}

func countCorners() {
	for y := 0; y < len(terr)+1; y++ {
		for x := 0; x < len(terr[0])+1; x++ {
			ul := valAt(x-1, y-1)
			ur := valAt(x, y-1)
			ll := valAt(x-1, y)
			lr := valAt(x, y)

			if ul == ur && ul == ll && ul == lr {
				continue // mid-region, no corners
			}
			if ul == ur {
				if ll == ul {
					cMap[XY{x - 1, y - 1}] += 1
					cMap[XY{x, y}] += 1
				} else if lr == ul {
					cMap[XY{x - 1, y - 1}] += 1
					cMap[XY{x - 1, y}] += 1
				} else if ll != lr {
					cMap[XY{x - 1, y}] += 1
					cMap[XY{x, y}] += 1
				}
			} else if ll == lr {
				if ul == ll {
					cMap[XY{x - 1, y}] += 1
					cMap[XY{x, y - 1}] += 1
				} else if ur == ll {
					cMap[XY{x - 1, y}] += 1
					cMap[XY{x - 1, y - 1}] += 1
				} else if ul != ur {
					cMap[XY{x - 1, y - 1}] += 1
					cMap[XY{x, y - 1}] += 1
				}
			} else if ul == ll {
				if ul == ur {
					cMap[XY{x - 1, y - 1}] += 1
					cMap[XY{x, y}] += 1
				} else if ul == lr {
					cMap[XY{x - 1, y - 1}] += 1
					cMap[XY{x, y - 1}] += 1
				} else if ur != lr {
					cMap[XY{x, y - 1}] += 1
					cMap[XY{x, y}] += 1
				}
			} else if ur == lr {
				if ul == ur {
					cMap[XY{x, y - 1}] += 1
					cMap[XY{x - 1, y}] += 1
				} else if ll == ur {
					cMap[XY{x, y - 1}] += 1
					cMap[XY{x - 1, y - 1}] += 1
				} else if ul != ll {
					cMap[XY{x - 1, y - 1}] += 1
					cMap[XY{x - 1, y}] += 1
				}
			} else {
				cMap[XY{x - 1, y - 1}] += 1
				cMap[XY{x, y - 1}] += 1
				cMap[XY{x - 1, y}] += 1
				cMap[XY{x, y}] += 1
			}
		}
	}
}

func scoreRegions() int {
	total := 0
	for _, v := range regionMap {
		area := v.Len()
		corners := 0
		for e := v.Front(); e != nil; e = e.Next() {
			corners += cMap[e.Value.(XY)]
		}
		score := area * corners
		total += score
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
	terr = lines
	initSeen(lines)
	scan(lines)
	countCorners()
	// fmt.Println(cMap)
	fmt.Println(scoreRegions())
}
