package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
  Time, Record int
}

func parseInput(f *os.File) []Race {
  scanner := bufio.NewScanner(f)
  
  var times, records []int

  scanner.Scan()
  fields := strings.Split(scanner.Text(), " ")
  for i := 1; i < len(fields); i++ {
    if fields[i] == "" || fields[i] == " " {
      continue
    }
    n, _ := strconv.Atoi(fields[i])
    times = append(times, n)
  }

  scanner.Scan()
  fields = strings.Split(scanner.Text(), " ")
  for i := 1; i < len(fields); i++ {
    if fields[i] == "" || fields[i] == " " {
      continue
    }
    n, _ := strconv.Atoi(fields[i])
    records = append(records, n)
  }

  var races []Race
  for i := 0; i < len(times); i++ {
    races = append(races, Race{Time: times[i], Record: records[i]})
  }
  return races
}

func computeDistance(race Race) int {
  waysToBeatTheRecord := 0
  for speed := 1; speed < race.Time-1; speed++ {
    distance := (race.Time - speed) * speed
    if distance > race.Record {
      waysToBeatTheRecord++
    }
  }
  return waysToBeatTheRecord
}

func part1(races []Race) {
  totalWinningPossibilities := 1
  for _, race := range races {
    waysToBeatTheRecord := computeDistance(race)
    totalWinningPossibilities *= waysToBeatTheRecord
  }
  fmt.Println("The possibilities to win are:", totalWinningPossibilities)
}

func part2(races []Race) {
  timeStr, recordStr := "", ""
  for _, r := range races {
    timeStr += strconv.Itoa(r.Time)
    recordStr += strconv.Itoa(r.Record)
  }

  time, _ := strconv.Atoi(timeStr)
  record, _ := strconv.Atoi(recordStr)
  race := Race{Time: time, Record: record}

  waysToBeatTheRecord := computeDistance(race)
  fmt.Println("The possibilities to win are:", waysToBeatTheRecord)
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  races := parseInput(f)
 
  part1(races)
  part2(races)
}
