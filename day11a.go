package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	ival int
	sval string
}

func parse(input string) *list.List {
	result := list.New()
	svals := strings.Split(input, " ")
	for _, sval := range svals {
		ival, _ := strconv.Atoi(sval)
		result.PushBack(Node{ival, sval})
	}
	return result
}

func printList(l *list.List) {
	current := l.Front()
	for current != nil {
		fmt.Print(current.Value.(Node).sval, " ")
		current = current.Next()
	}
	fmt.Println("")
}

func process(e *list.Element, l *list.List) *list.Element {
	node := e.Value.(Node)
	if node.ival == 0 {
		e.Value = Node{1, "1"}
		return e
	} else if len(node.sval)%2 == 0 {
		lstr := string(node.sval[:len(node.sval)/2])
		rstr := string(node.sval[(len(node.sval) / 2):])

		ival, _ := strconv.Atoi(lstr)
		sval := strconv.Itoa(ival)
		e.Value = Node{ival, sval}

		newival, _ := strconv.Atoi(rstr)
		newsval := strconv.Itoa(newival)
		return l.InsertAfter(Node{newival, newsval}, e)
	} else {
		ival := node.ival * 2024
		sval := strconv.Itoa(ival)
		e.Value = Node{ival, sval}
		return e
	}
}

var uniques = make(map[int]bool)
var seedMap = make(map[int]map[int]int)

func iterate(l *list.List, iterations int, seed int) *list.List {
	if seedMap[seed] == nil {
		seedMap[seed] = make(map[int]int)
	}
	for i := 0; i < iterations; i++ {
		for e := l.Front(); e != nil; e = e.Next() {
			e = process(e, l)
		}
	}
	for e := l.Front(); e != nil; e = e.Next() {
		uniques[e.Value.(Node).ival] = true
		seedMap[seed][e.Value.(Node).ival] += 1
	}
	return l
}

func explode(l *list.List) {
	total := 0
	for e := l.Front(); e != nil; e = e.Next() {
		onelist := list.New()
		node := e.Value.(Node)
		onelist.PushBack(node)

		iterate(onelist, 25, node.ival)
		x := onelist.Len()
		total += x
	}
	for k, _ := range uniques {
		_, ok := seedMap[k]
		if !ok {
			onelist := list.New()
			node := Node{k, strconv.Itoa(k)}
			onelist.PushBack(node)
			iterate(onelist, 25, node.ival)
		}
	}
	for k, _ := range uniques {
		_, ok := seedMap[k]
		if !ok {
			onelist := list.New()
			node := Node{k, strconv.Itoa(k)}
			onelist.PushBack(node)
			iterate(onelist, 25, node.ival)
		}
	}
	fmt.Println(len(seedMap))
}

func deriveLength75(l *list.List) int {
	total := 0
	for e := l.Front(); e != nil; e = e.Next() {
		ival := e.Value.(Node).ival
		for aval, acount := range seedMap[ival] {
			for bval, bcount := range seedMap[aval] {
				for _, ccount := range seedMap[bval] {
					total += ccount * bcount * acount
				}
			}
		}
	}
	return total
}

func main() {

	var data []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	l := parse(data[0])
	printList(l)
	explode(l)
	fmt.Println(deriveLength75(l))
}
