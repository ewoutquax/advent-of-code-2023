package day02cubeconundrum

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "02"

type CubeSet struct {
	CubeRed   int // Number of RED cubes in this draw
	CubeBlue  int // Number of BLUE cubes in this draw
	CubeGreen int // Number of GREEN cubes in this draw
}

func (d CubeSet) Power() int {
	return d.CubeRed * d.CubeBlue * d.CubeGreen
}

type Game struct {
	Id    int
	Draws []CubeSet
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	var validSet = CubeSet{
		CubeRed:   12,
		CubeGreen: 13,
		CubeBlue:  14,
	}

	lines := utils.ReadFileAsLines(inputFile)
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, SumValidGames(lines, validSet))
}

func solvePart2(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumMinCubePower(lines))

}

func SumMinCubePower(lines []string) (sum int) {
	for _, line := range lines {
		sum += MinCubesPerGame(ParseLine(line)).Power()
	}

	return
}

func MinCubesPerGame(game Game) CubeSet {
	var minDraw = CubeSet{
		CubeRed:   0,
		CubeBlue:  0,
		CubeGreen: 0,
	}

	for _, draw := range game.Draws {
		if minDraw.CubeRed < draw.CubeRed {
			minDraw.CubeRed = draw.CubeRed
		}
		if minDraw.CubeGreen < draw.CubeGreen {
			minDraw.CubeGreen = draw.CubeGreen
		}
		if minDraw.CubeBlue < draw.CubeBlue {
			minDraw.CubeBlue = draw.CubeBlue
		}
	}

	return minDraw
}

func SumValidGames(lines []string, validSet CubeSet) (sum int) {
	for _, line := range lines {
		var game Game = ParseLine(line)
		if ValidGame(game, validSet) {
			sum += game.Id
		}
	}

	return
}

func ValidGame(game Game, validSet CubeSet) bool {
	for _, draw := range game.Draws {
		if draw.CubeRed > validSet.CubeRed ||
			draw.CubeBlue > validSet.CubeBlue ||
			draw.CubeGreen > validSet.CubeGreen {
			return false
		}
	}

	return true
}

func ParseLine(line string) (game Game) {
	parts := strings.Split(line, ": ")
	subparts := strings.Split(parts[0], " ")
	game.Id = utils.ConvStrToI(subparts[1])

	draws := strings.Split(parts[1], "; ")
	for _, rawDraw := range draws {
		var cubeSet = CubeSet{
			CubeRed:   0,
			CubeBlue:  0,
			CubeGreen: 0,
		}
		for _, cube := range strings.Split(rawDraw, ", ") {
			cubeParts := strings.Split(cube, " ")
			amount := utils.ConvStrToI(cubeParts[0])
			switch cubeParts[1] {
			case "red":
				cubeSet.CubeRed = amount
			case "blue":
				cubeSet.CubeBlue = amount
			case "green":
				cubeSet.CubeGreen = amount
			default:
				panic("unknown color")
			}
		}
		game.Draws = append(game.Draws, cubeSet)
	}

	return
}
