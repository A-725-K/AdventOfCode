package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Node struct {
  Label string
  Neighbors []int
}
type Graph map[int]Node

func (g Graph) String() string {
  n := len(g)
  s := ""
  for i := 0; i < n; i++ {
    node := g[i]
    s += fmt.Sprintf("[%s (%d)] -> ", node.Label, i)
    for _, vv := range node.Neighbors {
      s += fmt.Sprintf("%s (%d)", g[vv].Label, vv) + ", "
    }
    s += "\n"
  }
  return s
}

func parseInput(f *os.File) Graph {
  graph := make(Graph)
  scanner := bufio.NewScanner(f)
  neighsByIdx := [][]string{}
  node2int := make(map[string]int)
  idx := 0
  for scanner.Scan() {
    fields := strings.Split(strings.ReplaceAll(scanner.Text(), ":", ""), " ")
    graph[idx] = Node{Label: fields[0], Neighbors: []int{}}
    node2int[fields[0]] = idx
    neighsByIdx = append(neighsByIdx, fields[1:])
    idx++
  }
  for i, neighs := range neighsByIdx {
    for _, v := range neighs {
      if _, ok := node2int[v]; !ok {
        graph[idx] = Node{Label: v, Neighbors: []int{}}
        node2int[v] = idx
        idx++
      }
      if e, ok := graph[i]; ok {
        e.Neighbors = append(e.Neighbors, node2int[v])
        graph[i] = e
      }
      if e, ok := graph[node2int[v]]; ok {
        e.Neighbors = append(e.Neighbors, i)
        graph[node2int[v]] = e
      }
    }
  }
  return graph
}

func maxElement(lst []int) int {
  maxEl, maxIdx := 0, 0
  for i, el := range lst {
    if el > maxEl {
      maxEl = el
      maxIdx = i
    }
  }
  if maxIdx > 0 {
    return maxIdx
  }

  return len(lst)-1
}

func minPair(x, y int, xL, yL []int) (int, []int) {
  if x < y {
    return x, xL
  }
  return y, yL
}

// Minimum cut solver, Stoer-Wagner algorithm, sources:
// - https://en.wikipedia.org/wiki/Stoer%E2%80%93Wagner_algorithm
// - https://github.com/kth-competitive-programming/kactl/blob/main/content/graph/GlobalMinCut.h
func (g Graph) StoerWagner() (int, []int) {
  n := len(g)

  best := math.MaxInt
  var bestLst []int
  co := make([][]int, n)

  // transform graph to adjacency matrix representation and initialize "co"
  var adj [][]int
  for i := 0; i < n; i++ {
    adj = append(adj, make([]int, n))
    co[i] = []int{i}
  }
  for u, node := range g {
    for _, v := range node.Neighbors {
      adj[u][v] = 1
      adj[v][u] = 1
    }
  }

  for ph := 1; ph < n; ph++ {
    w := make([]int, n)
    copy(w, adj[0])
    s, t := 0, 0
    for it := 0; it < n - ph; it++ {
      w[t] = math.MinInt
      s, t = t, maxElement(w)
      for i := 0; i < n; i++ {
        w[i] += adj[t][i]
      }
    }
    best, bestLst = minPair(best, w[t]-adj[t][t], bestLst, co[t])
    for _, el := range co[t] {
      co[s] = append(co[s], el)
    }
    for i := 0; i < n; i++ {
      adj[s][i] += adj[t][i]
      adj[i][s] = adj[s][i]
    }
    adj[0][t] = math.MinInt
  }

  return best, bestLst
}

func part1(graph Graph) {
  numOfWiresToCut, minimumCut := graph.StoerWagner()
  if numOfWiresToCut != 3 {
    panic("The problem requires more than 3 wires")
  }
  connectedComponent1 := len(minimumCut)
  connectedComponent2 := len(graph) - connectedComponent1
  prod := connectedComponent1 * connectedComponent2
  fmt.Println("The product of the 2 disconnected groups is:", prod)
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  part1(parseInput(f))
}
