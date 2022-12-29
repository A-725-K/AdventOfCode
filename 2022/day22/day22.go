package main

// WHY 2 DIFFERENT CUBE SHAPES ?!?!?!?!?! ;)
//
// mini_input
//     1
// 2 3 4
//     5 6
//
// input
//   1 2
//   3
// 4 5
// 6

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "AdventOfCode/ds"
)

const (
  NONE = -1
  EMPTY = 8
  WALL = 9

  LEFT = 2
  RIGHT = 0
  DOWN = 1
  UP = 3

  ROW_MULT = 1000
  COL_MULT = 4
  RIGHT_MULT = 0
  DOWN_MULT = 1
  LEFT_MULT = 2
  UP_MULT = 3

  // CUBE_SIZE = 4 // mini_input
  CUBE_SIZE = 50 // input
)

type Map struct {
  maze [][]int
  rows, cols int
}

func (m Map) String() string {
  s := ""
  for i := 0; i < m.rows; i++ {
    for j := 0; j < m.cols; j++ {
      var c string
      switch m.maze[i][j] {
      case NONE:
        c = " "
      case EMPTY:
        c = "."
      case WALL:
        c = "#"
      case RIGHT:
        c = ">"
      case LEFT:
        c = "<"
      case UP:
        c = "^"
      case DOWN:
        c = "V"
      default:
        panic("Cannot render this cell")
      }
      s += c
    }
    s += "\n"
  }
  return s
}

func (m Map) Display() {
  fmt.Println(m)
}

func parseInput(f *os.File) (Map, string) {
  scanner := bufio.NewScanner(f)

  var instructions string
  var m Map

  isMazeFinished := false
  maxCols := 0
  for scanner.Scan() {
    line := scanner.Text()
    if line == "" {
      isMazeFinished = true
      continue
    }

    if !isMazeFinished {
      maxCols = ds.Max(maxCols, len(line))
      var row []int
      for _, c := range line {
        switch c {
        case ' ':
          row = append(row, NONE)
        case '.':
          row = append(row, EMPTY)
        case '#':
          row = append(row, WALL)
        default:
          panic("Character not known")
        }
      }
      m.maze = append(m.maze, row)
    } else {
      instructions = scanner.Text()
    }
  }

  // pad all the rows
  for i := range m.maze {
    for len(m.maze[i]) < maxCols {
      m.maze[i] = append(m.maze[i], NONE)
    }
  }

  m.rows = len(m.maze)
  m.cols = maxCols

  return m, instructions
}

func (m Map) findStartingY() int {
  for i, c := range m.maze[0] {
    if c == EMPTY {
      return i
    }
  }
  panic("Cannot find starting position")
}

type Instruction struct {
  steps, nextDirection int
}

func (i Instruction) String() string {
  var s string
  switch i.nextDirection {
  case LEFT:
    s = "LEFT"
  case RIGHT:
    s = "RIGHT"
  case NONE:
    s = ""
  default:
    panic("Direction not known")
  }
  return fmt.Sprintf("Walk %d steps %s", i.steps, s)
}

func parseNextInstruction(
  instructionList string,
  nextIdx int,
) (Instruction, int) {
  n := len(instructionList)
  if nextIdx > n-1 {
    return Instruction{-1, -1}, -1
  }

  var instr Instruction
  for i := nextIdx; i < n; i++ {
    num, err := strconv.Atoi(instructionList[i:i+1])
    if err != nil {
      switch instructionList[i:i+1] {
      case "R":
        instr.nextDirection = RIGHT
      case "L":
        instr.nextDirection = LEFT
      default:
        panic("Direction not known")
      }
      return instr, i+1
    }
    instr.steps = (instr.steps*10) + num
  }

  instr.nextDirection = NONE
  return instr, n
}

func rotate(direction, nextDirection int) int {
  switch direction {
  case UP:
    switch nextDirection {
    case LEFT:
      return LEFT
    case RIGHT:
      return RIGHT
    case NONE:
      return NONE
    default:
      panic("Next direction not known")
    }

  case DOWN:
    switch nextDirection {
    case LEFT:
      return RIGHT
    case RIGHT:
      return LEFT
    case NONE:
      return NONE
    default:
      panic("Next direction not known")
    }

  case LEFT:
    switch nextDirection {
    case LEFT:
      return DOWN
    case RIGHT:
      return UP
    case NONE:
      return NONE
    default:
      panic("Next direction not known")
    }

  case RIGHT:
    switch nextDirection {
    case LEFT:
      return UP
    case RIGHT:
      return DOWN
    case NONE:
      return NONE
    default:
      panic("Next direction not known")
    }

  default:
    panic("Cannot walk in this direction")
  }
}

