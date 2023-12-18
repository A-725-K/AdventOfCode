package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

  "container/heap"
)

type Grid struct {
  Data [][]int
  N, M int
}
type Item struct {
  HeatLoss, I, J, PrevUp, PrevDown, PrevLeft, PrevRight int
}
type PrioQueue []*Item

func (i Item) String() string {
  return fmt.Sprintf(
    "%d:%d:%d:%d:%d:%d",
    i.I, i.J, i.PrevUp, i.PrevDown, i.PrevLeft, i.PrevRight,
  )
}

func (pq PrioQueue) Len() int {
  return len(pq)
}

func (pq PrioQueue) IsEmpty() bool {
  return len(pq) == 0
}

func (pq PrioQueue) Less(i, j int) bool {
	return pq[i].HeatLoss < pq[j].HeatLoss
}

func (pq PrioQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PrioQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PrioQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*pq = old[0:n-1]
	return item
}

func (g Grid) String() string {
  s := ""
  for _, row := range g.Data {
    for _, n := range row {
      s += fmt.Sprintf("%d ", n)
    }
    s += "\n"
  }
  return s
}

func (g Grid) IsValid(i, j int) bool {
  return i >= 0 && i < g.N && j >= 0 && j < g.M
}

func parseInput(f *os.File) Grid {
  var grid Grid
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    var row []int
    for _, nStr := range scanner.Text() {
      n, _ := strconv.Atoi(string(nStr))
      row = append(row, n)
    }
    grid.Data = append(grid.Data, row)
  }
  grid.N = len(grid.Data)
  grid.M = len(grid.Data[0])
  return grid
}

func part1(grid Grid) {
  var q PrioQueue
  var newI, newJ, minHeatLoss int
  visited := make(map[string]bool)
  maxStepsInADirection := 3
  start := &Item{
    HeatLoss: 0, I: 0, J: 0,
    PrevUp: 0, PrevDown: 0, PrevLeft: 0, PrevRight: 0,
  }
  heap.Push(&q, start)

  for !q.IsEmpty() {
    curr := heap.Pop(&q).(*Item)

    // found the end
    if curr.I == grid.N-1 && curr.J == grid.M-1 {
      minHeatLoss = curr.HeatLoss

      // Djikstra should find the first time the path with minimum cost since
      // I am using a minimum heap
      break
    }

    key := curr.String()
    if _, alreadySeen := visited[key]; alreadySeen {
      continue
    }
    visited[key] = true

    // move up
    newI, newJ = curr.I-1, curr.J
    canTurnUp := curr.PrevUp < maxStepsInADirection && curr.PrevDown == 0
    if grid.IsValid(newI, newJ) && canTurnUp {
      newItem := &Item{
        HeatLoss: curr.HeatLoss+grid.Data[newI][newJ], I: newI, J: newJ,
        PrevUp: curr.PrevUp+1, PrevDown: 0, PrevLeft: 0, PrevRight: 0,
      }
      heap.Push(&q, newItem)
    }

    // move down
    newI, newJ = curr.I+1, curr.J
    canTurnDown := curr.PrevDown < maxStepsInADirection && curr.PrevUp == 0
    if grid.IsValid(newI, newJ) && canTurnDown {
      newItem := &Item{
        HeatLoss: curr.HeatLoss+grid.Data[newI][newJ], I: newI, J: newJ,
        PrevUp: 0, PrevDown: curr.PrevDown+1, PrevLeft: 0, PrevRight: 0,
      }
      heap.Push(&q, newItem)
    }

    // move left
    newI, newJ = curr.I, curr.J-1
    canTurnLeft := curr.PrevLeft < maxStepsInADirection && curr.PrevRight == 0
    if grid.IsValid(newI, newJ) && canTurnLeft {
      newItem := &Item{
        HeatLoss: curr.HeatLoss+grid.Data[newI][newJ], I: newI, J: newJ,
        PrevUp: 0, PrevDown: 0, PrevLeft: curr.PrevLeft+1, PrevRight: 0,
      }
      heap.Push(&q, newItem)
    }

    // move right
    newI, newJ = curr.I, curr.J+1
    canTurnRight := curr.PrevRight < maxStepsInADirection && curr.PrevLeft == 0
    if grid.IsValid(newI, newJ) && canTurnRight {
      newItem := &Item{
        HeatLoss: curr.HeatLoss+grid.Data[newI][newJ], I: newI, J: newJ,
        PrevUp: 0, PrevDown: 0, PrevLeft: 0, PrevRight: curr.PrevRight+1,
      }
      heap.Push(&q, newItem)
    }
  }
  fmt.Println("The minimum heat loss is:", minHeatLoss)
}

