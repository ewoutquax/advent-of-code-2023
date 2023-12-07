package day07camelcards

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "07"

type Strength uint

const CardValues string = "AKQJT98765432"
const CardValuesWithJoker string = "AKQT98765432J"

const (
	Unknown Strength = iota
	HighCard
	Pair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards    string
	Strength Strength
	Bid      int
	Rank     int
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	hands := ParseInput(utils.ReadFileAsLines(inputFile))
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, TotalWinning(hands, false))
}

func solvePart2(inputFile string) {
	hands := ParseInputWithJoker(utils.ReadFileAsLines(inputFile))
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, TotalWinning(hands, true))
}

func TotalWinning(hands []Hand, withJoker bool) (total int) {
	ranked := RankHands(hands, withJoker)

	for _, h := range ranked {
		total += h.Rank * h.Bid
	}

	return
}

func RankHands(hands []Hand, withJoker bool) []Hand {
	var cardValues string

	if withJoker {
		cardValues = CardValuesWithJoker
	} else {
		cardValues = CardValues
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Strength < hands[j].Strength {
			return true
		}
		if hands[i].Strength > hands[j].Strength {
			return false
		}
		var idx int = 0
		for hands[i].Cards[idx] == hands[j].Cards[idx] {
			idx++
		}
		cardValueI := strings.Index(cardValues, string(hands[i].Cards[idx]))
		cardValueJ := strings.Index(cardValues, string(hands[j].Cards[idx]))

		return cardValueI > cardValueJ
	})

	var rankedHands []Hand
	for i, h := range hands {
		h.Rank = i + 1
		rankedHands = append(rankedHands, h)
	}

	return rankedHands
}

func ParseInput(lines []string) (hands []Hand) {
	for _, line := range lines {
		parts := strings.Split(line, " ")
		hands = append(hands, Hand{
			Cards:    parts[0],
			Strength: HandStrength(parts[0], false),
			Bid:      utils.ConvStrToI(parts[1]),
			Rank:     0,
		})
	}

	return
}

func ParseInputWithJoker(lines []string) (hands []Hand) {
	for _, line := range lines {
		parts := strings.Split(line, " ")
		hands = append(hands, Hand{
			Cards:    parts[0],
			Strength: HandStrength(parts[0], true),
			Bid:      utils.ConvStrToI(parts[1]),
			Rank:     0,
		})
	}

	return
}

func HandStrength(hand string, withJoker bool) Strength {
	var cardCounts = make(map[string]int, 5)

	for _, card := range strings.Split(hand, "") {
		if _, exists := cardCounts[card]; !exists {
			cardCounts[card] = 0
		}
		cardCounts[card]++
	}

	var counts = []int{0}
	var nrJokers int = 0
	for card, count := range cardCounts {
		if withJoker && card == "J" {
			nrJokers += count
		} else {
			counts = append(counts, count)
		}
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	counts[0] += nrJokers

	switch {
	case counts[0] == 5:
		return FiveOfAKind
	case counts[0] == 4:
		return FourOfAKind
	case counts[0] == 3 && counts[1] == 2:
		return FullHouse
	case counts[0] == 3:
		return ThreeOfAKind
	case counts[0] == 2 && counts[1] == 2:
		return TwoPair
	case counts[0] == 2:
		return Pair
	default:
		return HighCard
	}
}
