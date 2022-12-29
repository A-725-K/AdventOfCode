package main

import (
	"os"
	"fmt"
	"bufio"
	"AdventOfCode/ds"
	s "AdventOfCode/ds/set"
	stack "AdventOfCode/ds/stack"
	c "AdventOfCode/types/coord"
)

// DIRECTION3D: all possible directions in 3D space
var DIRECTION3D = []c.Coord3D{
	{X: 1, Y: 0, Z: 0},
	{X: 0, Y: 1, Z: 0},
	{X: 0, Y: 0, Z: 1},
	{X: -1, Y: 0, Z: 0},
	{X: 0, Y: -1, Z: 0},
	{X: 0, Y: 0, Z: -1},
}

func parseInput(f *os.File) []c.Coord3D {
	scanner := bufio.NewScanner(f)

	var cubes []c.Coord3D
	for scanner.Scan() {
		line := scanner.Text()
		cubes = append(cubes, c.FromKey3D(line))
	}

	return cubes
}

func part1(cubes []c.Coord3D) {
	adjacent := make(map[string]*s.Set[string])

	for _, c1 := range cubes {
		for _, c2 := range cubes {
			if c1.Equals(c2) {
				continue
			}

			if c1.IsAdjacent(c2) {
				if _, ok := adjacent[c1.ToKey()]; ok {
					adjacent[c1.ToKey()].Add(c2.ToKey())
				} else {
					ss := s.NewSet[string]()
					adjacent[c1.ToKey()] = &ss
					adjacent[c1.ToKey()].Add(c2.ToKey())
				}
				if _, ok := adjacent[c2.ToKey()]; ok {
					adjacent[c2.ToKey()].Add(c1.ToKey())
				} else {
					ss := s.NewSet[string]()
					adjacent[c2.ToKey()] = &ss
					adjacent[c2.ToKey()].Add(c1.ToKey())
				}
			}
		}
	}

	totalSides := len(cubes) * 6
	adjacentSides := 0
	for _, adjList := range adjacent {
		adjacentSides += adjList.Size()
	}
	visibleSides := totalSides - adjacentSides

	fmt.Println("In total there are", visibleSides, "visible sides of cubes")
}

func isPresent(cubes []c.Coord3D, node c.Coord3D) bool {
	for _, c := range cubes {
		if c.Equals(node) {
			return true
		}
	}
	return false
}

// Thanks William Y. Feng for your clear explanation of the 2nd part of the
// problem, withouth your help I would not have been able to implement this
// algorithm
func floodFill(
	cubes []c.Coord3D,
	node c.Coord3D,
	minXYZ, maxXYZ int,
	memo *map[string]int,
) int {
	if v, ok := (*memo)[node.String()]; ok {
		return v
	}

	toVisit := stack.NewStack[c.Coord3D]()
	toVisit.Push(node)

	alreadyChecked := s.NewSet[string]()
	for !toVisit.IsEmpty() {
		currNode := toVisit.Pop()
		if isPresent(cubes, currNode) {
			continue
		}

		if !currNode.IsInside(minXYZ, maxXYZ) {
			(*memo)[node.String()] = 1
			return 1
		}

		if _, ok := alreadyChecked[currNode.ToKey()]; ok {
			continue
		}
		alreadyChecked.Add(currNode.ToKey())

		for _, d := range DIRECTION3D {
			toVisit.Push(currNode.Move(d))
		}
	}

	(*memo)[node.String()] = 0
	return 0
}

func part2(cubes []c.Coord3D) {
	minXYZ, maxXYZ := 999, 0
	for _, c := range cubes {
		minXYZ = ds.Min(minXYZ, c.X)
		minXYZ = ds.Min(minXYZ, c.Y)
		minXYZ = ds.Min(minXYZ, c.Z)

		maxXYZ = ds.Max(maxXYZ, c.X)
		maxXYZ = ds.Max(maxXYZ, c.Y)
		maxXYZ = ds.Max(maxXYZ, c.Z)
	}

	memo := make(map[string]int)
	reachable := 0
	for _, c := range cubes {
		for _, d := range DIRECTION3D {
			reachable += floodFill(cubes, c.Move(d), minXYZ, maxXYZ, &memo)
		}
	}

	fmt.Println("There are actually", reachable, "sides")
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cubes := parseInput(f)

	part1(cubes)
	part2(cubes)
}

