package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
  "AdventOfCode/ds"
)

const (
  EMPTY uint8 = 0
  ROCK uint8 = 1
  SAND uint8 = 2

  MAX_ROWS = 180    // input
  MAX_COLS = 2000
  // MAX_COLS = 100     // input_part1
  // MAX_ROWS = 12     // mini_input
  // MAX_COLS = 30     // mini_input

  X_NORM = 0 // input
  // X_NORM = 485 // mini_input and input_part1
)

type Cave [MAX_ROWS][MAX_COLS]uint8

func drawRocks(cave *Cave, xS, yS, xE, yE int) {
  var start, end int
  if xS == xE {
    if yS > yE {
      start, end = yE, yS
    } else {
      start, end = yS, yE
    }
    for i := start; i <= end; i++ {
      (*cave)[i][xS] = ROCK
    }
  } else {
    if xS > xE {
      start, end = xE, xS
    } else {
      start, end = xS, xE
    }
    for j := start; j <= end; j++ {
      (*cave)[yS][j] = ROCK
    }
  }
}

func parseInput(f *os.File) (Cave, int) {
  scanner := bufio.NewScanner(f)
 
  var cave Cave
  for i := 0; i < MAX_ROWS; i++ {
    for j := 0; j < MAX_COLS; j++ {
      cave[i][j] = EMPTY
    }
  }

  var maxY int
  for scanner.Scan() {
    run := strings.Split(scanner.Text(), " -> ")
    xStart, xEnd, yStart, yEnd := -1, -1, -1, -1
    var err error
    for idx, r := range run {
      coord := strings.Split(r, ",")
      if idx == 0 {
        xStart, err = strconv.Atoi(coord[0])
        if err != nil {
          panic("Cannot convert xStart coord")
        }
        xStart -= X_NORM // normalize col
        yStart, err = strconv.Atoi(coord[1])
        if err != nil {
          panic("Cannot convert yStart coord")
        }
        
        maxY = ds.Max(maxY, yStart)
      } else {
        xEnd, err = strconv.Atoi(coord[0])
        if err != nil {
          panic("Cannot convert xEnd coord")
        }
        xEnd -= X_NORM // normalize col
        yEnd, err = strconv.Atoi(coord[1])
        if err != nil {
          panic("Cannot convert yEnd coord")
        }

        maxY = ds.Max(maxY, yEnd)
       
        drawRocks(&cave, xStart, yStart, xEnd, yEnd)
        xStart, yStart = xEnd, yEnd
      }
    }
  }

  return cave, maxY
}

func displayCave(cave Cave) {
  for i := 0; i < MAX_ROWS; i++ {
    for j := 0; j < MAX_COLS; j++ {
      var c string
      switch cave[i][j] {
      case EMPTY:
        c = "."
      case ROCK:
        c = "#"
      case SAND:
        c = "o"
      default:
        panic("Value not known")
      }
      fmt.Print(c)
    }
    fmt.Println()
  }
  fmt.Println("\n")
}

func canFallDown(cave Cave, x, y int) bool {
  return y+1 < MAX_ROWS && cave[y+1][x] == EMPTY
}
func canFallLeft(cave Cave, x, y int) bool {
  return x-1 >= 0 &&
         x-1 <= MAX_COLS-1 &&
         cave[y+1][x-1] == EMPTY
}
func canFallRight(cave Cave, x, y int) bool {
  return x+1 >= 0 &&
         x+1 < MAX_COLS &&
         cave[y+1][x+1] == EMPTY
}
func isFallingForever(x, y int) bool {
  return x < 0 ||
         x >= MAX_COLS ||
         y+1 >= MAX_ROWS
}
func canMove(cave Cave, x, y int) bool {
  return !isFallingForever(x, y) &&
         (canFallDown(cave, x, y) ||
         canFallLeft(cave, x, y) ||
         canFallRight(cave, x, y))
}

func makeSandFall(cave *Cave) bool {
  currX, currY := 500 - X_NORM, 0

  for canMove(*cave, currX, currY) {
    if canFallDown(*cave, currX, currY) {
      currY++
    } else if canFallLeft(*cave, currX, currY) {
      currY++
      currX--
    } else if canFallRight(*cave, currX, currY) {
      currY++
      currX++
    }
  }

  if isFallingForever(currX, currY) {
    return false
  }

  // fmt.Printf(">> Sand grain in (%d,%d)\n", currY, currX)
  (*cave)[currY][currX] = SAND
  
  return currX != 500 - X_NORM || currY != 0
}

func part1(cave Cave) {
  sandGrains := 0
  for {
    if !makeSandFall(&cave) {
      break
    }
    sandGrains++
    // displayCave(cave)
  }
  displayCave(cave)

  fmt.Println("In total", sandGrains, "sand grains fell down\n")
}

func part2(cave Cave, maxY int) {
  sandGrains := 0

  for j := 0; j < MAX_COLS; j++ {
    cave[maxY+2][j] = ROCK
  }

  for {
    if !makeSandFall(&cave) {
      break
    }
    sandGrains++
    if sandGrains%100 == 0 {
      fmt.Println("Already", sandGrains, "sand grains...")
    }
    // displayCave(cave)
  }
  sandGrains++ // missing the top one
  // displayCave(cave)

  fmt.Println("In total", sandGrains, "sand grains fell down")
}

func main() {
  // f, err := os.Open("mini_input")
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  cave, maxY := parseInput(f)
  part1(cave)
  part2(cave, maxY)
}

