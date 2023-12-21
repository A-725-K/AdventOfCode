package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
  X, Y, Z int
}
type Brick []Point
type BrickWall []Brick

func (p Point) String() string {
  return fmt.Sprintf("(%d, %d, %d)", p.X, p.Y, p.Z)
}

func (b Brick) String() string {
  return fmt.Sprintf("%s -> %s", b[0].String(), b[1].String())
}

func (b Brick) ToKey() string {
  return fmt.Sprintf(
    "%d:%d:%d:%d:%d:%d",
    b[0].X, b[0].Y, b[0].Z,
    b[1].X, b[1].Y, b[1].Z,
  )
}

func FromKey(s string) Brick {
  fields := strings.Split(s, ":")
  x0, _ := strconv.Atoi(fields[0])
  y0, _ := strconv.Atoi(fields[1])
  z0, _ := strconv.Atoi(fields[2])
  x1, _ := strconv.Atoi(fields[3])
  y1, _ := strconv.Atoi(fields[4])
  z1, _ := strconv.Atoi(fields[5])
  return Brick{
    Point{x0, y0, z0},
    Point{x1, y1, z1},
  }
}

func parseInput(f *os.File) BrickWall {
  var bricks BrickWall
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    fields := strings.Split(scanner.Text(), "~")
    p0 := strings.Split(fields[0], ",")
    x0, _ := strconv.Atoi(p0[0])
    y0, _ := strconv.Atoi(p0[1])
    z0, _ := strconv.Atoi(p0[2])
    p1 := strings.Split(fields[1], ",")
    x1, _ := strconv.Atoi(p1[0])
    y1, _ := strconv.Atoi(p1[1])
    z1, _ := strconv.Atoi(p1[2])
    brick := Brick{
      Point{x0, y0, z0},
      Point{x1, y1, z1},
    }
    bricks = append(bricks, brick)
  }
  return bricks
}

func maxI(x, y int) int {
  if x > y {
    return x
  }
  return y
}

func minI(x, y int) int {
  if x < y {
    return x
  }
  return y
}

func (brick Brick) Overlap(brickBelow Brick) bool {
  // check if overlapping rectangles: max(s1, s2) <= min(e1, e2)
  // both vertical and horizonal ranges should overlap
  return maxI(brick[0].X, brickBelow[0].X) <= minI(brick[1].X, brickBelow[1].X) &&
     maxI(brick[0].Y, brickBelow[0].Y) <= minI(brick[1].Y, brickBelow[1].Y)
}

func (brick Brick) IsOnTop(brickBelow Brick) bool {
  return brick[0].Z == brickBelow[1].Z + 1 && brick.Overlap(brickBelow)
}

func (bricks *BrickWall) LetBricksFall() {
  cmp := func (i, j int) bool {
    return (*bricks)[i][0].Z < (*bricks)[j][0].Z
  }
  sort.SliceStable(*bricks, cmp)


  for zIdx, brick := range *bricks {
    minZ := 1
    for _, brickBelow := range (*bricks)[:zIdx] {
      if brick.Overlap(brickBelow) {
        // it has to stop 1 level above the brick it stops on top
        minZ = maxI(minZ, brickBelow[1].Z + 1)
      }
    }
    brick[1].Z -= brick[0].Z - minZ
    brick[0].Z = minZ
  }

  // sort the bricks again because after falling they might have ended up in a
  // different order
  sort.SliceStable(*bricks, cmp)
}

