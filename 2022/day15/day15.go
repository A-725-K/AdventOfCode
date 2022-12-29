package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
  "AdventOfCode/ds"
  c "AdventOfCode/types/coord"
  r "AdventOfCode/types/range"
)

func parseInput(f *os.File) ([]c.CoordWithDist, []c.CoordWithDist) {
  scanner := bufio.NewScanner(f)

  var sensors, beacons []c.CoordWithDist
  for scanner.Scan() {
    fields := strings.Split(scanner.Text(), " ")
    
    xB, err := strconv.Atoi(strings.TrimSuffix(strings.Split(fields[8], "=")[1], ","))
    if err != nil {
      panic("Cannot convert x coord of beacon")
    }
    yB, err := strconv.Atoi(strings.Split(fields[9], "=")[1])
    if err != nil {
      panic("Cannot convert y coord of beacon")
    }
    beacon := c.CoordWithDist{X: xB, Y: yB, Distance: 0}
    beacons = append(beacons, beacon)

    xS, err := strconv.Atoi(strings.TrimSuffix(strings.Split(fields[2], "=")[1], ","))
    if err != nil {
      panic("Cannot convert x coord of sensor")
    }
    yS, err := strconv.Atoi(strings.TrimSuffix(strings.Split(fields[3], "=")[1], ":"))
    if err != nil {
      panic("Cannot convert y coord of sensor")
    }
    sensor := c.CoordWithDist{X: xS, Y: yS, Distance: 0}
    sensor.Distance = sensor.ManhattanDist(beacon)
    sensors = append(sensors, sensor)
  }

  return sensors, beacons
}

func addCoverage(grid *[30][50]string, d int, s c.CoordWithDist) {
  howMany := d
  ii := 0
  for howMany > 0 {
    for j := 0; j < howMany; j++ {
      x, y := s.X - s.Distance + ii + j + 4, s.Y - ii + 4
      if x < 0 || x >= 50 || y < 0 || y >= 30 {
        continue
      }
      (*grid)[y][x] = "#"
    }
    ii++
    howMany -= 2
  }
  
  howMany = d
  ii = 0
  for howMany > 0 {
    for j := 0; j < howMany; j++ {
      x, y := s.X - s.Distance + ii + j + 4, s.Y + ii + 4
      if x < 0 || x >= 50 || y < 0 || y >= 30 {
        continue
      }
      (*grid)[y][x] = "#"
    }
    ii++
    howMany -= 2
  }
}

func printSensors(sensors, beacons []c.CoordWithDist) {
  fmt.Println("Sensors:")
  for i, s := range sensors {
    fmt.Printf("(%d, %d) --> connected to beacon (%d, %d) at distance %d\n",
      s.X, s.Y,
      beacons[i].X, beacons[i].Y,
      s.Distance,
    )
  }
}

func part1(sensors, beacons []c.CoordWithDist) {
  // interestingLine := 10 // mini_input
  interestingLine := 2000000 // input
  posWithNoBeacons := 0

  // var grid [30][50]string
  // for i := 0; i < 30; i++ {
  //   for j := 0; j < 50; j++ {
  //     grid[i][j] = "."
  //   }
  // }
  // for _, s := range sensors {
  //   grid[s.Y+4][s.X+4] = "S"
  // }
  // for _, b := range beacons {
  //   grid[b.Y+4][b.X+4] = "B"
  // }

  var ranges []r.Range
  for _, sensor := range sensors {
    diameter := (sensor.Distance * 2) + 1
    // addCoverage(&grid, diameter, sensor)
    if sensor.Y - sensor.Distance < interestingLine && interestingLine < sensor.Y + sensor.Distance {
      vertDist := ds.Abs(sensor.Y - interestingLine)
      radius := (diameter - vertDist*2 - 1) / 2
      rangeCoveredInInterestingLine := r.Range{
        Start: sensor.X - radius,
        End: sensor.X + radius,
      }
      ranges = append(ranges, rangeCoveredInInterestingLine)
    }
  }

  ranges = r.MergeRanges(ranges)
  for _, rng := range ranges {
    posWithNoBeacons += rng.End - rng.Start
  }

  // for _, s := range sensors {
  //   grid[s.Y+4][s.X+4] = "S"
  // }
  // for _, b := range beacons {
  //   grid[b.Y+4][b.X+4] = "B"
  // }
  // for i := 0; i < 30; i++ {
  //   for j := 0; j < 50; j++ {
  //     fmt.Print(grid[i][j])
  //   }
  //   fmt.Println()
  // }

  fmt.Println("The slots without beacons in line", interestingLine, "are", posWithNoBeacons)
}

func inspectRanges(ranges []r.Range) int {
  if len(ranges) == 2 {
    if emptySpace := ranges[0].Start - ranges[1].End; emptySpace > 1 {
      return ranges[0].Start - 1
    }
  }
  return -1
}

func part2(sensors, beacons []c.CoordWithDist) {
  // topHeight := 20 // mini_input
  topHeight := 4000000 // input

  var solX, solY int
  for interestingLine := 0; interestingLine < topHeight; interestingLine++ {
    var ranges []r.Range
    for _, sensor := range sensors {
      diameter := (sensor.Distance * 2) + 1
      if sensor.Y - sensor.Distance < interestingLine && interestingLine < sensor.Y + sensor.Distance {
        vertDist := ds.Abs(sensor.Y - interestingLine)
        radius := (diameter - vertDist*2 - 1) / 2
        rangeCoveredInInterestingLine := r.Range{
          Start: sensor.X - radius,
          End: sensor.X + radius,
        }
        ranges = append(ranges, rangeCoveredInInterestingLine)
      }
    }

    ranges = r.MergeRanges(ranges)
    if solX = inspectRanges(ranges); solX >= 0 {
      solY = interestingLine
      break
    }
  }

  tuningFrequency := solX * topHeight + solY
  fmt.Printf("The only possible position for the beacon is (%d,%d)\n", solX, solY)
  fmt.Println("The tuning frequency of the beacon is:", tuningFrequency)
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  sensors, beacons := parseInput(f)
  part1(sensors, beacons)
  part2(sensors, beacons)
}

