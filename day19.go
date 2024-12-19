package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var known map[string]bool = make(map[string]bool)

func search(frag string) bool {
	if known[frag] == true {
		return true
	}
	if len(frag) == 1 {
		return false
	}
	cutpoint := len(frag) / 2
	first, second := frag[:cutpoint], frag[cutpoint:]
	if search(first) && search(second) {
		known[frag] = true
		return true
	}
	for i := -3; i < 3; i++ {
		cp := cutpoint + i
		if cp > 0 && cp < len(frag) {
			first, second = frag[:cutpoint+i], frag[cutpoint+i:]
			if search(first) && search(second) {
				known[frag] = true
				return true
			}
		}
	}
	return false
}

func searchAll(designs []string, patterns []string) int {
	total := 0
	for _, v := range designs {
		if search(v) {
			fmt.Print(".")
			total++
		}
	}
	fmt.Print("\n")
	// keys := slices.Collect(maps.Keys(known))
	// fmt.Println(keys)
	return total
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	patternLine := scanner.Text()
	scanner.Scan()

	patterns := make([]string, 0)
	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	pstrs := strings.Split(patternLine, ",")
	for _, v := range pstrs {
		patterns = append(patterns, strings.TrimSpace(v))
	}

	for _, p := range patterns {
		known[p] = true
	}

	// fmt.Println(patterns, lines)
	fmt.Println(searchAll(lines, patterns))

	// test := "abcdefghijklm"
	// fmt.Println("0:1", test[0:1]) a
	// fmt.Println("1:1", test[1:1]) <blank>
	// fmt.Println("1:2", test[1:2]) b
	// fmt.Println(":4", "4:", test[:4], test[4:]) abcd efghijklm
	// fmt.Println("len/2: ", test[:len(test)/2], test[len(test)/2:]) abcdef ghijklm
	// fmt.Printf("halfslice is %T\n", test[:len(test)/2]) string
	// fmt.Printf("oneslice is %T\n", test[1:1]) string

}
