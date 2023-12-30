package day23alongwalk_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-23-a-long-walk"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var universe Universe = ParseInput(testInput())

	assert := assert.New(t)
	assert.Len(universe.Locations, 213)

	coordinateStart := Coordinate{1, 0}
	coordinateEnd := Coordinate{21, 22}

	assert.True(universe.Locations[coordinateStart].IsStart)
	assert.False(universe.Locations[coordinateStart].IsEnd)
	assert.False(universe.Locations[coordinateEnd].IsStart)
	assert.True(universe.Locations[coordinateEnd].IsEnd)

	assert.Equal(DirectionNone, universe.Locations[coordinateStart].ForcedDirection)
	assert.Equal(DirectionNone, universe.Locations[coordinateEnd].ForcedDirection)

	coordinateSlopeRight := Coordinate{10, 3}
	assert.False(universe.Locations[coordinateSlopeRight].IsStart)
	assert.False(universe.Locations[coordinateSlopeRight].IsEnd)
	assert.Equal(DirectionRight, universe.Locations[coordinateSlopeRight].ForcedDirection)
}

func TestFindLongestPath(t *testing.T) {
	u := ParseInput(testInput())

	path := u.FindLongestPath()

	assert.Equal(t, path.NrSteps, 95)

	assert.True(t, false)
}

func testInput() []string {
	return []string{
		"#.#####################",
		"#.......#########...###",
		"#######.#########.#.###",
		"###.....#.>.>.###.#.###",
		"###v#####.#v#.###.#.###",
		"###.>...#.#.#.....#...#",
		"###v###.#.#.#########.#",
		"###...#.#.#.......#...#",
		"#####.#.#.#######.#.###",
		"#.....#.#.#.......#...#",
		"#.#####.#.#.#########v#",
		"#.#...#...#...###...>.#",
		"#.#.#v#######v###.###v#",
		"#...#.>.#...>.>.#.###.#",
		"#####v#.#.###v#.#.###.#",
		"#.....#...#...#.#.#...#",
		"#.#########.###.#.#.###",
		"#...###...#...#...#.###",
		"###.###.#.###v#####v###",
		"#...#...#.#.>.>.#.>.###",
		"#.###.###.#.###.#.#v###",
		"#.....###...###...#...#",
		"#####################.#",
	}
}
