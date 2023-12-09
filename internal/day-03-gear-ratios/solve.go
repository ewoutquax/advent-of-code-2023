package day03gearratios

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "03"
const GearSymbol string = "*"

type Location struct {
	X int
	Y int
}

type Symbol struct {
	Location
	Value string

	NextToNumbers []Number
}

func (s Symbol) isGear() bool {
	return s.Value == GearSymbol
}

type Number struct {
	Location
	Value  int
	Length int
}

type Universe struct {
	Numbers []Number  // List of all the numbers
	Symbols []*Symbol // List of all the symbols

	SymbolLocations map[Location]*Symbol // List of all the Symbols, indexed by their location
}

func (u Universe) SumEngineParts() (sum int) {
	for _, number := range u.Numbers {
		if number.IsNextToSymbol(u) {
			sum += number.Value
		}
	}

	return
}

func (u Universe) SumGearRatios() (sum int) {
	u.LinkSymbolsToNumbers()
	for _, symbol := range u.SymbolLocations {
		if symbol.isGear() && len(symbol.NextToNumbers) == 2 {
			sum += symbol.NextToNumbers[0].Value * symbol.NextToNumbers[1].Value
		}
	}
	return
}

func (u Universe) LinkSymbolsToNumbers() {
	for _, number := range u.Numbers {
		for _, symbol := range number.NextToSymbols(u) {
			symbol.NextToNumbers = append(symbol.NextToNumbers, number)
		}
	}
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	universe := ParseInput(utils.ReadFileAsLines(inputFile))
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, universe.SumEngineParts())
}

func solvePart2(inputFile string) {
	universe := ParseInput(utils.ReadFileAsLines(inputFile))
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, universe.SumGearRatios())
}

func (n Number) IsNextToSymbol(u Universe) bool {
	return len(n.NextToSymbols(u)) != 0
}

func (n Number) NextToSymbols(u Universe) (symbols []*Symbol) {
	var loc Location

	// left side
	loc.Y = n.Y
	loc.X = n.X - 1
	if s, exists := u.SymbolLocations[loc]; exists {
		symbols = append(symbols, s)
	}
	// right side
	loc.X = n.X + n.Length
	if s, exists := u.SymbolLocations[loc]; exists {
		symbols = append(symbols, s)
	}

	// Check the top- and bottomside of the number
	for vX := -1; vX <= n.Length; vX++ {
		loc.X = n.X + vX

		// top side
		loc.Y = n.Y - 1
		if s, exists := u.SymbolLocations[loc]; exists {
			symbols = append(symbols, s)
		}
		// bottom side
		loc.Y = n.Y + 1
		if s, exists := u.SymbolLocations[loc]; exists {
			symbols = append(symbols, s)
		}
	}

	return
}

func ParseInput(lines []string) (u Universe) {
	var currentNumber = Number{}

	// Find all numbers
	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if strings.Contains("0123456789", char) {
				if currentNumber.Value == 0 {
					currentNumber.X = x
					currentNumber.Y = y
					currentNumber.Length = 0
				}
				currentNumber.Value *= 10
				currentNumber.Value += utils.ConvStrToI(char)
				currentNumber.Length++
			} else {
				if currentNumber.Value != 0 {
					u.Numbers = append(u.Numbers, currentNumber)
					currentNumber = Number{}
				}
				// And find the symbols
				if char != "." {
					u.Symbols = append(u.Symbols, &Symbol{Location{x, y}, char, []Number{}})
				}
			}
		}
		// End of the line also marks end of a number
		if currentNumber.Value != 0 {
			u.Numbers = append(u.Numbers, currentNumber)
			currentNumber = Number{}
		}
	}

	u.SymbolLocations = make(map[Location]*Symbol, len(u.Symbols))
	for _, symbol := range u.Symbols {
		u.SymbolLocations[symbol.Location] = symbol
	}

	return
}
