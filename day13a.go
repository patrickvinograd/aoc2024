package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

type Game struct {
	ax     int
	ay     int
	bx     int
	by     int
	prizex int
	prizey int
}

// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400

// 94a + 22b = 8400
// 34a + 67b = 5400

func play(game Game) (int64, bool) {

	// Create the matrix A and vector b.
	// In this example, we will solve the equations
	// 2x + y = 3
	// x - 3y = 5

	A := mat.NewDense(2, 2, []float64{float64(game.ax), float64(game.bx), float64(game.ay), float64(game.by)})
	b := mat.NewVecDense(2, []float64{float64(game.prizex), float64(game.prizey)})

	// Solve the equations using the Solve function.
	var x mat.VecDense
	if err := x.SolveVec(A, b); err != nil {
		fmt.Println(err)
		return 0, false
	}
	// fmt.Println(x)
	r1 := int64(math.Round(x.AtVec(0)))
	r2 := int64(math.Round(x.AtVec(1)))
	if r1*int64(game.ax)+r2*int64(game.bx) == int64(game.prizex) && r1*int64(game.ay)+r2*int64(game.by) == int64(game.prizey) {
		fmt.Println(r1, r2, 3*r1+r2)
		return 3*r1 + r2, true
	} else {
		// fmt.Println("lost")
		return -1, false
	}
}

func playAll(games []Game) int64 {
	total := int64(0)
	for _, game := range games {
		tokens, won := play(game)
		if won {
			fmt.Println("Won game:", game, tokens)
			total += int64(tokens)
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
		// fmt.Println(line)
		matches := re_a.FindStringSubmatch(line)
		// fmt.Println(matches)
		ax, _ := strconv.Atoi(string(matches[1]))
		ay, _ := strconv.Atoi(string(matches[2]))

		scanner.Scan()
		line = scanner.Text()
		// fmt.Println(line)
		matches = re_b.FindStringSubmatch(line)
		// fmt.Println(matches)
		bx, _ := strconv.Atoi(string(matches[1]))
		by, _ := strconv.Atoi(string(matches[2]))

		scanner.Scan()
		line = scanner.Text()
		// fmt.Println(line)
		matches = re_p.FindStringSubmatch(line)
		// fmt.Println(matches)
		px, _ := strconv.Atoi(string(matches[1]))
		py, _ := strconv.Atoi(string(matches[2]))

		game := Game{ax, ay, bx, by, px + 10000000000000, py + 10000000000000}
		// game := Game{ax, ay, bx, by, px, py}
		data = append(data, game)

		scanner.Scan()
		line = scanner.Text()
		// fmt.Println(line)
	}

	// fmt.Println(data)
	fmt.Println(playAll(data))

}
