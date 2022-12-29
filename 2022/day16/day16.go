package main

import (
	"os"
	"fmt"
	"sort"
	"bufio"
	"strconv"
	"strings"
	"AdventOfCode/ds"
	graph "AdventOfCode/ds/graph"
)

const (
	TIME_LIMIT_WITH_ELEPHANT = 26
	TIME_LIMIT               = 30
	OPEN_VALVE               = 1
	WALK_TUNNEL              = 1
)

func parseInput(f *os.File) graph.Graph {
	scanner := bufio.NewScanner(f)

	g := graph.NewGraph()
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		fieldsLen := len(fields)
		name := fields[1]
		rate, err := strconv.Atoi(strings.TrimSuffix(strings.Split(fields[4], "=")[1], ";"))
		if err != nil {
			panic("Cannot convert rate: " + fields[4])
		}
		var neighborhood []string
		for i := 0; i < fieldsLen-9; i++ {
			neighborhood = append(neighborhood, strings.TrimSuffix(fields[9+i], ","))
		}
		g[name] = graph.NewNode(name, rate, neighborhood)
	}

	return g
}

// FIXME: BACKTRACKING APPROACH NOT WORKING PROPERLY... THE MISTAKE IS THAT IT IS
//       	NOT TRYING TO SOLVE WITH BOTH OPEN AND CLOSE VALVE, BUT ONLY ONE PATH
// func findBestPath(g Graph, currNode string, prevNode string, maxPressure *int, currPressure int, timeLimit int, valveOpened *map[string]bool) {
//   if timeLimit < 1 {
//     *maxPressure = ds.Max(*maxPressure, currPressure)
//     return
//   }
//
//   hasOpenValve := false
//   if g[currNode].value > 0 && !(*valveOpened)[currNode] {
//     hasOpenValve = true
//     (*valveOpened)[currNode] = true
//
//     timeLimit -= OPEN_VALVE
//     currPressure += g[currNode].value*timeLimit
//   }
//
//   for _, n := range g[currNode].neighbors {
//     if n == prevNode {
//       if len(g[currNode].neighbors) == 1 {
//         findBestPath(g, n, currNode, maxPressure, currPressure, timeLimit-WALK_TUNNEL, valveOpened)
//       }
//     } else {
//       findBestPath(g, n, currNode, maxPressure, currPressure, timeLimit-WALK_TUNNEL, valveOpened)
//     }
//   }
//
//   if hasOpenValve {
//     (*valveOpened)[currNode] = false
//   }
// }

func makeKey(s string, i int, m map[string]bool, n int) string {
	var tmp []string
	for k, v := range m {
		if v {
			tmp = append(tmp, k)
		}
	}
	sort.Strings(tmp)
	si := strconv.Itoa(i)
	sn := strconv.Itoa(n)
	return s + si + strings.Join(tmp, ":") + sn
}

// DID NOT KNOW THIS KIND OF PROBLEM COULD HAVE BEEN SOLVED WITH DYNAMIC
// PROGRAMMING (DP), HAD TO LOOK FOR THE SOLUTION
// FROM "Competitive Programming 3":
//
//	==> DP is primarily used to solve optimization problems and counting problems. If you
//	encounter a problem that says “minimize this” or “maximize that” or “count the ways to
//	do that”, then there is a (high) chance that it is a DP problem. Most DP problems in
//	programming contests only ask for the optimal/total value and not the optimal solution
//	itself, which often makes the problem easier to solve by removing the need to backtrack and
//	produce the solution.
//
// FIXME: MAKE UTIL TO RUN MEMOIZED FUNCTIONS SIMILAR TO PYTHON3
//			  @functools.lru_cache (not really possible)
func findBestPath(
	g graph.Graph,
	currNode string,
	timeLimit int,
	valveOpened map[string]bool,
	memo *map[string]int,
	elephants int,
) int {
	if timeLimit <= 1 {
		var res int
		if elephants == 0 {
			res = 0
		} else {
			res = findBestPath(g, "AA", TIME_LIMIT_WITH_ELEPHANT, valveOpened, memo, elephants-1)
		}
		(*memo)[makeKey(currNode, timeLimit, valveOpened, elephants)] = res
		return res
	}

	if pressure, ok := (*memo)[makeKey(currNode, timeLimit, valveOpened, elephants)]; ok {
		return pressure
	}

	maxPressure := -1
	if g[currNode].Value > 0 && !valveOpened[currNode] {
		valveOpened[currNode] = true
		maxPressure = (timeLimit-OPEN_VALVE)*g[currNode].Value + findBestPath(g, currNode, timeLimit-WALK_TUNNEL, valveOpened, memo, elephants)
		valveOpened[currNode] = false
	}

	for _, n := range g[currNode].Neighbors {
		maxPressure = ds.Max(maxPressure, findBestPath(g, n, timeLimit-WALK_TUNNEL, valveOpened, memo, elephants))
	}

	(*memo)[makeKey(currNode, timeLimit, valveOpened, elephants)] = maxPressure
	return maxPressure
}

func part1(g graph.Graph) {
	valveOpened := make(map[string]bool)
	for n := range g {
		if g[n].Value == 0 {
			valveOpened[n] = true
		} else {
			valveOpened[n] = false
		}
	}

	memo := make(map[string]int)
	pressureReleased := findBestPath(g, "AA", TIME_LIMIT, valveOpened, &memo, 0)
	fmt.Println("The total pressure released is", pressureReleased)
}

func part2(g graph.Graph) {
	valveOpened := make(map[string]bool)
	for n := range g {
		valveOpened[n] = false
	}

	memo := make(map[string]int)
	pressureReleased := findBestPath(g, "AA", TIME_LIMIT_WITH_ELEPHANT, valveOpened, &memo, 1)
	fmt.Println("The total pressure released with another elephant is", pressureReleased)
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g := parseInput(f)
	// g.VisitGraph("AA")
	part1(g)
	part2(g)
}

