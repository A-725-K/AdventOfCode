package main

import (
  "io"
  "os"
  "fmt"
  "sort"
  "bufio"
  "strings"
  "strconv"
	
	queue "AdventOfCode/ds/queue"
)

const (
  FEW_ROUNDS = 20
  ROUNDS = 10000
  COOL_OFF = 3
)

type Monkey struct {
  items *queue.Queue[int]
  op func (int) int
  test func (int) bool
  toThrow map[bool]int
  monkeyBusiness int
  div int
}

func NewMonkey(itms *queue.Queue[int], opr func(int) int, tst func(int) bool, tt map[bool]int, div int) Monkey {
  return Monkey{itms, opr, tst, tt, 0, div}
}

func readTestLine(line string) int {
  fields := strings.Split(line, " ")
  last := len(fields) - 1
  n, err := strconv.Atoi(fields[last])
  if err != nil {
    panic("Cannot convert test result")
  }
  return n
}

func readItems(line string) *queue.Queue[int] {
  line = strings.TrimSpace(line)
  fields := strings.Split(line, " ")[2:]

  q := queue.NewQueue[int]()
  for _, nStr := range fields {
    off := len(nStr) - 1
    if nStr[off] != ',' {
      off++
    }

    n, err := strconv.Atoi(nStr[:off])
    if err != nil {
      fmt.Println("ERROR:", nStr[:off])
      panic("Cannot convert item")
    }
    q.Enqueue(n)
  }

  return q
}

func readOperationLine(line string) (func(int) int) {
  line = strings.TrimSpace(line)
  fields := strings.Split(line, " ")[4:]
  
  n, err := strconv.Atoi(fields[1])
  if err != nil {
    n = -1
  }

  switch fields[0] {
  case "+":
    if n < 0 {
      return func(x int) int {
        return x + x
      }
    }
    return func(x int) int {
      return x + n
    }
  case "*":
    if n < 0 {
      return func(x int) int {
        return x * x
      }
    }
    return func(x int) int {
      return x * n
    }
  default:
    panic("Operation not known")
  }
}

func parseInput(f *os.File) []Monkey {
  var lines []string
  f.Seek(0, io.SeekStart)
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  n := len(lines)
  var monkeys []Monkey
  for i := 0; i < n; i += 7 {
    var itms *queue.Queue[int]
    tt := make(map[bool]int)

    itms = readItems(lines[i+1])
    opr := readOperationLine(lines[i+2])
    div := readTestLine(lines[i+3])
    tst := func (x int) bool {
      return x%div == 0
    }
    tt[true] = readTestLine(lines[i+4])
    tt[false] = readTestLine(lines[i+5])

    monkeys = append(monkeys, NewMonkey(itms, opr, tst, tt, div))
  }

  return monkeys
}

func printRound(monkeys []Monkey, round int) {
  fmt.Println("After round", round)
  for idx, m := range monkeys {
    fmt.Printf("Monkey (%d): %v\tbusiness: %d\n", idx, m.items.GetQueue(), m.monkeyBusiness)
  }
  fmt.Println()
}

func part1(monkeys []Monkey) {
  for i := 1; i <= FEW_ROUNDS; i++ {
    for idx, m := range monkeys {
      monkeys[idx].monkeyBusiness += m.items.GetSize()
      for !m.items.IsEmpty() {
        item := m.items.Dequeue()
        item = m.op(item)
        item /= COOL_OFF
        key := m.test(item)
        monkeys[m.toThrow[key]].items.Enqueue(item)
      }
    }
    // printRound(monkeys, i)
  }

  var monkeyBusinesses []int
  for _, m := range monkeys {
    monkeyBusinesses = append(monkeyBusinesses, m.monkeyBusiness)
  }
  sort.Ints(monkeyBusinesses)

  n := len(monkeyBusinesses)
  fmt.Println("The monkey business value is:", monkeyBusinesses[n-1] * monkeyBusinesses[n-2])
}

// // IT WORKS BUT TAKES TOO MUCH MEMORY AND TIME :( [using library math/big]
// func part2(monkeys []Monkey) {
//   for i := 1; i <= ROUNDS; i++ {
//     for idx, m := range monkeys {
//       monkeys[idx].monkeyBusiness += m.items.size
//       for !m.items.IsEmpty() {
//         item := m.items.Dequeue()
//         item = m.op(item)
//         key := m.test(item)
//         monkeys[m.toThrow[key]].items.Enqueue(item)
//       }
//     }
//     if i%100 == 0 {
//       fmt.Printf("Round %d...\n", i)
//     }
//     if i == 1 || i == 20 || i%1000 == 0 {
//       fmt.Printf("=== After round %d ===\n", i)
//       for idx, m := range monkeys {
//         fmt.Println("Monkey", idx, "inspected items", m.monkeyBusiness, "times")
//       }
//       fmt.Println()
//       printRound(monkeys, i)
//     }
//   }
//
//   var monkeyBusinesses []int
//   for _, m := range monkeys {
//     monkeyBusinesses = append(monkeyBusinesses, m.monkeyBusiness)
//   }
//   sort.Ints(monkeyBusinesses)
//
//   n := len(monkeyBusinesses)
//   fmt.Println("The monkey business value is:", monkeyBusinesses[n-1] * monkeyBusinesses[n-2])
// }

func part2(monkeys []Monkey) {
  // there is no need to carry a number greater than the LCM of the monkeys'
  // dividends, because the math modulo LCM is the same since we do not
  // divide anymore by COOL_OFF
  lcm := 1
  for _, m := range monkeys {
    lcm *= m.div
  }

  for i := 1; i <= ROUNDS; i++ {
    for idx, m := range monkeys {
      monkeys[idx].monkeyBusiness += m.items.GetSize()
      for !m.items.IsEmpty() {
        item := m.items.Dequeue()
        item = m.op(item)
        key := m.test(item)
        monkeys[m.toThrow[key]].items.Enqueue(item % lcm)
      }
    }
    // if i == 1 || i == 20 || i%1000 == 0 {
    //   fmt.Printf("=== After round %d ===\n", i)
    //   for idx, m := range monkeys {
    //     fmt.Println("Monkey", idx, "inspected items", m.monkeyBusiness, "times")
    //   }
    //   fmt.Println()
    //   printRound(monkeys, i)
    // }
  }
  
  var monkeyBusinesses []int
  for _, m := range monkeys {
    monkeyBusinesses = append(monkeyBusinesses, m.monkeyBusiness)
  }
  sort.Ints(monkeyBusinesses)

  n := len(monkeyBusinesses)
  fmt.Println("The monkey business value is:", monkeyBusinesses[n-1] * monkeyBusinesses[n-2])
}


func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  part1(parseInput(f))
  part2(parseInput(f))
}

