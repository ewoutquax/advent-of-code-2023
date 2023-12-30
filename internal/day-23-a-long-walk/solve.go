package day23alongwalk

import (
	"fmt"
	"maps"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/benchmark"
	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "23"

var bench benchmark.Benchmark

func init() {
	register.Day(Day+"a", solvePart1)
	// register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	u := ParseInput(utils.ReadFileAsLines(inputFile))
	path := u.FindLongestPath()

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, path.NrSteps-1)
}

func solvePart2(inputFile string) {
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, 0)

	// Too low:
	// --------
	//     2189
	//     2190
}

type Direction uint

const (
	DirectionNone Direction = iota
	DirectionUp
	DirectionRight
	DirectionDown
	DirectionLeft
)

type Coordinate struct {
	X int
	Y int
}

type Location struct {
	Coordinate

	IsStart         bool
	IsEnd           bool
	ForcedDirection Direction
	nrSteps         int // steps taken since the start
}

type Path struct {
	VisitedLocations          map[Coordinate]bool // Previous visited locations
	currentLocation           Location            // Where are we now
	currentForbiddenDirection Direction           // The direction we can't go, to prevent backtracking
	isActive                  bool                // Can we still reach valid locations from the current position
	isEndReached              bool                // We made it to the end
	NrSteps                   int                 // Nr steps taken since the start
}

func (p *Path) findValidNextLocations(u *Universe) (locs map[Direction]Location) {
	vectors := map[Direction][2]int{
		DirectionLeft:  {-1, 0},
		DirectionRight: {1, 0},
		DirectionUp:    {0, -1},
		DirectionDown:  {0, 1},
	}

	var isEndReached bool = false
	locs = map[Direction]Location{
		p.currentForbiddenDirection: p.currentLocation,
	}
	for !isEndReached && len(locs) == 1 {
		for currentForbiddenDirection, currentLoc := range locs {
			delete(locs, currentForbiddenDirection)

			for vectorDirection, vector := range vectors {
				c := Coordinate{
					X: currentLoc.X + vector[0],
					Y: currentLoc.Y + vector[1],
				}
				loc, locExists := u.Locations[c]
				_, locVisited := p.VisitedLocations[c]
				// if locExists && !locVisited && vectorDirection != currentForbiddenDirection && loc.ForcedDirection != contraDirection(vectorDirection) {
				if locExists && !locVisited && vectorDirection != currentForbiddenDirection {
					loc.nrSteps = currentLoc.nrSteps + 1
					locs[contraDirection(vectorDirection)] = loc
					if loc.IsEnd {
						isEndReached = true
					}
				}
			}
			if len(locs) > 1 {
				p.VisitedLocations[currentLoc.Coordinate] = true
			}
		}
	}

	return
}

func (p Path) moveToLoc(d Direction, l Location) Path {
	fmt.Printf("moveToLoc: %v\n", l)
	path := Path{
		VisitedLocations:          make(map[Coordinate]bool),
		currentLocation:           l,
		currentForbiddenDirection: d,
		isActive:                  true,
		isEndReached:              l.IsEnd,
		NrSteps:                   l.nrSteps,
	}

	bench.Start("Path/copyVisitedLocations")
	maps.Copy(path.VisitedLocations, p.VisitedLocations)
	bench.Stop("Path/copyVisitedLocations")

	path.VisitedLocations[l.Coordinate] = true

	return path
}

type Universe struct {
	Locations map[Coordinate]Location
}

func (u *Universe) FindLongestPath() Path {
	bench.Start("FindLongestPath")

	bench.Start("InitFirstPath")
	var paths = make([]Path, 0, len(u.Locations))
	var longestPath Path

	var start Location
	for _, l := range u.Locations {
		if l.IsStart {
			start = l
		}
	}
	path := Path{
		VisitedLocations:          map[Coordinate]bool{start.Coordinate: true},
		currentLocation:           start,
		currentForbiddenDirection: DirectionUp,
		isEndReached:              false,
		NrSteps:                   0,
	}
	paths = append(paths, path)
	bench.Stop("InitFirstPath")

	for len(paths) > 0 {
		var newPaths = make([]Path, 0, len(u.Locations))

		for _, p := range paths {
			bench.Start("findValidNextLocations")
			nextLocs := p.findValidNextLocations(u)
			bench.Stop("findValidNextLocations")
			for nextForbiddenDirection, nextLoc := range nextLocs {
				bench.Start("Path/moveToLoc")
				path := p.moveToLoc(nextForbiddenDirection, nextLoc)
				bench.Stop("Path/moveToLoc")
				bench.Start("appendPathToNewPaths")
				newPaths = append(newPaths, path)
				bench.Stop("appendPathToNewPaths")

				if nextLoc.IsEnd && longestPath.NrSteps < path.NrSteps {
					fmt.Println("we found a new, longer path to the end")
					longestPath = path
				}
			}
		}

		bench.Start("copyNewPathsToPaths")
		paths = newPaths
		bench.Stop("copyNewPathsToPaths")
	}

	bench.Stop("FindLongestPath")
	fmt.Printf("bench.Report(): %v\n", bench.Report())

	return longestPath
}

func ParseInput(lines []string) Universe {
	bench.Start("ParseInput")

	var maxY int = len(lines)
	var maxX int = len(lines[0])

	var universe = Universe{
		Locations: make(map[Coordinate]Location, maxX*maxY),
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			var isValidLoc bool = true
			loc := Location{
				Coordinate: Coordinate{
					X: x,
					Y: y,
				},
				IsStart: y == 0,
				IsEnd:   y == maxY-1,
			}

			switch char {
			case "#":
				isValidLoc = false
			case "^":
				loc.ForcedDirection = DirectionUp
			case ">":
				loc.ForcedDirection = DirectionRight
			case "v":
				loc.ForcedDirection = DirectionDown
			case "<":
				loc.ForcedDirection = DirectionLeft
			case ".":
				loc.ForcedDirection = DirectionNone
			}

			if isValidLoc {
				universe.Locations[loc.Coordinate] = loc
			}
		}
	}

	bench.Stop("ParseInput")

	return universe
}

func contraDirection(d Direction) Direction {
	return map[Direction]Direction{
		DirectionUp:    DirectionDown,
		DirectionRight: DirectionLeft,
		DirectionDown:  DirectionUp,
		DirectionLeft:  DirectionRight,
	}[d]
}
