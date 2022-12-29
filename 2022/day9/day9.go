package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
  "AdventOfCode/ds"
  c "AdventOfCode/types/coord"
)

type Move struct {
  steps int
  direction string
}

func printGrid(allPos []c.Coord) {
  N := 10
  mmap := make(map[string]bool)
  for i := N; i >= -N; i-- {
    for j := -N; j < N; j++ {
      if i == 0 && j == 0 {
        fmt.Print("s")
      } else {
        found := false
        for h, p := range allPos {
          if p.X == j && p.Y == i {
            var c string
            if h == 0 {
              c = "H"
            } else {
              c = strconv.Itoa(h)
            }
            found = true
            if _, ok := mmap[ds.ToKey(i, j)]; !ok {
              fmt.Print(c)
              mmap[ds.ToKey(i, j)] = true
            }
          }
        }
        if !found {
          fmt.Print(".")
        }
      }
    }
    fmt.Println()
  }
}

func parseInput(f *os.File) []Move {
  scanner := bufio.NewScanner(f)
  var moves []Move
  
  for scanner.Scan() {
    line := scanner.Text()
    fields := strings.Split(line, " ")
    steps, err := strconv.Atoi(fields[1])
    if err != nil {
      panic("Cannot convert steps")
    }
    moves = append(moves, Move{steps, fields[0]})
  }

  return moves
}

// func (p1 c.Coord) isAdjacent(p2 c.Coord) bool {
//   same := p2.X == p1.X && p2.Y == p1.Y
//   hor := (p2.X == (p1.X + 1) || p2.X == (p1.X - 1)) && p2.Y == p1.Y
//   ver := (p2.Y == (p1.Y + 1) || p2.Y == (p1.Y - 1)) && p2.X == p1.X
//   diag := (p2.X == (p1.X + 1) && p2.Y == (p1.Y + 1)) ||
//           (p2.X == (p1.X + 1) && p2.Y == (p1.Y - 1)) ||
//           (p2.X == (p1.X - 1) && p2.Y == (p1.Y + 1)) ||
//           (p2.X == (p1.X - 1) && p2.Y == (p1.Y - 1))
//   return same || hor || ver || diag
//
//   // SMARTER!!!!
//   // return abs(p1.X - p2.X) <= 1 && abs(p1.Y - p2.Y) <= 1
// }
//
// func horOff(pos1, pos2 c.Coord) int {
//   if pos1.X > pos2.X {
//     return 1
//   }
//   if pos1.X < pos2.X {
//     return -1
//   }
//   return 0
// }
//
// func verOff(pos1, pos2 c.Coord) int {
//   if pos1.Y > pos2.Y {
//     return 1
//   }
//   if pos1.Y < pos2.Y {
//     return -1
//   }
//   return 0
// }
//
// // FIXME: FIX THIS VERSION.... IT WORKS ONLY FOR THE FIRST PART
// // CASE THAT BREAKS THE ALGORITHM:
//
// 4 | | | | |
// 3 | |H|1| |
// 2 | | | | |
// 1 | | | |2|
// 0 | | | |3|
//    0 1 2 3
//
// EXPECTED MOVE: [2] (3,1) -> (3, 2)
// MY MOVE: [2] (3, 1) -> (2, 2)
//
//
// func moveRope(m Move, allPos *[]c.Coord, tailPos *map[string]bool) {
//   n := len(*allPos)
//   for i := 0; i < m.steps; i++ {
//     switch m.direction {
//     case "U":
//         (*allPos)[0].Y++
//         for pos := 1; pos < n; pos++ {
//           if !(*allPos)[pos].isAdjacent((*allPos)[pos-1]) {
//             (*allPos)[pos].Y++
//             (*allPos)[pos].X += horOff((*allPos)[pos-1], (*allPos)[pos])
//           }
//         }
//     case "D":
//         (*allPos)[0].Y--
//         for pos := 1; pos < n; pos++ {
//           if !(*allPos)[pos].isAdjacent((*allPos)[pos-1]) {
//             (*allPos)[pos].Y--
//             (*allPos)[pos].X += horOff((*allPos)[pos-1], (*allPos)[pos])
//           }
//         }
//     case "R":
//         (*allPos)[0].X++
//         for pos := 1; pos < n; pos++ {
//           if !(*allPos)[pos].isAdjacent((*allPos)[pos-1]) {
//             (*allPos)[pos].X++
//             (*allPos)[pos].Y += verOff((*allPos)[pos-1], (*allPos)[pos])
//           }
//         }
//     case "L":
//         (*allPos)[0].X--
//         for pos := 1; pos < n; pos++ {
//           if !(*allPos)[pos].isAdjacent((*allPos)[pos-1]) {
//             (*allPos)[pos].X--
//             (*allPos)[pos].Y += verOff((*allPos)[pos-1], (*allPos)[pos])
//           }
//         }
//     default:
//       panic("Direction not known")
//     }
//
//     (*tailPos)[toKey((*allPos)[n-1].X, (*allPos)[n-1].Y)] = true
//   }
// }

func moveRope(m Move, allPos *[]c.Coord, tailPos *map[string]bool) {
  n:= len(*allPos)

  for i := 0; i < m.steps; i++ {
    switch m.direction {
    case "U":
        (*allPos)[0].Y++
    case "D":
        (*allPos)[0].Y--
    case "R":
        (*allPos)[0].X++
    case "L":
        (*allPos)[0].X--
    default:
      panic("Direction not known")
    }

    for pos := 1; pos < n; pos++ {
      dx, dy := (*allPos)[pos].MoveCloser((*allPos)[pos-1])
      (*allPos)[pos].X += dx
      (*allPos)[pos].Y += dy
    }

    (*tailPos)[ds.ToKey((*allPos)[n-1].X, (*allPos)[n-1].Y)] = true
  }
}

func part1(moves []Move) {
  allPos := []c.Coord{{X: 0, Y: 0}, {X: 0, Y: 0}}
  
  tailPos := make(map[string]bool)
  tailPos[ds.ToKey(allPos[0].X, allPos[0].Y)] = true
  
  for _, m := range moves {
    moveRope(m, &allPos, &tailPos)
  }

  fmt.Println("In total the tail visited", len(tailPos), "locations")
}

func part2(moves []Move) {
  var allPos []c.Coord
  for i := 0; i < 10; i++ {
    allPos = append(allPos, c.Coord{X: 0, Y: 0})
  }

  tailPos := make(map[string]bool)
  tailPos[ds.ToKey(allPos[0].X, allPos[0].Y)] = true
  
  for _, m := range moves {
    moveRope(m, &allPos, &tailPos)
  }

  fmt.Println("In total the tail visited", len(tailPos), "locations")
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }

  moves := parseInput(f)
  part1(moves)
  part2(moves)
}

