package main

import (
	"bufio"
	"fmt"
	"os"
)

type CoordArray [2]int
type Maze [][]rune
type Set map[string]bool
type Queue struct {
  Queue []CoordArray
  Size int
}

func NewSet() Set {
  return make(map[string]bool)
}

func FromKeys(els []string) Set {
  s := NewSet()
  for _, el := range els {
    s.Add(el)
  }
  return s
}

func (s Set) Contains(el string) bool {
  _, ok := s[el]
  return ok
}

func (s *Set) Add(el string) {
  (*s)[el] = true
}

func (s *Set) Intersect(other Set) {
  for k := range *s {
    if !other.Contains(k) {
      delete(*s, k)
    }
  }
}

func (s *Set) Union(other Set) {
  for k := range other {
    (*s).Add(k)
  }
}

func (m Maze) GetTile(c CoordArray) string {
  return string(m[c[1]][c[0]])
}

func (m Maze) String() string {
  s := ""
  n := len(m)
  for i, r := range m {
    for _, c := range r {
      s += string(c)
    }
    if i < n-1 {
      s += "\n"
    }
  }
  return s
}

func (m Maze) findS() (int, int) {
  for i, row := range m {
    for j, el := range row {
      if string(el) == "S" {
        return i, j
      }
    }
  }
  return -1, -1
}

func parseInput(f *os.File) Maze {
  var maze Maze
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    var row []rune
    for _, c := range scanner.Text() {
      row = append(row, c)
    }
    maze = append(maze, row)
  }
  return maze
}

func (d CoordArray) ToString() string {
  return fmt.Sprintf("%d:%d", d[0], d[1])
}

func NewQueue() Queue {
  return Queue{Queue: []CoordArray{}, Size: 0}
}

func (q *Queue) Enqueue(coord CoordArray) {
  q.Queue = append(q.Queue, coord)
  q.Size++
}

func (q *Queue) Dequeue() CoordArray {
  defer func() {
    q.Size--
    q.Queue = q.Queue[1:]
  }()
  return q.Queue[0]
}

func (q Queue) IsEmpty() bool {
  return q.Size == 0
}

func part1(maze Maze) (Set, string, CoordArray) {
  n, m := len(maze), len(maze[0])

  startY, startX := maze.findS()
  whatIsS := FromKeys([]string{"-", "|", "L", "7", "F", "J"})

  q := NewQueue()
  start := CoordArray{startX, startY}
  q.Enqueue(start)

  loop := NewSet()
  loop.Add(start.ToString())
  for !q.IsEmpty() {
    pos := q.Dequeue()
    currTile := maze.GetTile(pos)
    // fmt.Println("visiting:", pos[0], pos[1])

    // move RIGHT
    if pos[0] < m-1 {
      nextTilePos := CoordArray{pos[0]+1, pos[1]}
      nextTile := maze.GetTile(nextTilePos)
      _, isInLoop := loop[nextTilePos.ToString()]
      if !isInLoop && (
        currTile == "S" ||
        currTile == "-" ||
        currTile == "F" ||
        currTile == "L") &&
      (
        // nextTile == "S" ||
        nextTile == "-" ||
        nextTile == "J" ||
        nextTile == "7") {
        loop.Add(nextTilePos.ToString())
        q.Enqueue(nextTilePos)
        if currTile == "S" {
          possibleS := FromKeys([]string{"-", "L", "F"})
          whatIsS.Intersect(possibleS)
        }
      }
    }

    // move LEFT
    if pos[0] > 0 {
      nextTilePos := CoordArray{pos[0]-1, pos[1]}
      nextTile := maze.GetTile(nextTilePos)
      _, isInLoop := loop[nextTilePos.ToString()]
      if !isInLoop && (
        currTile == "S" ||
        currTile == "-" ||
        currTile == "J" ||
        currTile == "7") &&
      (
        // nextTile == "S" ||
        nextTile == "-" ||
        nextTile == "L" ||
        nextTile == "F") {
        loop.Add(nextTilePos.ToString())
        q.Enqueue(nextTilePos)
        if currTile == "S" {
          possibleS := FromKeys([]string{"-", "7", "J"})
          whatIsS.Intersect(possibleS)
        }
      }
    }

    // move DOWN
    if pos[1] < n-1 {
      nextTilePos := CoordArray{pos[0], pos[1]+1}
      nextTile := maze.GetTile(nextTilePos)
      _, isInLoop := loop[nextTilePos.ToString()]
      if !isInLoop && (
        currTile == "S" ||
        currTile == "|" ||
        currTile == "7" ||
        currTile == "F") &&
      (
        // nextTile == "S" ||
        nextTile == "|" ||
        nextTile == "L" ||
        nextTile == "J") {
        loop.Add(nextTilePos.ToString())
        q.Enqueue(nextTilePos)
        if currTile == "S" {
          possibleS := FromKeys([]string{"|", "7", "F"})
          whatIsS.Intersect(possibleS)
        }
      }
    }

    // move UP
    if pos[1] > 0 {
      nextTilePos := CoordArray{pos[0], pos[1]-1}
      nextTile := maze.GetTile(nextTilePos)
      _, isInLoop := loop[nextTilePos.ToString()]
      if !isInLoop && (
        currTile == "S" ||
        currTile == "|" ||
        currTile == "J" ||
        currTile == "L") &&
      (
        // nextTile == "S" ||
        nextTile == "|" ||
        nextTile == "F" ||
        nextTile == "7") {
        loop.Add(nextTilePos.ToString())
        q.Enqueue(nextTilePos)
        if currTile == "S" {
          possibleS := FromKeys([]string{"|", "J", "L"})
          whatIsS.Intersect(possibleS)
        }
      }
    }
  }

  fmt.Println("The farthest tile is", (len(loop)+1) / 2, "steps away")

  // extract real value of S tile from the map
  valueOfS := ""
  // only one value in the map, but Go does not index it :(
  for k := range whatIsS {
    valueOfS = k
  }
  return loop, valueOfS, start
}

// Thanks to Hyper-Neutrino! :)
func part2(maze Maze, loop Set) {
  outside := NewSet()
  for y, row := range maze {
    isInsideLoop := false
    isGoingUp := false
    for x, tile := range row {
      if tile == '|' {
        isInsideLoop = !isInsideLoop
      } else if tile == 'L' || tile == 'F' {
        isGoingUp = tile == 'L'
      } else if tile == '7' || tile == 'J' {
        if (isGoingUp && tile != 'J') || (!isGoingUp && tile != '7') {
          isInsideLoop = !isInsideLoop
        }
        isGoingUp = false
      }
      if !isInsideLoop {
        outside.Add(CoordArray{x, y}.ToString())
      }
    }
  }

  loop.Union(outside)
  inLoop := len(maze) * len(maze[0]) - len(loop)

  fmt.Println("There are", inLoop, "tiles in the loop")
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  maze := parseInput(f)
 
  loop, valueOfS, s := part1(maze)
  // replace S
  maze[s[1]][s[0]] = rune(valueOfS[0])
  // replace garbage pipes
  for y, row := range maze {
    for x := range row {
      // if _, ok := loop[Coord{x, y}.ToString()]; !ok {
      if !loop.Contains(CoordArray{x, y}.ToString()) {
        maze[y][x] = '.'
      }
    }
  }
  part2(maze, loop)
}
