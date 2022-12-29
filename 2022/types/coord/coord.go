package types

import (
  "strings"
  "strconv"
  "AdventOfCode/ds"
)

type Coord struct {
  X, Y int
}

type Coord64 struct {
  X, Y int64
}

func (p1 *Coord) MoveCloser(p2 Coord) (int, int) {
  dx, dy := 0, 0
  deltaX := p2.X - p1.X
  deltaY := p2.Y - p1.Y

  if ds.Abs(deltaX) == 2 && ds.Abs(deltaY) == 2 {
    dx = deltaX / 2
    dy = deltaY / 2
  } else if ds.Abs(deltaX) == 2 {
    dx = deltaX / 2
    dy = deltaY
  } else if ds.Abs(deltaY) == 2 {
    dx = deltaX
    dy = deltaY / 2
  }

  return dx, dy
}

func (c Coord) ToKey() string {
  return strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y)
}

func FromKey(s string) Coord {
  coords := strings.Split(s, ",")
  x, err := strconv.Atoi(coords[0])
  if err != nil {
    panic("Cannot convert X")
  }
  y, err := strconv.Atoi(coords[1])
  if err != nil {
    panic("Cannot convert Y")
  }
  return Coord{X:x, Y:y}
}

func MinMax(cc []Coord) (int, int, int, int) {
  minX, minY, maxX, maxY := cc[0].X, cc[0].Y, cc[0].X, cc[0].Y
  for _, c := range cc {
    minX = ds.Min(minX, c.X)
    maxX = ds.Max(maxX, c.X)
    minY = ds.Min(minY, c.Y)
    maxY = ds.Max(maxY, c.Y)
  }
  return minX, minY, maxX, maxY
}
func MinMax64(cc []Coord64) (int64, int64, int64, int64) {
  minX, minY, maxX, maxY := cc[0].X, cc[0].Y, cc[0].X, cc[0].Y
  for _, c := range cc {
    minX = ds.Min(minX, c.X)
    maxX = ds.Max(maxX, c.X)
    minY = ds.Min(minY, c.Y)
    maxY = ds.Max(maxY, c.Y)
  }
  return minX, minY, maxX, maxY
}
