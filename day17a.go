package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
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

var iptr int = 0
var rega int
var regb int
var regc int
var output = make([]int, 0)

type State struct {
	iptr   int
	rega   int
	regb   int
	regc   int
	output []int
}

var instructions = make([]int, 0)
var rev []int

var ftable = map[int]func(int, *State) bool{
	0: adv,
	1: bxl,
	2: bst,
	3: jnz,
	4: bxc,
	5: out,
	6: bdv,
	7: cdv,
}

func combo(operand int, state *State) int {
	if operand == 0 || operand == 1 || operand == 2 || operand == 3 {
		return operand
	} else if operand == 4 {
		return state.rega
	} else if operand == 5 {
		return state.regb
	} else if operand == 6 {
		return state.regc
	} else if operand == 7 {
		panic("unknown operand")
	}
	panic("unknown operand")
}

func adv(operand int, state *State) bool {
	num := state.rega
	denom := math.Pow(2, float64(combo(operand, state)))
	state.rega = int(float64(num) / denom)
	return true
}

func bxl(operand int, state *State) bool {
	state.regb = state.regb ^ operand
	return true
}

func bst(operand int, state *State) bool {
	state.regb = (combo(operand, state)%8 + 8) % 8
	return true
}

func jnz(operand int, state *State) bool {
	if state.rega == 0 {
		return true
	} else {
		state.iptr = operand
		return false
	}
}

func bxc(operand int, state *State) bool {
	state.regb = state.regb ^ state.regc
	return true
}

func out(operand int, state *State) bool {
	val := (combo(operand, state)%8 + 8) % 8
	state.output = append(state.output, val)
	return true
}

func bdv(operand int, state *State) bool {
	num := state.rega
	denom := math.Pow(2, float64(combo(operand, state)))
	state.regb = int(float64(num) / denom)
	return true
}

func cdv(operand int, state *State) bool {
	num := state.rega
	denom := math.Pow(2, float64(combo(operand, state)))
	state.regc = int(float64(num) / denom)
	return true
}

func execute(state *State) {
	for true {
		if state.iptr >= len(instructions) {
			return
		}
		opcode := instructions[state.iptr]
		operand := instructions[state.iptr+1]
		fn := ftable[opcode]
		adv := fn(operand, state)
		if adv {
			state.iptr = state.iptr + 2
		}
		// fmt.Println(state)
	}
}

func metaExecute() {
	candidates := list.New()
	candidates.PushBack(43)
	candidates.PushBack(47)

	for e := candidates.Front(); e != nil; e = candidates.Front() {
		candidates.Remove(e)
		cval := e.Value.(int)
		for ival := cval - 10; ival < cval+10; ival++ {
			state := State{0, int(ival), regb, regc, make([]int, 0)}
			execute(&state)
			slices.Reverse(state.output)
			if slices.Compare(rev, state.output) == 0 {
				fmt.Println("match! ", ival, rev, state.output)
				return
			}
			if slices.Compare(rev[0:len(state.output)], state.output) == 0 {
				fmt.Println("submatch! ", ival, rev, state.output)
				candidates.PushBack(ival * 8)
			}
			if len(state.output) > len(rev) {
				break
			}

		}
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	re_a := regexp.MustCompile(`Register A: (\d+)`)
	re_b := regexp.MustCompile(`Register B: (\d+)`)
	re_c := regexp.MustCompile(`Register C: (\d+)`)
	re_p := regexp.MustCompile(`Program: (.*)`)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		matches := re_a.FindStringSubmatch(line)
		// fmt.Println(matches)
		rega, _ = strconv.Atoi(string(matches[1]))

		scanner.Scan()
		line = scanner.Text()
		// fmt.Println(line)
		matches = re_b.FindStringSubmatch(line)
		// fmt.Println(matches)
		regb, _ = strconv.Atoi(string(matches[1]))

		scanner.Scan()
		line = scanner.Text()
		// fmt.Println(line)
		matches = re_c.FindStringSubmatch(line)
		// fmt.Println(matches)
		regc, _ = strconv.Atoi(string(matches[1]))

		scanner.Scan()
		line = scanner.Text()

		scanner.Scan()
		line = scanner.Text()
		// fmt.Println(line)
		matches = re_p.FindStringSubmatch(line)
		// fmt.Println(matches)
		istrs := strings.Split(matches[1], ",")
		for _, s := range istrs {
			ic, _ := strconv.Atoi(s)
			instructions = append(instructions, ic)
		}
		rev = make([]int, len(instructions))
		copy(rev, instructions)
		slices.Reverse(rev)
	}
	fmt.Println(rega, regb, regc, instructions, iptr)
	fmt.Println("rev: ", rev)
	metaExecute()

}
