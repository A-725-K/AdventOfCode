package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
  Line string
  Blocks []int
}

// Proud to got to the part 1 by myself!!! :D I thought it would have been
// sufficient to generate all the permutations efficiently using some bit-wise
// logic :(
// 
// func getPossibleCombinations(unknownIdxs []int) [][]string {
//   var possibleCombinations [][]string
//   // fmt.Println(len(unknownIdxs))
//   n := len(unknownIdxs)
//   size := int(math.Pow(2, float64(n)))
//   for i := 0; i < size; i++ {
//     var testCase []string
//     for j := 0; j < n; j++ {
//       mask := 1 << j
//       if (mask & i) == 0 {
//         testCase = append(testCase, ".")
//       } else {
//         testCase = append(testCase, "#")
//       }
//     }
//     possibleCombinations = append(possibleCombinations, testCase)
//   }
//   // fmt.Println(possibleCombinations, len(possibleCombinations))
//   return possibleCombinations
// }
//
// func (m Machine) IsValid() bool {
//   blockCount := 0
//   blockIdx := 0
//   for _, c := range m.Line {
//     if c == '#' {
//       blockCount++
//     }
//     if c == '.' && blockCount > 0 {
//       if blockIdx == len(m.Blocks) {
//         break
//       }
//       if blockCount != m.Blocks[blockIdx] {
//         return false
//       }
//       blockCount = 0
//       blockIdx++
//     }
//   }
//   if blockIdx == len(m.Blocks) && blockCount == 0 {
//     return true
//   }
//   if blockIdx < len(m.Blocks)-1 {
//     return false
//   }
//   if blockIdx == len(m.Blocks)-1 && blockCount == m.Blocks[blockIdx] {
//     return true
//   }
//
//   return false
// }
//
// func (m Machine) GetIdxOfUnknowns() []int {
//   var unknownIdxs []int
//   for i, c := range m.Line {
//     if c == '?' {
//       unknownIdxs = append(unknownIdxs, i)
//     }
//   }
//   // fmt.Println(unknownIdxs)
//   return unknownIdxs
// }
//
// func (m *Machine) SubstituteValues(unknownIdxs []int, combination []string) {
//   newLine := ""
//   unknownIdx := 0
//   for i := 0; i < len(m.Line); i++ {
//     if unknownIdx < len(combination) && i == unknownIdxs[unknownIdx] {
//       newLine += combination[unknownIdx]
//       unknownIdx++
//     } else {
//       newLine += string(m.Line[i])
//     }
//   }
//   m.Line = newLine
// }
//
// func (m Machine) FindCombinations() int {
//   unknownIdxs := m.GetIdxOfUnknowns()
//   possibleCombinations := getPossibleCombinations(unknownIdxs)
//
//   validCombinations := 0
//   for _, c := range possibleCombinations {
//     m.SubstituteValues(unknownIdxs, c)
//     if m.IsValid() {
//       validCombinations++
//     }
//   }
//   return validCombinations
// }
// 
// func part1Slow(machines []Machine) {
//   possibleCombinations := 0
//   for _, m := range machines {
//     // fmt.Println(m)
//     possibleCombinations += m.FindCombinations()
//   }
//   fmt.Println("There are", possibleCombinations, "possible combinations")
// }

func parseInput(f *os.File) []Machine {
  var machines []Machine
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    fields := strings.Split(scanner.Text(), " ")
    var blocks []int
    for _, nStr := range strings.Split(fields[1], ",") {
      n, _ := strconv.Atoi(nStr)
      blocks = append(blocks, n)
    }
    machines = append(machines, Machine{Line: fields[0], Blocks: blocks})
  }
  return machines
}

// Once again, thanks hyper-neutrino, I am sure we will meet again...
// N.B. the memoization part I figured out by myself, even if it was for free
// with the rest of the algorithm in place :D
func CountArrangements(cfg string, nums []int, memo *map[string]int64) int64 {
  memoKey := ""
  for _, n := range nums {
    memoKey += fmt.Sprintf("%d:", n)
  }
  memoKey += cfg

  if res, ok := (*memo)[memoKey]; ok {
    return res
  }

  if cfg == "" {
    // if the line has been scanned entirely, only if we have no more blocks
    // the arrangement counts as valid
    if len(nums) == 0 {
      return 1
    }
    return 0
  }

  if len(nums) == 0 {
    // if there are no more blocks to examine but we still have # in the line
    // it means that there are additional blocks that we should not consider
    if strings.Contains(cfg, "#") {
      return 0
    }
    return 1
  }

  combinations := int64(0)

  // we consider ? to be a .
  if cfg[0] == '.' || cfg[0] == '?' {
    combinations += CountArrangements(cfg[1:], nums, memo)
  }

  // we consider ? to be a #
  if cfg[0] == '#' || cfg[0] == '?' {
    // we hit the start of the block and there are 3 conditions to consider:
    //  - there must be enough machines left in the line
    //  - there should not be any working machin in the first nums[0], because
    //    the broken ones are contiguous
    //  - the next machine has to be operational or there should be no more
    //    machines after
    enoughSpace := nums[0] <= len(cfg)
    allBroken := enoughSpace && !strings.Contains(cfg[:nums[0]], ".")
    nextIsOperational := allBroken && (nums[0] == len(cfg) || cfg[nums[0]] != '#')

    if nextIsOperational {
      newCfg := ""
      if nums[0]+1 < len(cfg) {
        newCfg = cfg[nums[0]+1:]
      }
      combinations += CountArrangements(newCfg, nums[1:], memo)
    }
  }

  (*memo)[memoKey] = combinations
  return combinations
}

func part1Fast(machines []Machine) {
  possibleCombinations := int64(0)
  memo := make(map[string]int64)
  for _, m := range machines {
    // fmt.Println(m)
    possibleCombinations += CountArrangements(m.Line, m.Blocks, &memo)
  }
  fmt.Println("There are", possibleCombinations, "possible combinations")
}

func part2(machines []Machine) {
  for i, m := range machines {
    newLine := ""
    newBlocks := []int{}
    for i := 0; i < 5; i++ {
      newLine += m.Line + "?"
      newBlocks = append(newBlocks, m.Blocks...)
    }
    machines[i].Line = newLine[:len(newLine)-1] // remove last '?'
    machines[i].Blocks = newBlocks
  }
  // fmt.Println(machines)

  memo := make(map[string]int64)
  possibleCombinations := int64(0)
  for _, m := range machines {
    possibleCombinations += CountArrangements(m.Line, m.Blocks, &memo)
  }
  fmt.Println("There are", possibleCombinations, "possible combinations")
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  machines := parseInput(f)
 
  // part1Slow(machines)
  part1Fast(machines)
  part2(machines)
}
