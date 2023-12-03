package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PartNumber struct {
  Num, XStart, XEnd, Y int
}

type Symbol struct {
  Glyph string
  X, Y int
}

type Grid struct {
  PartNums []PartNumber
  Symbols []Symbol
}

func parseInput(f *os.File) (Grid) {
  scanner := bufio.NewScanner(f)
  i := 0
  var symbols []Symbol
  var partNums []PartNumber
  for scanner.Scan() {
    xStart, xEnd, y := -1, -1, -1
    s := ""
    j := 0
    for _, c := range scanner.Text() {
      if c >= '0' && c <= '9' {
        if len(s) == 0 {
          y = i
          xStart = j
        }
        s += string(c)
      } else {
        if len(s) > 0 {
          y = i
          xEnd = j-1
          num, _ := strconv.Atoi(s)
          newPn := PartNumber{Num: num, XStart: xStart, XEnd: xEnd, Y: y}
          partNums = append(partNums, newPn)
          s = ""
          xStart, xEnd, y = -1, -1, -1
        }
        if c != '.' {
          symbols = append(symbols, Symbol{Glyph: string(c), X: j, Y: i})
        }
      }
      j++
    }
    if len(s) > 0 {
      y = i
      xEnd = j-1
      num, _ := strconv.Atoi(s)
      newPn := PartNumber{Num: num, XStart: xStart, XEnd: xEnd, Y: y}
      partNums = append(partNums, newPn)
    }

    i++
  }
return Grid{PartNums: partNums, Symbols: symbols}
}

func isPart(pn PartNumber, symbols []Symbol) (bool, bool) {
  xLeft, yLeft, xRight, yRight := pn.XStart-1, pn.Y-1, pn.XEnd+1, pn.Y+1
  for _, s := range symbols {
    if s.X >= xLeft && s.X <= xRight && s.Y >= yLeft && s.Y <= yRight {
      isStar := false
      if s.Glyph == "*" {
        isStar = true
      }
      return true, isStar
    }
  }
  return false, false
}

func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func hasNumbersClose(s Symbol, pns []PartNumber) (bool, []int) {
  closeN := 0
  var nums []int
  for _, pn := range pns {
    // fmt.Println("pn =", pn.Num, fmt.Sprintf("(%d,%d)", pn.XStart, pn.Y), fmt.Sprintf("(%d,%d)", pn.XEnd, pn.Y))
    if (abs(pn.XEnd - s.X) < 2 || abs(pn.XStart - s.X) < 2) && abs(pn.Y-s.Y) < 2 {
      nums = append(nums, pn.Num)
      closeN++
    }
  }
  // fmt.Println(fmt.Sprintf("s = (%d,%d)", s.X, s.Y), closeN)
  if closeN > 0 {
    return true, nums
  }
  return false, []int{}
}

func part1(grid Grid) ([]PartNumber) {
  var onlyParts []PartNumber
  sumOfParts := 0
  for _, n := range grid.PartNums {
    if ok, isStar := isPart(n, grid.Symbols); ok {
      sumOfParts += n.Num
      if isStar {
        onlyParts = append(onlyParts, n)
      }
    }
  }

  fmt.Println("The sum of part numbers is:", sumOfParts)
  return onlyParts
}

func part2(grid Grid) {
  sumOfGearRatio := 0

  for _, s := range grid.Symbols {
    if ok, nums := hasNumbersClose(s, grid.PartNums); ok {
      // fmt.Println(nums)
      if (len(nums) > 1) {
        for i := 0; i < len(nums)-1; i++ {
          for j := i+1; j< len(nums); j++ {
            sumOfGearRatio += nums[i] * nums[j]
          }
        }
      }
    }
  }

  fmt.Println("The sum of gear ratio is:", sumOfGearRatio)
}

func preparePart2() {

}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  grid := parseInput(f)

  onlyParts := part1(grid)
  var onlyStars []Symbol
  for _, s := range grid.Symbols {
    if s.Glyph == "*" {
      onlyStars = append(onlyStars, s)
    }
  }
  grid.Symbols = onlyStars
  grid.PartNums = onlyParts
  part2(grid)
}
