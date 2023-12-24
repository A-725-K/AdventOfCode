package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type HailStone struct {
  X, Y, Z float64
  Vx, Vy, Vz float64
}

func parseInput(f *os.File) []HailStone {
  var hailstones []HailStone
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    fields := strings.Split(
      strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), "@", ","), " ", ""),
      ",",
    )
    x, _ := strconv.ParseFloat(fields[0], 64)
    y, _ := strconv.ParseFloat(fields[1], 64)
    z, _ := strconv.ParseFloat(fields[2], 64)
    vx, _ := strconv.ParseFloat(fields[3], 64)
    vy, _ := strconv.ParseFloat(fields[4], 64)
    vz, _ := strconv.ParseFloat(fields[5], 64)

    hailstones = append(hailstones, HailStone{
      X: x, Y: y, Z: z,
      Vx: vx, Vy: vy, Vz:vz,
    })
  }
  return hailstones
}

func (h0_0 HailStone) ComputeIntersection(h1_0 HailStone, i, j float64) bool {
  // compute the next point in the trajectory of the first stone
  h0_1 := HailStone{
    X: h0_0.X + h0_0.Vx,
    Y: h0_0.Y + h0_0.Vy,
    Z: h0_0.Z + h0_0.Vz,
    Vx: h0_0.Vx, Vy: h0_0.Vy, Vz: h0_0.Vz,
  }
  A0 := h0_1.Y - h0_0.Y
  B0 := h0_0.X - h0_1.X
  C0 := A0*h0_0.X + B0*h0_0.Y

  // compute the next point in the trajectory of the second stone
  h1_1 := HailStone{
    X: h1_0.X + h1_0.Vx,
    Y: h1_0.Y + h1_0.Vy,
    Z: h1_0.Z + h1_0.Vz,
    Vx: h1_0.Vx, Vy: h1_0.Vy, Vz: h1_0.Vz,
  }
  A1 := h1_1.Y - h1_0.Y
  B1 := h1_0.X - h1_1.X
  C1 := A1*h1_0.X + B1*h1_0.Y

  det := A0*B1 - A1*B0
  // det == 0 ---> the trajectories are parallel
  if det == 0 {
    return false
  }

  // compute the intersection
  xNew := (B1*C0 - B0*C1) / det
  yNew := (A0*C1 - A1*C0) / det

  // check if the intersection is in range and if will happen in the future
  if i <= xNew && xNew <= j && i <= yNew && yNew <= j &&
    ((xNew >= h0_0.X && h0_0.Vx >= 0) || (xNew < h0_0.X && h0_0.Vx < 0)) &&
    ((yNew >= h0_0.Y && h0_0.Vy >= 0) || (yNew < h0_0.Y && h0_0.Vy < 0)) &&
    ((xNew >= h1_0.X && h1_0.Vx >= 0) || (xNew < h1_0.X && h1_0.Vx < 0)) &&
    ((yNew >= h1_0.Y && h1_0.Vy >= 0) || (yNew < h1_0.Y && h1_0.Vy < 0)) &&
    det > 0 { // det > 0 ===> in the future
    return true
  }

  return false
}

func part1(hailstones []HailStone) {
  countIntersectionInArea := 0

  for i, h1 := range hailstones {
    for j, h2 := range hailstones {
      if i == j {
        continue
      }

      if h1.ComputeIntersection(h2, 200000000000000, 400000000000000) {
        countIntersectionInArea++
      }
    }
  }

  fmt.Println(
    "In the designed area there will be",
    countIntersectionInArea, "intersections",
  )
}

func part2(filename string) {
  // to run this command it is needed to have a virtual environemnt with Sympy
  // library installed and accessible
  cmd, err := exec.Command("python3", "day24.py", filename).Output()
  if err != nil {
    panic(err)
  }
  fmt.Println(string(cmd))
}

func main() {
  filename := os.Args[1]
  f, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  hailstones := parseInput(f)

  part1(hailstones)
  part2(filename)
}
