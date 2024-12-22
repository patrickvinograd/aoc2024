package main

import (
	"bufio"
	"fmt"
	"math"
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

var mf = map[string]func(Coord) Coord{
	"^": func(pt Coord) Coord { return Coord{pt.x, pt.y - 1} },
	"v": func(pt Coord) Coord { return Coord{pt.x, pt.y + 1} },
	"<": func(pt Coord) Coord { return Coord{pt.x - 1, pt.y} },
	">": func(pt Coord) Coord { return Coord{pt.x + 1, pt.y} },
}

type Coord struct {
	x int
	y int
}

type Move struct {
	start string
	end   string
}

func permutations(arr []byte) [][]byte {
	var helper func([]byte, int)
	res := [][]byte{}

	helper = func(arr []byte, n int) {
		if n == 1 {
			tmp := make([]byte, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

var numpad = make(map[string]Coord)
var dpad = make(map[string]Coord)

func planNumMove(start string, end string) string {
	result := ""
	sc := numpad[start]
	ec := numpad[end]
	// fmt.Println(start, end, sc, ec)
	curx := sc.x
	cury := sc.y
	if sc.x == 0 && ec.y == 3 {
		for i := 0; i < ec.x-sc.x; i++ {
			result += ">"
			curx++
		}
	}
	if sc.y == 3 && ec.x == 0 {
		for i := 0; i < sc.y-ec.y; i++ {
			result += "^"
			cury--
		}
	}
	if cury > ec.y {
		diff := cury - ec.y
		for i := 0; i < diff; i++ {
			result += "^"
			cury--
		}
	} else if cury < ec.y {
		diff := ec.y - cury
		for i := 0; i < diff; i++ {
			result += "v"
			cury++
		}
	}
	if curx > ec.x {
		diff := curx - ec.x
		for i := 0; i < diff; i++ {
			result += "<"
			curx--
		}
	} else if curx < ec.x {
		diff := ec.x - curx
		for i := 0; i < diff; i++ {
			result += ">"
			curx++
		}
	}
	result += "A"
	return result
}

func valid(perm []byte, move Move) bool {
	pos := numpad[move.start]
	for _, b := range perm {
		pos = mf[string(b)](pos)
		if pos.x == 0 && pos.y == 3 {
			return false
		}
	}
	return true
}

func bestPerm(perms [][]byte, move Move) (string, int) {
	best := math.MaxInt
	bestMoves := ""
	for _, perm := range perms {
		if !valid(perm, move) {
			continue
		}
		layer2moves, _ := planDcode(string(perm)+"A", "A")
		// fmt.Println(layer2moves)
		layer3moves, _ := planDcode(layer2moves, "A")
		// fmt.Println(layer3moves)
		if len(layer3moves) < best {
			best = len(layer3moves)
			bestMoves = layer3moves
		}
	}
	return bestMoves, best
}

var bestNumMoves = make(map[Move]int)

func findBestNumMove(start string, end string) (string, int) {
	fmt.Println("planning best ", start, end)
	moves := ""
	sc := numpad[start]
	ec := numpad[end]
	if ec.x > sc.x {
		moves += strings.Repeat(">", ec.x-sc.x)
	} else if ec.x < sc.x {
		moves += strings.Repeat("<", sc.x-ec.x)
	}
	if ec.y > sc.y {
		moves += strings.Repeat("v", ec.y-sc.y)
	} else if ec.y < sc.y {
		moves += strings.Repeat("^", sc.y-ec.y)
	}
	fmt.Println("abstract moves", moves)
	perms := permutations([]byte(moves))
	bm, score := bestPerm(perms, Move{start, end})
	fmt.Println(start, end, score, bm)
	bestNumMoves[Move{start, end}] = score
	return bm, score
}

func planNumcode(code string, start string) (moves string, score int) {
	fmt.Println("\nOUTER", code)
	result := ""
	current := start
	for _, elem := range strings.Split(code, "") {
		next := string(elem)
		moves, _ := findBestNumMove(current, next)
		result += moves
		current = next
	}
	s := getScore(result, code)
	return result, s
}

func planDMove(start string, end string) string {
	result := ""
	sc := dpad[start]
	ec := dpad[end]
	// fmt.Println(start, end, sc, ec)
	curx := sc.x
	cury := sc.y
	if sc.x == 0 && ec.y == 0 {
		for i := 0; i < ec.x-sc.x; i++ {
			result += ">"
			curx++
		}
	}
	if sc.y == 0 && ec.x == 0 {
		for i := 0; i < sc.y-ec.y; i++ {
			result += "v"
			cury++
		}
	}
	if cury > ec.y {
		diff := cury - ec.y
		for i := 0; i < diff; i++ {
			result += "^"
			cury--
		}
	} else if cury < ec.y {
		diff := ec.y - cury
		for i := 0; i < diff; i++ {
			result += "v"
			cury++
		}
	}
	if curx > ec.x {
		diff := curx - ec.x
		for i := 0; i < diff; i++ {
			result += "<"
			curx--
		}
	} else if curx < ec.x {
		diff := ec.x - curx
		for i := 0; i < diff; i++ {
			result += ">"
			curx++
		}
	}
	result += "A"
	// fmt.Println(start, end, result)
	return result
}

func planDcode(code string, start string) (moves string, end string) {
	result := ""
	current := start
	for _, elem := range strings.Split(code, "") {
		next := string(elem)
		result = result + planDMove(current, next)
		current = next
	}
	return result, current
}

func getScore(moves string, code string) int {
	numpart, _ := strings.CutSuffix(code, "A")
	num, _ := strconv.Atoi(numpart)
	result := num * len(moves)
	fmt.Println(len(moves), num, result, "\n")
	return result

}

func makePads() {
	numpad["7"] = Coord{0, 0}
	numpad["8"] = Coord{1, 0}
	numpad["9"] = Coord{2, 0}
	numpad["4"] = Coord{0, 1}
	numpad["5"] = Coord{1, 1}
	numpad["6"] = Coord{2, 1}
	numpad["1"] = Coord{0, 2}
	numpad["2"] = Coord{1, 2}
	numpad["3"] = Coord{2, 2}
	numpad["0"] = Coord{1, 3}
	numpad["A"] = Coord{2, 3}

	dpad["^"] = Coord{1, 0}
	dpad["A"] = Coord{2, 0}
	dpad["<"] = Coord{0, 1}
	dpad["v"] = Coord{1, 1}
	dpad[">"] = Coord{2, 1}
}

func main() {

	var codes []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		codes = append(codes, line)
	}
	makePads()
	total := 0
	// for i := 1; i < len(codes[0]); i++ {
	// 	findBestNumMove(string(codes[0][i-1]), string(codes[0][i]))
	// }
	for _, code := range codes {
		_, score := planNumcode(code, "A")
		total += score
	}
	fmt.Println(total)

}
