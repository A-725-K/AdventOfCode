package main

import (
  "io"
  "os"
  "fmt"
  "bufio"
  "strings"
)

const (
  LOSE = 0
  DRAFT = 3
  WIN = 6
)

// part 1
var ROCKS = [...]string{"A", "X"}
var PAPERS = [...]string{"B", "Y"}
var SCISSORS = [...]string{"C", "Z"}

// part 2
var MY_MOVES = [...]string{"X", "Y", "Z"}
var OUTCOMES = map[int]string {
  LOSE: "X",
  DRAFT: "Y",
  WIN: "Z",
}

func isRock (move string) bool {
  return move == ROCKS[0] || move == ROCKS[1]
}
func isPaper (move string) bool {
  return move == PAPERS[0] || move == PAPERS[1]
}
func isScissor (move string) bool {
  return move == SCISSORS[0] || move == SCISSORS[1]
}

func rockScissorPaper(move1 string, move2 string) int {
  if (isRock(move1)) {
    if (isPaper(move2)) {
      return LOSE
    }
    if (isScissor(move2)) {
      return WIN
    }
    return DRAFT
  }

  if (isScissor(move1)) {
    if (isRock(move2)) {
      return LOSE
    }
    if (isPaper(move2)) {
      return WIN
    }
    return DRAFT
  }

  if (isPaper(move1)) {
    if (isScissor(move2)) {
      return LOSE
    }
    if (isRock(move2)) {
      return WIN
    }
    return DRAFT
  }

  panic("Move not known")
}

func getInitialScore(myMove string) int {
  if (isRock(myMove)) {
    return 1
  }
  if (isPaper(myMove)) {
    return 2
  }
  if (isScissor(myMove)) {
    return 3
  }
  panic("Move not valid")
}

func playRound(myMove string, opponentMove string) int {
  return getInitialScore(myMove) + rockScissorPaper(myMove, opponentMove)
}

func guessRound(opponentMove string, desiredOutcome string) int {
  for _, myMove := range MY_MOVES {
    possibleOutcome := rockScissorPaper(myMove, opponentMove)
    if (OUTCOMES[possibleOutcome] == desiredOutcome) {
      return getInitialScore(myMove) + possibleOutcome
    }
  }
  panic("Outcome not possible")
}

func part1 (f *os.File) {
  scanner := bufio.NewScanner(f)
  
  totalScore := 0
  for scanner.Scan() {
    moves := strings.Split(scanner.Text(), " ")
    myMove, opponentMove := moves[1], moves[0]
    totalScore += playRound(myMove, opponentMove)
  }

  fmt.Println("Your total score is:", totalScore)
}

func part2 (f *os.File) {
  scanner := bufio.NewScanner(f)
  
  totalScore := 0
  for scanner.Scan() {
    moves := strings.Split(scanner.Text(), " ")
    desiredOutcome, opponentMove := moves[1], moves[0]
    totalScore += guessRound(opponentMove, desiredOutcome)
  }

  fmt.Println("Your total score is:", totalScore)
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  part1(f)
  f.Seek(0, io.SeekStart)
  part2(f)
}
