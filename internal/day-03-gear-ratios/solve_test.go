package day03gearratios_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-03-gear-ratios"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	u := ParseInput(testInput())

	assert := assert.New(t)
	assert.Equal("day03gearratios.Universe", fmt.Sprintf("%s", reflect.TypeOf(u)))
	assert.Len(u.Numbers, 10)
	assert.Equal(467, u.Numbers[0].Value)
	assert.Len(u.Symbols, 6)
	assert.Equal("*", u.Symbols[0].Value)
	assert.Len(u.SymbolLocations, 6)
}

func TestIsNextToSymbol(t *testing.T) {
	u := ParseInput(testInput())

	assert := assert.New(t)
	assert.True(u.Numbers[0].IsNextToSymbol(u))  // via left side
	assert.True(u.Numbers[3].IsNextToSymbol(u))  // via bottom
	assert.False(u.Numbers[5].IsNextToSymbol(u)) // isn't next to a symbol
}

func TestSumEngineParts(t *testing.T) {
	u := ParseInput(testInput())

	assert.Equal(t, 4361, u.SumEngineParts())
}

func TestSumGearRation(t *testing.T) {
	u := ParseInput(testInput())

	assert.Equal(t, 467835, u.SumGearRatios())
}

func testInput() []string {
	return []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
}
