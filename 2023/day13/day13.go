package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pattern [][]rune

func (p Pattern) Transpose() Pattern {
  n, m := len(p), len(p[0])
  var t Pattern
  for i := 0; i < m; i++ {
    t = append(t, make([]rune, n))
  }
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      t[j][i] = p[i][j]
    }
  }
  return t
}

func (p Pattern) String() string {
  s := ""
  for _, row := range p {
    for _, ch := range row {
      s += string(ch) + " "
    }
    s += "\n"
  }
  return s
}

func parseInput(f *os.File) []Pattern {
  var patterns []Pattern
  scanner := bufio.NewScanner(f)
  var pattern Pattern
  end := true
  for {
    scanner.Scan()
    if scanner.Text() == "" {
      if end {
        break
      }
      end = true
      patterns = append(patterns, pattern)
      pattern = Pattern{}
      continue
    }
    end = false
    var row []rune
    for _, c := range scanner.Text() {
      row = append(row, c)
    }
    pattern = append(pattern, row)
  }
  return patterns
}
func reverse(v Pattern) Pattern {
  var res Pattern
  for i := len(v)-1; i >= 0; i-- {
    res = append(res, v[i])
  }
  return res
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

func (p Pattern) FindMirror() int {
  for i := 1; i < len(p); i++ {
    above := reverse(p[:i])
    below := p[i:]

    minIdx := min(len(below), len(above))
    above = above[:minIdx]
    below = below[:minIdx]

    if above.String() == below.String() {
      return i
    }
  }
  return 0
}

func (p Pattern) FindMirrorWithSmudge() int {
  for i := 1; i < len(p); i++ {
    above := reverse(p[:i])
    below := p[i:]

    minIdx := min(len(below), len(above))

    // for each block above and below, count how many different characters
    // there are, we can tolerate at most 1
    //  ______
    // |x x x| x x
    // |x x x|
    // |_____|
    numOfMismatches := 0
    for i := 0; i < minIdx; i++ {
      minLen := min(len(above[i]), len(below[i]))
      for j := 0; j < minLen; j++ {
        if above[i][j] != below[i][j] {
          numOfMismatches++
        }
      }
    }

    // we have exactly one smudge on the mirror
    if numOfMismatches == 1 {
      return i
    }
  }
  return 0
}


// Cannot understand which cases are missing... :(
// func (p Pattern) FindMirror() int {
//   start, end := 1, len(p)-1
//   found := false
//   for start < end {
//     fmt.Println(start, end)
//     if string(p[start]) == string(p[end]) {
//       start++
//       end--
//       found = true
//       continue
//     } else {
//       found = false
//     }
//     start++
//   }
//   fmt.Println(1, found, start, end, len(p)-1)
//   found = found || (start < len(p)-1 && string(p[start]) == string(p[len(p)-1]))
//   found = found && start != end
//   if found {
//     return start
//   }
//
//   start, end = 0, len(p)-2
//   found = false
//   for start < end {
//     fmt.Println(start, end)
//     if string(p[start]) == string(p[end]) {
//       start++
//       end--
//       found = true
//       continue
//     } else {
//       found = false
//     }
//     end--
//   }
//   fmt.Println(2, found, start, end)
//   found = found && (end == 0 || (end > 0 && string(p[end]) == string(p[end-1])))
//   if found {
//     return start
//   }
//   return 0
// }

// :(
func part1(patterns []Pattern) {
  totalOfReflections := 0
  for _, p := range patterns {
    // fmt.Println(p.Transpose())
    mirrorIdx := p.FindMirror()
    // fmt.Print(i, ") row: ", mirrorIdx)
    totalOfReflections += 100 * mirrorIdx
    mirrorIdx = p.Transpose().FindMirror()
    // fmt.Println(" col:", mirrorIdx)
    totalOfReflections += mirrorIdx
    // fmt.Println("__________________")
  }
  fmt.Println("The total reflections are:", totalOfReflections)
}

func part2(patterns []Pattern) {
  totalOfReflectionsWithSmudge := 0
  for _, p := range patterns {
    // fmt.Println(p.Transpose())
    mirrorIdx := p.FindMirrorWithSmudge()
    // fmt.Print(i, ") row: ", mirrorIdx)
    totalOfReflectionsWithSmudge += 100 * mirrorIdx
    mirrorIdx = p.Transpose().FindMirrorWithSmudge()
    // fmt.Println(" col:", mirrorIdx)
    totalOfReflectionsWithSmudge += mirrorIdx
    // fmt.Println("__________________")
  }
  fmt.Println(
    "The total reflections after cleaning the smudge are:",
    totalOfReflectionsWithSmudge,
  )
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  patterns := parseInput(f)
  // patterns = patterns[29:30]
  // fmt.Println(patterns[0])
 
  part1(patterns)
  part2(patterns)
}
