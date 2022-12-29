package main

import (
	"os"
	"fmt"
	"bufio"
	// "AdventOfCode/ds"
	queue "AdventOfCode/ds/queue"
	c "AdventOfCode/types/coord"
)

type Maze [][]int

type Coord struct {
	x, y int
}

func parseInput(f *os.File) (Maze, Coord, Coord) {
	scanner := bufio.NewScanner(f)

	var maze Maze
	var start, end Coord
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for j, c := range line {
			toAppend := int(c - 'a')
			if c == 'S' {
				toAppend = 0
				start = Coord{i, j}
			} else if c == 'E' {
				toAppend = 25
				end = Coord{i, j}
			}
			row = append(row, toAppend)
		}
		maze = append(maze, row)
		i++
	}
	return maze, start, end
}

func (m Maze) display() {
	rows, cols := len(m), len(m[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%2v ", m[i][j])
		}
		fmt.Println()
	}
}

func isSafe(maze *Maze, visited *[][]bool, currPos, newCoord Coord) bool {
	rows, cols := len(*maze), len((*maze)[0])
	return newCoord.x >= 0 && newCoord.x < rows &&
		newCoord.y >= 0 && newCoord.y < cols &&
		((*maze)[newCoord.x][newCoord.y]-(*maze)[currPos.x][currPos.y]) <= 1 &&
		!(*visited)[newCoord.x][newCoord.y]
}

// // NOTE: IT WORKS ON THE SMALL INPUT BUT EXPLODE IN TIME ON THE REAL INPUT
// // TIME COMPLEXITY: O(4^(N*M)) BACKTRACKING
// func shortestPath(maze *Maze, start, end Coord, visited *[][]bool, min_dist *int, dist int) {
//   if start.x == end.x && start.y == end.y {
//     *min_dist = ds.Min(dist, *min_dist)
//     fmt.Println("Found a path of:", *min_dist)
//     return
//   }
//
//   (*visited)[start.x][start.y] = true
//
//   east := Coord{start.x + 1, start.y}
//   if isSafe(maze, visited, start, east) {
//     shortestPath(maze, east, end, visited, min_dist, dist + 1)
//   }
//
//   west := Coord{start.x - 1, start.y}
//   if isSafe(maze, visited, start, west) {
//     shortestPath(maze, west, end, visited, min_dist, dist + 1)
//   }
//
//   north := Coord{start.x, start.y + 1}
//   if isSafe(maze, visited, start, north) {
//     shortestPath(maze, north, end, visited, min_dist, dist + 1)
//   }
//
//   south := Coord{start.x, start.y - 1}
//   if isSafe(maze, visited, start, south) {
//     shortestPath(maze, south, end, visited, min_dist, dist + 1)
//   }
//
//   (*visited)[start.x][start.y] = false
// }

func shortestPath(maze Maze, start, end Coord) int {
	rows, cols := len(maze), len(maze[0])

	var visited [][]bool
	for i := 0; i < rows; i++ {
		var row []bool
		for j := 0; j < cols; j++ {
			row = append(row, false)
		}
		visited = append(visited, row)
	}

	visited[start.x][start.y] = true
	q := queue.NewQueue[c.CoordWithDist]()
	q.Enqueue(c.CoordWithDist{X: start.x, Y: start.y, Distance: 0})

	neighbors := []Coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for !q.IsEmpty() {
		currPos := q.Dequeue()

		if currPos.X == end.x && currPos.Y == end.y {
			return currPos.Distance
		}

		for _, n := range neighbors {
			newPos := Coord{currPos.X + n.x, currPos.Y + n.y}
			if isSafe(&maze, &visited, Coord{currPos.X, currPos.Y}, newPos) {
				visited[newPos.x][newPos.y] = true
				q.Enqueue(c.CoordWithDist{
					X: newPos.x,
					Y: newPos.y,
					Distance: currPos.Distance + 1,
				})
			}
		}
	}

	return -1
}

func part1(maze Maze, start, end Coord) {
	shortestPathLen := shortestPath(maze, start, end)
	fmt.Println("Shortest path consists of", shortestPathLen, "steps")
}

func part2(maze Maze, end Coord) {
	rows, cols := len(maze), len(maze[0])

	var shortestPaths []int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if maze[i][j] == 0 {
				shortestPathLen := shortestPath(maze, Coord{i, j}, end)
				if shortestPathLen > 0 {
					shortestPaths = append(shortestPaths, shortestPathLen)
				}
			}
		}
	}

	minDistance := shortestPaths[0]
	for _, sp := range shortestPaths {
		if sp < minDistance {
			minDistance = sp
		}
	}
	fmt.Println("The very shortest path is", minDistance, "steps")
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	maze, start, end := parseInput(f)
	// fmt.Println("START:", start, "\tEND:", end)
	// fmt.Println("MAZE:", len(maze), "X", len(maze[0]))
	// maze.display()
	part1(maze, start, end)
	part2(maze, end)
}
