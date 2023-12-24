package day21stepcounter_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-21-step-counter"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var universe Universe = ParseInput(testInput())

	assert := assert.New(t)
	assert.Len(universe.Rocks, 40)
	assert.Equal(10, universe.MaxX)
	assert.Equal(10, universe.MaxY)
	assert.Equal(5, universe.Start.X)
	assert.Equal(5, universe.Start.Y)
}

func TestReachableLocations(t *testing.T) {
	u := ParseInput(testInput())

	reachableLocations := u.ReachableLocations(6)
	assert.Len(t, reachableLocations, 29)
}

func TestReachableLocationsInEvenSteps(t *testing.T) {
	u := ParseInput(testInput())

	reachableLocations := u.ReachableLocations(6)

	count := 0
	for _, inEvenSteps := range reachableLocations {
		if inEvenSteps {
			count++
		}
	}
	assert.Equal(t, 16, count)
}

func testInput() []string {
	return []string{
		"...........",
		".....###.#.",
		".###.##..#.",
		"..#.#...#..",
		"....#.#....",
		".##..S####.",
		".##..#...#.",
		".......##..",
		".##.#.####.",
		".##..##.##.",
		"...........",
	}
}