func part2(grid Grid) {
  var q PrioQueue
  var newI, newJ int
  minHeatLoss := math.MaxInt
  visited := make(map[string]bool)
  minStespInADirection, maxStepsInADirection := 4, 10
  start := &Item{
    HeatLoss: 0, I: 0, J: 0,
    PrevUp: 0, PrevDown: 0, PrevLeft: 0, PrevRight: 0,
  }
  heap.Push(&q, start)

  for !q.IsEmpty() {
    curr := heap.Pop(&q).(*Item)

    // found the end
    if curr.I == grid.N-1 && curr.J == grid.M-1 {
      minHeatLoss = curr.HeatLoss

      // Djikstra should find the first time the path with minimum cost since
      // I am using a minimum heap
      break
    }

    key := curr.String()
    if _, alreadySeen := visited[key]; alreadySeen {
      continue
    }
    visited[key] = true

    for steps := minStespInADirection; steps <= maxStepsInADirection; steps ++ {
      canTurnVertically := curr.PrevDown == 0 && curr.PrevUp == 0
      canTurnHorizontally := curr.PrevRight == 0 && curr.PrevLeft == 0

      // move up
      newI, newJ = curr.I-steps, curr.J
      if grid.IsValid(newI, newJ) && canTurnVertically {
        newHeatLoss := curr.HeatLoss
        for i := curr.I-1; i >= newI; i-- {
          newHeatLoss += grid.Data[i][curr.J]
        }
        newItem := &Item{
          HeatLoss: newHeatLoss, I: newI, J: newJ,
          PrevUp: steps, PrevDown: 0, PrevLeft: 0, PrevRight: 0,
        }
        heap.Push(&q, newItem)
      }

      // move down
      newI, newJ = curr.I+steps, curr.J
      if grid.IsValid(newI, newJ) && canTurnVertically {
        newHeatLoss := curr.HeatLoss
        for i := curr.I+1; i <= newI; i++ {
          newHeatLoss += grid.Data[i][curr.J]
        }
        newItem := &Item{
          HeatLoss: newHeatLoss, I: newI, J: newJ,
          PrevUp: 0, PrevDown: steps, PrevLeft: 0, PrevRight: 0,
        }
        heap.Push(&q, newItem)
      }

      // move left
      newI, newJ = curr.I, curr.J-steps
      if grid.IsValid(newI, newJ) && canTurnHorizontally {
        newHeatLoss := curr.HeatLoss
        for j := curr.J-1; j >= newJ; j-- {
          newHeatLoss += grid.Data[curr.I][j]
        }
        newItem := &Item{
          HeatLoss: newHeatLoss, I: newI, J: newJ,
          PrevUp: 0, PrevDown: 0, PrevLeft: steps, PrevRight: 0,
        }
        heap.Push(&q, newItem)
      }

      // move right
      newI, newJ = curr.I, curr.J+steps
      if grid.IsValid(newI, newJ) && canTurnHorizontally {
        newHeatLoss := curr.HeatLoss
        for j := curr.J+1; j <= newJ; j++ {
          newHeatLoss += grid.Data[curr.I][j]
        }
        newItem := &Item{
          HeatLoss: newHeatLoss, I: newI, J: newJ,
          PrevUp: 0, PrevDown: 0, PrevLeft: 0, PrevRight: curr.PrevRight+steps,
        }
        heap.Push(&q, newItem)
      }
    }
  }
  fmt.Println("The minimum heat loss with ultra crucibles is:", minHeatLoss)
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  grid := parseInput(f)
 
  part1(grid)
  part2(grid)
}
