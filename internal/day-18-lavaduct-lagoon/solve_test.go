package day18lavaductlagoon_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-18-lavaduct-lagoon"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var instructions []Instruction = ParseInput(testInput())

	assert := assert.New(t)
	assert.Len(instructions, 14)
	assert.Equal(Right, instructions[0].Direction)
	assert.Equal(6, instructions[0].Length)
	assert.Equal(ColorCode("70c710"), instructions[0].Code)
}

func TestBuildMap(t *testing.T) {
	var myMap Universe = BuildMap()

	assert.Len(t, myMap.Locations, 1)
	assert.Equal(t, 0, myMap.Locations[0].X)
	assert.Equal(t, 0, myMap.Locations[0].Y)
}

func TestExecInstructions(t *testing.T) {
	instructions := ParseInput(testInput())

	myMap := BuildMap()
	myMap.ExecInstructions(instructions)

	assert.Len(t, myMap.Locations, 15)
}

func TestMapSurface(t *testing.T) {
	instructions := ParseInput(testInput())

	myMap := BuildMap()
	myMap.ExecInstructions(instructions)

	assert.Equal(t, 62, myMap.Surface())
}

func TestParseInputWithHex(t *testing.T) {
	var instructions []Instruction = ParseInputWithHex(testInput())

	assert := assert.New(t)
	assert.Len(instructions, 14)
	assert.Equal(Right, instructions[0].Direction)
	assert.Equal(461937, instructions[0].Length)

	assert.True(false)
}

func TestMapSurfaceWithHex(t *testing.T) {
	instructions := ParseInputWithHex(testInput())

	myMap := BuildMap()
	myMap.ExecInstructions(instructions)

	assert.Equal(t, 952408144115, myMap.Surface())
}

func testInput() []string {
	return []string{
		"R 6 (#70c710)",
		"D 5 (#0dc571)",
		"L 2 (#5713f0)",
		"D 2 (#d2c081)",
		"R 2 (#59c680)",
		"D 2 (#411b91)",
		"L 5 (#8ceee2)",
		"U 2 (#caa173)",
		"L 1 (#1b58a2)",
		"U 2 (#caa171)",
		"R 2 (#7807d2)",
		"U 3 (#a77fa3)",
		"L 2 (#015232)",
		"U 2 (#7a21e3)",
	}
}
