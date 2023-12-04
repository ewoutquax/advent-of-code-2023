package day04scratchcards_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-04-scratchcards"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var c Card = ParseInput(testInput()[0])

	assert := assert.New(t)
	assert.Len(c.WinningNumbers, 5)
	assert.Len(c.PlayingNumbers, 8)
	assert.Equal(1, c.NrCopies)
}

func TestCardScore(t *testing.T) {
	testCases := map[string]int{
		testInput()[0]: 8,
		testInput()[1]: 2,
		testInput()[2]: 2,
		testInput()[3]: 1,
		testInput()[4]: 0,
		testInput()[5]: 0,
	}

	for line, expectedScore := range testCases {
		card := ParseInput(line)
		assert.Equal(t, expectedScore, card.CardScore())
	}
}

func TestWinWithCopies(t *testing.T) {
	u := BuildUniverse(testInput())
	u.WinWithCopies()
	assert.Equal(t, 30, u.CountCopies())
}

func testInput() []string {
	return []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
}
