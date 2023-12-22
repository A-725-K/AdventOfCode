package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/openacid/slimarray/polyfit"
	// "runtime"
	// "time"
)

type Garden [][]rune
type CoordArr [3]int // 5 in the failed attempt
type Queue struct {
  Queue []CoordArr
  Size int
}

func (c CoordArr) ToKey() string {
  return fmt.Sprintf("%d:%d", c[0], c[1])
}

func FromKey(s string) CoordArr {
  fields := strings.Split(s, ":")
  c0, _ := strconv.Atoi(fields[0])
  c1, _ := strconv.Atoi(fields[1])
  return CoordArr{c0, c1}
}

// XXX: this functions belong to the failed attempt
// func (c CoordArr) ToKeysBig(g Garden) (string, string) {
//   cp := CoordArr{c[0], c[1], c[2]}
//   g.NormalizeCoord(&cp)
//   gridKey := fmt.Sprintf("%d:%d", cp[3], cp[4])
//   coordKey := cp.ToKey()
//   return gridKey, coordKey
// }

func NewQueue() Queue {
  return Queue{Queue: []CoordArr{}, Size: 0}
}

func (q *Queue) Enqueue(el CoordArr) {
  q.Queue = append(q.Queue, el)
  q.Size++
}

func (q *Queue) Dequeue() CoordArr {
  defer func() {
    q.Size--
    q.Queue = q.Queue[1:]
  }()
  return q.Queue[0]
}

func (q Queue) IsEmpty() bool {
  return q.Size == 0
}

func (g Garden) Display() {
  for _, row := range g {
    for _, ch := range row {
      fmt.Print(string(ch) + " ")
    }
    fmt.Println()
  }
}

func (g Garden) IsValid(c CoordArr) bool {
  i, j := c[0], c[1]
  return i >= 0 && i < len(g) && j >= 0 && j < len(g[0]) && g[i][j] != '#'
}

// XXX: this functions belong to the failed attempt
func (g Garden) IsValidBig(c CoordArr) bool {
  cp := c
  g.NormalizeCoord(&cp)
  return g[cp[0]][cp[1]] != '#'
}

// XXX: these functions belong to the failed attempt
// func (g Garden) NormalizeCoord(c *CoordArr) {
//   n, m := len(g), len(g[0])
//   mod := func (a, b int) int {
// 	  var res int = a % b
// 	  if (res < 0 && b > 0) || (res > 0 && b < 0) {
// 		  return res + b
// 	  }
// 	  return res
//   }
//   (*c)[3] = (*c)[0]
//   (*c)[4] = (*c)[1]
//   var gridI, gridJ int
//   if (*c)[3] < 0 {
//     for (*c)[3] < 0 {
//       (*c)[3] += n
//       gridI--
//     }
//     (*c)[3] = gridI
//   } else {
//     (*c)[3] = (*c)[0] / n
//   }
//   if (*c)[4] < 0 {
//     for (*c)[4] < 0 {
//       (*c)[4] += m
//       gridJ--
//     }
//     (*c)[4] = gridJ
//   } else {
//     (*c)[4] = (*c)[1] / m
//   }
//   (*c)[0] = mod((*c)[0], n)
//   (*c)[1] = mod((*c)[1], m)
// }

func parseInput(f *os.File) (Garden, CoordArr) {
  var start CoordArr
  var garden Garden
  scanner := bufio.NewScanner(f)
  for i := 0; scanner.Scan(); i++ {
    var row []rune
    for j, ch := range scanner.Text() {
      row = append(row, ch)
      if ch == 'S' {
        start[0] = i
        start[1] = j
      }
    }
    garden = append(garden, row)
  }
  return garden, start
}

