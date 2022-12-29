package main

import (
	"os"
	"fmt"
	"sync"
	"bufio"
	"strconv"
	"strings"
	"AdventOfCode/ds"
)

type Forest [][]int

type SafeForest struct {
	mtx   sync.Mutex
	trees map[string]bool
	ch    chan int
}

func readGrid(f *os.File) (Forest, int) {
	scanner := bufio.NewScanner(f)
	var grid Forest
	for scanner.Scan() {
		line := scanner.Text()

		var intLine []int
		ints := strings.Split(line, "")
		for _, i := range ints {
			n, err := strconv.Atoi(i)
			if err != nil {
				panic("Cannot covert to int")
			}
			intLine = append(intLine, n)
		}
		grid = append(grid, intLine)
	}

	return grid, len(grid)
}

/*
|
V
*/
func visibleNS(forest Forest, n int, treesSeenSoFar *SafeForest) {
	visibleTrees := 0
	for j := 1; j < n-1; j++ {
		highestTree := forest[0][j]
		for i := 1; i < n-1; i++ {
			if forest[i][j] > highestTree {
				highestTree = forest[i][j]

				coord := ds.ToKey(i, j)
				(*treesSeenSoFar).mtx.Lock()
				if _, ok := (*treesSeenSoFar).trees[coord]; !ok {
					visibleTrees++
					(*treesSeenSoFar).trees[coord] = true
				}
				(*treesSeenSoFar).mtx.Unlock()
			}
		}
	}

	(*treesSeenSoFar).ch <- visibleTrees
}

/*
^
|
*/
func visibleSN(forest Forest, n int, treesSeenSoFar *SafeForest) {
	visibleTrees := 0
	for j := 1; j < n-1; j++ {
		highestTree := forest[n-1][j]
		for i := n - 2; i >= 1; i-- {
			if forest[i][j] > highestTree {
				highestTree = forest[i][j]

				coord := ds.ToKey(i, j)
				(*treesSeenSoFar).mtx.Lock()
				if _, ok := (*treesSeenSoFar).trees[coord]; !ok {
					visibleTrees++
					(*treesSeenSoFar).trees[coord] = true
				}
				(*treesSeenSoFar).mtx.Unlock()
			}
		}
	}

	(*treesSeenSoFar).ch <- visibleTrees
}

/*
->
*/
func visibleWE(forest Forest, n int, treesSeenSoFar *SafeForest) {
	visibleTrees := 0
	for i := 1; i < n-1; i++ {
		highestTree := forest[i][0]
		for j := 1; j < n-1; j++ {
			if forest[i][j] > highestTree {
				highestTree = forest[i][j]

				coord := ds.ToKey(i, j)
				(*treesSeenSoFar).mtx.Lock()
				if _, ok := (*treesSeenSoFar).trees[coord]; !ok {
					visibleTrees++
					(*treesSeenSoFar).trees[coord] = true
				}
				(*treesSeenSoFar).mtx.Unlock()
			}
		}
	}

	(*treesSeenSoFar).ch <- visibleTrees
}

/*
<-
*/
func visibleEW(forest Forest, n int, treesSeenSoFar *SafeForest) {
	visibleTrees := 0
	for i := 1; i < n-1; i++ {
		highestTree := forest[i][n-1]
		for j := n - 2; j >= 1; j-- {
			if forest[i][j] > highestTree {
				highestTree = forest[i][j]

				coord := ds.ToKey(i, j)
				(*treesSeenSoFar).mtx.Lock()
				if _, ok := (*treesSeenSoFar).trees[coord]; !ok {
					visibleTrees++
					(*treesSeenSoFar).trees[coord] = true
				}
				(*treesSeenSoFar).mtx.Unlock()
			}
		}
	}

	(*treesSeenSoFar).ch <- visibleTrees
}

func printForest(forest Forest, n int) {
	fmt.Println("Grid is", n, "x", n)
	for _, row := range forest {
		for _, el := range row {
			fmt.Printf("%d ", el)
		}
		fmt.Println()
	}
}

func computeScenicDistance(forest Forest, n, i, j int) int {
	top, bottom, left, right := 0, 0, 0, 0
	currHeight := forest[i][j]

	// look left
	for col := j - 1; col >= 0; col-- {
		left++
		if forest[i][col] >= currHeight {
			break
		}
	}
	// look right
	for col := j + 1; col < n; col++ {
		right++
		if forest[i][col] >= currHeight {
			break
		}
	}
	// look up
	for row := i - 1; row >= 0; row-- {
		top++
		if forest[row][j] >= currHeight {
			break
		}
	}
	// look down
	for row := i + 1; row < n; row++ {
		bottom++
		if forest[row][j] >= currHeight {
			break
		}
	}

	return top * bottom * left * right
}

/*
	N

W   E

	S
*/
func part1(forest Forest, n int) map[string]bool {
	treesSeenSoFar := SafeForest{trees: make(map[string]bool), ch: make(chan int)}
	visibleTrees := n*4 - 4

	go visibleNS(forest, n, &treesSeenSoFar)
	go visibleSN(forest, n, &treesSeenSoFar)
	go visibleWE(forest, n, &treesSeenSoFar)
	go visibleEW(forest, n, &treesSeenSoFar)

	for i := 0; i < 4; i++ {
		visibleTrees += <-treesSeenSoFar.ch
	}

	fmt.Println("There are", visibleTrees, "visible trees in the forest")
	return treesSeenSoFar.trees
}

func part2(forest Forest, n int, treesSeenSoFar map[string]bool) {
	maxScenicDistance := -1
	for k := range treesSeenSoFar {
		i, j := ds.FromKey(k)
		scenicDistance := computeScenicDistance(forest, n, i, j)
		if scenicDistance > maxScenicDistance {
			maxScenicDistance = scenicDistance
		}
	}
	fmt.Println("The biggest scenic distance is:", maxScenicDistance)
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	forest, n := readGrid(f)
	// printForest(forest, n)
	treesSeenSoFar := part1(forest, n)
	part2(forest, n, treesSeenSoFar)
}
