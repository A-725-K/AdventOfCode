package main

import (
  "io"
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
  r "AdventOfCode/types/range"
)

func NewRange(s string) r.Range {
  fields := strings.Split(s, "-")
  start, err := strconv.Atoi(fields[0])
  if err != nil {
    panic("Range not valid")
  }
  end, err := strconv.Atoi(fields[1])
  if err != nil {
    panic("Range not valid")
  }
  return r.Range{Start: start, End: end}
}

func part2(f *os.File) {
  scanner := bufio.NewScanner(f)

  numbrOfTotalOverlaps := 0
  for scanner.Scan() {
    line := scanner.Text()
    ranges := strings.Split(line, ",")
    fstRng, sndRng := NewRange(ranges[0]), NewRange(ranges[1])
    if fstRng.PartialOverlap(sndRng) || sndRng.PartialOverlap(fstRng) {
      numbrOfTotalOverlaps++
    }
  }
  fmt.Println("In total there are", numbrOfTotalOverlaps, "overlapping areas")
}

func part1(f *os.File) {
  scanner := bufio.NewScanner(f)

  overlappingRngs := 0
  for scanner.Scan() {
    line := scanner.Text()
    ranges := strings.Split(line, ",")
    fstRng, sndRng := NewRange(ranges[0]), NewRange(ranges[1])
    if fstRng.Overlap(sndRng) || sndRng.Overlap(fstRng) {
      overlappingRngs += 1
    }
  }
  fmt.Println("There are", overlappingRngs, "overlapping areas")
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

