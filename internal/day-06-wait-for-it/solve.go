package day06waitforit

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "06"

type Race struct {
	Time           int // Duration of the race in milliseconds
	RecordDistance int // target distance to beat
}

func (r Race) DistanceForSpeedUp(speedUpTime int) int {
	speed := speedUpTime
	raceTime := r.Time - speedUpTime

	return speed * raceTime
}

func (r Race) CountWins() (count int) {
	for time := 0; time <= r.Time; time++ {
		if r.DistanceForSpeedUp(time) > r.RecordDistance {
			count++
		}
	}

	return
}

type Universe struct {
	Races []Race
}

func (u Universe) ErrorMargin() (margin int) {
	var minTime, maxTime int
	margin = 1

	for _, r := range u.Races {
		minTime = 0
		for r.DistanceForSpeedUp(minTime) <= r.RecordDistance {
			minTime++
		}
		maxTime = r.Time
		for r.DistanceForSpeedUp(maxTime) <= r.RecordDistance {
			maxTime--
		}
		margin *= (maxTime - minTime + 1)
	}

	return
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInput(lines)
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, universe.ErrorMargin())
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	universe := ParseInputSpaceless(lines)
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, universe.ErrorMargin())
}

func ParseInput(lines []string) (u Universe) {
	re := regexp.MustCompile(`(\d+)+`)

	var times, distances []int
	for _, time := range re.FindAll([]byte(lines[0]), -1) {
		times = append(times, utils.ConvStrToI((string(time))))
	}
	for _, distance := range re.FindAll([]byte(lines[1]), -1) {
		distances = append(distances, utils.ConvStrToI((string(distance))))
	}

	for idx := 0; idx < len(times); idx++ {
		u.Races = append(u.Races, Race{
			Time:           times[idx],
			RecordDistance: distances[idx],
		})
	}

	return
}

func ParseInputSpaceless(lines []string) (u Universe) {
	time := strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", "")
	distance := strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", "")

	u.Races = append(u.Races, Race{
		Time:           utils.ConvStrToI(time),
		RecordDistance: utils.ConvStrToI(distance),
	})

	return
}
