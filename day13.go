package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Game struct {
	ax     int
	ay     int
	bx     int
	by     int
	prizex int
	prizey int
}

// Button A: X+72, Y+24
// Button B: X+25, Y+73
// Prize: X=17782, Y=9862

func play(game Game) (int, bool) {
	bestScore := 500
	sat := false
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i*game.ax+j*game.bx == game.prizex && i*game.ay+j*game.by == game.prizey {
				tokens := 3*i + j
				if tokens < bestScore {
					bestScore = tokens
					sat = true
				}
			}
		}
	}
	return bestScore, sat
}

func playAll(games []Game) int {
	total := 0
	for _, game := range games {
		tokens, won := play(game)
		if won {
			fmt.Println("Won game:", game, tokens)
			total += tokens
		}
	}
	return total
}

func main() {

	var data []Game
	scanner := bufio.NewScanner(os.Stdin)

	re_a := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	re_b := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	re_p := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		matches := re_a.FindStringSubmatch(line)
		fmt.Println(matches)
		ax, _ := strconv.Atoi(string(matches[1]))
		ay, _ := strconv.Atoi(string(matches[2]))

		scanner.Scan()
		line = scanner.Text()
		fmt.Println(line)
		matches = re_b.FindStringSubmatch(line)
		fmt.Println(matches)
		bx, _ := strconv.Atoi(string(matches[1]))
		by, _ := strconv.Atoi(string(matches[2]))

		scanner.Scan()
		line = scanner.Text()
		fmt.Println(line)
		matches = re_p.FindStringSubmatch(line)
		fmt.Println(matches)
		px, _ := strconv.Atoi(string(matches[1]))
		py, _ := strconv.Atoi(string(matches[2]))

		game := Game{ax, ay, bx, by, px, py}
		data = append(data, game)

		scanner.Scan()
		line = scanner.Text()
		fmt.Println(line)
	}

	fmt.Println(data)
	fmt.Println(playAll(data))

}
