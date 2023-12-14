package day14parabolicrefractiondish

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "14"

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	u := ParseInput(lines)
	u.RollRocks(North)
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, u.Weight())
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	u := ParseInput(lines)
	weight := RepeatCycles(1000000000, u)
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, weight)
}

type Location struct {
	X int
	Y int
}

func (l Location) toS() string {
	return fmt.Sprintf("%d,%d", l.X, l.Y)
}

type Rock struct {
	Location
	Rollable bool
}

type Universe struct {
	Rocks       map[Location]*Rock
	sortedRocks []*Rock

	MaxX int
	MaxY int
}

type Direction uint

const (
	North Direction = iota + 1
	West
	South
	East
)

func (u Universe) CycleTilts() {
	u.RollRocks(North)
	u.RollRocks(West)
	u.RollRocks(South)
	u.RollRocks(East)
}

func (u Universe) RollRocks(d Direction) {
	var loc Location

	var startX, startY, stepX, stepY int = 0, 0, 1, 1

	vX, vY := convertDirectionIntoVector(d)

	if vX == 1 {
		startX = u.MaxX
		stepX = -1
	}
	if vY == 1 {
		startY = u.MaxY
		stepY = -1
	}

	// Move rock in order of top to bottom
	for y := startY; y >= 0 && y <= u.MaxY; y += stepY {
		for x := startX; x >= 0 && x <= u.MaxX; x += stepX {
			loc = Location{x, y}
			if r, exists := u.Rocks[loc]; exists && r.Rollable {
				var locNext = Location{r.X + vX, r.Y + vY}
				var occupied bool

				_, occupied = u.Rocks[locNext]
				for !occupied &&
					locNext.X >= 0 && locNext.X <= u.MaxX &&
					locNext.Y >= 0 && locNext.Y <= u.MaxY {
					// Move the rock
					delete(u.Rocks, r.Location)
					u.Rocks[locNext] = r
					r.Location = locNext

					locNext = Location{r.X + vX, r.Y + vY}
					_, occupied = u.Rocks[locNext]
				}
			}
		}
	}
}

func convertDirectionIntoVector(d Direction) (int, int) {
	switch d {
	case North:
		return 0, -1
	case East:
		return 1, 0
	case South:
		return 0, 1
	case West:
		return -1, 0
	default:
		panic("No valid case found")
	}
}

func (u Universe) Weight() (sum int) {
	for _, r := range u.sortedRocks {
		if r.Rollable {
			sum += u.MaxY - r.Y + 1
		}
	}

	return
}

func (u Universe) hash() (hash string) {
	hashes := make([]string, 0)
	for _, r := range u.sortedRocks {
		hashes = append(hashes, r.Location.toS())
	}

	return strings.Join(hashes, ";")
}

func (u Universe) Draw() {
	for y := 0; y <= u.MaxY; y++ {
		for x := 0; x <= u.MaxX; x++ {
			loc := Location{x, y}
			if rock, exists := u.Rocks[loc]; exists {
				if rock.Rollable {
					fmt.Print("O")
				} else {
					fmt.Print("#")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func ParseInput(lines []string) (u Universe) {
	u.Rocks = make(map[Location]*Rock, 0)
	u.sortedRocks = make([]*Rock, 0)
	u.MaxY = len(lines) - 1
	u.MaxX = len(lines[0]) - 1

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			var rock Rock
			loc := Location{x, y}
			switch char {
			case "#":
				rock = Rock{Location: loc, Rollable: false}
			case "O":
				rock = Rock{Location: loc, Rollable: true}
			}
			if char != "." {
				u.Rocks[loc] = &rock
				u.sortedRocks = append(u.sortedRocks, &rock)
			}
		}
	}

	return
}

func RepeatCycles(repeat int32, u Universe) int {
	var scores = make(map[string]int32, 0)

	var step int32
	var cycleFound bool = false
	for step = 0; step < repeat; step++ {

		u.CycleTilts()

		hash := u.hash()

		if prevStep, exists := scores[hash]; !cycleFound && exists {
			cycleFound = true
			// fmt.Printf("solve: found repeating cycle with %d and %d\n", prevStep, step)
			cycles := repeat / (step - prevStep)
			step += (cycles - 1) * (step - prevStep)
			fmt.Printf("step updated to: %v\n", step)
		}
		scores[hash] = step

		fmt.Printf("solve: step %d: %s\n", step, u.hash())
	}

	return u.Weight()
}
