package day10pipemaze_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-10-pipe-maze"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var universe Universe = ParseInput(testInputSmallestMaze())

	topLeftPipe := universe.Pipes[Location{1, 1}]

	assert := assert.New(t)
	assert.Len(universe.Pipes, 8)
	assert.Equal("F", topLeftPipe.Symbol)
	assert.Len(topLeftPipe.Neighbours, 2)

	assert.Equal("|", topLeftPipe.Neighbours[0].Symbol)
	assert.Equal("-", topLeftPipe.Neighbours[1].Symbol)
}

func TestStartingPipe(t *testing.T) {
	assert := assert.New(t)

	u := ParseInput(testInputSmallMaze())
	startPipe := u.Pipes[Location{1, 1}]
	bottomRightPipe := u.Pipes[Location{4, 4}]
	leftCenterPipe := u.Pipes[Location{0, 2}]
	centerRightPipe := u.Pipes[Location{1, 2}]

	assert.True(startPipe.IsStart)
	assert.False(bottomRightPipe.IsStart)
	assert.False(leftCenterPipe.IsStart)
	assert.False(centerRightPipe.IsStart)
}

func TestLoopLength(t *testing.T) {
	u1 := ParseInput(testInputSmallMaze())
	assert.Equal(t, 7, u1.LoopLength())

	u2 := ParseInput(testInputComplexMaze())
	assert.Equal(t, 15, u2.LoopLength())
}

func TestStepsFarthersFromStart(t *testing.T) {
	u1 := ParseInput(testInputSmallMaze())
	assert.Equal(t, 4, u1.StepsFarthestFromStart())

	u2 := ParseInput(testInputComplexMaze())
	assert.Equal(t, 8, u2.StepsFarthestFromStart())

}

func TestSumEnclosedTiles(t *testing.T) {
	u := ParseInput(testInputEnclosedSimple())

	u.StepsFarthestFromStart() // Find all the pipes in the loop
	u.CSIEnhance()
	u.MarkReachableTiles()
	u.Draw()

	assert.Equal(t, 4, u.SumUnreachableTiles())
}

func testInputSmallestMaze() []string {
	return []string{
		".....",
		".F-7.",
		".|.|.",
		".L-J.",
		".....",
	}
}

func testInputSmallMaze() []string {
	return []string{
		"-L|F7",
		"7S-7|",
		"L|7||",
		"-L-J|",
		"L|-JF",
	}
}

func testInputComplexMaze() []string {
	return []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}
}

func testInputEnclosedSimple() []string {
	return []string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	}
}
