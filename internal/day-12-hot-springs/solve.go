package day12hotsprings

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "12"

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumRecordArrangements(lines))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumUnfoldedRecordArrangements(lines))
}

type Spring struct {
	Start int
	Size  int
}

func (s Spring) SpringMatchesMask(mask string) bool {
	for offset := 0; offset < s.Size; offset++ {
		char := mask[s.Start+offset]
		if char == '.' {
			return false
		}
	}

	// The submask right of the spring can not contain required fields, not covered by this rightmost spring
	if strings.Contains(mask[s.Start+s.Size:], "#") {
		return false
	}
	// Spring can't have a required field directly on its left
	if s.Start > 0 && mask[s.Start-1] == '#' {
		return false
	}

	return true
}

func (s Spring) canMoveRight(mask string) bool {
	return s.Start+s.Size < len(mask)
}

type Record struct {
	Mask                  string
	Springs               []Spring
	masksWithCombinations map[string]int // Cache known outcomes of masks with nr of springs
}

func (r Record) cacheKey() string {
	return fmt.Sprintf("%s-%d", r.Mask, len(r.Springs))
}

func (r Record) SubRecord(cutOffAt int) Record {
	subRecord := Record{
		masksWithCombinations: r.masksWithCombinations,
		Springs:               make([]Spring, 0),
	}

	if cutOffAt > 0 {
		subRecord.Mask = r.Mask[:cutOffAt]
	}

	for idx := 0; idx < len(r.Springs)-1; idx++ {
		subRecord.Springs = append(subRecord.Springs, r.Springs[idx])
	}

	return subRecord
}

func (r Record) SpringsMatchMask() bool {
	currentMask := string(r.cacheKey())

	for idx := len(r.Springs) - 1; idx >= 0; idx-- {
		s := r.Springs[idx]
		if !s.SpringMatchesMask(currentMask) {
			return false
		}
		currentMask = currentMask[:s.Start]
	}

	if strings.Contains(currentMask, "#") {
		return false
	}

	return true
}

func SumRecordArrangements(lines []string) (sum int) {
	for _, line := range lines {
		record := ParseLine(line)
		sum += CountArrangements(record)
	}
	return
}

func SumUnfoldedRecordArrangements(lines []string) (sum int) {

	for _, line := range lines {
		record := ParseLineUnfolded(line)
		sum += CountArrangements(record)
	}
	return
}

func CountArrangements(r Record) (count int) {
	if nrCombinations, exists := r.masksWithCombinations[r.cacheKey()]; exists {
		return nrCombinations
	}

	if len(r.Springs) == 0 {
		if strings.Contains(r.cacheKey(), "#") {
			return 0
		} else {
			return 1
		}
	}

	rightMostSpring := r.Springs[len(r.Springs)-1]

	var doContinue bool = true
	for doContinue {
		if rightMostSpring.SpringMatchesMask(r.Mask) {
			count += CountArrangements(r.SubRecord(rightMostSpring.Start - 1))
		}

		if rightMostSpring.canMoveRight(r.Mask) {
			rightMostSpring.Start += 1
		} else {
			doContinue = false
		}
	}
	r.masksWithCombinations[r.cacheKey()] = count

	return
}

func ParseLineUnfolded(line string) (r Record) {
	r.masksWithCombinations = make(map[string]int, 0)
	r.Springs = make([]Spring, 0)

	parts := strings.Split(line, " ")

	fullCounts := strings.Join([]string{
		parts[1],
		parts[1],
		parts[1],
		parts[1],
		parts[1],
	}, ",")
	counts := strings.Split(fullCounts, ",")

	cumulative := 0
	for _, count := range counts {
		r.Springs = append(r.Springs, Spring{
			Start: cumulative,
			Size:  utils.ConvStrToI(count),
		})
		cumulative += utils.ConvStrToI(count) + 1
	}

	r.Mask = strings.Join([]string{
		parts[0],
		parts[0],
		parts[0],
		parts[0],
		parts[0],
	}, "?")

	return
}

func ParseLine(line string) (r Record) {
	r.masksWithCombinations = make(map[string]int, 0)
	r.Springs = make([]Spring, 0)

	parts := strings.Split(line, " ")
	counts := strings.Split(parts[1], ",")
	cumulative := 0
	for _, count := range counts {
		r.Springs = append(r.Springs, Spring{
			Start: cumulative,
			Size:  utils.ConvStrToI(count),
		})
		cumulative += utils.ConvStrToI(count) + 1
	}

	r.Mask = parts[0]

	return
}
