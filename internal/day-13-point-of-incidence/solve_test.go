package day13pointofincidence_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-13-point-of-incidence"
	"github.com/stretchr/testify/assert"
)

func TestParseMape(t *testing.T) {
	var myMap Map = ParseMap(testInput()[0])

	assert.Len(t, myMap.Locations, 28)
	assert.Equal(t, 8, myMap.MaxX)
	assert.Equal(t, 6, myMap.MaxY)
}

func TestParseInput(t *testing.T) {
	var u Universe = ParseInput(testInput())

	assert.Len(t, u.Maps, 2)
}

func TestFindMirrorVertical(t *testing.T) {
	m := ParseMap(testInput()[0])

	m.FindMirror()

	assert := assert.New(t)
	assert.Equal(4, m.Mirror.AfterVertical)
	assert.Equal(5, m.Mirror.BeforeVertical)
	assert.Equal(5, m.Mirror.Score())
}

func TestFindMirrorVerticalBroken(t *testing.T) {
	m := ParseMap(testInputBrokenMap())

	m.FindMirror()

	assert := assert.New(t)
	assert.Equal(11, m.Mirror.AfterVertical)
	assert.Equal(12, m.Mirror.BeforeVertical)
	assert.Equal(12, m.Mirror.Score())
}

func TestFindMirrorHorizontal(t *testing.T) {
	m := ParseMap(testInput()[1])

	m.FindMirror()

	assert := assert.New(t)
	assert.Equal(3, m.Mirror.AfterHorizontal)
	assert.Equal(4, m.Mirror.BeforeHorizontal)
	assert.Equal(400, m.Mirror.Score())
}

func TestFindMirrorHorizontalWithSmudge(t *testing.T) {
	m := ParseMap(testInput()[0])

	m.FindMirrorWithSmudge()

	assert := assert.New(t)
	assert.Equal(2, m.Mirror.AfterHorizontal)
	assert.Equal(4, m.Mirror.BeforeHorizontal)
	assert.Equal(300, m.Mirror.Score())
}

func TestFindMirrorVerticalWithSmudge(t *testing.T) {
	m := ParseMap(testInput()[1])

	m.FindMirrorWithSmudge()

	assert := assert.New(t)
	assert.Equal(0, m.Mirror.AfterHorizontal)
	assert.Equal(1, m.Mirror.BeforeHorizontal)
	assert.Equal(100, m.Mirror.Score())
}

func TestSumMapScores(t *testing.T) {
	var universe Universe = ParseInput(testInput())
	assert.Equal(t, 405, universe.SumMapScores())
}

func TestMatchWithSmudge(t *testing.T) {
	assert := assert.New(t)

	var part1, part2, part3 []int
	var matches, withSmudge bool

	part1 = []int{1, 2, 4}
	part2 = []int{1, 2, 4, 5}
	part3 = []int{1, 2, 5}
	matches, withSmudge = MirrorsWithSmudge(part1, part2)
	assert.True(matches)
	assert.True(withSmudge)

	matches, withSmudge = MirrorsWithSmudge(part2, part1)
	assert.True(matches)
	assert.True(withSmudge)

	matches, withSmudge = MirrorsWithSmudge(part1, part1)
	assert.True(matches)
	assert.False(withSmudge)

	matches, withSmudge = MirrorsWithSmudge(part1, part3)
	assert.False(matches)
	assert.False(withSmudge)
}

func testInput() [][]string {
	return [][]string{{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	}, {
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}}
}

func testInputBrokenMap() []string {
	return []string{
		"##..#...#.###",
		".#.##.#...###",
		"##.#..###..##",
		"..........###",
		".#.###.....##",
		"#...#.#.#..##",
		"#.##.#####...",
		".#.##..###...",
		"...##..##.###",
		"##...........",
		"#####..####..",
		"#####..####..",
		"#............",
	}
}
