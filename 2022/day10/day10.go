package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
)

var INTERESTING_CYCLES = []int{20, 60, 100, 140, 180, 220}

const LINE_LEN = 40

func isInterestingCycle(c int) bool {
  return (c-20) % LINE_LEN == 0
}

type CPU struct {
  cycles int
  X int
}

func NewCPU() *CPU {
  return &CPU{1, 1}
}

func (cpu *CPU) SignalStrength() int {
  return cpu.X * cpu.cycles
}

var sss string = ""
func (cpu *CPU) drawCRT() {
  cpuCyclesNorm := (cpu.cycles - 1) % LINE_LEN

  var c string
  if cpuCyclesNorm >= cpu.X - 1 && cpuCyclesNorm <= cpu.X + 1 {
    c = "###"
  } else {
    c = "..."
  }
  if (cpu.cycles % LINE_LEN) == 0 {
    c += "\n"
  }

  fmt.Print(c)
}

func (cpu *CPU) noop() *int {
  var retval *int = nil

  cpu.drawCRT()
  cpu.cycles++
  if isInterestingCycle(cpu.cycles) {
    retval = new(int)
    *retval = cpu.SignalStrength()
  }
  return retval
}

func (cpu *CPU) addX(n int) *int {
  var retval *int = nil 

  cpu.drawCRT()
  cpu.cycles++
  if isInterestingCycle(cpu.cycles) {
    retval = new(int)
    *retval = cpu.SignalStrength()
  }
  cpu.drawCRT()
  cpu.cycles++
  cpu.X += n
  if isInterestingCycle(cpu.cycles) {
    retval = new(int)
    *retval = cpu.SignalStrength()
  }
  return retval
}

func part1_2(lines []string) {
  cpu := NewCPU()

  signals := []int{}
  var signal *int

  for _, line := range lines {
    cmds := strings.Split(line, " ")
    switch cmds[0] {
    case "addx":
      n, err := strconv.Atoi(cmds[1])
      if err != nil {
        panic("Cannot convert addx operand")
      }
      signal = cpu.addX(n)
    case "noop":
      signal = cpu.noop()
    default:
      panic("Instruction not known")
    }

    if signal != nil {
      signals = append(signals, *signal)
    }
  }

  sumOfSignalsStrenghts := 0
  for _, ss := range signals {
    sumOfSignalsStrenghts += ss
  }

  fmt.Println("The sum of signal strengths is:", sumOfSignalsStrenghts)
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  var lines []string

  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  part1_2(lines)
}