func (m Map) nextCoordInRow(row, j, direction int) int {
  var offset int
  switch direction {
  case LEFT:
    offset = -1
  case RIGHT:
    offset = 1
  default:
    panic("Row direction not known")
  }

  prevJ := j
  for {
    j += offset
    
    if j >= m.cols {
      j = 0
    }
    if j < 0 {
      j = m.cols-1
    }

    if m.maze[row][j] == EMPTY {
      return j
    }

    if m.maze[row][j] == WALL {
      return prevJ
    }
  }
}
func (m Map) nextCoordInCol(col, i, direction int) int {
  var offset int
  switch direction {
  case DOWN:
    offset = 1
  case UP:
    offset = -1
  default:
    panic("Col direction not known")
  }

  prevI := i
  for {
    i += offset
    
    if i >= m.rows {
      i = 0
    }
    if i < 0 {
      i = m.rows-1
    }

    if m.maze[i][col] == EMPTY {
      return i
    }

    if m.maze[i][col] == WALL {
      return prevI
    }
  }
}

func (m Map) nextMove(
  direction, x, y int,
  instr Instruction,
) (int, int, int) {
  for i := 0; i < instr.steps; i++ {
    if direction == UP || direction == DOWN {
      y = m.nextCoordInCol(x, y, direction)
    } else if direction == LEFT || direction == RIGHT {
      x = m.nextCoordInRow(y, x, direction)
    }
  }

  return x, y, rotate(direction, instr.nextDirection)
}

func part1(m Map, instructions string) {
  x, y := m.findStartingY(), 0
  prevDirection, direction := NONE, RIGHT

  nextInstr, nextInstrIdx := parseNextInstruction(instructions, 0)
  for nextInstrIdx > 0 {
    prevDirection = direction
    x, y, direction = m.nextMove(direction, x, y, nextInstr)

    // fmt.Println(nextInstr)
    // tmp := m.maze[y][x]
    // m.maze[y][x] = direction
    // m.Display()
    // fmt.Println()
    // m.maze[y][x] = tmp

    nextInstr, nextInstrIdx = parseNextInstruction(instructions, nextInstrIdx)
  }
  fmt.Printf("End in (%d, %d) %d\n", y, x, prevDirection)

  finalPassword := (y+1) * ROW_MULT +
                   (x+1) * COL_MULT +
                   prevDirection
  fmt.Println("The final password in the 2D map is:", finalPassword)
}

// idx -> map of the surface [][]
type Cube map[int][CUBE_SIZE][CUBE_SIZE]int

func NewCube() Cube {
  return make(map[int][CUBE_SIZE][CUBE_SIZE]int)
}

func (c Cube) String() string {
  s := ""
  for faceIdx, face := range c {
    s += fmt.Sprintf("===== Face %d =====\n", faceIdx)
    var f [][]int
    for i := 0; i < CUBE_SIZE; i++ {
      var r []int
      for j := 0; j < CUBE_SIZE; j++ {
        r = append(r, face[i][j])
      }
      f = append(f, r)
    }
    m := Map{maze: f, rows: CUBE_SIZE, cols: CUBE_SIZE}
    s += m.String() + "\n"
  }
  return s
}

func fromMap(m Map) Cube {
  idxFace := 1
  cube := NewCube()

  var positivesMap [][]int
  for _, r := range m.maze {
    var row []int
    for _, c := range r {
      if c > 0 {
        row = append(row, c)
      }
    }
    positivesMap = append(positivesMap, row)
  }

  for i := 0; i < len(positivesMap); i += CUBE_SIZE {
    var face [CUBE_SIZE][CUBE_SIZE]int
    for j := 0; j < len(positivesMap[i]); j++ {
      for k := 0; k < CUBE_SIZE; k++ {
        face[(i%CUBE_SIZE)+k][j%CUBE_SIZE] = positivesMap[i+k][j]
      }
      if (j+1)%CUBE_SIZE == 0 {
        cube[idxFace] = face
        idxFace++

        for ii := 0; ii < CUBE_SIZE; ii++ {
          for jj := 0; jj < CUBE_SIZE; jj++ {
            face[ii][jj] = 0
          } 
        }
      }
    }
  }

  return cube
}