func part1(
  bricks BrickWall,
) (BrickWall, map[string]BrickWall, map[string]BrickWall) {
  canBeRemoved := 0
  bricks.LetBricksFall()
  // fmt.Println(bricks)

  supports := make(map[string]BrickWall)
  areSupported := make(map[string]BrickWall)

  for _, b := range bricks {
    supports[b.ToKey()] = BrickWall{}
    areSupported[b.ToKey()] = BrickWall{}
  }

  for _, b := range bricks {
    for _, otherB := range bricks {
      if b[0].X == otherB[0].X && b[1].X == otherB[1].X &&
        b[0].Y == otherB[0].Y && b[1].Y == otherB[1].Y &&
        b[0].Z == otherB[0].Z && b[1].Z == otherB[1].Z {
        continue
      }

      if b.IsOnTop(otherB) {
        supports[otherB.ToKey()] = append(supports[otherB.ToKey()], b)
        areSupported[b.ToKey()] = append(areSupported[b.ToKey()], otherB)
      }
    }
  }
  // fmt.Println("are supported:", areSupported)
  // fmt.Println("supports:", supports)

  for _, b := range bricks {
    key := b.ToKey()
    bsSupp, _ := supports[key]

    flag := true
    for _, other := range bsSupp {
      otherKey := other.ToKey()
      otherSupports, _ := areSupported[otherKey]
      if len(otherSupports) < 2 {
        flag = false
        break
      }
    }
    if flag {
      canBeRemoved++
    }
  }

  fmt.Println("There are", canBeRemoved, "bricks that can be removed singularly")

  return bricks, supports, areSupported
}

type QueueWithoutDuplicates struct {
  Queue BrickWall
  Size int
  InQueue map[string]bool
}

func NewQueueWithoutDuplicates() QueueWithoutDuplicates {
  return QueueWithoutDuplicates{
    Queue: BrickWall{},
    Size: 0,
    InQueue: make(map[string]bool),
  }
}
func (q *QueueWithoutDuplicates) Enqueue(el Brick) {
  if _, ok := q.InQueue[el.ToKey()]; ok {
    return
  }
  q.Queue = append(q.Queue, el)
  q.Size++
  q.InQueue[el.ToKey()] = true
}

func (q *QueueWithoutDuplicates) Dequeue() Brick {
  defer func() {
    q.Size--
    q.Queue = q.Queue[1:]
  }()
  delete(q.InQueue, q.Queue[0].ToKey())
  return q.Queue[0]
}

func (q QueueWithoutDuplicates) IsEmpty() bool {
  return q.Size == 0
}

func (b Brick) Eq(other Brick) bool {
  return b[0].X == other[0].X && b[1].X == other[1].X &&
    b[0].Y == other[0].Y && b[1].Y == other[1].Y &&
    b[0].Z == other[0].Z && b[1].Z == other[1].Z
}

func areAllFallen(bricksBelow BrickWall, fallen map[string]bool) bool {
  for _, bb := range bricksBelow {
    if _, isFallen := fallen[bb.ToKey()]; !isFallen {
      return false
    }
  }
  return true
}

func (b Brick) ComputeChain(supports, areSupported map[string]BrickWall) int {
  totalFallingBricks := 0

  fallen := make(map[string]bool)
  q := NewQueueWithoutDuplicates()

  // mark all bricks below the current one as fallen
  for _, support := range areSupported[b.ToKey()] {
    fallen[support.ToKey()] = true
  }

  q.Enqueue(b)
  for !q.IsEmpty() {
    currBrick := q.Dequeue()
    if !areAllFallen(areSupported[currBrick.ToKey()], fallen) {
      continue
    }
    totalFallingBricks++
    fallen[currBrick.ToKey()] = true
    for _, brickAbove := range supports[currBrick.ToKey()] {
      q.Enqueue(brickAbove)
    }
  }

  // the brick we're removing should not be considered in the chain
  return totalFallingBricks - 1
}

func part2(bricks BrickWall, supports, areSupported map[string]BrickWall) {
  chainsTotal := 0
  for _, brick := range bricks {
    chainsTotal += brick.ComputeChain(supports, areSupported)
  }
  fmt.Println("The total number of bricks that fall in chain is:", chainsTotal)
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  bricks := parseInput(f)
 
  bricks, supports, areSupported := part1(bricks)
  part2(bricks, supports, areSupported)
}
