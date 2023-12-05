package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_INT64 int64 = 1 << 63 - 1

type Range struct {
  Source, Destination, Length int64
}

type SeedToLocation struct {
  SeedToSoil, SoilToFertilizer, FertilizerToWater, WaterToLight, LightToTemperature, TemperatureToHumidity, HumidityToLocation []Range
}

func mutate(seed int64, seedToLocation *SeedToLocation) int64 {
  for i := 0; i < 7; i++ {
    var lst []Range
    switch i {
    case 0:
      lst = seedToLocation.SeedToSoil
    case 1:
      lst = seedToLocation.SoilToFertilizer
    case 2:
      lst = seedToLocation.FertilizerToWater
    case 3:
      lst = seedToLocation.WaterToLight
    case 4:
      lst = seedToLocation.LightToTemperature
    case 5:
      lst = seedToLocation.TemperatureToHumidity
    case 6:
      lst = seedToLocation.HumidityToLocation
    }
    for _, rng := range lst {
      if seed >= rng.Source && seed < rng.Source + rng.Length {
        diff := seed - rng.Source
        seed = rng.Destination + diff
        break
      }
    }
  }
  return seed
}

func parseInput(f *os.File) ([]int64, SeedToLocation, []Range) {
  scanner := bufio.NewScanner(f)
  var seedToLocation SeedToLocation
  var seeds []int64
  scanner.Scan()
  fields := strings.Split(scanner.Text(), " ")
  for i := 1; i < len(fields); i++ {
    seed, _ := strconv.ParseInt(fields[i], 10, 64)
    seeds = append(seeds, seed)
  }
  scanner.Scan() // empty line
  var rngs []Range
  for i, s := range seeds {
    if i % 2 == 1 {
      rngs = append(rngs, Range{Source: seeds[i-1], Length: s, Destination: 0})
    }
  }

  for i := 0; i < 7; i++ {
    scanner.Scan()
    var lst []Range
    for scanner.Scan() {
      line := scanner.Text()
      if line == "" {
        break
      }
      fields = strings.Split(line, " ")
      src, _ := strconv.ParseInt(fields[1], 10, 64)
      dst, _ := strconv.ParseInt(fields[0], 10, 64)
      length, _ := strconv.ParseInt(fields[2], 10, 64)
      lst = append(lst, Range{Source: src, Destination: dst, Length: length})
    }
    switch i {
    case 0:
      seedToLocation.SeedToSoil = lst
    case 1:
      seedToLocation.SoilToFertilizer = lst
    case 2:
      seedToLocation.FertilizerToWater = lst
    case 3:
      seedToLocation.WaterToLight = lst
    case 4:
      seedToLocation.LightToTemperature = lst
    case 5:
      seedToLocation.TemperatureToHumidity = lst
    case 6:
      seedToLocation.HumidityToLocation = lst
    }
  }
  return seeds, seedToLocation, rngs
}

func part1(seeds []int64, seedToLocation *SeedToLocation) {
  minLocation := MAX_INT64
  for idx, s := range seeds {
    seeds[idx] = s
    location := mutate(s, seedToLocation)
    if location < minLocation {
      minLocation = location
    }
  }
  fmt.Println("The minimum distance location is:", minLocation)
}

func part2(sourceRanges []Range, seedToLocation *SeedToLocation) {
  minLocation := MAX_INT64

  subroutine := func (rng Range, idx int, c chan int64) {
    minLocation := MAX_INT64
    for seed := rng.Source; seed < rng.Source + rng.Length; seed++ {
      location := mutate(seed, seedToLocation)
      if location < minLocation {
        minLocation = location
      }
    }
    c <- minLocation
    fmt.Println("[!!] Finished range", idx)
  }

  c := make(chan int64)
  for j, srcrng := range sourceRanges {
    fmt.Printf("(%d) Looking at range [%d, %d]\n", j, srcrng.Source, srcrng.Source + srcrng.Length)
    go subroutine(srcrng, j, c)
  }
  for i := 0; i < len(sourceRanges); i++ {
    location := <-c
    if location < minLocation {
      minLocation = location
    }
  }

  fmt.Println("The minimum distance location is:", minLocation)
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  seeds, seedToLocation, ranges := parseInput(f)
  // fmt.Println(seeds, seedToLocation)
 
  part1(seeds, &seedToLocation)
  part2(ranges, &seedToLocation)
}
