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
		// fmt.Println("split:", lstr, rstr)

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

func iterate(l *list.List, iterations int) *list.List {
	for i := 0; i < iterations; i++ {
		for e := l.Front(); e != nil; e = e.Next() {
			e = process(e, l)
		}
		// printList(l)
	}
	return l
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
	iterate(l, 25)
	fmt.Println(l.Len())

}
