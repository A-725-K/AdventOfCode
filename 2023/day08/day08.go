package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
  Id string
  Neighbors map[string]string
}

type Graph map[string]Node

func (n Node) isFinal() bool {
  return n.Id[2] == 'Z'
}
func (n Node) isInitial() bool {
  return n.Id[2] == 'A'
}

func parseInput(f *os.File) (string, Graph) {
  scanner := bufio.NewScanner(f)
  scanner.Scan()
  directions := scanner.Text()
  scanner.Scan()
  graph := make(Graph)
  for scanner.Scan() {
    fields := strings.Split(scanner.Text(), "=")
    id := strings.Trim(fields[0], " ")
    neighbors := make(map[string]string)
    neighbors["left"] = fields[1][2:5]
    neighbors["right"] = fields[1][7:10]
    graph[id] = Node{Id: id, Neighbors: neighbors}
  }
  return directions, graph
}

func Gcd (a, b int64) int64 {
  if b == 0 {
    return a
  }
  return Gcd(b, a%b)
}

func Lcm(lst []int) int64 {
  lcm := int64(lst[0])
  n := len(lst)
  for i := 1; i < n; i++ {
    lcm = int64(lst[i]) * lcm / Gcd(int64(lst[i]), lcm)
  }
  return lcm
}

func part1(g Graph, directions string) {
  steps := 0
  currentDirectionIdx := 0
  currentNode := g["AAA"]
  for currentNode.Id != "ZZZ" {
    newDirection := directions[currentDirectionIdx % len(directions)]
    currentDirectionIdx++
    if newDirection == 'L' {
      currentNode = g[currentNode.Neighbors["left"]]
    } else {
      currentNode = g[currentNode.Neighbors["right"]]
    }
    steps++
  }

  fmt.Println("It takes", steps, "steps to get out the storm")
}

func part2(g Graph, directions string) {
  var onlyA []Node
  for _, node := range g {
    if node.isInitial() {
      onlyA = append(onlyA, node)
    }
  }

  var lst []int
  for i := range onlyA {
    currentDirectionIdx := 0
    steps := 0
    for !onlyA[i].isFinal() {
      newDirection := directions[currentDirectionIdx % len(directions)]
      currentDirectionIdx++
      if newDirection == 'L' {
        onlyA[i] = g[onlyA[i].Neighbors["left"]]
      } else {
        onlyA[i] = g[onlyA[i].Neighbors["right"]]
      }
      steps++
    }
    lst = append(lst, steps)
  }

  fmt.Println("It takes", Lcm(lst), "steps to get out the storm")
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  directions, graph := parseInput(f)
  part1(graph, directions)
  part2(graph, directions)
}