func (garden Garden) WalkInTheGarden(start CoordArr, steps int) int {
  var key string

  memo := make(map[string]bool)
  finalGardenPlots := make(map[string]bool)

  q := NewQueue()
  q.Enqueue(start)
  for !q.IsEmpty() {
    currPos := q.Dequeue()

    // only if it is at an even number of steps because you can bounce between
    // two tiles and that would be count as valid
    if currPos[2] % 2 == 0 {
      finalGardenPlots[currPos.ToKey()] = true
    }
    if currPos[2] >= steps {
      continue
    }
    
    // move up
    up := CoordArr{currPos[0]-1, currPos[1], currPos[2]+1}
    key = up.ToKey()
    if _, seen := memo[key]; garden.IsValid(up) && !seen {
      memo[key] = true
      q.Enqueue(up)
    }
    // move down
    down := CoordArr{currPos[0]+1, currPos[1], currPos[2]+1}
    key = down.ToKey()
    if _, seen := memo[key]; garden.IsValid(down) && !seen {
      memo[key] = true
      q.Enqueue(down)
    }
    // move left
    left := CoordArr{currPos[0], currPos[1]-1, currPos[2]+1}
    key = left.ToKey()
    if _, seen := memo[key]; garden.IsValid(left) && !seen {
      memo[key] = true
      q.Enqueue(left)
    }
    // move right
    right := CoordArr{currPos[0], currPos[1]+1, currPos[2]+1}
    key = right.ToKey()
    if _, seen := memo[key]; garden.IsValid(right) && !seen {
      memo[key] = true
      q.Enqueue(right)
    }
  }

  return len(finalGardenPlots)
}

func part1(garden Garden, start CoordArr, steps int) {
  finalGardenPlots := garden.WalkInTheGarden(start, steps)

  fmt.Println(
    "The elf can reach", finalGardenPlots, "garden plots in",
    steps, "steps",
  )
}

// OBS: even though this algorithm optimize the running time of part 1, it is
// too much resource intense. The memory check implemented to avoid topping up
// all the RAM has been triggered during the execution, before getting to a
// result :( The memory limit was set to 27 GB, which is WAY TOO MUCH. There
// has to exist a more efficient solution to this problem. I am a little bit sad
// because I though that the modulo approach would have optimized more the
// algorithm :(
//
// func limitMemoryUsage() {
//   var m runtime.MemStats
//   const GB_27 = 27 * 1024 * 1024 * 1024
//
//   for {
// 	  runtime.ReadMemStats(&m)
// 	  if m.Sys > GB_27 {
// 	    panic("Too much memory")
// 	  }
// 	  time.Sleep(1 * time.Second)
//   }
// }
// 
// func part2_FAILURE(garden Garden, start CoordArr, steps int) {
//   var gridKey, coordKey string
//
//   // now the cache has this format:
//   // - gridX:gridY -> key for identifying the coordinates of the square in the grid
//   //            consisting of the replicas of the initial garden
//   // - i:j         -> the coordinate of the point in the selected grid
//   memo := make(map[string]map[string]bool)
//   numOfFinalGardenPlots := int64(0)
//
//   q := NewQueue()
//   q.Enqueue(start)
//   for !q.IsEmpty() {
//     currPos := q.Dequeue()
//
//     // only if it is at an even number of steps because you can bounce between
//     // two tiles and that would be count as valid
//     if currPos[2] % 2 == 0 {
//       numOfFinalGardenPlots++
//       if numOfFinalGardenPlots % 10_000_000 == 0 {
//         fmt.Println(numOfFinalGardenPlots)
//       }
//     }
//     if currPos[2] >= steps {
//       continue
//     }
//
//     // move up
//     up := CoordArr{currPos[0]-1, currPos[1], currPos[2]+1}
//     gridKey, coordKey = up.ToKeysBig(garden)
//     _, gridExists := memo[gridKey]
//     if !gridExists {
//       memo[gridKey] = make(map[string]bool)
//     }
//     if _, seen := memo[gridKey][coordKey]; garden.IsValidBig(up) && !seen {
//       memo[gridKey][coordKey] = true
//       q.Enqueue(up)
//     }
//     // move down
//     down := CoordArr{currPos[0]+1, currPos[1], currPos[2]+1}
//     gridKey, coordKey = down.ToKeysBig(garden)
//     _, gridExists = memo[gridKey]
//     if !gridExists {
//       memo[gridKey] = make(map[string]bool)
//     }
//     if _, seen := memo[gridKey][coordKey]; garden.IsValidBig(down) && !seen {
//       memo[gridKey][coordKey] = true
//       q.Enqueue(down)
//     }
//     // move left
//     left := CoordArr{currPos[0], currPos[1]-1, currPos[2]+1}
//     gridKey, coordKey = left.ToKeysBig(garden)
//     _, gridExists = memo[gridKey]
//     if !gridExists {
//       memo[gridKey] = make(map[string]bool)
//     }
//     if _, seen := memo[gridKey][coordKey]; garden.IsValidBig(left) && !seen {
//       memo[gridKey][coordKey] = true
//       q.Enqueue(left)
//     }
//     // move right
//     right := CoordArr{currPos[0], currPos[1]+1, currPos[2]+1}
//     gridKey, coordKey = right.ToKeysBig(garden)
//     _, gridExists = memo[gridKey]
//     if !gridExists {
//       memo[gridKey] = make(map[string]bool)
//     }
//     if _, seen := memo[gridKey][coordKey]; garden.IsValidBig(right) && !seen {
//       memo[gridKey][coordKey] = true
//       q.Enqueue(right)
//     }
//   }
//
//   // this algorithm gives the result off-by-one, need to subtract 1
//   numOfFinalGardenPlots--
//   fmt.Println(
//     "The elf can reach", numOfFinalGardenPlots, "garden plots in",
//     steps, "steps in the real bigger garden",
//   )
// }

