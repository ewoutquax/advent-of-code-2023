package day11cosmicexpansion

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "11"

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines, 1)
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, universe.SumGalaxyDistance())
}

func solvePart2(inputFile string) {
	// too high
	// 598693677482
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines, 999999)
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, universe.SumGalaxyDistance())
}

type Location struct {
	X int
	Y int
}

type Galaxy struct {
	Location
}

type Universe struct {
	Galaxies map[Location]Galaxy
}

func (u Universe) SumGalaxyDistance() (sum int) {
	for locFrom, _ := range u.Galaxies {
		for locTo, _ := range u.Galaxies {
			if locTo != locFrom {
				sum += abs(locTo.X-locFrom.X) + abs(locTo.Y-locFrom.Y)
			}
		}
	}

	return sum / 2
}

func ParseInput(lines []string, expansion int) (u Universe) {
	var usedX = make(map[int]bool, len(lines[0]))
	var usedY = make(map[int]bool, len(lines))

	u.Galaxies = make(map[Location]Galaxy, 100)

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				usedX[x] = true
				usedY[y] = true

				loc := Location{x, y}
				u.Galaxies[loc] = Galaxy{
					Location: loc,
				}
			}
		}
	}

	// expand galaxy along rows
	for x := len(lines[0]) - 1; x >= 0; x-- {
		if _, exists := usedX[x]; !exists {
			fmt.Printf("solve: found unused row: %d\n", x)
			var moved = make(map[Location]bool, len(u.Galaxies))
			// Move all Galaxies beyond x
			for loc, galaxy := range u.Galaxies {
				_, exists := moved[loc]
				if loc.X > x && !exists {
					fmt.Printf("move galaxy at loc: %v\n", loc)
					u.Galaxies[Location{loc.X + expansion, loc.Y}] = galaxy
					delete(u.Galaxies, loc)
					moved[Location{loc.X + expansion, loc.Y}] = true
				}
			}
		}
	}

	// expand galaxy along columns
	for y := len(lines) - 1; y >= 0; y-- {
		if _, exists := usedY[y]; !exists {
			fmt.Printf("solve: found unused column: %d\n", y)
			var moved = make(map[Location]bool, len(u.Galaxies))
			// Move all Galaxies beyond x
			for loc, galaxy := range u.Galaxies {
				_, exists := moved[loc]
				if loc.Y > y && !exists {
					fmt.Printf("move galaxy at loc: %v\n", loc)
					u.Galaxies[Location{loc.X, loc.Y + expansion}] = galaxy
					delete(u.Galaxies, loc)
					moved[Location{loc.X, loc.Y + expansion}] = true
				}
			}
		}
	}

	return u
}

func abs(number int) int {
	if number < 0 {
		return 0 - number
	}

	return number
}
