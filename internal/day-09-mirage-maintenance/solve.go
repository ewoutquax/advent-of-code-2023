package day09miragemaintenance

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "09"

type Sequence struct {
	Numbers []int
}

func (s Sequence) ListDifferences() (next Sequence) {
	next.Numbers = make([]int, 0)
	for idx := 1; idx < len(s.Numbers); idx++ {
		next.Numbers = append(next.Numbers, s.Numbers[idx]-s.Numbers[idx-1])
	}

	return
}

type History struct {
	Sequences []Sequence
}

func (h History) PredictNextNumber() int {
	for h.Sequences[len(h.Sequences)-1].Numbers[len(h.Sequences[len(h.Sequences)-1].Numbers)-1] != 0 {
		h.Sequences = append(h.Sequences, h.Sequences[len(h.Sequences)-1].ListDifferences())
	}

	for idx := len(h.Sequences) - 1; idx > 0; idx-- {
		lastNumberIdxCurr := h.Sequences[idx].Numbers[len(h.Sequences[idx].Numbers)-1]
		lastNumberIdxPrev := h.Sequences[idx-1].Numbers[len(h.Sequences[idx-1].Numbers)-1]
		h.Sequences[idx-1].Numbers = append(h.Sequences[idx-1].Numbers, lastNumberIdxCurr+lastNumberIdxPrev)
	}

	return h.Sequences[0].Numbers[len(h.Sequences[0].Numbers)-1]
}

func (h History) PredictPreviousNumber() int {
	for h.Sequences[len(h.Sequences)-1].Numbers[len(h.Sequences[len(h.Sequences)-1].Numbers)-1] != 0 {
		h.Sequences = append(h.Sequences, h.Sequences[len(h.Sequences)-1].ListDifferences())
	}

	for idx := len(h.Sequences) - 1; idx > 0; idx-- {
		firstNumberIdxCurr := h.Sequences[idx].Numbers[0]
		firstNumberIdxPrev := h.Sequences[idx-1].Numbers[0]
		h.Sequences[idx-1].Numbers = append(
			[]int{firstNumberIdxPrev - firstNumberIdxCurr},
			h.Sequences[idx-1].Numbers...,
		)
	}

	return h.Sequences[0].Numbers[0]
}

type Universe struct {
	Histories []History
}

func (u Universe) SumNextPredictions() (sum int) {
	for _, h := range u.Histories {
		sum += h.PredictNextNumber()
	}

	return
}

func (u Universe) SumPreviousPredictions() (sum int) {
	for _, h := range u.Histories {
		sum += h.PredictPreviousNumber()
	}

	return
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	u := ParseInput(lines)
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, u.SumNextPredictions())
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	u := ParseInput(lines)
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, u.SumPreviousPredictions())
}

func ParseInput(lines []string) Universe {
	u := Universe{
		Histories: make([]History, 0),
	}

	for _, line := range lines {
		h := History{
			Sequences: make([]Sequence, 0),
		}
		s := Sequence{
			Numbers: make([]int, 0),
		}
		chars := strings.Split(line, " ")
		for _, char := range chars {
			s.Numbers = append(s.Numbers, utils.ConvStrToI(char))
		}
		h.Sequences = append(h.Sequences, s)
		u.Histories = append(u.Histories, h)
	}

	return u
}
