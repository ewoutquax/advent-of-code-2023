package day01trebuchet

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "01"

type NumberMatching map[string]string

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	var lines []string = utils.ReadFileAsLines(inputFile)
	matchings := NumbersBase()
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumCalibrations(lines, matchings))
}

func solvePart2(inputFile string) {
	var lines []string = utils.ReadFileAsLines(inputFile)
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumCalibrations(lines, MatchingsExtend()))
}

func SumCalibrations(lines []string, matchings NumberMatching) (total int) {
	for _, line := range lines {
		total += ExtractCalibration(line, matchings)
	}

	return
}

func ExtractCalibration(line string, matchings NumberMatching) int {
	var first, last string

	for idx := 0; idx < len(line); idx++ {
		substr := line[idx:]
		if found, number := findMatch(substr, matchings); found {
			if first == "" {
				first = number
			}
			last = number
		}
	}

	return utils.ConvStrToI(first + last)
}

func findMatch(char string, matchings NumberMatching) (bool, string) {
	for key, value := range matchings {
		if strings.Index(char, key) == 0 {
			return true, value
		}
	}
	return false, ""
}

func MatchingsExtend() NumberMatching {
	matchings := NumbersBase()
	for k, v := range NumbersWord() {
		matchings[k] = v
	}

	return matchings
}

func NumbersBase() NumberMatching {
	return NumberMatching{
		"0": "0",
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
		"5": "5",
		"6": "6",
		"7": "7",
		"8": "8",
		"9": "9",
	}
}

func NumbersWord() NumberMatching {
	return NumberMatching{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
}
