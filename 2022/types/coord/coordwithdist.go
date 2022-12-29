package types

import (
  "fmt"
  "strings"
  "strconv"
  "AdventOfCode/ds"
)

type CoordWithDist struct {
  X, Y, Distance int
}

func (p1 *CoordWithDist) ManhattanDist(p2 CoordWithDist) int {
  return ds.Abs(p1.X - p2.X) + ds.Abs(p1.Y - p2.Y)
}

func (c CoordWithDist) ToKeyDist() string {
  return fmt.Sprintf("%d+%d+%d", c.X, c.Y, c.Distance) 
}
func FromKeyDist(k string) CoordWithDist {
  fields := strings.Split(k, "+")

  x, err := strconv.Atoi(fields[0])
  if err != nil {
    panic("Cannot convert X")
  }
  y, err := strconv.Atoi(fields[1])
  if err != nil {
    panic("Cannot convert Y")
  }
  dist, err := strconv.Atoi(fields[2])
  if err != nil {
    panic("Cannot convert Distance")
  }

  return CoordWithDist{X: x, Y: y, Distance: dist}
}
