package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Grid struct {
  Data [][]rune
  N, M int
}

func parseInput(f *os.File) Grid {
  var grid Grid
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    var row []rune
    for _, ch := range scanner.Text() {
      row = append(row, ch)
    }
    grid.Data = append(grid.Data, row)
  }
  grid.N = len(grid.Data)
  grid.M = len(grid.Data[0])
  return grid
}

func (g Grid) IsValid(i, j int) bool {
  return i >= 0 && i < g.N && j >= 0 && j < g.M
}

func (g Grid) LaserBeam(i, j int, direction string, seen *map[string]bool) {
  // if out of bound exit
  if !g.IsValid(i, j) {
    return
  }

  // if a loop is detected exit
  seenKey := fmt.Sprintf("%d:%d:%s", i, j, direction)
  if _, alreadySeen := (*seen)[seenKey]; alreadySeen {
    return
  }

  // walk until it is possible
  for g.IsValid(i, j) && g.Data[i][j] == '.' {
    seenKey = fmt.Sprintf("%d:%d:%s", i, j, direction)
    if _, alreadySeen := (*seen)[seenKey]; !alreadySeen {
      (*seen)[seenKey] = true
    } else {
      return
    }
    switch direction {
    case ">":
      j++
    case "<":
      j--
    case "V":
      i++
    case "^":
      i--
    default:
      panic("Direction not possible: " + direction)
    }
  }

  // deal with special tiles
  if g.IsValid(i, j) {
    seenKey = fmt.Sprintf("%d:%d:%s", i, j, direction)
    (*seen)[seenKey] = true
    switch g.Data[i][j] {
    case '/':
      switch direction {
        case ">":
          g.LaserBeam(i-1, j, "^", seen)
        case "<":
          g.LaserBeam(i+1, j, "V", seen)
        case "^":
          g.LaserBeam(i, j+1, ">", seen)
        case "V":
          g.LaserBeam(i, j-1, "<", seen)
      }
    case '\\':
      switch direction {
        case ">":
          g.LaserBeam(i+1, j, "V", seen)
        case "<":
          g.LaserBeam(i-1, j, "^", seen)
        case "^":
          g.LaserBeam(i, j-1, "<", seen)
        case "V":
          g.LaserBeam(i, j+1, ">", seen)
      }
    case '-':
      switch direction {
        case ">":
          g.LaserBeam(i, j+1, ">", seen)
        case "<":
          g.LaserBeam(i, j-1, "<", seen)
        case "^":
          fallthrough
        case "V":
          g.LaserBeam(i, j-1, "<", seen)
          g.LaserBeam(i, j+1, ">", seen)
      }
    case '|':
      switch direction {
        case ">":
          fallthrough
        case "<":
          g.LaserBeam(i+1, j, "V", seen)
          g.LaserBeam(i-1, j, "^", seen)
        case "^":
          g.LaserBeam(i-1, j, "^", seen)
        case "V":
          g.LaserBeam(i+1, j, "V", seen)
      }
    default:
      panic("Impossible character: " + string(g.Data[i][j]))
    }
  }
}

func (g Grid) Display(seen *map[string]bool) {
  tilesSeen := make(map[string]bool)
  for k := range *seen {
    newKey := strings.Join(strings.Split(k, ":")[:2], ":")
    tilesSeen[newKey] = true
  }

  for i, row := range g.Data {
    for j, ch := range row {
      seenKey := fmt.Sprintf("%d:%d", i, j)
      if _, ok := tilesSeen[seenKey]; ok {
        fmt.Print("# ")
      } else {
        fmt.Print(string(ch) + " ")
      }
    }
    fmt.Println()
  }
}

func (g Grid) CountEnergizedTiles(seen map[string]bool) int {
  tilesSeen := make(map[string]bool)
  for k := range seen {
    newKey := strings.Join(strings.Split(k, ":")[:2], ":")
    tilesSeen[newKey] = true
  }
  return len(tilesSeen)
}

func (g Grid) ShootLaser(end, startTile int, direction string, c chan int) {
  maxEnergizedTiles := 0
  for i := 0; i < end; i++ {
    seen := make(map[string]bool)
    if direction == ">" || direction == "<" {
      g.LaserBeam(i, startTile, direction, &seen)
    } else {
      g.LaserBeam(startTile, i, direction, &seen)
    }
    energizedTiles := g.CountEnergizedTiles(seen)
    if energizedTiles > maxEnergizedTiles {
      maxEnergizedTiles = energizedTiles
    }
  }
  c <- maxEnergizedTiles
}

func part1(grid Grid) {
  seen := make(map[string]bool)
  grid.LaserBeam(0, 0, ">", &seen)
  // grid.Display(&seen)
  fmt.Println(
    "Eventually there are",
    grid.CountEnergizedTiles(seen),
    "energized tiles",
  )
}

func part2(grid Grid) {
  maxEnergizedTiles := 0
  c := make(chan int)

  fromLeft := func () {
    grid.ShootLaser(grid.N, 0, ">", c)
  }
  fromRight := func () {
    grid.ShootLaser(grid.N, grid.M-1, "<", c)
  }
  fromTop := func () {
    grid.ShootLaser(grid.M, 0, "V", c)
  }
  fromBottom := func () {
    grid.ShootLaser(grid.M, grid.N-1, "^", c)
  }

  go fromLeft()
  go fromRight()
  go fromTop()
  go fromBottom()

  for i := 0; i < 4; i++ {
    energizedTilesFromDirection := <-c
    if energizedTilesFromDirection > maxEnergizedTiles {
      maxEnergizedTiles = energizedTilesFromDirection
    }
  }

  fmt.Println(
    "The configuration energizing most tiles enables",
    maxEnergizedTiles, "tiles",
  )
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  grid := parseInput(f)
  // m := make(map[string]bool)
  // grid.Display(&m)
  // fmt.Println()
 
  part1(grid)
  part2(grid)
}
