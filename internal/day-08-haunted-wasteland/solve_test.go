package day08hauntedwasteland_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-08-haunted-wasteland"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	assert := assert.New(t)
	var universe Universe = ParseInput(testInput())
	universe.AddSingleRunner()

	assert.Len(universe.Runners, 1)
	assert.Equal(0, universe.Runners[0].NrSteps)
	assert.Equal("AAA", universe.Runners[0].CurrentNode.Label)

	assert.Len(universe.Directions, 2)
	assert.Equal("L", string(universe.Directions[1]))

	assert.Len(universe.Nodes, 7)
	assert.Equal("AAA", universe.Nodes["AAA"].Label)
	assert.Len(universe.Nodes["AAA"].NextNodes, 2)
	assert.Equal("BBB", universe.Nodes["AAA"].NextNodes[0].Label)

	assert.True(universe.Nodes["AAA"].IsStart)
	assert.True(universe.Nodes["ZZZ"].IsEnd)
	assert.False(universe.Nodes["AAA"].IsEnd)
	assert.False(universe.Nodes["ZZZ"].IsStart)
}

func TestFollowDirection(t *testing.T) {
	u := ParseInput(testInput())
	u.AddSingleRunner()

	runner := u.Runners[0]
	runner.FollowDirection(u)

	assert := assert.New(t)
	assert.Equal(1, runner.NrSteps)
	assert.Equal("CCC", runner.CurrentNode.Label)
}

func TestWalkToEnd(t *testing.T) {
	u1 := ParseInput(testInput())
	u1.AddSingleRunner()
	assert.Equal(t, 2, StepsTillEnd(u1))

	u2 := ParseInput(testInput2())
	u2.AddSingleRunner()
	assert.Equal(t, 6, StepsTillEnd(u2))
}

func TestAddMultipleRunners(t *testing.T) {
	u := ParseInput(testInput3())
	u.AddMultipleRunners()

	assert.Len(t, u.Runners, 2)
}

func TestWalkAllToEnd(t *testing.T) {
	u := ParseInput(testInput3())
	u.AddMultipleRunners()

	assert.Equal(t, 6, StepsAllTillEnd(u))
}

func testInput() [][]string {
	return [][]string{{
		"RL",
	}, {
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}}
}

func testInput2() [][]string {
	return [][]string{{
		"LLR",
	}, {
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}}
}

func testInput3() [][]string {
	return [][]string{{
		"LR",
	}, {
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	}}
}
