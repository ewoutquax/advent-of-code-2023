package day12hotsprings_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-12-hot-springs"
	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	var record Record = ParseLine(testInput()[0])

	assert.Len(t, record.Springs, 3)
	assert.Equal(t, 0, record.Springs[0].Start)
	assert.Equal(t, 1, record.Springs[0].Size)
	assert.Equal(t, 4, record.Springs[len(record.Springs)-1].Start)
	assert.Equal(t, 3, record.Springs[len(record.Springs)-1].Size)
}

func TestSpringsMatchMask(t *testing.T) {
	var record Record = ParseLine(testInput()[0])

	assert.True(t, record.SpringsMatchMask())
}

func TestSumRecordArrangements(t *testing.T) {
	assert.Equal(t, 21, SumRecordArrangements(testInput()))
}

func TestCountArrangements(t *testing.T) {
	var record Record
	testCases := map[int]int{
		0: 1,
		1: 4,
		2: 1,
		3: 1,
		4: 4,
		5: 10,
	}

	for idxInput, expectedCount := range testCases {
		record = ParseLine(testInput()[idxInput])
		assert.Equal(t, expectedCount, CountArrangements(record), idxInput)
	}
}

func TestCountUnfoldedArrangements(t *testing.T) {
	var record Record
	testCases := map[int]int{
		0: 1,
		1: 16384,
		2: 1,
		3: 16,
		4: 2500,
		5: 506250,
	}

	for idxInput, expectedCount := range testCases {
		record = ParseLineUnfolded(testInput()[idxInput])
		assert.Equal(t, expectedCount, CountArrangements(record), idxInput)
	}
}

func testInput() []string {
	return []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		//  '0----|----+---- 1,3,1,6',
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
		// '.+++.++.+...'
	}
}
