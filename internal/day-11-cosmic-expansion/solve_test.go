package day11cosmicexpansion_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-11-cosmic-expansion"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var u Universe = ParseInput(testInput(), 1)

	fmt.Printf("u: %v\n", u)

	assert := assert.New(t)
	assert.Len(u.Galaxies, 9)

	var exists bool
	_, exists = u.Galaxies[Location{0, 2}]
	assert.True(exists, "galaxy at 0,2")

	_, exists = u.Galaxies[Location{1, 5}]
	assert.False(exists, "moved galaxy at 1,5")
	_, exists = u.Galaxies[Location{1, 6}]
	assert.True(exists, "expanded galaxy at 1,6")

	_, exists = u.Galaxies[Location{4, 0}]
	assert.True(exists, "expanded galaxy at 4,0")
	_, exists = u.Galaxies[Location{3, 0}]
	assert.False(exists, "moved galaxy at 3,0")
}

func TestSumGalaxyDistances(t *testing.T) {
	u := ParseInput(testInput(), 1)

	assert.Equal(t, 374, u.SumGalaxyDistance())
}

func TestSumGalaxyDistancesTimes10(t *testing.T) {
	u := ParseInput(testInput(), 9)

	assert.Equal(t, 1030, u.SumGalaxyDistance())
}

func TestSumGalaxyDistancesTimes100(t *testing.T) {
	u := ParseInput(testInput(), 99)

	assert.Equal(t, 8410, u.SumGalaxyDistance())
}

func testInput() []string {
	return []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
}
