package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func parseInput(f *os.File) []string {
  var lines []string
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines
}

func parseAtIndex(line string, idx, strlen int) (string, error) {
  // is a digit
  if line[idx] >= '0' && line[idx] <= '9' {
    return string(line[idx]), nil
  }

  // it is spelled as word
  if idx + 2 < strlen {
    possible3LetterNumber := line[idx:idx+3]
    if possible3LetterNumber == "one" {
      return "1", nil
    }
    if possible3LetterNumber == "two" {
      return "2", nil
    }
    if possible3LetterNumber == "six" {
      return "6", nil
    }
  }
  if idx + 3 < strlen {
    possible4LetterNumber := line[idx:idx+4]
    if possible4LetterNumber == "four" {
      return "4", nil
    }
    if possible4LetterNumber == "five" {
      return "5", nil
    }
    if possible4LetterNumber == "nine" {
      return "9", nil
    }
    if possible4LetterNumber == "zero" {
      return "0", nil
    }
  }
  if idx + 4 < strlen {
    possible5LetterNumber := line[idx:idx+5]
    if possible5LetterNumber == "three" {
      return "3", nil
    }
    if possible5LetterNumber == "seven" {
      return "7", nil
    }
    if possible5LetterNumber == "eight" {
      return "8", nil
    }
  }
  
  // not a digit nor a well-spelled number
  return "", errors.New("no digit")
}

func part1(lines []string) {
  sum := 0
  for _, line := range lines {
    isFirst := true
    var firstDigit, lastDigit rune
    for _, c := range line {
      if c >= '0' && c <= '9' {
        if isFirst {
          firstDigit = c
          isFirst = false
        }
        lastDigit = c
      }
    }
    num, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))
    sum += num
  }
  fmt.Println("Total sum:", sum)
}

func part2(lines []string) {
  sum := 0
  for _, line := range lines {
    isFirst := true
    var firstDigit, lastDigit string
    strlen := len(line)
    i := 0
    for i < strlen {
      num, err := parseAtIndex(line, i, strlen)
      if err != nil {
        i++
        continue
      }
      if isFirst {
        firstDigit = num
        isFirst = false
      }
      lastDigit = num
      i += 1
    }
    num, _ := strconv.Atoi(firstDigit + lastDigit)
    sum += num
  }
  fmt.Println("Total sum:", sum)
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
