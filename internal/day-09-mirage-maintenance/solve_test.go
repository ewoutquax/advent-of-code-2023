package day09miragemaintenance_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-09-mirage-maintenance"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var universe Universe = ParseInput(testInput())

	assert := assert.New(t)
	assert.Len(universe.Histories, 3)
	assert.Equal(0, universe.Histories[0].Sequences[0].Numbers[0])
	assert.Equal(45, universe.Histories[len(universe.Histories)-1].Sequences[0].Numbers[len(universe.Histories[len(universe.Histories)-1].Sequences[0].Numbers)-1])
}

func TestNextSequence(t *testing.T) {
	u := ParseInput(testInput())
	s := u.Histories[0].Sequences[0]

	var next Sequence = s.ListDifferences()

	assert := assert.New(t)
	assert.Len(next.Numbers, 5)

	for _, number := range next.Numbers {
		assert.Equal(3, number)
	}
}

func TestPredictNextValue(t *testing.T) {
	u := ParseInput(testInput())

	assert := assert.New(t)
	assert.Equal(18, u.Histories[0].PredictNextNumber())
	assert.Equal(28, u.Histories[1].PredictNextNumber())
	assert.Equal(68, u.Histories[2].PredictNextNumber())
}

func TestSumNextPredictions(t *testing.T) {
	u := ParseInput(testInput())

	assert.Equal(t, 114, u.SumNextPredictions())
}

func TestSumPreviousPredictions(t *testing.T) {
	u := ParseInput(testInput())

	assert.Equal(t, 2, u.SumPreviousPredictions())
}

func TestPredictPreviousNumber(t *testing.T) {
	u := ParseInput(testInput())

	assert := assert.New(t)
	assert.Equal(5, u.Histories[2].PredictPreviousNumber())
}

func testInput() []string {
	return []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
}
