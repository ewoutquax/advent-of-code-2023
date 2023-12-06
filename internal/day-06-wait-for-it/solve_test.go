package day06waitforit_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-06-wait-for-it"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var universe Universe = ParseInput(testInput())

	assert.Len(t, universe.Races, 3)
	assert.Equal(t, 7, universe.Races[0].Time)
	assert.Equal(t, 200, universe.Races[2].RecordDistance)
}

func TestParseInputSpaceless(t *testing.T) {
	var universe Universe = ParseInputSpaceless(testInput())

	assert.Len(t, universe.Races, 1)
	assert.Equal(t, 71530, universe.Races[0].Time)
	assert.Equal(t, 940200, universe.Races[0].RecordDistance)
}

func TestRaceDistanceForSpeedUp(t *testing.T) {
	u := ParseInput(testInput())

	testCases := map[int]int{
		0: 0,
		1: 6,
		2: 10,
		3: 12,
		4: 12,
		5: 10,
		6: 6,
		7: 0,
	}

	for speedUpTime, expectedDistance := range testCases {
		assert.Equal(t, expectedDistance, u.Races[0].DistanceForSpeedUp(speedUpTime))
	}
}

func TestCountWins(t *testing.T) {
	u := ParseInput(testInput())

	assert.Equal(t, 4, u.Races[0].CountWins())
}

func TestErrorMargin(t *testing.T) {
	u := ParseInput(testInput())

	assert.Equal(t, 288, u.ErrorMargin())
}

func TestErrorMarginSpaceless(t *testing.T) {
	u := ParseInputSpaceless(testInput())

	assert.Equal(t, 71503, u.ErrorMargin())
}

func testInput() []string {
	return []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
}
