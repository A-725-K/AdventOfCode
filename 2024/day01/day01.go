package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(f *os.File) ([]int, []int) {
	var locIds1, locIds2 []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Split(strings.Replace(scanner.Text(), " ", "", 2), " ")
		locId1, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		locIds1 = append(locIds1, locId1)
		locId2, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		locIds2 = append(locIds2, locId2)
	}
	return locIds1, locIds2
}

func part1(locIds1, locIds2 []int) {
	sort.Ints(locIds1)
	sort.Ints(locIds2)

	diff := 0
	for i, id1 := range locIds1 {
		diff += int(math.Abs(float64(id1 - locIds2[i])))
	}
	fmt.Println("The diff is", diff)
}

func part2(locIds1, locIds2 []int) {
	diff := 0

	freq := make(map[int]int)
	for _, id2 := range locIds2 {
		if _, ok := freq[id2]; ok {
			freq[id2]++
		} else {
			freq[id2] = 1
		}
	}

	for _, id1 := range locIds1 {
		if weight, ok := freq[id1]; ok {
			diff += id1 * weight
		}
	}
	fmt.Println("The real diff is", diff)
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	locIds1, locIds2 := parseInput(f)

	part1(locIds1, locIds2)
	part2(locIds1, locIds2)
}
