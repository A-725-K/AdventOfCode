package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameSettings struct {
  Red, Blue, Green int
}

type Game map[int][]Round

type Pair struct {
  Amount int
  Color string
}

type Round []Pair

func parseInput(f *os.File) (Game) {
  game := make(map[int][]Round)
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    line := scanner.Text()
    fields := strings.Split(line, ":")
    id, _ := strconv.Atoi(strings.Split(fields[0], " ")[1])
    var rounds []Round
    for _, c := range strings.Split(fields[1], ";") {
      var cubes Round
      for _, cc := range strings.Split(c, ",") {
        cc = strings.TrimSpace(cc)
        cubesFields := strings.Split(cc, " ")
        amount, _ := strconv.Atoi(strings.TrimSpace(cubesFields[0]))
        cubes = append(cubes, Pair{Amount: amount, Color: cubesFields[1]})
      }
      rounds = append(rounds, cubes)
    }
    game[id] = rounds
  }
  return game
}

func part1(g Game) {
  gameSettings := GameSettings{Red: 12, Blue: 14, Green: 13}
  sumOfPossible := 0

  for id, turn := range g {
    skipRound := false
    for _, round := range turn {
      for _, cubes := range round {
        switch cubes.Color {
        case "red":
          if cubes.Amount > gameSettings.Red {
            skipRound = true
          }
        case "blue":
          if cubes.Amount > gameSettings.Blue {
            skipRound = true
          }
        case "green":
          if cubes.Amount > gameSettings.Green {
            skipRound = true
          }
        default:
          panic("Impossible color")
        }
        if skipRound {
          break
        }
      }
    }
    if !skipRound {
      sumOfPossible += id
    }
  }

  fmt.Println("The sum of possible games is:", sumOfPossible)
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func part2(g Game) {
  sumOfPowers := 0
  
  for _, turn := range g {
    gameSettings := GameSettings{Red: 0, Green: 0, Blue: 0}
    for _, round := range turn {
      for _, cubes := range round {
        switch cubes.Color {
        case "red":
          gameSettings.Red = max(gameSettings.Red, cubes.Amount)
        case "blue":
          gameSettings.Blue = max(gameSettings.Blue, cubes.Amount)
        case "green":
          gameSettings.Green = max(gameSettings.Green, cubes.Amount)
        default:
          panic("Impossible color")
        }
      }
    }
    power := gameSettings.Red * gameSettings.Blue * gameSettings.Green
    sumOfPowers += power
  }

  fmt.Println("The sum of powers is:", sumOfPowers)
}

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  game := parseInput(f)

  part1(game)
  part2(game)
}
