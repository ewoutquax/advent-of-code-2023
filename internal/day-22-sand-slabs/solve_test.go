package day22sandslabs_test

import (
	"fmt"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-22-sand-slabs"
	"github.com/stretchr/testify/assert"
)

func TestParseBrick(t *testing.T) {
	var brick *Brick = ParseBrick(testInput()[6])

	assert := assert.New(t)
	assert.Len(brick.Blocks, 2)
	assert.True(brick.Blocks[0].IsLowest)
	assert.False(brick.Blocks[1].IsLowest)
}

func TestParseInput(t *testing.T) {
	var universe Universe = ParseInput(testInput())

	assert.Len(t, universe.Bricks, 7)
}

func TestDropBricks(t *testing.T) {
	u := ParseInput(testInput())
	u.DropBricks()

	brick3 := u.Bricks[2]
	brick7 := u.Bricks[6]

	assert := assert.New(t)
	assert.Len(u.Bricks, 7)
	assert.Equal(2, brick3.Blocks[0].Z)
	assert.Equal(5, brick7.Blocks[0].Z)

	for i, brick := range u.Bricks {
		fmt.Printf("%d: brick: %v\n", i, brick)
	}
}

func TestConnectBricks(t *testing.T) {
	u := ParseInput(testInput())
	u.DropBricks()
	u.ConnectBricks()

	brickA := u.Bricks[0]
	brickB := u.Bricks[1]
	brickC := u.Bricks[2]
	brickD := u.Bricks[3]
	brickE := u.Bricks[4]
	brickG := u.Bricks[6]

	assert := assert.New(t)
	assert.Len(brickA.RestsOn, 0)
	assert.Len(brickB.RestsOn, 1)
	assert.Equal(brickA, brickB.RestsOn[0])
	assert.Len(brickC.RestsOn, 1)
	assert.Equal(brickA, brickC.RestsOn[0])
	assert.Len(brickD.RestsOn, 2)
	assert.Len(brickE.RestsOn, 2)
	assert.Len(brickG.RestsOn, 1)
}

func TestCountSoleSupporters(t *testing.T) {
	u := ParseInput(testInput())
	u.DropBricks()
	u.ConnectBricks()

	assert.Equal(t, 5, u.CountNotSoleSupporters())
}

func TestCountFallingBlocks(t *testing.T) {
	u := ParseInput(testInput())
	u.DropBricks()
	u.ConnectBricks()

	assert.Equal(t, 7, SumFallingBricks(&u))
}

func testInput() []string {
	return []string{
		"1,0,1~1,2,1",
		"0,0,2~2,0,2",
		"0,2,3~2,2,3",
		"0,0,4~0,2,4",
		"2,0,5~2,2,5",
		"0,1,6~2,1,6",
		"1,1,8~1,1,9",
	}
}
