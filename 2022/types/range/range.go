package types

import (
  "sort"
  stack "AdventOfCode/ds/stack"
)

type Range struct {
  Start, End int
}

func (r1 *Range) Overlap(r2 Range) bool {
  return r1.Start >= r2.Start && r1.End <= r2.End
}

func (r1 *Range) PartialOverlap(r2 Range) bool {
  return r1.Start <= r2.Start && r2.Start <= r1.End
}

func (r *Range) String() string {
  // day4
  // return ">Start: " + strconv.Itoa(r.Start) + "\tEnd: " + strconv.Itoa(r.End)

  // generic
  return "(" + string(r.Start) + "," + string(r.End) + ")"
}

func MergeRanges(ranges []Range) []Range {
  if len(ranges) == 0 {
    return ranges
  }

  sort.SliceStable(ranges, func (i, j int) bool {
    return ranges[i].Start < ranges[j].Start
  })

  s := stack.NewStack[Range]()
  s.Push(ranges[0])

  n := len(ranges)
  for i := 1; i < n; i++ {
    currRange := s.Top()
    if currRange.End < ranges[i].Start {
      s.Push(ranges[i])
    } else if currRange.End < ranges[i].End {
      currRange.End = ranges[i].End
      s.Pop()
      s.Push(currRange)
    }
  }

  var merged []Range
  for !s.IsEmpty() {
    merged = append(merged, s.Pop())
  }

  return merged
}
