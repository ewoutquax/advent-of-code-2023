package day10pipemaze

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "10"

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	u := ParseInput(lines)
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, u.StepsFarthestFromStart())
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	u := ParseInput(lines)
	u.StepsFarthestFromStart() // Find all the pipes in the loop
	u.CSIEnhance()
	u.MarkReachableTiles()
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, u.SumUnreachableTiles())
}

type Location struct {
	X int
	Y int
}

type Tile struct {
	Location
	LoopPipe   bool // Is this part of the loop
	Score      int  // Sum all count, divide by 9, and that's the answer of part 2
	IsEnclosed bool // Is this location surrounded by the loop? True by default
}

type Pipe struct {
	Symbol     string // How is this pipe bend
	IsStart    bool   // Is this the starting pipe?
	IsOnLoop   bool   // Is this piece on the loop? false by default
	Location          // X, Y-coordinations of the pipe
	Neighbours []*Pipe
}

func (p *Pipe) StepsToStart(comingFrom *Pipe) (onLoop bool, steps int) {
	if p.IsStart {
		// We found the end!
		return true, 0
	}

	if len(p.Neighbours) == 1 {
		// Dead end
		return false, 0
	}

	if p.Neighbours[0].Location != comingFrom.Location &&
		p.Neighbours[1].Location != comingFrom.Location {
		// This location doesn't point back to the previous location: invalid loop
		return false, 0
	}

	var nextNeighbour *Pipe = p.Neighbours[0]

	if nextNeighbour.Location == comingFrom.Location {
		nextNeighbour = p.Neighbours[1]
	}

	reachedStart, count := nextNeighbour.StepsToStart(p)
	if reachedStart {
		// Mark this piece of pipe as being part of the loop
		p.IsOnLoop = true
	}

	return reachedStart, count + 1
}

type Universe struct {
	Pipes         map[Location]*Pipe
	EnhancedTiles map[Location]*Tile // The pipes, with the loop marked, three times enhanced and surrounded by open location
}

func (u Universe) LoopLength() (maxLoopLength int) {
	var startingPipe Pipe

	for _, p := range u.Pipes {
		if p.IsStart {
			startingPipe = *p
		}
	}

	neighouringLocations := []Location{
		{startingPipe.X - 1, startingPipe.Y + 0},
		{startingPipe.X + 1, startingPipe.Y + 0},
		{startingPipe.X + 0, startingPipe.Y - 1},
		{startingPipe.X + 0, startingPipe.Y + 1},
	}
	for _, l := range neighouringLocations {
		if neighbouringPipe, exists := u.Pipes[l]; exists {
			if onLoop, loopLength := neighbouringPipe.StepsToStart(&startingPipe); onLoop {
				if maxLoopLength < loopLength {
					maxLoopLength = loopLength
				}
			}
		}
	}

	return maxLoopLength
}

func (u Universe) StepsFarthestFromStart() int {
	return (u.LoopLength() + 1) / 2
}

func (u *Universe) CSIEnhance() {
	var maxX, maxY int
	u.EnhancedTiles = make(map[Location]*Tile, 10000)

	for _, pipe := range u.Pipes {
		if pipe.IsOnLoop || pipe.IsStart {
			pipeX := pipe.X * 3
			pipeY := pipe.Y * 3
			if maxX < pipeX+2 {
				maxX = pipeX + 2
			}
			if maxY < pipeY+2 {
				maxY = pipeY + 2
			}
			// Create 9 empty tiles, with score 0
			for idxY := 0; idxY < 3; idxY++ {
				for idxX := 0; idxX < 3; idxX++ {
					tileLoc := Location{pipeX + idxX, pipeY + idxY}
					u.EnhancedTiles[tileLoc] = &Tile{
						Location:   tileLoc,
						LoopPipe:   false,
						Score:      0,
						IsEnclosed: false,
					}
				}
			}

			// Mark 3 tiles as part of the loop, depending on the pipe-symbol
			u.EnhancedTiles[Location{pipeX + 1, pipeY + 1}].LoopPipe = true

			switch pipe.Symbol {
			case "F":
				u.EnhancedTiles[Location{pipeX + 1, pipeY + 2}].LoopPipe = true
				u.EnhancedTiles[Location{pipeX + 2, pipeY + 1}].LoopPipe = true
			case "-":
				u.EnhancedTiles[Location{pipeX + 0, pipeY + 1}].LoopPipe = true
				u.EnhancedTiles[Location{pipeX + 2, pipeY + 1}].LoopPipe = true
			case "7":
				u.EnhancedTiles[Location{pipeX + 0, pipeY + 1}].LoopPipe = true
				u.EnhancedTiles[Location{pipeX + 1, pipeY + 2}].LoopPipe = true
			case "|":
				u.EnhancedTiles[Location{pipeX + 1, pipeY + 0}].LoopPipe = true
				u.EnhancedTiles[Location{pipeX + 1, pipeY + 2}].LoopPipe = true
			case "J":
				u.EnhancedTiles[Location{pipeX + 1, pipeY + 0}].LoopPipe = true
				u.EnhancedTiles[Location{pipeX + 0, pipeY + 1}].LoopPipe = true
			case "L":
				u.EnhancedTiles[Location{pipeX + 2, pipeY + 1}].LoopPipe = true
				u.EnhancedTiles[Location{pipeX + 1, pipeY + 0}].LoopPipe = true
			case "S":
				u.EnhancedTiles[Location{pipeX + 1, pipeY + 0}].LoopPipe = true
				u.EnhancedTiles[Location{pipeX + 1, pipeY + 2}].LoopPipe = true
				u.EnhancedTiles[Location{pipeX + 0, pipeY + 1}].LoopPipe = true
				u.EnhancedTiles[Location{pipeX + 2, pipeY + 1}].LoopPipe = true
			default:
				panic("No valid case found")
			}

		}
	}

	// Create empty tiles, as not-reachable and with score 1
	fmt.Printf("solve: MaxX / MaxY: %d / %d\n", maxX, maxY)
	for y := 0; y < maxY+3; y++ {
		for x := 0; x < maxX+3; x++ {
			loc := Location{x, y}
			if _, exists := u.EnhancedTiles[loc]; !exists {
				u.EnhancedTiles[loc] = &Tile{
					Location:   loc,
					LoopPipe:   false,
					Score:      1,
					IsEnclosed: true,
				}
			}
		}
	}
}

