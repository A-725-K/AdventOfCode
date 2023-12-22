package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid struct {
  Grid [][]rune
  N, M int
}

func (g Grid) Display() {
  for _, row := range g.Grid {
    for _, ch := range row {
      fmt.Print(string(ch) + "")
    }
    fmt.Println()
  }
}

func parseInput(f *os.File) Grid {
  var grid Grid
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    var row []rune
    for _, ch := range scanner.Text() {
      row = append(row, ch)
    }
    grid.Grid = append(grid.Grid, row)
  }
  grid.N = len(grid.Grid)
  grid.M = len(grid.Grid[0])
  return grid
}

func (g Grid) WalkInMaze(i, j, currLen int, maxLen *int, visited [][]bool, considerSlopes bool) bool {
  if i == g.N - 1 && j == g.M - 2 {
    if currLen > *maxLen {
      *maxLen = currLen
    }
    // false instead of true to continue to loop after finding the first result
    return false
  }

  // fmt.Println("-->", i, j, currLen)
  if visited[i][j] {
    return false
  }

  visited[i][j] = true
  finished := false

  if considerSlopes && g.Grid[i][j] == '>' {
    finished = g.WalkInMaze(i, j+1, currLen+1, maxLen, visited, considerSlopes)
  } else if considerSlopes && g.Grid[i][j] == 'v' {
    finished = g.WalkInMaze(i+1, j, currLen+1, maxLen, visited, considerSlopes)
  } else if considerSlopes && g.Grid[i][j] == '<' {
    finished = g.WalkInMaze(i, j-1, currLen+1, maxLen, visited, considerSlopes)
  } else if considerSlopes && g.Grid[i][j] == '^' {
    finished = g.WalkInMaze(i-1, j, currLen+1, maxLen, visited, considerSlopes)
  } else {
    if i-1 >= 0 && g.Grid[i-1][j] != '#' && !visited[i-1][j] {
      finished = finished || g.WalkInMaze(i-1, j, currLen+1, maxLen, visited, considerSlopes)
    }
    if i+1 < g.N && g.Grid[i+1][j] != '#' && !visited[i+1][j] {
      finished = finished || g.WalkInMaze(i+1, j, currLen+1, maxLen, visited, considerSlopes)
    }
    if j+1 < g.M && g.Grid[i][j+1] != '#' && !visited[i][j+1] {
      finished = finished || g.WalkInMaze(i, j+1, currLen+1, maxLen, visited, considerSlopes)
    }
    if j-1 >= 0 && g.Grid[i][j-1] != '#' && !visited[i][j-1] {
      finished = finished || g.WalkInMaze(i, j-1, currLen+1, maxLen, visited, considerSlopes)
    }
  }

  visited[i][j] = false

  return finished
}

func part1(grid Grid) {
  visited := make([][]bool, grid.N)
  for i := 0; i < grid.N; i++ {
    visited[i] = make([]bool, grid.M)
  }
  maxLen := 0
  grid.WalkInMaze(0, 1, 0, &maxLen, visited, true)
  fmt.Println("The longest path consists of", maxLen, "steps")
}

func part2(grid Grid) {
  visited := make([][]bool, grid.N)
  for i := 0; i < grid.N; i++ {
    visited[i] = make([]bool, grid.M)
  }
  maxLen := 0
  grid.WalkInMaze(0, 1, 0, &maxLen, visited, false)
  fmt.Println("The longest path without considering slopes consists of", maxLen, "steps")
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
