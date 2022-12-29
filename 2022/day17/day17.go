package main

import (
	"fmt"
	"os"
	"bufio"
	"AdventOfCode/ds"
	s "AdventOfCode/ds/set"
	c "AdventOfCode/types/coord"
)

const (
	WIDTH                 = 7
	SMALL_TETROMINS_COUNT = 2022
	BIG_TETROMINS_COUNT   = 1000000000000
	ROCK_NUM              = 5
)

type Rock []c.Coord64

func getNextRock(idx int, y int64) Rock {
	y += 3

	// see: rocks.txt
	switch idx {
	case 0:
		return Rock{
			c.Coord64{X: 2, Y: y},
			c.Coord64{X: 3, Y: y},
			c.Coord64{X: 4, Y: y},
			c.Coord64{X: 5, Y: y},
		}
	case 1:
		return Rock{
			c.Coord64{X: 3, Y: y},
			c.Coord64{X: 2, Y: y + 1},
			c.Coord64{X: 3, Y: y + 1},
			c.Coord64{X: 4, Y: y + 1},
			c.Coord64{X: 3, Y: y + 2},
		}
	case 2:
		return Rock{
			c.Coord64{X: 2, Y: y},
			c.Coord64{X: 3, Y: y},
			c.Coord64{X: 4, Y: y},
			c.Coord64{X: 4, Y: y + 1},
			c.Coord64{X: 4, Y: y + 2},
		}
	case 3:
		return Rock{
			c.Coord64{X: 2, Y: y},
			c.Coord64{X: 2, Y: y + 1},
			c.Coord64{X: 2, Y: y + 2},
			c.Coord64{X: 2, Y: y + 3},
		}
	case 4:
		return Rock{
			c.Coord64{X: 2, Y: y},
			c.Coord64{X: 3, Y: y},
			c.Coord64{X: 2, Y: y + 1},
			c.Coord64{X: 3, Y: y + 1},
		}
	default:
		panic("Cannot found this kind of rock")
	}
}

func (r *Rock) moveRight() {
	for i := range *r {
		(*r)[i].X++
	}
}
func (r *Rock) moveLeft() {
	for i := range *r {
		(*r)[i].X--
	}
}
func (r *Rock) moveDown() {
	for i := range *r {
		(*r)[i].Y--
	}
}
func (r *Rock) moveUp() {
	for i := range *r {
		(*r)[i].Y++
	}
}

func (r *Rock) tryMoveRight(fallenRocks *s.Set[string]) {
	r.moveRight()
	for _, c := range *r {
		if c.X >= WIDTH || overlap(c, fallenRocks) {
			r.moveLeft()
			return
		}
	}
	// fmt.Println("===> MOVE RIGHT")
}
func (r *Rock) tryMoveLeft(fallenRocks *s.Set[string]) {
	r.moveLeft()
	for _, c := range *r {
		if c.X < 0 || overlap(c, fallenRocks) {
			r.moveRight()
			return
		}
	}
	// fmt.Println("<=== MOVE LEFT")
}
func (r *Rock) tryMoveDown(fallenRocks *s.Set[string]) (bool, int64) {
	r.moveDown()
	for _, c := range *r {
		if c.Y < 0 || overlap(c, fallenRocks) {
			r.moveUp()

			maxY := int64(0)
			for _, c := range *r {
				maxY = ds.Max(maxY, c.Y)
			}
			return false, maxY + 1
		}
	}
	// fmt.Println("VVV MOVE DOWN")
	return true, -1
}

func overlap(c c.Coord64, fallenRocks *s.Set[string]) bool {
	for fallenRock := range *fallenRocks {
		frX, frY := ds.FromKey64(fallenRock)
		if frX == c.X && frY == c.Y {
			return true
		}
	}
	return false
}

func (r *Rock) moveHorizontally(direction byte, fallenRocks *s.Set[string]) {
	switch direction {
	case '>':
		r.tryMoveRight(fallenRocks)
	case '<':
		r.tryMoveLeft(fallenRocks)
	default:
		panic("Invalid gas jet direction")
	}
}

func makeRockFall(
	fallenRocks *s.Set[string],
	rockIdx int,
	gasJets string,
	jetIdx *int,
	maxY *int64,
) {
	n := len(gasJets)
	rock := getNextRock(rockIdx, *maxY)

	for {
		rock.moveHorizontally(gasJets[*jetIdx], fallenRocks)
		*jetIdx = (*jetIdx + 1) % n
		if ok, newY := rock.tryMoveDown(fallenRocks); !ok {
			*maxY = ds.Max(*maxY, newY)
			for _, c := range rock {
				fallenRocks.Add(ds.ToKey64(c.X, c.Y))
			}
			break
		}
	}
}

func printCave(fallenRocks s.Set[string]) {
	maxY := int64(0)
	for c := range fallenRocks {
		_, y := ds.FromKey64(c)
		maxY = ds.Max(maxY, y)
	}
	maxY++

	for i := maxY; i >= 0; i-- {
		fmt.Print("|")
		for j := int64(0); j < WIDTH; j++ {
			if _, ok := fallenRocks[ds.ToKey64(j, i)]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println()
}

func simulateTetris(gasJets string, n int64) {
	jetIdx, rockIdx, maxY := 0, 0, int64(0)
	fallenRocks := s.NewSet[string]()

	for i := int64(0); i < n; i++ {
		makeRockFall(&fallenRocks, rockIdx, gasJets, &jetIdx, &maxY)
		rockIdx = (rockIdx + 1) % ROCK_NUM
	}
	printCave(fallenRocks)

	fmt.Println("After", n, "rocks the tower is tall", maxY)
}

func part1(gasJets string) {
	simulateTetris(gasJets, SMALL_TETROMINS_COUNT)
}

type Pattern struct {
	Tetramin, Height int64
}

func simulateTetrisFindPattern(gasJets string, n int64) (finalHeight int64) {
	jetIdx, rockIdx, maxY := 0, 0, int64(0)
	fallenRocks := s.NewSet[string]()
	alreadySeen := make(map[string]Pattern)

	for i := int64(0); i < n; i++ {
		makeRockFall(&fallenRocks, rockIdx, gasJets, &jetIdx, &maxY)
		rockIdx = (rockIdx + 1) % ROCK_NUM

		// the key represents which rock block stopped moving at which gas jet index
		// if this pair is duplicated it means that it found a repetition
		key := fmt.Sprintf("%d-%d", (i-1)%ROCK_NUM, jetIdx-1)
		if hit, ok := alreadySeen[key]; ok {
			q, rem := ds.Divmod(n-i, i-hit.Tetramin)

			if rem == 0 {
				// maxY so far include the pre-pattern height
				finalHeight = maxY + (maxY - hit.Height)*q - 1
				break
			}
		} else {
			alreadySeen[key] = Pattern{Tetramin: i, Height: maxY}
		}
	}

	return finalHeight
}

func part2(gasJets string) {
	// MY IDEA OF ALGORITHM:
	// - there are patterns
	// - find the start and end of the pattern
	// - find what is the height in the pattern
	// - divide the big number by the length of the pattern (Divmod)
	// - compute what is the height of the pre-pattern + rest of previous
	//   division
	// - compute fast the total height with those numbers
	// Thanks https://github.com/carolinasolfernandez for a huge hint on how to
	// find the repetitions
	maxY := simulateTetrisFindPattern(gasJets, BIG_TETROMINS_COUNT)
	fmt.Println("The height after", BIG_TETROMINS_COUNT, "rocks is:", maxY)
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	gasJets := scanner.Text()

	part1(gasJets)
	part2(gasJets)
}

