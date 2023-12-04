package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
  Id int
  MagicNumbers, NumberList []int
}

func parseInput(f *os.File) []Card {
  var cards []Card
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    fields := strings.Split(scanner.Text(), " ")
    i := 1
    for fields[i] == "" || fields[i] == " " {
      i++
    }
    id, _ := strconv.Atoi(fields[i][:len(fields[i])-1])
    var magicNumbers, numberList []int
    for ; fields[i] != "|"; i++ {
      n, _ := strconv.Atoi(fields[i])
      if n == 0 {
        continue
      }
      magicNumbers = append(magicNumbers, n)
    }
    for i++; i < len(fields); i++ {
      n, _ := strconv.Atoi(fields[i])
      if n == 0 {
        continue
      }
      numberList = append(numberList, n)
    }
    cards = append(cards, Card{Id: id, MagicNumbers: magicNumbers, NumberList: numberList})
  }
  return cards
}

func (c *Card) scratch() int {
  intersection := 0
  hash := make(map[int]bool)
  for _, n := range c.MagicNumbers {
    hash[n] = true
  }
  for _, n := range c.NumberList {
    if _, ok := hash[n]; ok {
      intersection++
    }
  }
  return intersection
}

func part1(cards []Card) {
  sumOfWinningCards := 0
  for _, c := range cards {
    winning := c.scratch()
    if winning > 0 {
       sumOfWinningCards += int(math.Pow(2, float64(winning)-1))
    }
  }

  fmt.Println("The sum of winning cards is:", sumOfWinningCards)
}

func part2(cards []Card) {
  scratchCards := 0
  winningCards := make(map[int]int)

  for _, c := range cards {
    winningCards[c.Id] = 1
  }
  for _, c := range cards {
    scratchCards += winningCards[c.Id]
    winning := c.scratch()

    for i := 0; i < winning; i++ {
      value := 1
      winningCards[c.Id +i + 1] += (value * winningCards[c.Id])
      value *= 2
    }
  }

  fmt.Println("In total I have", scratchCards, "scratch cards.")
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  cards := parseInput(f)
 
  part1(cards)
  part2(cards)
}
