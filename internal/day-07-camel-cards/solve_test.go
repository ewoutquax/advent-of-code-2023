package day07camelcards_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-07-camel-cards"
	"github.com/stretchr/testify/assert"
)

func TestRankHand(t *testing.T) {
	testCases := map[string]Strength{
		"AAAAA": FiveOfAKind,
		"AA8AA": FourOfAKind,
		"23332": FullHouse,
		"TTT98": ThreeOfAKind,
		"23432": TwoPair,
		"A23A4": Pair,
		"23456": HighCard,
	}

	for hand, expectedStrength := range testCases {
		assert.Equal(t, expectedStrength, HandStrength(hand, false), hand)
	}
}

func TestRankHandWithJoker(t *testing.T) {
	testCases := map[string]Strength{
		"32T3K": Pair,
		"T55J5": FourOfAKind,
		"KK677": TwoPair,
		"KTJJT": FourOfAKind,
		"QQQJA": FourOfAKind,
		"1123J": ThreeOfAKind,
		"54JA5": ThreeOfAKind,
		"4J42Q": ThreeOfAKind,
	}

	for hand, expectedStrength := range testCases {
		assert.Equal(t, expectedStrength, HandStrength(hand, true), hand)
	}
}

func TestParseInput(t *testing.T) {
	var hands []Hand = ParseInput(testInput())

	assert.Len(t, hands, 5)
	assert.Equal(t, "32T3K", hands[0].Cards)
	assert.Equal(t, 684, hands[1].Bid)
	assert.Equal(t, TwoPair, hands[2].Strength)
}

func TestRankHands(t *testing.T) {
	hands := ParseInput(testInput())
	ranked := RankHands(hands, false)

	fmt.Printf("ranked: %v\n", ranked)

	handByCards := make(map[string]Hand, len(hands))
	for _, hand := range ranked {
		handByCards[hand.Cards] = hand
	}
	assert := assert.New(t)
	assert.Equal(5, handByCards["QQQJA"].Rank)
	assert.Equal(1, handByCards["32T3K"].Rank)
}

func TestTotalWinning(t *testing.T) {
	hands := ParseInput(testInput())

	assert.Equal(t, 6440, TotalWinning(hands, false))
}

func TestTotalWinningWithJoker(t *testing.T) {
	hands := ParseInputWithJoker(testInput())

	assert.Equal(t, 5905, TotalWinning(hands, true))
}

func testInput() []string {
	return []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
}
