package main

import (
  "io"
  "os"
  "fmt"
  "bufio"
  "strconv"
)

func part1 (f *os.File) {
  scanner := bufio.NewScanner(f)

  maxSoFar := 0
  currCalories := 0
  for scanner.Scan() {
    line := scanner.Text()
    if line == "" {
      if currCalories > maxSoFar {
        maxSoFar = currCalories
      }
      currCalories = 0
      continue
    }
    calories, err := strconv.Atoi(line)
    if err != nil {
      panic(err)
    }
    currCalories += calories
  }

  fmt.Println("The elf with most calories carries:", maxSoFar, "calories")
}

func part2 (f *os.File) {
  scanner := bufio.NewScanner(f)

  fst, snd, trd := -1, -1, -1
  currCalories := 0
  for scanner.Scan() {
    line := scanner.Text()
    if line == "" {
      if currCalories > fst {
        trd, snd, fst = snd, fst, currCalories
      } else if currCalories > snd {
        trd, snd = snd, currCalories
      } else if currCalories > trd {
        trd = currCalories
      }
      currCalories = 0
      continue
    }
    calories, err := strconv.Atoi(line)
    if err != nil {
      panic(err)
    }
    currCalories += calories
  }

  fmt.Println("\nThe first elf with most calories carries:", fst, "calories")
  fmt.Println("The second elf with most calories carries:", snd, "calories")
  fmt.Println("The third elf with most calories carries:", trd, "calories")
  fmt.Println("In total, they are carrying:", fst + snd + trd, "calories")
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()
  
  part1(f)
  f.Seek(0, io.SeekStart)
  part2(f)
}

