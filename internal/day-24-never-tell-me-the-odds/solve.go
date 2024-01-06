package day24nevertellmetheodds

import (
	"fmt"
	"regexp"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "24"

type Coordinate [2]float32

type Hailstone struct {
	X int
	Y int
	Z int

	Vx int
	Vy int
	Vz int
}

func init() {
	register.Day(Day+"a", solvePart1)
	// register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	hailstones := ParseInput(lines)

	count := HailstonesCollidingInsideArea(hailstones, 200000000000000, 400000000000000)
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, count)
}

func solvePart2(inputFile string) {
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, 0)
}

func HailstonesCollidingInsideArea(hailstones []Hailstone, minValue, maxValue int) (count int) {
	for len(hailstones) > 0 {
		currentHailstone := hailstones[0]
		hailstones = hailstones[1:]

		for _, comparingHailstone := range hailstones {
			if hailstoneCollideInFutureInArea(
				currentHailstone, comparingHailstone,
				float32(minValue), float32(maxValue),
			) {
				count++
			}
		}
	}

	return
}

func hailstoneCollideInFutureInArea(currentHailstone, comparingHailstone Hailstone, minValue, maxValue float32) bool {
	areColliding, coordinates := FindFutureIntersection(currentHailstone, comparingHailstone)

	return areColliding &&
		coordinates[0] >= minValue && coordinates[0] < maxValue &&
		coordinates[1] >= minValue && coordinates[1] < maxValue
}

// X1 + T1(Vx1) = X2 + T2(Vx2) =>
// T1(Vx1) = T2(Vx2) + X2 - X1 =>
// T1 = (T2(Vx2) + X2 - X1) / Vx1
//
// T1 = (T2(Vx2) + X2 - X1) / Vx1 && T1 = (T2(Vy2) + Y2 - Y1) / Vy1 =>
// (T2(Vx2) + X2 - X1) / Vx1 = (T2(Vy2) + Y2 - Y1) / Vy1 =>
// (T2(Vx2) + X2 - X1) * Vy1 = (T2(Vy2) + Y2 - Y1) * Vx1 =>
// (T2(Vx2) + X2 - X1) * Vy1 = (T2(Vy2) + Y2 - Y1) * Vx1 =>
// T2(Vx2 * Vy1) + X2 * Vy1 - X1 * Vy1 = T2(Vy2 * Vx1) + Y2 * Vx1 - Y1 * Vx1 =>
// T2(Vx2 * Vy1) - T2(Vy2 * Vx1) = Y2 * Vx1 - Y1 * Vx1 - X2 * Vy1 + X1 * Vy1 =>
// T2(Vx2 * Vy1 - Vy2 * Vx1) = Y2 * Vx1 - Y1 * Vx1 - X2 * Vy1 + X1 * Vy1 =>
// T2 = (Y2 * Vx1 - Y1 * Vx1 - X2 * Vy1 + X1 * Vy1) / (Vx2 * Vy1 - Vy2 * Vx1)
func FindFutureIntersection(hailstoneA, hailstoneB Hailstone) (bool, Coordinate) {
	// Find the moment, where hailstoneB will intersect with hailstoneA
	var T1, T2 float32

	divider := float32(hailstoneB.Vx*hailstoneA.Vy - hailstoneB.Vy*hailstoneA.Vx)

	if divider != 0 {
		T2 = float32(
			(hailstoneB.Y-hailstoneA.Y)*hailstoneA.Vx+
				(hailstoneA.X-hailstoneB.X)*hailstoneA.Vy) / divider

		T1 = (T2*float32(hailstoneB.Vx) + float32(hailstoneB.X-hailstoneA.X)) / float32(hailstoneA.Vx)

		if T1 >= 0 && T2 >= 0 {
			// Both timestamps are in the future; calculate the X and Y-coordinate
			return true, Coordinate{
				float32(hailstoneB.X) + T2*float32(hailstoneB.Vx),
				float32(hailstoneB.Y) + T2*float32(hailstoneB.Vy),
			}
		}
	}

	return false, Coordinate{0, 0}
}

func ParseInput(lines []string) []Hailstone {
	var hailstones []Hailstone = make([]Hailstone, 0, len(lines))

	for _, line := range lines {
		hailstones = append(hailstones, ParseLine(line))
	}

	return hailstones
}

func ParseLine(line string) Hailstone {
	ex := regexp.MustCompile(`(-?\d+)+`)
	matches := ex.FindAllString(line, -1)

	return Hailstone{
		X:  utils.ConvStrToI(matches[0]),
		Y:  utils.ConvStrToI(matches[1]),
		Z:  utils.ConvStrToI(matches[2]),
		Vx: utils.ConvStrToI(matches[3]),
		Vy: utils.ConvStrToI(matches[4]),
		Vz: utils.ConvStrToI(matches[5]),
	}
}
