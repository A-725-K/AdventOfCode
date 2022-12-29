package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"AdventOfCode/ds"
	b64 "encoding/base64"
	set "AdventOfCode/ds/set"
	stack "AdventOfCode/ds/stack"
)

const (
	TIME_LIMIT      = 24
	LONG_TIME_LIMIT = 32
)

type Factory struct {
	// production costs
	OreRobotOreCost        uint8
	ClayRobotOreCost       uint8
	ObsidianRobotOreCost   uint8
	ObsidianRobotClayCost  uint8
	GeodeRobotOreCost      uint8
	GeodeRobotObsidianCost uint8

	// workers
	OreRobots      uint8
	ClayRobots     uint8
	ObsidianRobots uint8
	GeodeRobots    uint8

	// materials
	OreQty      uint8
	ClayQty     uint8
	ObsidianQty uint8
	GeodeQty    uint8

	// values
	Idx          uint8
	QualityLevel int
}

func NewFactory(oroc, croc, obroc, orcc, groc, grobc, idx uint8) *Factory {
	return &Factory{
		OreRobotOreCost:        oroc,
		ClayRobotOreCost:       croc,
		ObsidianRobotOreCost:   obroc,
		ObsidianRobotClayCost:  orcc,
		GeodeRobotOreCost:      groc,
		GeodeRobotObsidianCost: grobc,

		OreRobots:      1,
		ClayRobots:     0,
		ObsidianRobots: 0,
		GeodeRobots:    0,

		OreQty:      0,
		ClayQty:     0,
		ObsidianQty: 0,
		GeodeQty:    0,

		Idx:          idx,
		QualityLevel: 0,
	}
}

func (f *Factory) Copy() *Factory {
	return &Factory{
		OreRobotOreCost:        f.OreRobotOreCost,
		ClayRobotOreCost:       f.ClayRobotOreCost,
		ObsidianRobotOreCost:   f.ObsidianRobotOreCost,
		ObsidianRobotClayCost:  f.ObsidianRobotClayCost,
		GeodeRobotOreCost:      f.GeodeRobotOreCost,
		GeodeRobotObsidianCost: f.GeodeRobotObsidianCost,

		OreRobots:      f.OreRobots,
		ClayRobots:     f.ClayRobots,
		ObsidianRobots: f.ObsidianRobots,
		GeodeRobots:    f.GeodeRobots,

		OreQty:      f.OreQty,
		ClayQty:     f.ClayQty,
		ObsidianQty: f.ObsidianQty,
		GeodeQty:    f.GeodeQty,

		Idx:          f.Idx,
		QualityLevel: f.QualityLevel,
	}
}

func (f *Factory) String() string {
	return fmt.Sprintf(
		"\nBlueprint: %d\n  OreR: %d ore\n  ClayR: %d ore\n  ObsR: %d ore %d clay\n  GeoR: %d ore %d obsidian\n\n  OreR #: %d\n  ClayR #: %d\n  ObsR #: %d\n  GeoR #: %d\n\n  Ore $: %d\n  Clay $: %d\n  Obsidian $: %d\n  Geode $: %d\n\n  Quality Level: %d\n",
		f.Idx,
		f.OreRobotOreCost,
		f.ClayRobotOreCost,
		f.ObsidianRobotOreCost,
		f.ObsidianRobotClayCost,
		f.GeodeRobotOreCost,
		f.GeodeRobotObsidianCost,
		f.OreRobots,
		f.ClayRobots,
		f.ObsidianRobots,
		f.GeodeRobots,
		f.OreQty,
		f.ClayQty,
		f.ObsidianQty,
		f.GeodeQty,
		f.QualityLevel,
	)
}

