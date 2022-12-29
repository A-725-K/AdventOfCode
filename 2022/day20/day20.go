package main

import (
  "io"
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
  "AdventOfCode/ds"
  set "AdventOfCode/ds/set"
  pair "AdventOfCode/types/pair"
)

const (
  ROUNDS = 10
  SECRET_KEY int64 = 811589153
  GROOVE_COORD_X = 1000
  GROOVE_COORD_Y = 2000
  GROOVE_COORD_Z = 3000
)

func parseInput(f *os.File) []pair.Pair[int, int] {
  scanner := bufio.NewScanner(f)

  i := 0
  var nums []pair.Pair[int, int]
  for scanner.Scan() {
    n, err := strconv.Atoi(scanner.Text())
    if err != nil {
      panic("Cannot convert a number")
    }
    nums = append(nums, pair.NewPair[int, int](n, i))
    i++
  }

  return nums
}

func ToKey(p pair.Pair[int, int]) string {
  return fmt.Sprintf("%d:%d", p.Fst(), p.Snd())
}
func FromKey(s string) pair.Pair[int, int] {
  fields := strings.Split(s, ":")
  n, err := strconv.Atoi(fields[0])
  if err != nil {
    panic("Cannot convert n")
  }
  idx, err := strconv.Atoi(fields[1])
  if err != nil {
    panic("Cannot convert idx")
  }
  return pair.NewPair[int, int](n, idx)
}

func findNext(nums []pair.Pair[int, int], alreadyMoved set.Set[string]) int {
  min, minIdx := 999999, -1
  for idx, n := range nums {
    if !alreadyMoved.Contains(ToKey(n)) {
      if n.Snd() < min {
        min = n.Snd()
        minIdx = idx
      }
    }
  }
  if minIdx == -1 {
    return -1
  }
  return minIdx
}

func moveRight(nums *[]pair.Pair[int, int], startIdx int, steps int64, mod int) {
  endIdx := int(ds.Mod(int64(startIdx) + steps, int64(mod-1)))

  // fmt.Println("R: startIdx =", startIdx, "\tendIdx =", endIdx, "\tsteps =", steps)

  if endIdx > startIdx {
    for i := startIdx; i < endIdx; i++ {
      (*nums)[i], (*nums)[i+1] = (*nums)[i+1], (*nums)[i]
    }
  } else if endIdx < startIdx {
    moveLeft(nums, startIdx, int64(startIdx-endIdx), mod)
  }
}
func moveLeft(nums *[]pair.Pair[int, int], startIdx int, steps int64, mod int) {
  endIdx := int(ds.Mod(int64(startIdx) - steps, int64(mod-1)))

  // fmt.Println("L: startIdx =", startIdx, "\tendIdx =", endIdx, "\tsteps =", steps)

  if endIdx > startIdx {
    moveRight(nums, startIdx, int64(endIdx-startIdx), mod)
  } else if endIdx < startIdx {
    for i := startIdx; i > endIdx; i-- {
      (*nums)[i], (*nums)[i-1] = (*nums)[i-1], (*nums)[i]
    }
  }
}

func printNums(nums []pair.Pair[int, int], scale int64) {
  n := len(nums)
  fmt.Print("[")
  for i, c := range nums {
    s := " "
    if i == n-1 {
      s = ""
    }
    fmt.Print(int64(c.Fst()) * scale, s)
  }
  fmt.Println("]")
}

func part1(nums []pair.Pair[int, int]) {
  alreadyMoved := set.NewSet[string]()

  n := len(nums)
  nextIdx := findNext(nums, alreadyMoved);
  for nextIdx >= 0 {
    currElem := nums[nextIdx]
    sign := ds.Sign(currElem.Fst())
    steps := int64(ds.Abs(currElem.Fst()))
    
    // fmt.Printf("===== MOVING [%d] =====\n", currElem.Fst())
    if sign > 0 {
      moveRight(&nums, nextIdx, steps, n)
    } else if sign < 0 {
      moveLeft(&nums, nextIdx, steps, n)
    }
    // printNums(nums, int64(1))
    alreadyMoved.Add(ToKey(currElem))
    nextIdx = findNext(nums, alreadyMoved)
  }

  zeroIdx := 0
  for i, p := range nums {
    if p.Fst() == 0 {
      zeroIdx = i
      break
    }
  }

  groveCoordSum := nums[(zeroIdx + GROOVE_COORD_X)%n].Fst() +
                   nums[(zeroIdx + GROOVE_COORD_Y)%n].Fst() +
                   nums[(zeroIdx + GROOVE_COORD_Z)%n].Fst()

  fmt.Println("The grove coordinates sum is:", groveCoordSum)
}

func part2(nums []pair.Pair[int, int]) {
  n := len(nums)

  for rnd := 0; rnd < ROUNDS; rnd++ {
    // fmt.Println("+++++ ROUND", rnd+1, "+++++")
    alreadyMoved := set.NewSet[string]()

    nextIdx := findNext(nums, alreadyMoved);
    for nextIdx >= 0 {
      currElem := nums[nextIdx]
      sign := ds.Sign(currElem.Fst())
      steps := ds.Abs(int64(currElem.Fst()) * SECRET_KEY)
     
      // fmt.Printf("===== MOVING [%d] =====\n", int64(currElem.Fst()) * SECRET_KEY)
      if sign > 0 {
        moveRight(&nums, nextIdx, steps, n)
      } else if sign < 0 {
        moveLeft(&nums, nextIdx, steps, n)
      }
      // printNums(nums, SECRET_KEY)
      alreadyMoved.Add(ToKey(currElem))
      nextIdx = findNext(nums, alreadyMoved)
    }
    // printNums(nums)
  }

  zeroIdx := 0
  for i, p := range nums {
    if p.Fst() == 0 {
      zeroIdx = i
      break
    }
  }

  groveCoordSum := int64(nums[(zeroIdx + GROOVE_COORD_X)%n].Fst() +
                    nums[(zeroIdx + GROOVE_COORD_Y)%n].Fst() +
                    nums[(zeroIdx + GROOVE_COORD_Z)%n].Fst()) *
                    SECRET_KEY

  fmt.Println("The grove coordinates sum is:", groveCoordSum)
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  part1(parseInput(f))
  f.Seek(0, io.SeekStart)
  part2(parseInput(f))
}
