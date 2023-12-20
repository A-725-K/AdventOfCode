package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range [2]int64
type Part struct {
  X, M, A, S int
}
type Rule struct {
  Logic func(Part) string
  Op, Letter, Destination string
  Num int64
}
type Workflow map[string][]Rule
type Queue struct {
  Queue []QueueItem
  Size int
}
type QueueItem struct {
  Label string
  XRange, MRange, ARange, SRange Range
  Path []string
}

func NewQueue() Queue {
  return Queue{Queue: []QueueItem{}, Size: 0}
}

func (q *Queue) Enqueue(qi QueueItem) {
  q.Queue = append(q.Queue, qi)
  q.Size++
}

func (q *Queue) Dequeue() QueueItem {
  defer func() {
    q.Size--
    q.Queue = q.Queue[1:]
  }()
  return q.Queue[0]
}

func (q Queue) IsEmpty() bool {
  return q.Size == 0
}

func (p Part) Value() int {
  return p.X + p.M + p.A + p.S
}

func (p Part) ApplyWorkflow(workflows map[string][]Rule) int {
  currWorkflow := "in"

  for currWorkflow != "A" && currWorkflow != "R" {
    workflow := workflows[currWorkflow]
    for _, rule := range workflow {
      currWorkflow = rule.Logic(p)
      if currWorkflow != "" {
         break
      }
    }
  }

  if currWorkflow == "R" {
    return 0
  }
  return p.Value()
}

func (r Rule) String() string {
  return fmt.Sprintf(
    "Op: %s\nLetter: %s\nDestination: %s\nNum: %d\n",
    r.Op, r.Letter, r.Destination, r.Num,
  )
}

func parseInput(f *os.File) (map[string][]Rule, []Part) {
  workflows := make(map[string][]Rule)
  var parts []Part
  scanner := bufio.NewScanner(f)
  nameRegex := regexp.MustCompile("([a-z]+){([^}]+)}")
  ruleRegex := regexp.MustCompile("([xmas])([<>])(\\d+):([a-zAR]+)")
  partRegex := regexp.MustCompile("{x=(\\d+),m=(\\d+),a=(\\d+),s=(\\d+)}")
  scanWorkflows := true
  for scanner.Scan() {
    if scanWorkflows {
      line := scanner.Text()
      if line == "" {
        scanWorkflows = false
        continue
      }
      fields := nameRegex.FindStringSubmatch(line)
      workflowName := fields[1]
      var rules []Rule
      for _, rule := range strings.Split(fields[2], ",") {
        ruleFields := ruleRegex.FindStringSubmatch(rule)
        if len(ruleFields) == 0 {
          rules = append(
            rules,
            Rule{
              Logic: func(p Part) string {
                return rule
              },
              Op: "",
              Num: 0,
              Letter: "",
              Destination: rule,
            },
          )
        } else {
          n, _ := strconv.Atoi(ruleFields[3])
          rules = append(
            rules,
            Rule{
              Logic: func(p Part) string {
                switch ruleFields[1] {
                case "x":
                  if (ruleFields[2] == ">" && p.X > n) ||
                    (ruleFields[2] == "<" && p.X < n) {
                    return ruleFields[4]
                  }
                case "m":
                  if (ruleFields[2] == ">" && p.M > n) ||
                    (ruleFields[2] == "<" && p.M < n) {
                    return ruleFields[4]
                  }
                case "a":
                  if (ruleFields[2] == ">" && p.A > n) ||
                    (ruleFields[2] == "<" && p.A < n) {
                    return ruleFields[4]
                  }
                case "s":
                  if (ruleFields[2] == ">" && p.S > n) ||
                    (ruleFields[2] == "<" && p.S < n) {
                    return ruleFields[4]
                  }
                  default:
                  panic("Character not known: " + ruleFields[1])
                }
                return ""
              },
              Op: ruleFields[2],
              Num: int64(n),
              Letter: ruleFields[1],
              Destination: ruleFields[4],
            })
        }
        workflows[workflowName] = rules
      }
    } else {
      partFields := partRegex.FindStringSubmatch(scanner.Text())
      x, _ := strconv.Atoi(partFields[1])
      m, _ := strconv.Atoi(partFields[2])
      a, _ := strconv.Atoi(partFields[3])
      s, _ := strconv.Atoi(partFields[4])
      parts = append(parts, Part{X: x, M: m, A: a, S: s})
    }
  }
  return workflows, parts 
}