func mod(a, b int) int {
	var res int = a % b
	if (res < 0 && b > 0) || (res > 0 && b < 0) {
		return res + b
	}
	return res
}

func (g Garden) NormalizeCoord(c *CoordArr) {
  n, m := len(g), len(g[0])
  (*c)[0] = mod((*c)[0], n)
  (*c)[1] = mod((*c)[1], m)
}

func (g Garden) SingleStep(positions *map[string]bool) {
  newPos := make(map[string]bool)

  for p := range *positions {
    cp := FromKey(p)

    // up
    up := CoordArr{cp[0]-1, cp[1]}
    if g.IsValidBig(up) {
      newPos[up.ToKey()] = true
    }
    // down
    down := CoordArr{cp[0]+1, cp[1]}
    if g.IsValidBig(down) {
      newPos[down.ToKey()] = true
    }
    // left
    left := CoordArr{cp[0], cp[1]-1}
    if g.IsValidBig(left) {
      newPos[left.ToKey()] = true
    }
    // right
    right := CoordArr{cp[0], cp[1]+1}
    if g.IsValidBig(right) {
      newPos[right.ToKey()] = true
    }
  }

  (*positions) = newPos
}

// Thanks a lot Anshuman Dash :D
// Assumptions:
//   - no blocking rocks on the row and the column of the starting point
//   - the initial grid is a square, i.e. len(row) == len(col)
//
// Facts:
//   - it takes 65 steps to map the entire starting region, or in other words,
//     to walk from the middle of the grid to every border
//   - corollary of the first fact: it takes 131 steps to go from the center of
//     a square to the center of the next one in every direction
//   - there is an underlying quadratic relation when growing:
//       f(x) := how many squares are visited at time 65 + 131x
//   - every time we map a new region, we map #prev+4 squares:
//                                           O
//                          O              O X O
//             O          O X O          O X X X O
//   O --->  O X O ---> O X X X O ---> O X X X X X O
//             O          O X O          O X X X O
//                          O              O X O
//                                           O
//   +1       +4           +8               +12
// Video explanation here:
//   - https://www.youtube.com/watch?v=xHIQ2zHVSjM
//   - https://www.youtube.com/watch?v=yVJ-J5Nkios (follow up)
func part2(garden Garden, start CoordArr, steps int) {
  xs := []float64{0, 1, 2}
  ys := []float64{}
  deg := 2

  // value of x, also a nice round integer: 202300
  xVar := float64((steps-65)/131)

  positions := make(map[string]bool)
  positions[start.ToKey()] = true

  for s := 0; s < (65 + 131*2 + 1); s++ {
    // we hit one of the values of f(x), we append it to Ys
    if s%131 == 65 {
      ys = append(ys, float64(len(positions)))
    }
    garden.SingleStep(&positions)
  }

  // create a polynomial of degree 2 to interpolate the points
  linearLeastSquareModel := NewFit(xs, ys, deg)
  polynomial := linearLeastSquareModel.Solve()

  // evaluate the quadratic model in the X
  result := float64(0)
	pow := float64(1)
	for _, coef := range polynomial {
		result += coef * pow
		pow *= xVar
	}
	// fmt.Println(xs, ys, polynomial)

  fmt.Println(
    "The elf can reach", int64(result), "garden plots in",
    steps, "steps in the real bigger garden",
  )
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  garden, start := parseInput(f)
  // fmt.Println(start)
  // garden.Display()

  // go limitMemoryUsage()
 
  part1(garden, start, 64)
  part2(garden, start, 26501365)
}
