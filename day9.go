package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	isData bool
	id     int
}

func expand(data string) []Node {
	result := make([]Node, 0)
	isData := true
	id := 0
	for _, v := range data {
		count, _ := strconv.Atoi(string(v))
		for i := 0; i < count; i++ {
			n := Node{isData, id}
			result = append(result, n)
		}
		if isData {
			id++
		}
		isData = !isData
	}
	return result
}

func printNodes(nodes []Node) {
	for _, v := range nodes {
		if v.isData {
			fmt.Print(v.id)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("")
}

func defrag(nodes []Node) []Node {
	bPtr := 0
	ePtr := len(nodes) - 1
	for bPtr < ePtr {
		for nodes[bPtr].isData {
			bPtr++
		}
		for !nodes[ePtr].isData {
			ePtr--
		}
		if ePtr <= bPtr { // make sure we don't over-correct
			break
		}
		nodes[bPtr] = Node{true, nodes[ePtr].id}
		nodes[ePtr] = Node{false, -1}
		// printNodes(nodes)
	}
	return nodes
}

func checksum(nodes []Node) int {
	result := 0
	for i := 0; i < len(nodes); i++ {
		if nodes[i].isData {
			result += i * nodes[i].id
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
