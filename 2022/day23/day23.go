package main

import (
  "os"
  "fmt"
  "bufio"
  c "AdventOfCode/types/coord"
)

const ROUNDS = 10

func checkNorth(elf c.Coord, otherPos map[string]int) bool {
  ne := c.Coord{X: elf.X+1, Y: elf.Y-1}
  nw := c.Coord{X: elf.X-1, Y: elf.Y-1}
  nn := c.Coord{X: elf.X,   Y: elf.Y-1}

  _, neOk := otherPos[ne.ToKey()]
  _, nwOk := otherPos[nw.ToKey()]
  _, nnOk := otherPos[nn.ToKey()]

  return !neOk && !nwOk && !nnOk
}
func checkSouth(elf c.Coord, otherPos map[string]int) bool {
  se := c.Coord{X: elf.X+1, Y: elf.Y+1}
  sw := c.Coord{X: elf.X-1, Y: elf.Y+1}
  ss := c.Coord{X: elf.X,   Y: elf.Y+1}

  _, seOk := otherPos[se.ToKey()]
  _, swOk := otherPos[sw.ToKey()]
  _, ssOk := otherPos[ss.ToKey()]

  return !seOk && !swOk && !ssOk
}
func checkEast(elf c.Coord, otherPos map[string]int) bool {
  se := c.Coord{X: elf.X+1, Y: elf.Y+1}
  ne := c.Coord{X: elf.X+1, Y: elf.Y-1}
  ee := c.Coord{X: elf.X+1, Y: elf.Y}

  _, seOk := otherPos[se.ToKey()]
  _, neOk := otherPos[ne.ToKey()]
  _, eeOk := otherPos[ee.ToKey()]

  return !seOk && !neOk && !eeOk
}
func checkWest(elf c.Coord, otherPos map[string]int) bool {
  sw := c.Coord{X: elf.X-1, Y: elf.Y-1}
  nw := c.Coord{X: elf.X-1, Y: elf.Y+1}
  ww := c.Coord{X: elf.X-1, Y: elf.Y}

  _, swOk := otherPos[sw.ToKey()]
  _, nwOk := otherPos[nw.ToKey()]
  _, wwOk := otherPos[ww.ToKey()]

  return !swOk && !nwOk && !wwOk
}
func checkAll(elf c.Coord, otherPos map[string]int) bool {
  return checkNorth(elf, otherPos) &&
         checkSouth(elf, otherPos) &&
         checkEast(elf, otherPos) &&
         checkWest(elf, otherPos)
}

func idxToNSEW(i int) string {
  switch i {
  case 0:
    return "NORTH"
  case 1:
    return "SOUTH"
  case 2:
    return "WEST"
  case 3:
    return "EAST"
  default:
    panic("Impossible")
  }
}

func makeSuggestions(
  round int,
  otherPos map[string]int,
) map[string]int {
  checkCoordsFuncs := []func(c.Coord, map[string]int)bool{
    checkNorth,
    checkSouth,
    checkWest,
    checkEast,
  }
  coordsToAdd := []c.Coord{
    {X: 0,  Y: -1},
    {X: 0,  Y: 1},
    {X: -1, Y: 0},
    {X: 1,  Y: 0},
  }
  offset := len(otherPos) + 1

  suggestions := make(map[string]int)
  for elfPos, elfIdx := range otherPos {
    for i := 0; i < 4; i++ {
      idx := ((round%4)+i)%4
      // fmt.Println("Checking idx =", idxToNSEW(idx))
      elf := c.FromKey(elfPos)

      // if all 8 tiles around are free, don't move
      if checkAll(elf, otherPos) {
        break
      }
      if checkCoordsFuncs[idx](elf, otherPos) {
        suggestion := c.Coord{
          X: elf.X+coordsToAdd[idx].X,
          Y: elf.Y+coordsToAdd[idx].Y,
        }
        suggestionKey := suggestion.ToKey()
        if v, ok := suggestions[suggestionKey]; !ok {
          // fmt.Println(elf.ToKey(), "===>", suggestionKey)
          suggestions[suggestionKey] = elfIdx
        } else {
          if v > 0 {
            // in this way I am sure the value will become negative
            suggestions[suggestionKey] -= offset
          }
        }
        break
      }
    }
  }

  return suggestions
}

func parseInput(f *os.File) []c.Coord {
  scanner := bufio.NewScanner(f)

  var elves []c.Coord
  i := 0
  for scanner.Scan() {
    line := scanner.Text()
    for j, ch := range line {
      if ch == '#' {
        elves = append(elves, c.Coord{X: j, Y: i})
      }
    }
    i++
  }

  return elves
}

func printElves(elvesPos map[string]int) {
  var elves []c.Coord
  for k, _ := range elvesPos {
    elves = append(elves, c.FromKey(k))
  }
  minX, minY, maxX, maxY := c.MinMax(elves)

  for y := minY; y <= maxY; y++ {
    for x := minX; x <= maxX; x++ {
      pos := c.Coord{X: x, Y: y}
      if _, ok := elvesPos[pos.ToKey()]; ok {
        fmt.Print("#")
      } else {
        fmt.Print(".")
      }
    }
    fmt.Println()
  }
}

func part1(elves []c.Coord) {
  for rnd := 0; rnd < ROUNDS; rnd++ {
    otherPos := make(map[string]int)
    for idx, elf := range elves {
      otherPos[elf.ToKey()] = idx
    }

    // fmt.Println("===== ROUND", rnd, "=====")
    // printElves(otherPos)
    // fmt.Println()

    suggestions := makeSuggestions(rnd, otherPos)
    // fmt.Println(suggestions)
    // fmt.Println(elves)
    for newPosStr, elfIdx := range suggestions {
      if elfIdx >= 0 {
        // fmt.Println("Moving", elves[elfIdx], "to", newPosStr)
        elves[elfIdx] = c.FromKey(newPosStr)
      }
    }
  }
  
  // fmt.Println("===== END =====")
  // otherPos := make(map[string]int)
  // for idx, elf := range elves {
  //   otherPos[elf.ToKey()] = idx
  // }
  // printElves(otherPos)
  
  minX, minY, maxX, maxY := c.MinMax(elves)
  emptyTiles := (maxX-minX+1) * (maxY-minY+1) - len(elves)
  fmt.Println("After", ROUNDS, "rounds there are", emptyTiles, "empty tiles")
}

func part2 (elves []c.Coord) {
  allInPos := false
  rnd := 0
  for !allInPos {
    otherPos := make(map[string]int)
    for idx, elf := range elves {
      otherPos[elf.ToKey()] = idx
    }

    suggestions := makeSuggestions(rnd, otherPos)
    for newPosStr, elfIdx := range suggestions {
      if elfIdx >= 0 {
        elves[elfIdx] = c.FromKey(newPosStr)
      }
    }
    allInPos = len(suggestions) == 0
    rnd++
  }
  // fmt.Println("===== END =====")
  // otherPos := make(map[string]int)
  // for idx, elf := range elves {
  //   otherPos[elf.ToKey()] = idx
  // }
  // printElves(otherPos)

  fmt.Println("Elves are settled after", rnd, "rounds")
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  elves := parseInput(f)
  var elves2 []c.Coord
  for _, e := range elves {
    elves2 = append(elves2, e)
  }
  
  part1(elves)
  part2(elves2)
}