func min64(x, y int64) int64 {
  if x < y {
    return x
  }
  return y
}

func max64 (x, y int64) int64 {
  if x > y {
    return x
  }
  return y
}

func part1(workflows map[string][]Rule, parts []Part) {
  sumOfAcceptedParts := 0
  for _, p := range parts {
    sumOfAcceptedParts += p.ApplyWorkflow(workflows)
  }
  fmt.Println("The sum of all accepted parts is:", sumOfAcceptedParts)
}

func part2(workflows map[string][]Rule) {
  q := NewQueue()
  validCombinations := int64(0)

  currItem := QueueItem{
    Label: "in",
    XRange: Range{1, 4000},
    MRange: Range{1, 4000},
    ARange: Range{1, 4000},
    SRange: Range{1, 4000},
    Path: []string{},
  }
  q.Enqueue(currItem)

  for !q.IsEmpty() {
    currItem = q.Dequeue()

    // rejected, ignore this state
    if currItem.Label == "R" {
      continue
    }

    // accepted, end state, compute the value of the current solution
    if currItem.Label == "A" {
      numOfCombinations := int64(currItem.XRange[1]-currItem.XRange[0]+1) *
        int64(currItem.MRange[1]-currItem.MRange[0]+1) *
        int64(currItem.ARange[1]-currItem.ARange[0]+1) *
        int64(currItem.SRange[1]-currItem.SRange[0]+1)
      validCombinations += numOfCombinations
      continue
    }

    for _, rule := range workflows[currItem.Label] {
      newItem := QueueItem{
        Label: rule.Destination,
        XRange: currItem.XRange,
        MRange: currItem.MRange,
        ARange: currItem.ARange,
        SRange: currItem.SRange,
        Path: append(currItem.Path, currItem.Label),
      }

      switch rule.Letter {
      case "x":
        if rule.Op == ">" {
          newItem.XRange[0] = max64(currItem.XRange[0], rule.Num+1)
          currItem.XRange[1] = newItem.XRange[0]-1
        } else if rule.Op == "<" {
          newItem.XRange[1] = min64(currItem.XRange[1], rule.Num-1)
          currItem.XRange[0] = newItem.XRange[1]+1
        }
      case "m":
        if rule.Op == ">" {
          newItem.MRange[0] = max64(currItem.MRange[0], rule.Num+1)
          currItem.MRange[1] = newItem.MRange[0]-1
        } else if rule.Op == "<" {
          newItem.MRange[1] = min64(currItem.MRange[1], rule.Num-1)
          currItem.MRange[0] = newItem.MRange[1]+1
        }
      case "a":
        if rule.Op == ">" {
          newItem.ARange[0] = max64(currItem.ARange[0], rule.Num+1)
          currItem.ARange[1] = newItem.ARange[0]-1
        } else if rule.Op == "<" {
          newItem.ARange[1] = min64(currItem.ARange[1], rule.Num-1)
          currItem.ARange[0] = newItem.ARange[1]+1
        }
      case "s":
        if rule.Op == ">" {
          newItem.SRange[0] = max64(currItem.SRange[0], rule.Num+1)
          currItem.SRange[1] = newItem.SRange[0]-1
        } else if rule.Op == "<" {
          newItem.SRange[1] = min64(currItem.SRange[1], rule.Num-1)
          currItem.SRange[0] = newItem.SRange[1]+1
        }
      }
      q.Enqueue(newItem)
    }
  }

  fmt.Println("There are", validCombinations, "valid combinations")
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  workflows, parts := parseInput(f)

  part1(workflows, parts)
  part2(workflows)
}
