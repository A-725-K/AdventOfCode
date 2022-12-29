package main

import (
  "os"
  "fmt"
  "bufio"
  "AdventOfCode/ds"
  set "AdventOfCode/ds/set"
  queue "AdventOfCode/ds/queue"
  coord "AdventOfCode/types/coord"
)

const (
  NORTH = 0
  SOUTH = 1
  EAST = 2
  WEST = 3
  WALL = 4
)

func parseInput(f *os.File) (set.Set[string], int, int) {
  scanner := bufio.NewScanner(f)
  
  row, cols := 0, 0
  storm := set.NewSet[string]()
  for scanner.Scan() {
    line := scanner.Text()
    cols = len(line)

    for col, c := range line {
      switch c {
      case '>':
        storm.Add(coord.CoordWithDist{X: col, Y: row, Distance: EAST}.ToKeyDist())
      case '<':
        storm.Add(coord.CoordWithDist{X: col, Y: row, Distance: WEST}.ToKeyDist())
      case '^':
        storm.Add(coord.CoordWithDist{X: col, Y: row, Distance: NORTH}.ToKeyDist())
      case 'v':
        storm.Add(coord.CoordWithDist{X: col, Y: row, Distance: SOUTH}.ToKeyDist())
      case '#':
        storm.Add(coord.CoordWithDist{X: col, Y: row, Distance: WALL}.ToKeyDist())
      default:
        continue
      }
    }
    row++
  }

  return storm, row, cols
}

func isValid(cc coord.CoordWithDist, invalid set.Set[string], rows, cols int) bool {
  if cc.X < 0 || cc.X > cols || cc.Y < 0 || cc.Y > rows {
    return false
  }
  for invKey := range invalid {
    inv := coord.FromKeyDist(invKey)
    if inv.X == cc.X && inv.Y == cc.Y {
      return false
    }
  }
  return true
}

func computeStorms(storms *map[int]set.Set[string], rows, cols int) {
  for t := 1; t < rows*cols; t++ {
    prevStorm := (*storms)[t-1]
    currStorm := set.NewSet[string]()
    for stKey := range prevStorm {
      st := coord.FromKeyDist(stKey)
      switch st.Distance {
        case NORTH:
          nextY := st.Y-1
          if nextY == 0 {
            nextY = rows-2
          }
          currStorm.Add(coord.CoordWithDist{X: st.X, Y: nextY, Distance: NORTH}.ToKeyDist())
        case SOUTH:
          nextY := st.Y+1
          if nextY == rows-1 {
            nextY = 1
          }
          currStorm.Add(coord.CoordWithDist{X: st.X, Y: nextY, Distance: SOUTH}.ToKeyDist())
        case EAST:
          nextX := st.X+1
          if nextX == cols-1 {
            nextX = 1
          }
          currStorm.Add(coord.CoordWithDist{X: nextX, Y: st.Y, Distance: EAST}.ToKeyDist())
        case WEST:
          nextX := st.X-1
          if nextX == 0 {
            nextX = cols-2
          }
          currStorm.Add(coord.CoordWithDist{X: nextX, Y: st.Y, Distance: WEST}.ToKeyDist())
        case WALL:
          currStorm.Add(coord.CoordWithDist{X: st.X, Y: st.Y, Distance: WALL}.ToKeyDist())
      }
    }
    (*storms)[t] = currStorm
  }
}

func display(storm set.Set[string], rows, cols int) {
  for y := 0; y < rows; y++ {
    for x := 0; x < cols; x++ {
      if isValid(coord.CoordWithDist{X: x, Y: y, Distance: 0}, storm, rows, cols) {
        fmt.Print(".")
      } else {
        fmt.Print("#")
      }
    }
    fmt.Println()
  }
  fmt.Println()
}

func walkInTheBlizzard(
  storms map[int]set.Set[string],
  rows, cols int,
  ends []coord.Coord,
) int {
  lcm := ds.Lcm(rows, cols)
  numberOfSteps := len(ends)

  q := queue.NewQueue[coord.CoordWithDist]()
  q.Enqueue(coord.CoordWithDist{X: 1, Y: 0, Distance: 0})
  
  step := 0
  alreadyVisited := set.NewSet[string]()
  for !q.IsEmpty() {
    currPos := q.Dequeue()
    currStorm := storms[(currPos.Distance+1)%lcm]

    if alreadyVisited.Contains(currPos.ToKeyDist()) {
      continue
    }
    alreadyVisited.Add(currPos.ToKeyDist())

    // fmt.Printf("VISITING (%d,%d) at time %d\n", currPos.X, currPos.Y, currPos.Distance)
    // display(currStorm, rows, cols)

    if currPos.X == ends[step].X && currPos.Y == ends[step].Y {
      for !q.IsEmpty() {
        q.Dequeue()
      }
      q.Enqueue(coord.CoordWithDist{
        X: ends[step].X,
        Y: ends[step].Y,
        Distance: currPos.Distance,
      })
      step++
      if step == numberOfSteps {
        return currPos.Distance
      }
    }

    north := coord.CoordWithDist{
      X: currPos.X,
      Y: currPos.Y-1,
      Distance: currPos.Distance+1,
    }
    if isValid(north, currStorm, rows, cols) {
      q.Enqueue(north)
    }

    south := coord.CoordWithDist{
      X: currPos.X,
      Y: currPos.Y+1,
      Distance: currPos.Distance+1,
    }
    if isValid(south, currStorm, rows, cols) {
      q.Enqueue(south)
    }

    west := coord.CoordWithDist{
      X: currPos.X-1,
      Y: currPos.Y,
      Distance: currPos.Distance+1,
    }
    if isValid(west, currStorm, rows, cols) {
      q.Enqueue(west)
    }

    east := coord.CoordWithDist{
      X: currPos.X+1,
      Y: currPos.Y,
      Distance: currPos.Distance+1,
    }
    if isValid(east, currStorm, rows, cols) {
      q.Enqueue(east)
    }

    wait := coord.CoordWithDist{
      X: currPos.X,
      Y: currPos.Y,
      Distance: currPos.Distance+1,
    }
    if isValid(wait, currStorm, rows, cols) {
      q.Enqueue(wait)
    }
  }

  // Cannot find a path to the end
  return -1
}

func part1(storms map[int]set.Set[string], rows, cols int) {
  ends := []coord.Coord{{X: cols-2, Y: rows-1}}
  minNumOfSteps := walkInTheBlizzard(storms, rows, cols, ends)
  fmt.Println("The minimum number of steps to avoid the blizzard is:", minNumOfSteps)
}

func part2(storms map[int]set.Set[string], rows, cols int) {
  ends := []coord.Coord{
    coord.Coord{X: cols-2, Y: rows-1},
    coord.Coord{X: 1, Y: 0},
    coord.Coord{X: cols-2, Y: rows-1},
  }
  minNumOfSteps := walkInTheBlizzard(storms, rows, cols, ends)
  fmt.Println("The minimum number of steps to go back for the snack is:", minNumOfSteps)
}

func main() {
  f, err := os.Open("mini_input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  storm, rows, cols := parseInput(f)

  storms := make(map[int]set.Set[string])
  storms[0] = storm
  computeStorms(&storms, rows, cols)

  part1(storms, rows, cols)
  part2(storms, rows, cols)
}