func (f *Factory) CanProduceOreRobot() bool {
	return f.OreQty >= f.OreRobotOreCost
}
func (f *Factory) ProduceOreRobot() {
	f.OreQty -= f.OreRobotOreCost
	f.OreRobots++
}
func (f *Factory) CanProduceClayRobot() bool {
	return f.OreQty >= f.ClayRobotOreCost
}
func (f *Factory) ProduceClayRobot() {
	f.OreQty -= f.ClayRobotOreCost
	f.ClayRobots++
}
func (f *Factory) CanProduceObisdianRobot() bool {
	return f.OreQty >= f.ObsidianRobotOreCost && f.ClayQty >= f.ObsidianRobotClayCost
}
func (f *Factory) ProduceObsidianRobot() {
	f.OreQty -= f.ObsidianRobotOreCost
	f.ClayQty -= f.ObsidianRobotClayCost
	f.ObsidianRobots++
}
func (f *Factory) CanProduceGeodeRobot() bool {
	return f.OreQty >= f.GeodeRobotOreCost && f.ObsidianQty >= f.GeodeRobotObsidianCost
}
func (f *Factory) ProduceGeodeRobot() {
	f.OreQty -= f.GeodeRobotOreCost
	f.ObsidianQty -= f.GeodeRobotObsidianCost
	f.GeodeRobots++
}

func (f *Factory) ProduceMaterials() {
	f.OreQty += f.OreRobots
	f.ClayQty += f.ClayRobots
	f.ObsidianQty += f.ObsidianRobots
	f.GeodeQty += f.GeodeRobots
}

// Using uint8 instead of int because it saves space
func (f *Factory) ComputeQualityLevel(time uint8) {
	s := stack.NewStack[[9]uint8]()
	s.Push(
		[9]uint8{
			f.OreQty, f.ClayQty, f.ObsidianQty, f.GeodeQty,
			f.OreRobots, f.ClayRobots, f.ObsidianRobots, f.GeodeRobots,
			time,
		},
	)

	memo := set.NewSet[string]()
	maxQuality, geodesOpened := 0, uint8(0)

	// This is needed to understand if I really need to build a robot of a 
	// certain type. If I have already enough robots per material, do not need
	// to build a new one, it will only slow down the process. Look also at the
	// shoulBuild*Robot variables below
	var allCosts [3]uint8
	maxOre := ds.Max(
		f.OreRobotOreCost,
		ds.Max(
			f.ClayRobotOreCost,
			ds.Max(
				f.ObsidianRobotOreCost,
				f.GeodeRobotOreCost,
			),
		),
	)
	allCosts[0] = maxOre
	allCosts[1] = f.ObsidianRobotClayCost
	allCosts[2] = f.GeodeRobotObsidianCost

	for !s.IsEmpty() {
		currState := s.Pop()

		if currState[8] == 0 {
			if int(int(currState[3])*int(f.Idx)) > maxQuality {
				maxQuality = int(int(currState[3]) * int(f.Idx))
				geodesOpened = currState[3]
			}
			continue
		}

		k := fmt.Sprintf(
			"%d:%d:%d:%d:%d:%d:%d:%d:%d",
			currState[8],
			currState[3], currState[2], currState[1], currState[0],
			currState[7], currState[6], currState[5], currState[4],
		)
		// save a shorter key to save space: ~5GB of RAM and ~32s to run
		memoKey := b64.StdEncoding.EncodeToString([]byte(k))[:23]

		if memo.Contains(memoKey) {
			continue
		}
		memo.Add(memoKey)

		s.Push(
			[9]uint8{
				currState[0] + currState[4], currState[1] + currState[5], currState[2] + currState[6], currState[3] + currState[7],
				currState[4], currState[5], currState[6], currState[7],
				currState[8] - 1,
			},
		)

		if currState[0] >= f.GeodeRobotOreCost && currState[2] >= f.GeodeRobotObsidianCost {
			s.Push(
				[9]uint8{
					currState[0] + currState[4] - f.GeodeRobotOreCost, currState[1] + currState[5], currState[2] + currState[6] - f.GeodeRobotObsidianCost, currState[3] + currState[7],
					currState[4], currState[5], currState[6], currState[7] + 1,
					currState[8] - 1,
				},
			)
			// If I can build a Geode robot, this is the best choice for this turn,
			// prune other possible solutions
			continue
		}

		shouldBuildOreRobot := currState[4] < allCosts[0]
		if shouldBuildOreRobot && currState[0] >= f.OreRobotOreCost {
			s.Push(
				[9]uint8{
					currState[0] + currState[4] - f.OreRobotOreCost, currState[1] + currState[5], currState[2] + currState[6], currState[3] + currState[7],
					currState[4] + 1, currState[5], currState[6], currState[7],
					currState[8] - 1,
				},
			)
		}

		shouldBuildClayRobot := currState[5] < allCosts[1]
		if shouldBuildClayRobot && currState[0] >= f.ClayRobotOreCost {
			s.Push(
				[9]uint8{
					currState[0] + currState[4] - f.ClayRobotOreCost, currState[1] + currState[5], currState[2] + currState[6], currState[3] + currState[7],
					currState[4], currState[5] + 1, currState[6], currState[7],
					currState[8] - 1,
				},
			)
		}

		shouldBuildObsidianRobot := currState[6] < allCosts[2]
		if shouldBuildObsidianRobot &&
			 currState[0] >= f.ObsidianRobotOreCost &&
			 currState[1] >= f.ObsidianRobotClayCost {
			s.Push(
				[9]uint8{
					currState[0] + currState[4] - f.ObsidianRobotOreCost, currState[1] + currState[5] - f.ObsidianRobotClayCost, currState[2] + currState[6], currState[3] + currState[7],
					currState[4], currState[5], currState[6] + 1, currState[7],
					currState[8] - 1,
				},
			)
		}
	}

	f.QualityLevel = maxQuality
	f.GeodeQty = geodesOpened
}

