package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	isData bool
	id     int
	size   int
}

func expand(data string) *list.List {
	result := list.New()
	isData := true
	id := 0
	for _, v := range data {
		count, _ := strconv.Atoi(string(v))
		n := Node{isData, id, count}
		result.PushBack(n)
		if isData {
			id++
		}
		isData = !isData
	}
	return result
}

func printNodes(nodes *list.List) {
	for e := nodes.Front(); e != nil; e = e.Next() {
		for i := 0; i < e.Value.(Node).size; i++ {
			if e.Value.(Node).isData {
				fmt.Print(e.Value.(Node).id)
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println("")
}

func findSpace(nodes *list.List, size int, id int) *list.Element {
	for e := nodes.Front(); e != nil; e = e.Next() {
		if e.Value.(Node).isData {
			if e.Value.(Node).id == id {
				return nil
			}
		} else {
			if e.Value.(Node).size >= size {
				return e
			}
		}
	}
	return nil
}

func swap(nodes *list.List, block *list.Element, space *list.Element) *list.Element {
	bn := block.Value.(Node)
	sn := space.Value.(Node)
	spaceDelta := sn.size - bn.size
	nodes.InsertBefore(bn, space)
	// fmt.Println(spaceDelta)
	if spaceDelta > 0 {
		nodes.InsertBefore(Node{false, -1, spaceDelta}, space)
	}
	nodes.Remove(space)

	newspace := bn.size
	pn := block.Prev().Value.(Node)
	if !pn.isData {
		newspace += pn.size
		nodes.Remove(block.Prev())
	}
	if block.Next() != nil && !block.Next().Value.(Node).isData {
		newspace += block.Next().Value.(Node).size
		nodes.Remove(block.Next())
	}
	swapsie := nodes.InsertBefore(Node{false, -1, newspace}, block)
	nodes.Remove(block)
	return swapsie
}

func defrag(nodes *list.List) *list.List {
	candidate := nodes.Back()
	for candidate != nil {
		node := candidate.Value.(Node)
		if node.isData {
			space := findSpace(nodes, node.size, node.id)
			if space != nil {
				candidate = swap(nodes, candidate, space)
			}
		}
		candidate = candidate.Prev()
	}
	return nodes
}

func checksum(nodes *list.List) int {
	result := 0
	index := 0
	for e := nodes.Front(); e != nil; e = e.Next() {
		for i := 0; i < e.Value.(Node).size; i++ {
			if e.Value.(Node).isData {
				result += index * e.Value.(Node).id
			}
			index++
		}
	}
	return result
}

func main() {

	var data []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	nodes := expand(data[0])
	// printNodes(nodes)
	nodes = defrag(nodes)
	// printNodes(nodes)
	fmt.Println(checksum(nodes))
}
