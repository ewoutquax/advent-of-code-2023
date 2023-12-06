package day05ifyougaveaseedafertilizer_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-05-if-you-gave-a-seed-a-fertilizer"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)
	var universe Universe = ParseInput(testInput())

	assert.Len(universe.Seeds, 4)
	assert.Equal(79, universe.Seeds[0])

	assert.Len(universe.Conversions, 7)
	assert.Equal(98, universe.Conversions[0].Ranges[0].From)
	assert.Equal(-48, universe.Conversions[0].Ranges[0].Delta)
}

func TestConvertCategory(t *testing.T) {
	u := ParseInput(testInput())

	var convertedCategory CategoryAmount
	convertedCategory, _ = u.ConvertCategory(CategoryAmount{"seed", 79})
	assert.Equal(t, Category("soil"), convertedCategory.Category)
	assert.Equal(t, 81, convertedCategory.Amount)

	var convertedCategory2 CategoryAmount
	convertedCategory2, _ = u.ConvertCategory(CategoryAmount{"fertilizer", 53})
	assert.Equal(t, Category("water"), convertedCategory2.Category)
	assert.Equal(t, 49, convertedCategory2.Amount)
}

func TestConvertCategoryToLocation(t *testing.T) {
	u := ParseInput(testInput())

	location, _ := ConvertSeedToLocation(79, u)
	assert.Equal(t, 82, location)
}

func TestFindNearestLocation(t *testing.T) {
	u := ParseInput(testInput())

	assert.Equal(t, 35, FindNearestLocation(u))
}

func TestFindNearestLocationWithRanges(t *testing.T) {
	u := ParseInputWithSeedRanges(testInput())

	assert.Equal(t, 46, FindNearestLocationWithRanges(u))
}

func testInput() [][]string {
	return [][]string{{
		"seeds: 79 14 55 13",
	}, {
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
	}, {
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
	}, {
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
	}, {
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
	}, {
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
	}, {
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
	}, {
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}}
}
