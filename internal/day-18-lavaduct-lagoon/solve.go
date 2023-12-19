package day18lavaductlagoon

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "18"

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	instructions := ParseInput(utils.ReadFileAsLines(inputFile))

	myMap := BuildMap()
	myMap.ExecInstructions(instructions)
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, myMap.Surface())
}

func solvePart2(inputFile string) {
	instructions := ParseInputWithHex(utils.ReadFileAsLines(inputFile))

	myMap := BuildMap()
	myMap.ExecInstructions(instructions)
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, myMap.Surface())
}

type Direction string
type ColorCode string

const (
	Up    Direction = "U"
	Down  Direction = "D"
	Right Direction = "R"
	Left  Direction = "L"
)

type Instruction struct {
	Direction
	Length int
	Code   ColorCode
}

type Location struct {
	X int
	Y int
}

type Universe struct {
	Locations    []Location
	Instructions []Instruction
}

func (u *Universe) ExecInstructions(instructions []Instruction) {
	for _, instruction := range instructions {
		vector := map[Direction][2]int{
			Up:    {0, -1},
			Down:  {0, 1},
			Left:  {-1, 0},
			Right: {1, 0},
		}[instruction.Direction]

		currentLocation := u.Locations[len(u.Locations)-1]
		currentLocation.X += vector[0] * (instruction.Length)
		currentLocation.Y += vector[1] * (instruction.Length)

		u.Locations = append(u.Locations, currentLocation)
	}

	u.Instructions = instructions
}

func (m Universe) Surface() (surface int) {
	for idx := 0; idx < len(m.Locations)-1; idx++ {
		fmt.Printf("solve: Locations: %v -> %v\n", m.Locations[idx], m.Locations[idx+1])
		surface += m.Locations[idx].X * m.Locations[idx+1].Y
		surface -= m.Locations[idx].Y * m.Locations[idx+1].X
	}

	perimeter := 0
	for _, instruction := range m.Instructions {
		perimeter += instruction.Length
	}

	fmt.Printf("(abs(surface) / 2): %v\n", (abs(surface) / 2))
	fmt.Printf("(perimeter / 2): %v\n", (perimeter / 2))

	return abs(surface)/2 + perimeter/2 + 1
}

func BuildMap() Universe {
	m := Universe{
		Locations: make([]Location, 0),
	}

	m.Locations = []Location{{0, 0}}

	return m
}

func ParseInput(lines []string) []Instruction {
	var instructions = make([]Instruction, 0)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		i := Instruction{
			Direction: Direction(parts[0]),
			Length:    utils.ConvStrToI(parts[1]),
			Code:      ColorCode(parts[2][2 : len(parts[2])-1]),
		}
		instructions = append(instructions, i)
	}

	return instructions
}

func ParseInputWithHex(lines []string) []Instruction {
	var instructions = make([]Instruction, 0)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		i := Instruction{
			Direction: convHexToDirection(parts[2][2 : len(parts[2])-1]),
			Length:    convHexToLength(parts[2][2 : len(parts[2])-1]),
			Code:      ColorCode(parts[2][2 : len(parts[2])-1]),
		}
		instructions = append(instructions, i)
	}

	return instructions
}

func convHexToDirection(s string) Direction {
	chars := strings.Split(s, "")
	switch chars[len(chars)-1] {
	case "0":
		return Right
	case "1":
		return Down
	case "2":
		return Left
	case "3":
		return Up
	default:
		panic("No valid case found")
	}
}

func convHexToLength(s string) int {
	return convHexToInt(s[0 : len(s)-1])
}

func convHexToInt(hex string) int {
	result, _ := strconv.ParseInt(hex, 16, 64)
	return int(result)
}

func abs(number int) int {
	if number < 0 {
		return 0 - number
	}
	return number
}
