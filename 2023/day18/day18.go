package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
  Direction, Color string
  Steps int
}
type CoordArray [2]int
type Segment [2]CoordArray

func (c CoordArray) String() string {
  return fmt.Sprintf("%d:%d", c[0], c[1])
}

func String2Coord(s string) CoordArray {
  f := strings.Split(s, ":")
  i, _ := strconv.Atoi(f[0])
  j, _ := strconv.Atoi(f[1])
  return CoordArray{i, j}
}

func (i Instruction) GetRealInstruction() Instruction {
  var realSteps int64
  var realDirection string

  switch i.Color[5] {
  case '0':
    realDirection = "R"
  case '1':
    realDirection = "D"
  case '2':
    realDirection = "L"
  case '3':
    realDirection = "U"
  default:
    panic("Wrong instruction!")
  }

  realSteps, _ = strconv.ParseInt(i.Color[:5], 16, 32)

  return Instruction{Direction: realDirection, Steps: int(realSteps)}
}

func parseInput(f *os.File) []Instruction {
  var instructions []Instruction
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    fields := strings.Split(scanner.Text(), " ")
    steps, _ := strconv.Atoi(fields[1])
    instructions = append(instructions, Instruction{Steps: steps, Direction: fields[0], Color: fields[2][2:8]})
  }
  return instructions
}

func displayGrid(grid [][]bool, n, m int) {
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      if grid[i][j] {
        fmt.Print("#")
      } else {
        fmt.Print(".")
      }
    }
    fmt.Println()
  }
  fmt.Println()
}

func part1(instructions []Instruction) {
  border := make(map[string]bool)
  currentPoint := CoordArray{0, 0}
  border[currentPoint.String()] = true
  var borderSegments []Segment

  for _, instruction := range instructions {
    var segment Segment
    segment[0] = CoordArray{currentPoint[0], currentPoint[1]}
    for s := 0; s < instruction.Steps; s++ {
      switch instruction.Direction {
      case "R":
        currentPoint[1]++
      case "L":
        currentPoint[1]--
      case "U":
        currentPoint[0]--
      case "D":
        currentPoint[0]++
      default:
        panic("Unexpected direction")
      }
      border[currentPoint.String()] = true
    }
    segment[1] = CoordArray{currentPoint[0], currentPoint[1]}
    borderSegments = append(borderSegments, segment)
  }

  // normalize the matrix
  minI, minJ := math.MaxInt, math.MaxInt
  for b := range border {
    c := String2Coord(b)
    if c[0] < minI {
      minI = c[0]
    }
    if c[1] < minJ {
      minJ = c[1]
    }
  }
  
  // normalize border
  newBorder := make(map[string]bool)
  for b := range border {
    c := String2Coord(b)
    c[0] += -minI
    c[1] += -minJ
    newBorder[c.String()] = true
  }
  // normalize border segments
  for i := 0; i < len(borderSegments); i++ {
    borderSegments[i][0][0] += -minI
    borderSegments[i][1][0] += -minI

    borderSegments[i][0][1] += -minJ
    borderSegments[i][1][1] += -minJ
  }

  n, m := 0, 0
  for b := range newBorder {
    c := String2Coord(b)
    if c[0] > n {
      n = c[0]
    }
    if c[1] > m {
      m = c[1]
    }
  }
  n++
  m++

  grid := make([][]bool, n)
  for i := 0; i < n; i++ {
    grid[i] = make([]bool, m)
  }
  for b := range newBorder {
    c := String2Coord(b)
    grid[c[0]][c[1]] = true
  }
  // displayGrid(grid, n, m)

  // Ray casting algorithm: https://en.wikipedia.org/wiki/Point_in_polygon
  rayIntersectSegment := func (i, j int, side Segment) bool {
    var a, b CoordArray
    if side[0][0] < side[1][0] {
      a, b = side[0], side[1]
    } else {
      a, b = side[1], side[0]
    }

    for i == a[0] || i == b[0] {
      i++
    }
    if i < a[0] || i > b[0] {
      return false
    }

    if a[1] > b[1] {
      if j > a[1] {
        return false
      }
      if j < b[1] {
        return true
      }
    } else {
      if j > b[1] {
        return false
      }
      if j < a[1] {
        return true
      }
    }

    return float64((i-a[0])/(j-a[1])) >= float64((b[0]-a[0])/(b[1]-a[1]))
  }
  isInternal := func (borderSegments []Segment, i, j, n, m int) bool {
    isIn := false
    for _, side := range borderSegments {
      if rayIntersectSegment(i, j, side) {
        isIn = !isIn
      }
    }
    return isIn
  }

  // count the internal + border tiles
  cubicMetersOfLava := 0
  var internals []CoordArray
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      if _, isBorder := newBorder[CoordArray{i, j}.String()]; isBorder {
        cubicMetersOfLava++
        continue
      }

      isIn := isInternal(borderSegments, i, j, n, m)
      if isIn {
        internals = append(internals, CoordArray{i, j})
        cubicMetersOfLava++
      }
    }
  }

  for _, ip := range internals {
    grid[ip[0]][ip[1]] = true
  }

  // displayGrid(grid, n, m)
  fmt.Println("The lagoon can contain", cubicMetersOfLava, "m^3 of lava")
}

func part2(instructions []Instruction) {
  var pointsOnBorder []CoordArray
  currentPoint := CoordArray{0, 0}
  pointsOnBorder = append(pointsOnBorder, currentPoint)
  perimeter := float64(0)

  for _, instruction := range instructions {
    realInstruction := instruction.GetRealInstruction()
    switch realInstruction.Direction {
    case "R":
      currentPoint[1] += realInstruction.Steps
    case "L":
      currentPoint[1] -= realInstruction.Steps
    case "U":
      currentPoint[0] -= realInstruction.Steps
    case "D":
      currentPoint[0] += realInstruction.Steps
    default:
      panic("Unexpected direction")
    }
    pointsOnBorder = append(pointsOnBorder, currentPoint)
    perimeter += float64(realInstruction.Steps)
  }

  // Shoelace Formula: https://en.wikipedia.org/wiki/Shoelace_formula
  // Determine the area of a simple polygon from the set of vertices
  // A = 1/2 * sum_i=1->n((y[i] + y[i+1]) * (x[i] - x[i+1]))
  areaInternal := float64(0)
  n := len(pointsOnBorder)
  for i := 0; i < n; i++ {
    areaInternal += float64(pointsOnBorder[i][1] + pointsOnBorder[(i+1)%n][1]) *
      float64(pointsOnBorder[i][0] - pointsOnBorder[(i+1)%n][0])
  }
  areaInternal /= -2 // instead of abs function divide by -2 instead of 2

  // Pick's theorem: https://en.wikipedia.org/wiki/Pick%27s_theorem
  // Compute area starting from the set of vertices in terms of the number of
  // integer points within it and on its boundary
  // A = interior_points - perimeter/2 - 1
  // BE CAREFUL: adjust rounding error !!!
  cubicMetersOfLava := int64(areaInternal+1) + (int64(perimeter+1)/2+1) - 1

  fmt.Println("The lagoon actually can contain", cubicMetersOfLava, "m^3 of lava")
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  instructions := parseInput(f)
 
  part1(instructions)
  part2(instructions)
}
