package main

import (
	"io"
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"

	stack "AdventOfCode/ds/stack"
)

var crates []*stack.Stack[string]

func printCrates() {
	for i, s := range crates {
		ss := "crate [" + strconv.Itoa(i) + "]:"
		for _, e := range s.GetStack() {
			ss += " " + e
		}
		fmt.Println(ss)
	}
}

func setup() {
	crates = []*stack.Stack[string]{}
	for i := 0; i < 9; i++ {
		crates = append(crates, stack.NewStack[string]())
	}
	crates[0].Push("Z")
	crates[0].Push("J")
	crates[0].Push("G")

	crates[1].Push("Q")
	crates[1].Push("L")
	crates[1].Push("R")
	crates[1].Push("P")
	crates[1].Push("W")
	crates[1].Push("F")
	crates[1].Push("V")
	crates[1].Push("C")

	crates[2].Push("F")
	crates[2].Push("P")
	crates[2].Push("M")
	crates[2].Push("C")
	crates[2].Push("L")
	crates[2].Push("G")
	crates[2].Push("R")

	crates[3].Push("L")
	crates[3].Push("F")
	crates[3].Push("B")
	crates[3].Push("W")
	crates[3].Push("P")
	crates[3].Push("H")
	crates[3].Push("M")

	crates[4].Push("G")
	crates[4].Push("C")
	crates[4].Push("F")
	crates[4].Push("S")
	crates[4].Push("V")
	crates[4].Push("Q")

	crates[5].Push("W")
	crates[5].Push("H")
	crates[5].Push("J")
	crates[5].Push("Z")
	crates[5].Push("M")
	crates[5].Push("Q")
	crates[5].Push("T")
	crates[5].Push("L")

	crates[6].Push("H")
	crates[6].Push("F")
	crates[6].Push("S")
	crates[6].Push("B")
	crates[6].Push("V")

	crates[7].Push("F")
	crates[7].Push("J")
	crates[7].Push("Z")
	crates[7].Push("S")

	crates[8].Push("M")
	crates[8].Push("C")
	crates[8].Push("D")
	crates[8].Push("P")
	crates[8].Push("F")
	crates[8].Push("H")
	crates[8].Push("B")
	crates[8].Push("T")
}

func move9001(qty, from, to int) {
	var tmp []string
	for i := 0; i < qty; i++ {
		tmp = append(tmp, crates[from].Pop())
	}
	for i := 0; i < qty; i++ {
		crates[to].Push(tmp[qty-i-1])
	}
}

func move9000(qty, from, to int) {
	for i := 0; i < qty; i++ {
		crates[to].Push(crates[from].Pop())
	}
}

func parseAndExecCommands(scanner *bufio.Scanner, moveFunc func(int, int, int)) {
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		cmds := strings.Split(scanner.Text(), " ")
		qty, err := strconv.Atoi(cmds[1])
		if err != nil {
			panic("Cannot convert to int [qty]")
		}
		from, err := strconv.Atoi(cmds[3])
		if err != nil {
			panic("Cannot convert to int [from]")
		}
		to, err := strconv.Atoi(cmds[5])
		if err != nil {
			panic("Cannot convert to int [to]")
		}
		moveFunc(qty, from-1, to-1)
	}
}

func part1(f *os.File) {
	scanner := bufio.NewScanner(f)

	parseAndExecCommands(scanner, move9000)
	result := ""
	for _, s := range crates {
		result += s.Top()
	}
	fmt.Println("The string eventually is:", result)
}

func part2(f *os.File) {
	scanner := bufio.NewScanner(f)

	parseAndExecCommands(scanner, move9001)
	result := ""
	for _, s := range crates {
		result += s.Top()
	}
	fmt.Println("The string eventually is:", result)
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	setup()
	part1(f)
	f.Seek(0, io.SeekStart)
	setup()
	part2(f)
}

