package day14parabolicrefractiondish_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-14-parabolic-refraction-dish"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var universe Universe = ParseInput(testInput())

	var ctrRollableRock int = 0
	for _, r := range universe.Rocks {
		if r.Rollable {
			ctrRollableRock++
		}
	}

	assert.Equal(t, 18, ctrRollableRock)
	assert.Len(t, universe.Rocks, 35)
	assert.Equal(t, 9, universe.MaxX)
	assert.Equal(t, 9, universe.MaxY)
}

func TestRollRocksNorth(t *testing.T) {
	u := ParseInput(testInput())
	u.RollRocks(North)
	u.Draw()

	assert.Len(t, u.Rocks, 35)
	assert.Equal(t, 136, u.Weight())
}

func TestTiltCycle(t *testing.T) {
	u := ParseInput(testInput())

	u.CycleTilts()
	u.Draw()
	u.CycleTilts()
	u.Draw()
	u.CycleTilts()
	u.Draw()

	assert.Len(t, u.Rocks, 35)
	assert.Equal(t, 136, u.Weight())
}

func TestRepeatCycles(t *testing.T) {
	u := ParseInput(testInput())

	weight := RepeatCycles(1000000000, u)

	assert.Equal(t, 64, weight)
}

func testInput() []string {
	return []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}
}
