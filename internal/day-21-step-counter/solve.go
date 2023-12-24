package day21stepcounter

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "21"

func init() {
	register.Day(Day+"a", solvePart1)
	// register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, ReachableInEvenSteps(&universe))
}

func solvePart2(inputFile string) {
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, 0)
}

type Coordinate struct {
	X int
	Y int
}

type Location struct {
	Coordinate
	reachedInEvenSteps bool
}

type Universe struct {
	Rocks map[Coordinate]bool
	Start Location

	MaxX int
	MaxY int
}

func ReachableInEvenSteps(u *Universe) (count int) {
	reachableLocations := u.ReachableLocations(64)

	for _, inEvenSteps := range reachableLocations {
		if inEvenSteps {
			count++
		}
	}

	return
}

func (u *Universe) ReachableLocations(maxSteps int) map[Coordinate]bool {
	var reachedLocations = make(map[Coordinate]bool, u.MaxX*u.MaxY)
	var visitCoordinates = []Coordinate{u.Start.Coordinate}
	var newCoordinates []Coordinate

	for step := 1; step <= maxSteps; step++ {
		newCoordinates = make([]Coordinate, 0)
		// fmt.Printf("step %d: locations to visit: %v\n", step, visitCoordinates)

		for _, loc := range visitCoordinates {
			for _, vector := range vectors() {
				newLoc := Location{
					Coordinate: Coordinate{
						X: loc.X + vector[0],
						Y: loc.Y + vector[1],
					},
					reachedInEvenSteps: step%2 == 0,
				}

				_, rockFound := u.Rocks[newLoc.Coordinate]
				_, isReached := reachedLocations[newLoc.Coordinate]
				if !(rockFound || isReached) {
					reachedLocations[newLoc.Coordinate] = newLoc.reachedInEvenSteps
					newCoordinates = append(newCoordinates, newLoc.Coordinate)
				}
			}
		}

		// u.Draw(reachedLocations, step)
		// fmt.Printf("step %d: new found coordinates: %v\n", step, newCoordinates)
		visitCoordinates = make([]Coordinate, len(newCoordinates))
		copy(visitCoordinates, newCoordinates)
	}

	return reachedLocations
}

func (u *Universe) Draw(reachedLocations map[Coordinate]bool, step int) {
	fmt.Printf("\nstep: %d\n", step)
	fmt.Println("------------")
	for y := 0; y <= u.MaxY; y++ {
		for x := 0; x <= u.MaxX; x++ {
			coor := Coordinate{x, y}
			if evenSteps, exists := reachedLocations[coor]; exists {
				if evenSteps {
					fmt.Print("O")
				} else {
					fmt.Print("x")
				}
			} else if _, exists := u.Rocks[coor]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func ParseInput(lines []string) Universe {
	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1

	u := Universe{
		Rocks: make(map[Coordinate]bool),
		MaxX:  maxX,
		MaxY:  maxY,
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			switch char {
			case "#":
				u.Rocks[Coordinate{x, y}] = true
			case "S":
				u.Start = Location{
					Coordinate:         Coordinate{x, y},
					reachedInEvenSteps: false,
				}
			case ".":
				continue
			default:
				panic("No valid case found")
			}
		}
	}

	return u
}

func vectors() [][2]int {
	return [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
}
