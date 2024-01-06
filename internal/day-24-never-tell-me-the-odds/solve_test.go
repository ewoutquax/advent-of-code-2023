package day24nevertellmetheodds_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-24-never-tell-me-the-odds"
	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	var hailstone Hailstone = ParseLine(testInput()[0])

	assert := assert.New(t)
	assert.Equal(19, hailstone.X)
	assert.Equal(13, hailstone.Y)
	assert.Equal(30, hailstone.Z)
	assert.Equal(-2, hailstone.Vx)
	assert.Equal(1, hailstone.Vy)
	assert.Equal(-2, hailstone.Vz)
}

func TestFindFutureIntersection(t *testing.T) {
	assert := assert.New(t)

	hailstoneA := Hailstone{
		X:  19,
		Y:  13,
		Z:  30,
		Vx: -2,
		Vy: 1,
		Vz: -2,
	}
	hailstoneB := Hailstone{
		X:  18,
		Y:  19,
		Z:  22,
		Vx: -1,
		Vy: -1,
		Vz: -2,
	}

	hailstoneC := Hailstone{
		X:  20,
		Y:  19,
		Z:  15,
		Vx: 1,
		Vy: -5,
		Vz: -3,
	}

	hailstoneD := Hailstone{
		X:  20,
		Y:  25,
		Z:  34,
		Vx: -2,
		Vy: -2,
		Vz: -4,
	}

	var result bool
	var coordinates Coordinate

	result, coordinates = FindFutureIntersection(hailstoneA, hailstoneB)
	assert.True(result)
	assert.InDelta(14.333, coordinates[0], 0.001)
	assert.InDelta(15.333, coordinates[1], 0.001)

	// Cross in the past for hailstoneA
	result, coordinates = FindFutureIntersection(hailstoneA, hailstoneC)
	assert.False(result)

	// Parallel lines return false
	result, coordinates = FindFutureIntersection(hailstoneB, hailstoneD)
	assert.False(result)
}

func TestHailstormsCollidingInsideArea(t *testing.T) {
	var hailstones []Hailstone = ParseInput(testInput())

	count := HailstonesCollidingInsideArea(hailstones, 7, 23)

	assert.Equal(t, 2, count)
}

func testInput() []string {
	return []string{
		"19, 13, 30 @ -2,  1, -2",
		"18, 19, 22 @ -1, -1, -2",
		"20, 25, 34 @ -2, -2, -4",
		"12, 31, 28 @ -1, -2, -1",
		"20, 19, 15 @  1, -5, -3",
	}
}
