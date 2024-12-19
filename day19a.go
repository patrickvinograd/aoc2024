package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var stems []string = make([]string, 0)
var ways map[string]int = make(map[string]int)

func search(frag string) int {
	if len(frag) == 0 {
		return 1
	}
	w, ok := ways[frag]
	if ok {
		return w
	}
	sum := 0
	for _, stem := range stems {
		rem, found := strings.CutPrefix(frag, stem)
		if found {
			sum += search(rem)
		}
	}
	ways[frag] = sum
	return sum
}

func searchAll(designs []string) int {
	total := 0
	for _, v := range designs {
		total += search(v)
		fmt.Print(".")
	}
	fmt.Print("\n")
	return total
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	patternLine := scanner.Text()
	scanner.Scan()

	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	pstrs := strings.Split(patternLine, ",")
	for _, v := range pstrs {
		p := strings.TrimSpace(v)
		stems = append(stems, p)
	}

	// fmt.Println(patterns, lines)
	fmt.Println(searchAll(lines))

	// test := "abcdefghijklm"
	// fmt.Println("0:1", test[0:1]) a
	// fmt.Println("1:1", test[1:1]) <blank>
	// fmt.Println("1:2", test[1:2]) b
	// fmt.Println(":4", "4:", test[:4], test[4:]) abcd efghijklm
	// fmt.Println("len/2: ", test[:len(test)/2], test[len(test)/2:]) abcdef ghijklm
	// fmt.Printf("halfslice is %T\n", test[:len(test)/2]) string
	// fmt.Printf("oneslice is %T\n", test[1:1]) string

}
