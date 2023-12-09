package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Timeline []int

func (t Timeline) computeDiff() Timeline {
  endIdx := len(t) - 1
  var values Timeline
  for endIdx > 0 {
    values = append(values, t[0])
    allZeros := true
    for i := 0; i < endIdx; i++ {
      t[i] = t[i+1] - t[i]
      if t[i] != 0 {
        allZeros = false
      }
    }
    if allZeros {
      break
    }
    endIdx--
  }
  return values
}

func (t Timeline) PredictForward() int {
  t.computeDiff()
  prediction := 0
  for _, n := range t {
    prediction += n
  }
  return prediction
}

func (t Timeline) PredictBackward() int {
  values := t.computeDiff()
  prediction := 0
  for endIdx := len(values)-1; endIdx >= 0; endIdx-- {
    prediction = values[endIdx] - prediction 
  }
  return prediction
}

func parseInput(f *os.File) []Timeline {
  var timelines []Timeline
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    var timeline Timeline
    for _, nStr := range strings.Split(scanner.Text(), " ") {
      n, _ := strconv.Atoi(nStr)
      timeline = append(timeline, n)
    }
    timelines = append(timelines, timeline)
  }
  return timelines
}

func part1(timelines []Timeline) {
  sumOfExtrapolatedValues := 0
  for _, tl := range timelines {
    sumOfExtrapolatedValues += tl.PredictForward()
  }
  fmt.Println(
    "The sum of extrapolated values in the future is",
    sumOfExtrapolatedValues,
  )
}

func part2(timelines []Timeline) {
  sumOfExtrapolatedValues := 0
  for _, tl := range timelines {
    sumOfExtrapolatedValues += tl.PredictBackward()
  }
  fmt.Println(
    "The sum of extrapolated values in the past is",
    sumOfExtrapolatedValues,
  )
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  timelines1 := parseInput(f)
  f.Seek(0, 0)
  timelines2 := parseInput(f)

  part1(timelines1)
  part2(timelines2)
}