func (u Universe) MarkReachableTiles() {
	var visitedLocations = make(map[Location]bool, len(u.EnhancedTiles))
	var reachedLocations = make(map[Location]bool, len(u.EnhancedTiles))
	var reachableLocations = make([]Location, 0)
	reachableLocations = append(reachableLocations, Location{0, 0})

	for len(reachableLocations) > 0 {
		currentLocation := reachableLocations[0]
		u.EnhancedTiles[currentLocation].IsEnclosed = false

		reachableLocations = reachableLocations[1:]
		visitedLocations[currentLocation] = true

		if !u.EnhancedTiles[currentLocation].LoopPipe {
			// Find next locations from this reachable locations; Do check if the location has already been visited
			for idxY := -1; idxY <= 1; idxY++ {
				for idxX := -1; idxX <= 1; idxX++ {
					tmpLocation := Location{currentLocation.X + idxX, currentLocation.Y + idxY}
					_, existsTile := u.EnhancedTiles[tmpLocation]
					_, existsVisited := visitedLocations[tmpLocation]
					_, existsReached := reachedLocations[tmpLocation]
					if existsTile && !existsVisited && !existsReached {
						reachableLocations = append(reachableLocations, tmpLocation)
						reachedLocations[tmpLocation] = true
					}
				}
			}
		}
	}
}

func (u Universe) Draw() {
	for y := 0; y < 30; y++ {
		for x := 0; x < 40; x++ {
			loc := Location{x, y}
			var printChar string
			if tile, exists := u.EnhancedTiles[loc]; exists {
				printChar = "."
				if tile.IsEnclosed {
					printChar = "I"
				}
				if tile.LoopPipe {
					printChar = "x"
				}
				fmt.Printf("%s", printChar)
			}
		}
		fmt.Printf("\n")
	}
}

func (u Universe) SumUnreachableTiles() (sum int) {
	for _, tile := range u.EnhancedTiles {
		if tile.IsEnclosed {
			sum += tile.Score
		}
	}

	return sum / 9
}

func ParseInput(lines []string) Universe {
	var u = Universe{
		Pipes: make(map[Location]*Pipe, 0),
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char != "." {
				pipe := Pipe{
					Symbol:   char,
					IsStart:  char == "S",
					IsOnLoop: false,
					Location: Location{x, y},
				}

				u.Pipes[pipe.Location] = &pipe
			}
		}
	}

	linkNeighbouringPipes(u)

	return u
}

func linkNeighbouringPipes(u Universe) {
	for loc, pipe := range u.Pipes {
		var possibleLocations = make([]Location, 0)

		switch pipe.Symbol {
		case "F":
			possibleLocations = append(possibleLocations, Location{loc.X, loc.Y + 1})
			possibleLocations = append(possibleLocations, Location{loc.X + 1, loc.Y})
		case "-":
			possibleLocations = append(possibleLocations, Location{loc.X - 1, loc.Y})
			possibleLocations = append(possibleLocations, Location{loc.X + 1, loc.Y})
		case "7":
			possibleLocations = append(possibleLocations, Location{loc.X - 1, loc.Y})
			possibleLocations = append(possibleLocations, Location{loc.X, loc.Y + 1})
		case "|":
			possibleLocations = append(possibleLocations, Location{loc.X, loc.Y - 1})
			possibleLocations = append(possibleLocations, Location{loc.X, loc.Y + 1})
		case "J":
			possibleLocations = append(possibleLocations, Location{loc.X, loc.Y - 1})
			possibleLocations = append(possibleLocations, Location{loc.X - 1, loc.Y})
		case "L":
			possibleLocations = append(possibleLocations, Location{loc.X + 1, loc.Y})
			possibleLocations = append(possibleLocations, Location{loc.X, loc.Y - 1})
		case "S":
			continue
		default:
			fmt.Printf("Unknown symbol: %v\n", pipe.Symbol)
			panic("Unknown symbol found")
		}

		for _, loc := range possibleLocations {
			if neighbour, exists := u.Pipes[loc]; exists {
				pipe.Neighbours = append(pipe.Neighbours, neighbour)
			}
		}
	}
}
