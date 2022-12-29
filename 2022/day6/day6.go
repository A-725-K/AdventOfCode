package main

import (
  "os"
  "fmt"
  "bufio"
)

func isUniq(buffer []byte, n int) bool {
  m := make(map[byte]int)
  cnt := 0
  for _, c := range buffer {
    if _, ok := m[c]; ok {
      m[c]++
    } else {
      m[c] = 1
      cnt++
    }
  }
  return cnt == n
}

func findStarter(s string, lenS, bufLen int) {
  idx := -1
  var smallBuffer []byte
  for i := 0; i < lenS; i++ {
    if i < bufLen {
      smallBuffer = append(smallBuffer, s[i])
    } else {
      smallBuffer[i%bufLen] = s[i]
      if isUniq(smallBuffer, bufLen) {
        idx = i+1
        break
      }
    }
  }
  fmt.Println("The index where the starter sequence starts is:", idx)
}

func part1(s string, n int) {
  findStarter(s, n, 4)
}

func part2(s string, n int) {
  findStarter(s, n, 14)
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  scanner := bufio.NewScanner(f)
  scanner.Scan()
  s := scanner.Text()
  n := len(s)

  part1(s, n)
  part2(s, n)
}