func (c Cube) changeFaceRow(
  prevFace, prevRow, sign int,
) (int, int, int, int) {
  switch prevFace {
  case 1:
    if sign > 0 {
      return prevRow, 0, RIGHT, 2
    }
    return CUBE_SIZE-1-prevRow, 0, RIGHT, 4

  case 2:
    if sign > 0 {
      return CUBE_SIZE-1-prevRow, CUBE_SIZE-1, LEFT, 5
    }
    return prevRow, CUBE_SIZE-1, LEFT, 1

  case 3:
    if sign > 0 {
      return CUBE_SIZE-1, prevRow, UP, 2
    }
    return 0, prevRow, DOWN, 4

  case 4:
    if sign > 0 {
      return prevRow, 0, RIGHT, 5
    }
    return CUBE_SIZE-1-prevRow, 0, RIGHT, 1

  case 5:
    if sign > 0 {
      return CUBE_SIZE-1-prevRow, CUBE_SIZE-1, LEFT, 2
    }
    return prevRow, CUBE_SIZE-1, LEFT, 4

  case 6:
    if sign > 0 {
      return CUBE_SIZE-1, prevRow, UP, 5
    }
    return 0, prevRow, DOWN, 1

  // MINI_INPUT, the big one has a different shape
  // case 1:
  //   if sign > 0 {
  //     return CUBE_SIZE-1-prevRow, CUBE_SIZE-1, LEFT, 6
  //   }
  //   return prevRow, 0, DOWN, 3
  //
  // case 2:
  //   if sign > 0 {
  //     return prevRow, 0, RIGHT, 3
  //   }
  //   return CUBE_SIZE-1, CUBE_SIZE-1-prevRow, UP, 6
  //
  // case 3:
  //   if sign > 0 {
  //     return prevRow, 0, RIGHT, 4
  //   }
  //   return prevRow, CUBE_SIZE-1, LEFT, 2
  //
  // case 4:
  //   if sign > 0 {
  //     return 0, CUBE_SIZE-1-prevRow, DOWN, 6
  //   }
  //   return prevRow, CUBE_SIZE-1, LEFT, 3
  //
  // case 5:
  //   if sign > 0 {
  //     return prevRow, 0, RIGHT, 6
  //   }
  //   return CUBE_SIZE-1, CUBE_SIZE-1-prevRow, UP, 3
  //
  // case 6:
  //   if sign > 0 {
  //     return CUBE_SIZE-1-prevRow, CUBE_SIZE-1, LEFT, 1
  //   }
  //   return prevRow, CUBE_SIZE-1, LEFT, 5

  default:
    panic("Impossible cube face")
  }
}
func (c Cube) nextCoordInRowCube(
  row, j, direction, face int,
) (int, int, int, int) {
  var offset int
  switch direction {
  case LEFT:
    offset = -1
  case RIGHT:
    offset = 1
  default:
    panic("Row direction not known")
  }
  
  prevRow, prevJ, prevDirection, prevFace := row, j, direction, face

  j += offset
  if j < 0 || j >= CUBE_SIZE {
    row, j, direction, face = c.changeFaceRow(prevFace, prevRow, j)
  }

  if c[face][row][j] == WALL {
    return prevRow, prevJ, prevDirection, prevFace
  }
  
  return row, j, direction, face
}
func (c Cube) changeFaceCol(
  prevFace, prevCol, sign int,
) (int, int, int, int) {
  switch prevFace {
  case 1:
    if sign > 0 {
      return 0, prevCol, DOWN, 3
    }
    return prevCol, 0, RIGHT, 6

  case 2:
    if sign > 0 {
      return prevCol, CUBE_SIZE-1, LEFT, 3
    }
    return CUBE_SIZE-1, prevCol, UP, 6

  case 3:
    if sign > 0 {
      return 0, prevCol, DOWN, 5
    }
    return CUBE_SIZE-1, prevCol, UP, 1

  case 4:
    if sign > 0 {
      return 0, prevCol, DOWN, 6
    }
    return prevCol, 0, RIGHT, 3

  case 5:
    if sign > 0 {
      return prevCol, CUBE_SIZE-1, LEFT, 6
    }
    return CUBE_SIZE-1, prevCol, UP, 3

  case 6:
    if sign > 0 {
      return 0, prevCol, DOWN, 2
    }
    return CUBE_SIZE-1, prevCol, UP, 4

  // MINI_INPUT, the big one has a different shape
  // case 1:
  //   if sign > 0 {
  //     return 0, prevCol, DOWN, 4
  //   }
  //   return 0, CUBE_SIZE-1-prevCol, DOWN, 2
  //
  // case 2:
  //   if sign > 0 {
  //     return CUBE_SIZE-1, CUBE_SIZE-1-prevCol, UP, 5
  //   }
  //   return 0, CUBE_SIZE-1-prevCol, DOWN, 1
  //
  // case 3:
  //   if sign > 0 {
  //     return CUBE_SIZE-1-prevCol, 0, RIGHT, 5
  //   }
  //   return prevCol, 0, RIGHT, 1
  //
  // case 4:
  //   if sign > 0 {
  //     return 0, prevCol, DOWN, 5
  //   }
  //   return CUBE_SIZE-1, prevCol, UP, 1
  //
  // case 5:
  //   if sign > 0 {
  //     return CUBE_SIZE-1, CUBE_SIZE-1-prevCol, UP, 2
  //   }
  //   return CUBE_SIZE-1, prevCol, UP, 4
  //
  // case 6:
  //   if sign > 0 {
  //     return CUBE_SIZE-1-prevCol, 0, RIGHT, 2
  //   }
  //   return CUBE_SIZE-1-prevCol, CUBE_SIZE-1, LEFT, 4

  default:
    panic("Impossible cube face")
  }
}
func (c Cube) nextCoordInColCube(
  i, col, direction, face int,
) (int, int, int, int) {
  var offset int
  switch direction {
  case UP:
    offset = -1
  case DOWN:
    offset = 1
  default:
    panic("Col direction not known")
  }
  
  prevI, prevCol, prevDirection, prevFace := i, col, direction, face

  i += offset
  if i < 0 || i >= CUBE_SIZE {
    i, col, direction, face = c.changeFaceCol(prevFace, prevCol, i)
  }

  if c[face][i][col] == WALL {
    return prevI, prevCol, prevDirection, prevFace
  }
  
  return i, col, direction, face
}

