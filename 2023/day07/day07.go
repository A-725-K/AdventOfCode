package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
  Cards []int
  Bid int64
}

const (
  fiveOfAKind = 7
  fourOfAKind = 6
  fullHouse = 5
  threeOfAKind = 4
  twoPair = 3
  onePair = 2
  highCard = 1

  DEBUG = false
)

func (h Hand) String() string {
  s := fmt.Sprintf("\nValue: %d\nHand: ", h.Bid)
  for _, c := range h.Cards {
    switch c {
      case 14:
        s += "A"
      case 13:
        s += "K"
      case 12:
        s += "Q"
      case 11:
        fallthrough
      case 1:
        s += "J"
      case 10:
        s += "T"
      default:
        s += fmt.Sprintf("%d", c)
    }
  }
  s += "\n"
  return s
}

func parseInput(f *os.File) []Hand {
  var hands []Hand
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    fields := strings.Split(scanner.Text(), " ")
    bid, _ := strconv.ParseInt(fields[1], 10, 64)
    hand := Hand{Cards: []int{}, Bid: bid}
    for _, c := range fields[0] {
      switch c {
      case 'A':
        hand.Cards = append(hand.Cards, 14)
      case 'K':
        hand.Cards = append(hand.Cards, 13)
      case 'Q':
        hand.Cards = append(hand.Cards, 12)
      case 'J':
        hand.Cards = append(hand.Cards, 11)
      case 'T':
        hand.Cards = append(hand.Cards, 10)
      default:
        hand.Cards = append(hand.Cards, int(c-'0'))
      }
    }
    hands = append(hands, hand)
  }
  return hands
}

func (h Hand) computeHandResult(withJoker bool) int {
  cardsCount := make(map[int]int)
  var keys []int

  for _, c := range h.Cards {
    if _, ok := cardsCount[c]; ok {
      keys = append(keys, c)
    }
    cardsCount[c]++
  }

  jokers := 0
  if withJoker {
    jokers = cardsCount[1]
    delete(cardsCount, 1)

    // corner case: all jokers
    if len(keys) == 0 {
      return fiveOfAKind
    }

    sort.SliceStable(keys, func(i, j int) bool{
      return cardsCount[keys[i]] > cardsCount[keys[j]]
    })
  }

  for _, k := range keys {
    count := cardsCount[k] + jokers
    if count == 5 {
      return fiveOfAKind
    }
    if count == 4 {
      return fourOfAKind
    }
    if count == 3 {
      if len(cardsCount) == 2 {
        return fullHouse
      }
      return threeOfAKind
    }
    if count == 2 {
      if len(cardsCount) == 2 {
        return fullHouse
      }
      if len(cardsCount) == 3 {
        return twoPair
      }
      return onePair
    }
  }
  return highCard
}

func printResult(res int) {
  s := "Result: "
  switch res {
  case fiveOfAKind:
    s += "5 of a kind"
  case fourOfAKind:
    s += "4 of a kind"
  case fullHouse:
    s += "Full house"
  case threeOfAKind:
    s += "3 of a kind"
  case twoPair:
    s += "2 pairs"
  case onePair:
    s += "1 pair"
  case highCard:
    s += "High card"
  }
  fmt.Println(s)
}

func getSorter(hands []Hand, withJokers bool) func (int, int) bool {
  return func (i, j int) bool {
    h1, h2 := hands[i], hands[j]
    h1Res, h2Res := h1.computeHandResult(withJokers), h2.computeHandResult(withJokers)

    if h1Res < h2Res {
      return true
    }
    if h1Res > h2Res {
      return false
    }
    for i := 0; i < 5; i++ {
      if h1.Cards[i] < h2.Cards[i] {
        return true
      }
      if h1.Cards[i] > h2.Cards[i] {
        return false
      }
    }
    return false
  }
}

func part1(hands []Hand) {
  sorter := getSorter(hands, false)
  sort.SliceStable(hands, sorter)

  if DEBUG {
    for _, h := range hands {
      fmt.Print(h)
      printResult(h.computeHandResult(true))
    }
  }

  totalWinnings := int64(0)
  for i, h := range hands {
    totalWinnings += int64(i+1) * h.Bid
  }
  fmt.Println("The total winning of this set of hands is:", totalWinnings)
}

func part2(hands []Hand) {
  for _, h := range hands {
    for i, c := range h.Cards {
      // demote Jokers
      if c == 11 {
        h.Cards[i] -= 10
      }
    }
  }

  sorter := getSorter(hands, true)
  sort.SliceStable(hands, sorter)

  if DEBUG {
    for _, h := range hands {
      fmt.Print(h)
      printResult(h.computeHandResult(true))
    }
  }

  totalWinnings := int64(0)
  for i, h := range hands {
    totalWinnings += int64(i+1) * h.Bid
  }
  fmt.Println("The total winning of this set of hands is:", totalWinnings)
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  hands := parseInput(f)
 
  part1(hands)
  part2(hands)
}
