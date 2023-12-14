package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"flag"
)

type Grid [][]rune

func (g Grid) String() string {
  s := ""
  for _, row := range g {
    for _, ch := range row {
      s += string(ch) + " "
    }
    s += "\n"
  }
  return s
}

func (g *Grid) Swap(i1, j1, i2, j2 int) {
  tmp := (*g)[i1][j1]
  (*g)[i1][j1] = (*g)[i2][j2]
  (*g)[i2][j2] = tmp
}

func (g *Grid) RollNorth() {
  n, m := len(*g), len((*g)[0])

  for j := 0; j < m; j++ {
    lastBlock := 0
    for i := 1; i < n; i++ {
      if (*g)[i][j] == 'O' {
        for upper := i; upper > lastBlock && (*g)[upper-1][j] == '.'; upper-- {
          g.Swap(upper, j, upper-1, j)
        }
      }
    }
    lastBlock++
  }
}

func (g *Grid) RollSouth() {
  n, m := len(*g), len((*g)[0])

  for j := 0; j < m; j++ {
    lastBlock := n-1
    for i := n-2; i >=0; i-- {
      if (*g)[i][j] == 'O' {
        for lower := i; lower < lastBlock && (*g)[lower+1][j] == '.'; lower++ {
          g.Swap(lower, j, lower+1, j)
        }
      }
    }
    lastBlock--
  }
}

func (g *Grid) RollEast() {
  n, m := len(*g), len((*g)[0])

  for i := 0; i < n; i++ {
    lastBlock := m-1
    for j := m-2; j >= 0; j-- {
      if (*g)[i][j] == 'O' {
        for right := j; right < lastBlock && (*g)[i][right+1] == '.'; right++ {
          g.Swap(i, right, i, right+1)
        }
      }
    }
    lastBlock--
  }
}

func (g *Grid) RollWest() {
  n, m := len(*g), len((*g)[0])

  for i := 0; i < n; i++ {
    lastBlock := 0
    for j := 1; j < m; j++ {
      if (*g)[i][j] == 'O' {
        for left := j; left > lastBlock && (*g)[i][left-1] == '.'; left-- {
          g.Swap(i, left, i, left-1)
        }
      }
    }
    lastBlock++
  }
}

func (g Grid) ComputeLoad() int {
  totalLoad := 0
  n := len(g)
  for i, row := range g {
    for _, ch := range row {
      if ch == 'O' {
        totalLoad += n-i
      }
    }
  }
  return totalLoad
}

func (g Grid) Copy() Grid {
  var gCopy Grid
  for _, row := range g {
    var gridRow []rune
    for _, ch := range row {
      gridRow = append(gridRow, ch)
    }
    gCopy = append(gCopy, gridRow)
  }
  return gCopy
}

func (g Grid) SlowPrint(dir string, animate bool) {
  if !animate {
    return
  }
  n, m := len(g), len(g[0])
  fmt.Print(strings.Repeat("\x1BD", n*m))
  fmt.Print(strings.Repeat("\x1B[1A", 100))
  fmt.Println(g)
  fmt.Println("Rolling:", dir)
  time.Sleep(time.Millisecond * 1200)
}

func (g *Grid) Cycle(animate bool) {
  g.SlowPrint(">", animate)
  g.RollNorth()
  g.SlowPrint("^", animate)
  g.RollWest()
  g.SlowPrint("<", animate)
  g.RollSouth()
  g.SlowPrint("V", animate)
  g.RollEast()
}

func parseInput(f *os.File) Grid {
  var grid Grid
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    var row []rune
    for _, ch := range scanner.Text() {
      row = append(row, ch)
    }
    grid = append(grid, row)
  }
  return grid
}

func part1(grid Grid) {
  grid.RollNorth()
  fmt.Println("The total load on the grid is:", grid.ComputeLoad())
}

func part2(grid Grid, animate bool) {
  orig := grid.Copy()

  memo := make(map[string]int)
  cycles := 1000000000
  var firstSeenIdx, lastSeenIdx int

  for i := 0; i < cycles; i++ {
    grid.Cycle(animate)
    key := grid.String()
    if idx, ok := memo[key]; ok {
      firstSeenIdx = idx
      lastSeenIdx = i
      break
    }
    memo[key] = i
  }
  recurringLen := lastSeenIdx - firstSeenIdx
  cyclesNeeded := ((cycles - firstSeenIdx) % recurringLen) + firstSeenIdx

  for i := 0; i < cyclesNeeded; i++ {
    orig.Cycle(animate)
  }
  
  fmt.Println(
    "The total load on the grid after",
    cycles, "cycles is:", grid.ComputeLoad(),
  )
}

func main() {
  var animate bool
  var filename string
  flag.StringVar(&filename, "file", "input", "Input file")
  flag.BoolVar(&animate, "animate", false, "Display an animation")
  flag.Parse()

  f, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  grid := parseInput(f)
 
  orig := grid.Copy()
  part1(orig)
  part2(grid, animate)
}
