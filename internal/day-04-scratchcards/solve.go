package day04scratchcards

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "04"

type Card struct {
	CardId         int
	NrCopies       int
	WinningNumbers []int
	PlayingNumbers []int
}

func (c Card) nrWinningNumbers() (count int) {
	for _, number := range c.PlayingNumbers {
		if inSlice(c.WinningNumbers, number) {
			count++
		}
	}

	return
}

func (c Card) CardScore() int {
	power := c.nrWinningNumbers()

	if power == 0 {
		return 0
	} else {
		return int(math.Pow(float64(2), float64(power-1)))
	}
}

type Universe struct {
	Cards map[int]*Card
}

func (u *Universe) WinWithCopies() {
	for cardId := 1; cardId <= len(u.Cards); cardId++ {
		card := u.Cards[cardId]
		count := card.nrWinningNumbers()
		for idx := 1; idx <= count; idx++ {
			u.Cards[cardId+idx].NrCopies += card.NrCopies
		}
	}
}

func (u Universe) CountCopies() (count int) {
	for _, card := range u.Cards {
		count += card.NrCopies
	}

	return
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)

	var score int = 0
	for _, line := range lines {
		card := ParseInput(line)
		score += card.CardScore()
	}

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, score)

}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	u := BuildUniverse(lines)
	u.WinWithCopies()

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, u.CountCopies())
}

func BuildUniverse(lines []string) Universe {
	var universe = Universe{
		Cards: make(map[int]*Card, len(lines)),
	}

	for _, line := range lines {
		card := ParseInput(line)
		universe.Cards[card.CardId] = &card
	}

	return universe
}

func ParseInput(line string) (c Card) {
	parts := strings.Split(line, ": ")
	subparts := strings.Split(parts[0], " ")
	rawCardId := strings.TrimPrefix(subparts[len(subparts)-1], " ")
	c.CardId = utils.ConvStrToI(rawCardId)

	re := regexp.MustCompile(`(\d+)+`)
	numberParts := strings.Split(parts[1], " | ")

	for _, number := range re.FindAll([]byte(numberParts[0]), -1) {
		c.WinningNumbers = append(c.WinningNumbers, utils.ConvStrToI(string(number)))
	}
	for _, number := range re.FindAll([]byte(numberParts[1]), -1) {
		c.PlayingNumbers = append(c.PlayingNumbers, utils.ConvStrToI(string(number)))
	}

	c.NrCopies = 1

	return
}

func inSlice(haystack []int, needle int) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}

	return false
}
