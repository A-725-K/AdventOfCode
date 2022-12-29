package main

import (
  "io"
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
  queue "AdventOfCode/ds/queue"

  // lib to eval expression
  "gopkg.in/Knetic/govaluate.v3"
)

const (
  ROOT = "root"
  HUMAN = "humn"
)

type Op struct {
  key, left, right, op string
}

func parseInput(f *os.File) (*queue.Queue[Op], map[string]int64) {
  scanner := bufio.NewScanner(f)

  ops := queue.NewQueue[Op]()
  res := make(map[string]int64)
  for scanner.Scan() {
    fields := strings.Split(scanner.Text(), " ")
    varName := strings.TrimSuffix(fields[0], ":")

    if len(fields) == 2 {
      value, err := strconv.ParseInt(fields[1], 10, 64)
      if err != nil {
        panic("Cannot convert value")
      }
      res[varName] = value
    } else {
      ops.Enqueue(Op{
        key: varName,
        left: fields[1],
        op: fields[2],
        right: fields[3],
      })
    }
  }

  return ops, res
}

func part1(operations *queue.Queue[Op], results map[string]int64) int64 {
  for !operations.IsEmpty() {
    currOp := operations.Dequeue()

    vL, okL := results[currOp.left]
    vR, okR := results[currOp.right]
    if okL && okR {
      switch currOp.op {
      case "+":
        results[currOp.key] = vL + vR
      case "-":
        results[currOp.key] = vL - vR
      case "*":
        results[currOp.key] = vL * vR
      case "/":
        results[currOp.key] = vL / vR
      default:
        panic("Operation not known")
      }
    } else {
      operations.Enqueue(currOp)
    }
  }

  fmt.Println("The value of root is:", results[ROOT])

  return results[ROOT]
}

func substitutionMethod(atom string, undef map[string]string) string {
  for !strings.Contains(atom, undef[HUMAN]) {
    for u, def := range undef {
      atom = strings.ReplaceAll(atom, u, "(" + def + ")")
    }
  }
  return atom
}

func solveEquation(root Op, undef map[string]string, results map[string]int64, end int64) int64 {
  // Substitute results already known, FMT: var1+var2
  for u, def := range undef {
    if u == HUMAN {
      continue
    }
    if _, ok := results[def[:4]]; ok {
      undef[u] = strings.ReplaceAll(
        def,
        def[:4],
        strconv.FormatInt(results[def[:4]], 10),
      )
    }
    if _, ok := results[def[5:]]; ok {
      undef[u] = strings.ReplaceAll(
        def,
        def[5:],
        strconv.FormatInt(results[def[5:]], 10),
      )
    }
  }

  // Substitute numbers in root too
  if _, ok := results[root.left]; ok {
    root.left = strings.ReplaceAll(
      root.left,
      root.left,
      strconv.FormatInt(results[root.left], 10),
    )
  }
  if _, ok := results[root.right]; ok {
    root.right = strings.ReplaceAll(
      root.right,
      root.right,
      strconv.FormatInt(results[root.right], 10),
    )
  }

  left := substitutionMethod(root.left, undef)
  right := root.right // the variable is always on the LEFT side

  start := int64(0)
  for {
    mid := int64((end+start) / 2)

    parameters := make(map[string]interface{})
    parameters[undef[HUMAN]] = int64(mid)

    expression, _ := govaluate.NewEvaluableExpression(left)
    tmp, err := expression.Evaluate(parameters)
    if tmp == nil {
      panic(err)
    }
    leftEq := int64(tmp.(float64))

    expression, _ = govaluate.NewEvaluableExpression(right)
    tmp, err = expression.Evaluate(parameters)
    if tmp == nil {
      panic(err)
    }
    rightEq := int64(tmp.(float64))

    if leftEq == rightEq {
      return mid
    }

    // merge-sort style to drop drastically the space of values to explore
    if leftEq < rightEq {
      end = mid
    } else {
      start = mid
    }
  }
}

func part2(operations *queue.Queue[Op], results map[string]int64, n int64) {
  // HUMAN(humn) is me, no more a monkey
  delete(results, HUMAN)

  // store values with unknown variables
  undef := make(map[string]string)
  undef[HUMAN] = "X"

  for len(operations.GetQueue()) > 1 {
    currOp := operations.Dequeue()

    if currOp.key == ROOT {
      currOp.op = "="
      operations.Enqueue(currOp)
      continue
    }

    if _, ok := undef[currOp.left]; ok {
      s := currOp.left + currOp.op + currOp.right
      undef[currOp.key] = s
    } else if _, ok := undef[currOp.right]; ok {
      s := currOp.left + currOp.op + currOp.right
      undef[currOp.key] = s
    } else {
      vL, okL := results[currOp.left]
      vR, okR := results[currOp.right]
      if okL && okR {
        switch currOp.op {
        case "+":
          results[currOp.key] = vL + vR
        case "-":
          results[currOp.key] = vL - vR
        case "*":
          results[currOp.key] = vL * vR
        case "/":
          results[currOp.key] = vL / vR
        default:
          panic("Operation not known")
        }
      } else {
        operations.Enqueue(currOp)
      }
    }
  }

  root := operations.Dequeue()
  X := solveEquation(root, undef, results, n)

  fmt.Println("The number to yell to solve the equation is:", X)
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  n := part1(parseInput(f))
  f.Seek(0, io.SeekStart)
  ops, res := parseInput(f)
  part2(ops, res, n)
}

