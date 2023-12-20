package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Module struct {
  Type, Label string
  State bool
  ConnectedComponent []string
  LastImpulses map[string]string
}
type Circuit map[string]Module
type Queue struct {
  Queue []QueueItem
  Size int
}
type QueueItem struct {
  ImpulseType, Dest, Source string
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

func (m Module) String() string {
  state := "off"
  if m.State {
    state = "on"
  }
  s := "Module         : " + m.Label + "\n"
  s += "Type               : " + m.Type + "\n"
  s += fmt.Sprintf("Last Impulse       : %s\n", m.LastImpulses)
  s += "State              : " + state + "\n"
  s += "Connected Component: " + strings.Join(m.ConnectedComponent, ",") + "\n\n"
  return s
}

func (c *Circuit) Reset() {
  for _, module := range *c {
    for k := range module.LastImpulses {
      (*c)[module.Label].LastImpulses[k] = "low"
    }
  }
}

func parseInput(f *os.File) Circuit {
  circuit := make(Circuit)
  scanner := bufio.NewScanner(f)
  
  moduleRegexp := regexp.MustCompile("(&|%)?([a-z]+) -> ([a-z, ]*)")
  for scanner.Scan() {
    moduleFields := moduleRegexp.FindStringSubmatch(scanner.Text())
    var module Module
    module.Label = strings.Trim(moduleFields[2], " ")
    switch moduleFields[1] {
    case "%":
      module.Type = "flipflop"
      module.State = false
    case "&":
      module.Type = "conjunction"
      module.LastImpulses = make(map[string]string)
    case "":
      module.Type = "broadcaster"
    }

    for _, cc := range strings.Split(moduleFields[3], ",") {
      if cc == "" {
        continue
      }
      module.ConnectedComponent = append(
        module.ConnectedComponent,
        strings.Trim(cc, " "),
      )
    }

    circuit[moduleFields[2]] = module
  }

  for _, module := range circuit {
    if module.Type == "conjunction" {
      for _, other := range circuit {
        if other.Label == module.Label {
          continue
        }
        for _, cc := range other.ConnectedComponent {
          if cc == module.Label {
            module.LastImpulses[other.Label] = "low"
            break
          }
        }
      }
    }
  } 

  return circuit
}

func part1(circuit Circuit) {
  lowImpulses, highImpulses := 0, 0
  q := NewQueue()

  for i := 0; i < 1000; i++ {
    q.Enqueue(QueueItem{ImpulseType: "low", Dest: "broadcaster", Source: ""})
    for !q.IsEmpty() {
      impulse := q.Dequeue()
      module, ok := circuit[impulse.Dest]

      // only-output state
      if !ok {
        if impulse.ImpulseType == "low" {
          lowImpulses++
        } else {
          highImpulses++
        }
        continue
      }

      // TODO: simplify this logic in functions
      if impulse.ImpulseType == "low" {
        lowImpulses++

        if module.Type == "flipflop" {
          module.State = !module.State
          newImpulse := "low"
          if module.State {
            newImpulse = "high"
          }
          for _, cc := range module.ConnectedComponent {
            q.Enqueue(QueueItem{
              ImpulseType: newImpulse,
              Dest: cc,
              Source: module.Label,
            })
          }
        } else if module.Type == "conjunction" {
          module.LastImpulses[impulse.Source] = impulse.ImpulseType
          // compute new impulse
          newImpulse := "low"
          for _, li := range module.LastImpulses {
            if li == "low" {
              newImpulse = "high"
              break
            }
          }
          for _, cc := range module.ConnectedComponent {
            q.Enqueue(QueueItem{
              ImpulseType: newImpulse,
              Dest: cc, Source:
              module.Label,
            })
          }
        } else if module.Type == "broadcaster" {
          for _, cc := range module.ConnectedComponent {
            q.Enqueue(QueueItem{
              ImpulseType: impulse.ImpulseType,
              Dest: cc,
              Source: module.Label,
            })
          }
        }
      } else if impulse.ImpulseType == "high" {
        highImpulses++

        if module.Type == "conjunction" {
          module.LastImpulses[impulse.Source] = impulse.ImpulseType
          // compute new impulse
          newImpulse := "low"
          for _, li := range module.LastImpulses {
            if li == "low" {
              newImpulse = "high"
              break
            }
          }
          for _, cc := range module.ConnectedComponent {
            q.Enqueue(QueueItem{
              ImpulseType: newImpulse,
              Dest: cc,
              Source: module.Label,
            })
          }
        } else if module.Type == "broadcaster" {
          for _, cc := range module.ConnectedComponent {
            q.Enqueue(QueueItem{
              ImpulseType: impulse.ImpulseType,
              Dest: cc,
              Source: module.Label,
            })
          }
        }
      }
      circuit[impulse.Dest] = module
    }
  }

  // fmt.Println("Low impulses:", lowImpulses)
  // fmt.Println("High impulses:", highImpulses)
  fmt.Println("The total number of pulses sent is:", lowImpulses * highImpulses)
}

// TODO: TO GENERALIZE!!! PROBABLY ONLY 1 FUNCTION FOR BOTH PARTS
func part2(circuit Circuit) {
  q := NewQueue()

  lowImpulseNotFound := true
  // TODO: make those variables a map !!
  var bqRange, ltRange, qhRange, vzRange [2]int64
  for i := int64(0); lowImpulseNotFound; i++ {
    q.Enqueue(QueueItem{ImpulseType: "low", Dest: "broadcaster", Source: ""})
    for !q.IsEmpty() {
      impulse := q.Dequeue()
      module, ok := circuit[impulse.Dest]

      // only-output state
      if !ok {
        lastImpulses := circuit[impulse.Source].LastImpulses
        if lastImpulses["bq"] == "high" && bqRange[0] * bqRange[1] == 0 {
          if bqRange[0] == 0 {
            bqRange[0] = i
          } else if i != bqRange[0] {
            bqRange[1] = i
          }
        }
        if lastImpulses["lt"] == "high" && ltRange[0] * ltRange[1] == 0 {
          if ltRange[0] == 0 {
            ltRange[0] = i
          } else if i != ltRange[0] {
            ltRange[1] = i
          }
        }
        if lastImpulses["qh"] == "high" && qhRange[0] * qhRange[1] == 0 {
          if qhRange[0] == 0 {
            qhRange[0] = i
          } else if i != qhRange[0] {
            qhRange[1] = i
          }
        }
        if lastImpulses["vz"] == "high" && vzRange[0] * vzRange[1] == 0 {
          if vzRange[0] == 0 {
            vzRange[0] = i
          } else if i != vzRange[0] {
            vzRange[1] = i
          }
        }
        lowImpulseNotFound = !(
          bqRange[0] * bqRange[1] > 0 &&
          ltRange[0] * ltRange[1] > 0 &&
          qhRange[0] * qhRange[1] > 0 &&
          vzRange[0] * vzRange[1] > 0)
        if !lowImpulseNotFound {
          break
        }
        continue
      }

      if impulse.ImpulseType == "low" {
        if module.Type == "flipflop" {
          module.State = !module.State
          newImpulse := "low"
          if module.State {
            newImpulse = "high"
          }
          for _, cc := range module.ConnectedComponent {
            q.Enqueue(QueueItem{
              ImpulseType: newImpulse,
              Dest: cc,
              Source: module.Label,
            })
          }
        } else if module.Type == "conjunction" {
          module.LastImpulses[impulse.Source] = impulse.ImpulseType
          // compute new impulse
          newImpulse := "low"
          for _, li := range module.LastImpulses {
            if li == "low" {
              newImpulse = "high"
              break
            }
          }
          for _, cc := range module.ConnectedComponent {
            q.Enqueue(QueueItem{
              ImpulseType: newImpulse,
              Dest: cc,
              Source: module.Label,
            })
          }
        } else if module.Type == "broadcaster" {
          for _, cc := range module.ConnectedComponent {
            q.Enqueue(QueueItem{
              ImpulseType: impulse.ImpulseType,
              Dest: cc,
              Source: module.Label,
            })
          }
        }
      } else if impulse.ImpulseType == "high" {
        if module.Type == "conjunction" {
          module.LastImpulses[impulse.Source] = impulse.ImpulseType
          // compute new impulse
          newImpulse := "low"
          for _, li := range module.LastImpulses {
            if li == "low" {
              newImpulse = "high"
              break
            }
          }
          for _, cc := range module.ConnectedComponent {
            q.Enqueue(QueueItem{
              ImpulseType: newImpulse,
              Dest: cc,
              Source: module.Label,
            })
          }
        } else if module.Type == "broadcaster" {
          for _, cc := range module.ConnectedComponent {
            q.Enqueue(QueueItem{
              ImpulseType: impulse.ImpulseType,
              Dest: cc,
              Source: module.Label,
            })
          }
        }
      }
      circuit[impulse.Dest] = module
    }
  }

  pulsesCount := (bqRange[1]-bqRange[0]) *
    (ltRange[1]-ltRange[0]) *
    (qhRange[1]-qhRange[0]) *
    (vzRange[1]-vzRange[0])
  fmt.Println(
    "The number of pulses sent before having a low in output is:",
    pulsesCount,
  )
}


func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  circuit := parseInput(f)
 
  part1(circuit)
  circuit.Reset()
  part2(circuit)
}
