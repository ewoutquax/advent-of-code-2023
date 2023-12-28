package day22sandslabs

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "22"

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	u := ParseInput(utils.ReadFileAsLines(inputFile))
	u.DropBricks()
	u.ConnectBricks()

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, u.CountNotSoleSupporters())
}

func solvePart2(inputFile string) {
	u := ParseInput(utils.ReadFileAsLines(inputFile))
	u.DropBricks()
	u.ConnectBricks()

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, SumFallingBricks(&u))
}

type Coordinate struct {
	X int
	Y int
	Z int
}

func (c Coordinate) toKey() BlockKey {
	return BlockKey(fmt.Sprintf("%d,%d,%d", c.X, c.Y, c.Z))
}

type Block struct {
	Coordinate
	IsLowest bool
}

type BlockKey string

func (b Block) toKey() BlockKey { return b.Coordinate.toKey() }

type Brick struct {
	Blocks []Block

	Supports []*Brick // Which bricks are directly above by this brick
	RestsOn  []*Brick // Which bricks are directly below this brick

	isSoleSupporter bool // The brick above this one rests only on this block
	willFall        bool // When a certain brick is removed, this one will fall as well
}

type Universe struct {
	Bricks []*Brick
}

func (u *Universe) DropBricks() {
	occupiedCoordinates := make(map[BlockKey]bool, 0)
	var anyBrickHasDropped bool = true

	// Collect all block-coordinates into a single map
	// for easy checking if a block can fall into a coordinate
	for _, brick := range u.Bricks {
		for _, block := range brick.Blocks {
			occupiedCoordinates[block.toKey()] = true
		}
	}

	// Keep dropping blocks, until none will move anymore
	for anyBrickHasDropped {
		anyBrickHasDropped = false

		for _, brick := range u.Bricks {
			var canBrickDrop bool = true

			for canBrickDrop {
				for _, block := range brick.Blocks {
					if block.IsLowest {
						if block.Z <= 1 {
							canBrickDrop = false
						}
						key := Coordinate{X: block.X, Y: block.Y, Z: block.Z - 1}.toKey()
						if _, occupied := occupiedCoordinates[key]; occupied {
							canBrickDrop = false
						}
					}
				}

				// All the blocks at the bottom layer of the brick have nothing underneath: the brick can drop
				if canBrickDrop {
					anyBrickHasDropped = true
					var newBlocks = make([]Block, 0)
					for _, block := range brick.Blocks {
						delete(occupiedCoordinates, block.toKey())
						block.Coordinate.Z--
						occupiedCoordinates[block.toKey()] = true
						newBlocks = append(newBlocks, block)
					}
					copy(brick.Blocks, newBlocks)
				}
			}
		}
	}
}

func (u *Universe) ConnectBricks() {
	// Build a list of the coordinates of each block, linked to their brick
	var coorBrick = make(map[BlockKey]*Brick)
	for _, brick := range u.Bricks {
		for _, block := range brick.Blocks {
			coorBrick[block.toKey()] = brick
		}
	}

	for _, brick := range u.Bricks {
		var restings = make(map[*Brick]bool, 0)

		for _, block := range brick.Blocks {
			if block.IsLowest {
				coorBelow := Coordinate{
					X: block.X,
					Y: block.Y,
					Z: block.Z - 1,
				}
				brickBelow, exists := coorBrick[coorBelow.toKey()]
				if exists {
					restings[brickBelow] = true
				}
			}
		}

		for brickBelow := range restings {
			brick.RestsOn = append(brick.RestsOn, brickBelow)
			brickBelow.Supports = append(brickBelow.Supports, brick)
		}
	}

	u.markSoleSupporters()
}

func (u *Universe) markSoleSupporters() {
	// When a brick rests on only one other brick, then that other brick is a sole supporter
	for _, brick := range u.Bricks {
		if len(brick.RestsOn) == 1 {
			brick.RestsOn[0].isSoleSupporter = true
		}
	}
}

func (u *Universe) CountNotSoleSupporters() (count int) {
	// Count all bricks, that are NOT sole supporters
	for _, b := range u.Bricks {
		if !b.isSoleSupporter {
			count++
		}
	}

	return
}

func SumFallingBricks(u *Universe) (sum int) {
	for _, removedBrick := range u.Bricks {
		for _, b := range u.Bricks {
			b.willFall = false
		}

		removedBrick.willFall = true
		var fallingBricks = make([]*Brick, 0)
		fallingBricks = append(fallingBricks, removedBrick)

		var newFallers bool = true
		for newFallers {
			newFallers = false
			for _, fallingBrick := range fallingBricks {
				for _, potentialFallingBrick := range fallingBrick.Supports {
					var willBlockFall bool = true
					for _, brickBelowPotentialFallingBrick := range potentialFallingBrick.RestsOn {
						if brickBelowPotentialFallingBrick.willFall == false {
							willBlockFall = false
						}
					}
					if willBlockFall && potentialFallingBrick.willFall == false {
						potentialFallingBrick.willFall = true
						fallingBricks = append(fallingBricks, potentialFallingBrick)
						newFallers = true
					}
				}
			}
		}

		sum += len(fallingBricks) - 1
	}

	return
}

func ParseInput(lines []string) (u Universe) {
	u.Bricks = make([]*Brick, 0)

	for _, line := range lines {
		u.Bricks = append(u.Bricks, ParseBrick(line))
	}

	return
}

func ParseBrick(line string) *Brick {
	var brick = Brick{
		Blocks:          make([]Block, 0),
		RestsOn:         make([]*Brick, 0),
		Supports:        make([]*Brick, 0),
		isSoleSupporter: false,
	}

	fromTo := strings.Split(line, "~")
	var coordinatesFrom, coordinatesTo []int

	for _, char := range strings.Split(fromTo[0], ",") {
		coordinatesFrom = append(coordinatesFrom, utils.ConvStrToI(char))
	}
	for _, char := range strings.Split(fromTo[1], ",") {
		coordinatesTo = append(coordinatesTo, utils.ConvStrToI(char))
	}

	minZ := coordinatesFrom[2]
	if minZ > coordinatesTo[2] {
		minZ = coordinatesTo[2]
	}
	for x := coordinatesFrom[0]; x <= coordinatesTo[0]; x++ {
		for y := coordinatesFrom[1]; y <= coordinatesTo[1]; y++ {
			for z := coordinatesFrom[2]; z <= coordinatesTo[2]; z++ {
				brick.Blocks = append(brick.Blocks, Block{
					Coordinate: Coordinate{x, y, z},
					IsLowest:   z == minZ,
				})
			}
		}
	}

	return &brick
}
