package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(f *os.File) []string {
  scanner := bufio.NewScanner(f)
  scanner.Scan()
  return strings.Split(scanner.Text(), ",")
}

type Lens struct {
  Label string
  FocalValue int
}

func (l Lens) GetPower(boxIdx, slot int) int{
  return (boxIdx+1) * (slot+1) * l.FocalValue
}

func hash(s string) int {
  value := 0
  for _, c := range s {
    value += int(c)
    value *= 17
    value %= 256
  }
  return value
}

func part1(instructions []string) {
  initializationResult := 0
  for _, ins := range instructions {
    initializationResult += hash(ins)
  }
  fmt.Println("The sum of all instructions is:", initializationResult)
}

func part2(instructions []string) {
  var boxes [256][]Lens
  for _, ins := range instructions {
    var opIdx int
    for i, ch := range ins {
      if ch == '-' || ch == '=' {
        opIdx = i
        break
      }
    }
    label := ins[:opIdx]
    boxIdx := hash(label)
    if ins[opIdx] == '-' {
      splitIdx := -1
      for i, lens := range boxes[boxIdx] {
        if lens.Label == label {
          splitIdx = i
          break
        }
      }
      if splitIdx >= 0 {
        endIdx := len(boxes[boxIdx])
        var toAppend []Lens
        if splitIdx+1 < endIdx {
          toAppend = boxes[boxIdx][splitIdx+1:]
        }
        boxes[boxIdx] = append(
          boxes[boxIdx][:splitIdx],
          toAppend...,
        )
      }
    } else if ins[opIdx] == '=' {
      focalValue, _ := strconv.Atoi(ins[opIdx+1:])
      newLens := Lens{Label: label, FocalValue: focalValue}
      alreadyPresent := false
      for i, lens := range boxes[boxIdx] {
        if lens.Label == newLens.Label {
          boxes[boxIdx][i].FocalValue = newLens.FocalValue
          alreadyPresent = true
          break
        }
      }
      if !alreadyPresent {
        boxes[boxIdx] = append(boxes[boxIdx], newLens)
      }
    } else {
      panic("Invalid op")
    }
  }

  focusingPower := 0
  for boxIdx, box := range boxes {
    for slot, lens := range box {
      focusingPower += lens.GetPower(boxIdx, slot)
    }
  }
  fmt.Println("The focusing power of the whole configuration is:", focusingPower)
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  instructions := parseInput(f)
 
  part1(instructions)
  part2(instructions)
}
