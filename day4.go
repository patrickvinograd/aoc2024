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

func checkLoc(lines []string, xmax int, ymax int, xp int, yp int, val byte) bool {
	if yp < 0 || yp >= ymax {
		return false
	}
	if xp < 0 || xp >= xmax {
		return false
	}
	return (lines[yp][xp] == val)
}

// X = 88
// M = 77
// A = 65
// S = 83
func search(lines []string, xmax int, ymax int, x int, y int, xmul int, ymul int) bool {

	if checkLoc(lines, xmax, ymax, x, y, 88) &&
		checkLoc(lines, xmax, ymax, x+(xmul*1), y+(ymul*1), 77) &&
		checkLoc(lines, xmax, ymax, x+(xmul*2), y+(ymul*2), 65) &&
		checkLoc(lines, xmax, ymax, x+(xmul*3), y+(ymul*3), 83) {
		return true
	}
	return false
}

func findXmas(lines []string) int {
	total := 0
	xmax := len(lines[0])
	ymax := len(lines)
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			for xi := -1; xi <= 1; xi++ {
				for yi := -1; yi <= 1; yi++ {
					// fmt.Println(x, y, xi, yi)
					if search(lines, xmax, ymax, x, y, xi, yi) {
						// fmt.Println(x, y, xi, yi)
						total++
					}
				}
			}
		}
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
	fmt.Println(findXmas(lines))
}
