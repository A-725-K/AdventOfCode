package main

import (
  "io"
  "os"
  "fmt"
  "bufio"
)

func removeDuplicates(s string) string {
  letters := make(map[rune]bool)
  for _, c := range s {
    letters[c] = true
  }
  sNoDups := ""
  for c := range letters {
    sNoDups += string(c)
  }
  return sNoDups
}

func intersectMultiple(group [3]string) rune {
  n := len(group)
  intersection := make(map[rune]int)
  for _, s := range group {
    sNoDups := removeDuplicates(s)
    for _, c := range sNoDups {
      if _, ok := intersection[c]; ok {
        intersection[c]++
      } else {
        intersection[c] = 1
      }
    }
  }
  for c, count := range intersection {
    if count == n {
      return c
    }
  }
  panic("No intersection")
}

func intersect(s1, s2 string) rune {
    intersection := make(map[rune]bool)
    for _, c := range s1 {
      intersection[c] = true
    }
    for _, c := range s2 {
      if _, ok := intersection[c]; ok {
        return c
      }
    }
    panic("No intersection")
}

func computeValue(c rune) int {
  if c >= 'a' && c <= 'z' {
    return int(c - 'a' + 1)
  }
  if c >= 'A' && c <= 'Z' {
    return int(c - 'A' + 1) + 26
  }
  panic("Character not valid")
}

func part2(f *os.File) {
  scanner := bufio.NewScanner(f)
  
  i, totalValue := 0, 0
  var group [3]string
  for scanner.Scan() {
    line := scanner.Text()
    group[i%3] = line
    i++
    if i%3 == 0 {
      commonElem := intersectMultiple(group)
      totalValue += computeValue(commonElem)
    }
  }
  fmt.Println("The total value of the duplicate elements is:", totalValue)
}

func part1(f *os.File) {
  scanner := bufio.NewScanner(f)
  
  totalValue := 0
  for scanner.Scan() {
    line := scanner.Text()
    halfLen := len(line) / 2
    fstComp, sndComp := line[:halfLen], line[halfLen:]
    commonElem := intersect(fstComp, sndComp)
    totalValue += computeValue(commonElem)
  }
  fmt.Println("The total value of the duplicate elements is:", totalValue)
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

