package main

import (
	"bufio"
	"os"
)

func parseInput(f *os.File) []string {
  var lines []string
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines
}

func part1(lines []string) {

}

func part2(lines []string) {

}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  lines := parseInput(f)
 
  part1(lines)
  part2(lines)
}
