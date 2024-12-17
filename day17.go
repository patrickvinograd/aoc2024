package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
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

var rega int
var regb int
var regc int

var instructions = make([]int, 0)

// var output = make([]int, 0)
var output = make([]string, 0)
var iptr int = 0

var ftable = map[int]func(int) bool{
	0: adv,
	1: bxl,
	2: bst,
	3: jnz,
	4: bxc,
	5: out,
	6: bdv,
	7: cdv,
}

func combo(operand int) int {
	if operand == 0 || operand == 1 || operand == 2 || operand == 3 {
		return operand
	} else if operand == 4 {
		return rega
	} else if operand == 5 {
		return regb
	} else if operand == 6 {
		return regc
	} else if operand == 7 {
		panic("unknown operand")
	}
	panic("unknown operand")
}

func adv(operand int) bool {
	num := rega
	denom := math.Pow(2, float64(combo(operand)))
	rega = int(float64(num) / denom)
	return true
}

func bxl(operand int) bool {
	regb = regb ^ operand
	return true
}

func bst(operand int) bool {
	regb = (combo(operand)%8 + 8) % 8
	return true
}

func jnz(operand int) bool {
	if rega == 0 {
		return true
	} else {
		iptr = operand
		return false
	}
}

func bxc(operand int) bool {
	regb = regb ^ regc
	return true
}

func out(operand int) bool {
	val := (combo(operand)%8 + 8) % 8
	// fmt.Println("output", val)
	output = append(output, strconv.Itoa(val))
	return true
}

func bdv(operand int) bool {
	num := rega
	denom := math.Pow(2, float64(combo(operand)))
	regb = int(float64(num) / denom)
	return true
}

func cdv(operand int) bool {
	num := rega
	denom := math.Pow(2, float64(combo(operand)))
	regc = int(float64(num) / denom)
	return true
}

func execute() {
	for true {
		if iptr >= len(instructions) {
			return
		}
		opcode := instructions[iptr]
		operand := instructions[iptr+1]
		fn := ftable[opcode]
		adv := fn(operand)
		if adv {
			iptr = iptr + 2
		}
		fmt.Println(rega, regb, regc, instructions, iptr)

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
	}
	fmt.Println(rega, regb, regc, instructions, iptr)
	execute()
	fmt.Println(strings.Join(output, ","))

}