func parseFactory(fields []string) *Factory {
	idx, err := strconv.Atoi(strings.TrimSuffix(fields[1], ":"))
	if err != nil {
		panic("Cannot convert idx")
	}
	oroc, err := strconv.Atoi(fields[6])
	if err != nil {
		panic("Cannot convert oroc")
	}
	croc, err := strconv.Atoi(fields[12])
	if err != nil {
		panic("Cannot convert croc")
	}
	obroc, err := strconv.Atoi(fields[18])
	if err != nil {
		panic("Cannot convert obroc")
	}
	orcc, err := strconv.Atoi(fields[21])
	if err != nil {
		panic("Cannot convert orcc")
	}
	groc, err := strconv.Atoi(fields[27])
	if err != nil {
		panic("Cannot convert groc")
	}
	grobc, err := strconv.Atoi(fields[30])
	if err != nil {
		panic("Cannot convert grobc")
	}

	return NewFactory(
		uint8(oroc),
		uint8(croc),
		uint8(obroc),
		uint8(orcc),
		uint8(groc),
		uint8(grobc),
		uint8(idx),
	)
}

func parseInput(f *os.File) []*Factory {
	scanner := bufio.NewScanner(f)

	var factories []*Factory
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		factories = append(factories, parseFactory(fields))
	}

	return factories
}

func part1(factories []*Factory) {
	totQualityLevel := 0
	for _, f := range factories {
		f.ComputeQualityLevel(TIME_LIMIT)
		fmt.Println(f.Idx, "====>", f.QualityLevel)
		totQualityLevel += int(f.QualityLevel)
	}

	fmt.Println("The global quality level is:", totQualityLevel)
}

func part2(factories []*Factory) {
	totQualityLevel := 1

	for _, f := range factories {
		f.ComputeQualityLevel(LONG_TIME_LIMIT)
		fmt.Println(f.Idx, "====>", f.GeodeQty)
		totQualityLevel *= int(f.GeodeQty)
	}

	fmt.Println("The global quality level is:", totQualityLevel)
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	factories := parseInput(f)
	var factories2 []*Factory
	for _, f := range factories[:3] {
		factories2 = append(factories2, f.Copy())
	}
	part1(factories)
	part2(factories2)
}

