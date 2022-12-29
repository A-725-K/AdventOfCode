package types

import (
  "strings"
  "strconv"
  "AdventOfCode/ds"
)

type Coord3D struct {
  X, Y, Z int
}

func (c Coord3D) String() string {
  return "(" + strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y) + "," + strconv.Itoa(c.Z) + ")"
}

func (c1 Coord3D) Equals(c2 Coord3D) bool {
  return c1.X == c2.X && c1.Y == c2.Y && c1.Z == c2.Z
}

func (c1 *Coord3D) IsAdjacent(c2 Coord3D) bool {
  return (c1.X == c2.X && c1.Y == c2.Y && ds.Abs(c1.Z - c2.Z) <= 1) ||
         (c1.X == c2.X && c1.Z == c2.Z && ds.Abs(c1.Y - c2.Y) <= 1) ||
         (c1.Y == c2.Y && c1.Z == c2.Z && ds.Abs(c1.X - c2.X) <= 1)
}

func (c1 Coord3D) IsInside(minC, maxC int) bool {
  return minC <= c1.X && c1.X <= maxC &&
         minC <= c1.Y && c1.Y <= maxC &&
         minC <= c1.Z && c1.Z <= maxC
}

func (c1 Coord3D) Move(c2 Coord3D) Coord3D {
  return Coord3D{X: c1.X+c2.X, Y: c1.Y+c2.Y, Z: c1.Z+c2.Z}
}

func (c Coord3D) ToKey() string {
  return strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y) + "," + strconv.Itoa(c.Z)
}

func FromKey3D(s string) Coord3D {
  coords := strings.Split(s, ",")
  x, err := strconv.Atoi(coords[0])
  if err != nil {
    panic("Cannot convert X")
  }
  y, err := strconv.Atoi(coords[1])
  if err != nil {
    panic("Cannot convert Y")
  }
  z, err := strconv.Atoi(coords[2])
  if err != nil {
    panic("Cannot convert Z")
  }
  return Coord3D{X:x, Y:y, Z:z}
}

