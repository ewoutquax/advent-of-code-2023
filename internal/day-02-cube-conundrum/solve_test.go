package day02cubeconundrum_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-02-cube-conundrum"
	"github.com/stretchr/testify/assert"
)

func TestSumValidGames(t *testing.T) {
	var validSet = CubeSet{
		CubeRed:   12,
		CubeBlue:  13,
		CubeGreen: 14,
	}

	assert.Equal(t, 8, SumValidGames(testInput(), validSet))
}

func TestValidGame(t *testing.T) {
	var validSet = CubeSet{
		CubeRed:   12,
		CubeBlue:  13,
		CubeGreen: 14,
	}

	assert := assert.New(t)
	assert.True(ValidGame(ParseLine(testInput()[0]), validSet))
	assert.True(ValidGame(ParseLine(testInput()[1]), validSet))
	assert.False(ValidGame(ParseLine(testInput()[2]), validSet))
	assert.False(ValidGame(ParseLine(testInput()[3]), validSet))
	assert.True(ValidGame(ParseLine(testInput()[4]), validSet))
}

func TestParseLine(t *testing.T) {
	line := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	expectedResult := Game{
		Id: 1,
		Draws: []CubeSet{
			{CubeRed: 4, CubeBlue: 3, CubeGreen: 0},
			{CubeRed: 1, CubeBlue: 6, CubeGreen: 2},
			{CubeRed: 0, CubeBlue: 0, CubeGreen: 2},
		},
	}

	assert.Equal(t, expectedResult, ParseLine(line))
}

func TestMinCubesPerGame(t *testing.T) {
	expectedDraw := CubeSet{
		CubeRed:   4,
		CubeBlue:  6,
		CubeGreen: 2,
	}

	assert.Equal(t, expectedDraw, MinCubesPerGame(ParseLine(testInput()[0])))
}

func TestDrawPower(t *testing.T) {
	draw := MinCubesPerGame(ParseLine(testInput()[0]))

	assert.Equal(t, 48, draw.Power())
}

func TestSumMinCubePower(t *testing.T) {
	assert.Equal(t, 2286, SumMinCubePower(testInput()))
}

func testInput() []string {
	return []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
}
