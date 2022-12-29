package main

import (
  "os"
  "fmt"
  "math"
  "bufio"
)

const BASE int64 = 5

func snafuToInt64(snafu string) int64 {
  lenS := len(snafu)
  var n int64 = 0
  for i := 0; i < lenS; i++ {
    factor := int64(math.Pow(float64(BASE), float64(lenS-i-1)))
    switch snafu[i] {
    case '2':
      n += 2 * factor
    case '1':
      n += factor
    case '0':
      n += 0
    case '-':
      n -= factor
    case '=':
      n -= 2 * factor
    default:
      panic("Invalid character")
    }
  }
  return n
}

func int64ToSnafu(n int64) string {
  snafu := ""

  fmt.Println(n)
  var carry, nextCarry int64 = 0, 0
  for n > 0 {
    carry = nextCarry
    nextDigit := n%BASE
    if nextDigit == 4 {
      nextDigit = -1
      nextCarry = 1
    } else if nextDigit == 3 {
      nextDigit = -2
      nextCarry = 1
    } else {
      nextCarry = 0
    }
    nextDigit += carry
    if nextDigit == 4 {
      nextDigit = -1
      nextCarry = 1
    } else if nextDigit == 3 {
      nextDigit = -2
      nextCarry = 1
    }
    // fmt.Println("N =>", n, "\tnextDigit =>", nextDigit, "\t\tsnafu =>", snafu)
    switch nextDigit {
    case 2:
      snafu = "2" + snafu
    case 1:
      snafu = "1" + snafu
    case 0:
      snafu = "0" + snafu
    case -1:
      snafu = "-" + snafu
    case -2:
      snafu = "=" + snafu
    default:
      fmt.Println(nextDigit)
      panic("Should not be possible, right?")
    }
    n /= BASE
  }

  if nextCarry > 0 {
    snafu = "1" + snafu
  }

  return snafu
}

func parseInput(f *os.File) []int64 {
  scanner := bufio.NewScanner(f)
  var nums []int64
  for scanner.Scan() {
    n := snafuToInt64(scanner.Text())
    nums = append(nums, n)
  }
  return nums
}

func part1(nums []int64) {
  var sum int64 = 0
  for _, n := range nums {
    sum += n
  }

  snafuAnswer := int64ToSnafu(sum)
  fmt.Println("The total amount of fuel in snafu is:", snafuAnswer)
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  nums := parseInput(f)

  part1(nums)
}

