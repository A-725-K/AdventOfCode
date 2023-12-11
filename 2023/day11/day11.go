package main

import (
	"bufio"
	"fmt"
	"os"
)

type Matrix [][]rune
type Coord64 struct {
  row, col, distance int64
}

func abs(x int64) int64 {
    if x < 0 {
      return -x
    }
    return x
  }

func parseInput(f *os.File) Matrix {
  var matrix [][]rune
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    var row []rune
    for _, c := range scanner.Text() {
      row = append(row, c)
    }
    matrix = append(matrix, row)
  }
  return matrix
}

func (m Matrix) display() {
  for _, row := range m {
    for _, c := range row {
      fmt.Print(string(c))
    }
    fmt.Println()
  }
}

// This abomination has been used to solve part 1... xD
// for the Queue type, have a look at 2023/day10.go
// func (m Matrix) ShortestPath(start, end Coord) int {
//   rows, cols := len(m), len(m[0])
//   var visited [][]bool
//   for _, row := range m {
//     visited = append(visited, make([]bool, len(row)))
//   }
//
//   q := NewQueue()
//   q.Enqueue(start)
//   visited[start.row][start.col] = true
//   for !q.IsEmpty() {
//     curr := q.Dequeue()
//
//     if curr.row == end.row && curr.col == end.col {
//       return curr.distance
//     }
//
//     if curr.row-1 >= 0 && !visited[curr.row-1][curr.col] {
//       q.Enqueue(Coord{row: curr.row-1, col: curr.col, distance: curr.distance+1})
//       visited[curr.row-1][curr.col] = true
//     }
//     if curr.row+1 < rows && !visited[curr.row+1][curr.col] {
//       q.Enqueue(Coord{row: curr.row+1, col: curr.col, distance: curr.distance+1})
//       visited[curr.row+1][curr.col] = true
//     }
//     if curr.col-1 >= 0 && !visited[curr.row][curr.col-1] {
//       q.Enqueue(Coord{row: curr.row, col: curr.col-1, distance: curr.distance+1})
//       visited[curr.row][curr.col-1] = true
//     }
//     if curr.col+1 < cols && !visited[curr.row][curr.col+1] {
//       q.Enqueue(Coord{row: curr.row, col: curr.col+1, distance: curr.distance+1})
//       visited[curr.row][curr.col+1] = true
//     }
//   }
//   return -1
// }
func (m Matrix) ShortestPath(start, end Coord64) int64 { 
  return abs(int64(end.row - start.row)) + abs(int64(end.col - start.col))
}

func (m Matrix) getGalaxies() []Coord64 {
var galaxies []Coord64
  for i, row := range m {
    for j, c := range row {
      if c == '#' {
        galaxies = append(galaxies, Coord64{row: int64(i), col: int64(j), distance: int64(0)})
      }
    }
  }
  return galaxies
}

func (m Matrix) findVoid() ([]int64, []int64) {
  var rowGaps, colGaps []int64
  rows, cols := len(m), len(m[0])
  for j := 0; j < cols; j++ {
    isEmptyCol := true
    isEmptyRow := true
    for i := 0; i < rows; i++ {
      if isEmptyCol && m[i][j] == '#' {
        isEmptyCol = false
      }
      if isEmptyRow && m[j][i] == '#' {
        isEmptyRow = false
      }
    }
    if isEmptyCol {
      colGaps = append(colGaps, int64(j))
    }
    if isEmptyRow {
      rowGaps = append(rowGaps, int64(j))
    }
  }
  return rowGaps, colGaps
}

func expand(galaxies *[]Coord64, growth int64, rowGaps, colGaps []int64) {
  growth-- // to allow passing nice and round numbers as parameters 
  for i, g := range *galaxies {
    rowExpIdx, colExpIdx := 0, 0
    for rowExpIdx < len(rowGaps) && g.row > rowGaps[rowExpIdx] {
      rowExpIdx++
    }
    for colExpIdx < len(colGaps) && g.col > colGaps[colExpIdx] {
      colExpIdx++
    }
    (*galaxies)[i] = Coord64{row: g.row + growth * int64(rowExpIdx), col: g.col + growth * int64(colExpIdx)}
  }
}

func lookAtTheUniverse(
  m Matrix,
  rowGaps, colGaps []int64,
  msg string,
  expansionFactor int64,
) {
  galaxies := m.getGalaxies()
  expand(&galaxies, expansionFactor, rowGaps, colGaps)
  numOfGalaxies := len(galaxies)
  distanceBetweenGalaxies := 0
  for i := 0; i < numOfGalaxies-1; i++ {
    for j := i+1; j < numOfGalaxies; j++ {
      distanceBetweenGalaxies += int(m.ShortestPath(galaxies[i], galaxies[j]))
    }
  }
  fmt.Println(msg, distanceBetweenGalaxies)
}

func part1(m Matrix, rowGaps, colGaps []int64) {
  lookAtTheUniverse(
    m, rowGaps, colGaps,
    "The total distance between galaxies is", 2,
  )
}

func part2(m Matrix, rowGaps, colGaps []int64) {
  lookAtTheUniverse(
    m, rowGaps, colGaps,
    "The total distance between far galaxies is:", 1_000_000,
  )
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  matrix := parseInput(f)
  rowGaps, colGaps := matrix.findVoid()
 
  part1(matrix, rowGaps, colGaps)
  part2(matrix, rowGaps, colGaps)
}