func (c Cube) nextMoveCube(
  direction, x, y, face int,
  instr Instruction,
) (int, int, int, int) {
  for i := 0; i < instr.steps; i++ {
    if direction == UP || direction == DOWN {
      y, x, direction, face = c.nextCoordInColCube(y, x, direction, face)
    } else if direction == LEFT || direction == RIGHT {
      y, x, direction, face = c.nextCoordInRowCube(y, x, direction, face)
    }
    // fmt.Printf("Moving in [%d] -> (%d, %d)\n", face, y, x)
  }

  return x, y, rotate(direction, instr.nextDirection), face
}

func part2(m Map, instructions string) {
  cube := fromMap(m)
  // fmt.Println(cube)
  prevDirection, direction := NONE, RIGHT
  x, y, face := 0, 0, 1

  nextInstr, nextInstrIdx := parseNextInstruction(instructions, 0)
  for nextInstrIdx > 0 {
    // fmt.Println(nextInstr)
    prevDirection = direction
    x, y, direction, face = cube.nextMoveCube(direction, x, y, face, nextInstr)
    
    // tmp, cpy := cube[face], cube[face]
    // cpy[y][x] = direction
    // cube[face] = cpy
    // fmt.Println(cube, "\n")
    // cube[face] = tmp

    nextInstr, nextInstrIdx = parseNextInstruction(instructions, nextInstrIdx)
  }

  fmt.Printf("End in [%d] -> (%d, %d) %d\n", face, y, x, prevDirection)

  finalPassword := (y+1+CUBE_SIZE*2) * ROW_MULT +
                   (x+1) * COL_MULT +
                   prevDirection

  fmt.Println("The final password in the 3D map is:", finalPassword)
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  maze, instructions := parseInput(f)
  part1(maze, instructions)
  part2(maze, instructions)
}

