package ds

import (
	"fmt"
	stack "AdventOfCode/ds/stack"
)

type Node struct {
	Name      string
	Value     int
	Neighbors []string
}

func (n *Node) String() string {
	return fmt.Sprintf("%s-[%d] -> %v", n.Name, n.Value, n.Neighbors)
}

func NewNode(n string, v int, ns []string) *Node {
	return &Node{Name: n, Value: v, Neighbors: ns}
}

// ----------------------------------------------------- //

type Graph map[string]*Node

func NewGraph() Graph {
	return make(map[string]*Node)
}

func (g *Graph) VisitGraph(start string) {
	visited := make(map[string]bool)
	for n := range *g {
		visited[n] = false
	}

	s := stack.NewStack[*Node]()
	s.Push((*g)[start])

	for !s.IsEmpty() {
		currNode := s.Pop()
		visited[currNode.Name] = true
		fmt.Println(currNode)

		for _, neighbor := range currNode.Neighbors {
			if !visited[neighbor] {
				s.Push((*g)[neighbor])
			}
		}
	}
	fmt.Println()
}

