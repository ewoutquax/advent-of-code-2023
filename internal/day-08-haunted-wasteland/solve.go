package day08hauntedwasteland

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "08"

type Node struct {
	Label     string
	NextNodes [2]*Node

	IsStart bool
	IsEnd   bool
}

type Runner struct {
	NrSteps     int
	CurrentNode Node
}

type Universe struct {
	Directions string
	Nodes      map[string]*Node
	Runners    []Runner
}

func (u *Universe) AddSingleRunner() {
	u.Runners = append(u.Runners, Runner{
		NrSteps:     0,
		CurrentNode: *u.Nodes["AAA"],
	})
}

func (u *Universe) AddMultipleRunners() {
	for _, node := range u.Nodes {
		if node.IsStart {
			u.Runners = append(u.Runners, Runner{
				NrSteps:     0,
				CurrentNode: *u.Nodes[node.Label],
			})
		}
	}
}

func (r *Runner) FollowDirection(u Universe) {
	var nextNode *Node

	idxDirection := r.NrSteps % len(u.Directions)
	direction := string(u.Directions[idxDirection])

	switch direction {
	case "L":
		nextNode = r.CurrentNode.NextNodes[0]
	case "R":
		nextNode = r.CurrentNode.NextNodes[1]
	default:
		panic("No valid case found")
	}

	r.CurrentNode = *nextNode
	r.NrSteps++
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	universe := ParseInput(blocks)
	universe.AddSingleRunner()
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, StepsTillEnd(universe))
}

func solvePart2(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	universe := ParseInput(blocks)
	universe.AddMultipleRunners()
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, StepsAllTillEnd(universe))
}

func StepsTillEnd(u Universe) int {
	r := u.Runners[0]
	for !r.CurrentNode.IsEnd {
		r.FollowDirection(u)
	}

	return r.NrSteps
}

func StepsAllTillEnd(u Universe) int {
	var steps = make([]int, 0)

	for _, r := range u.Runners {
		for !r.CurrentNode.IsEnd {
			r.FollowDirection(u)
		}
		steps = append(steps, r.NrSteps)
	}

	fmt.Printf("steps: %v\n", steps)

	var lcm int = 1
	for _, nrSteps := range steps {
		lcm = Lcm(lcm, nrSteps)
	}

	return lcm
}

func ParseInput(blocks [][]string) Universe {
	return Universe{
		Directions: parseDirections(blocks[0]),
		Nodes:      parseNodes(blocks[1]),
		Runners:    make([]Runner, 0),
	}

	// for _, node := range u.Nodes {
	//  if node.IsStart {
	//    u.Runners = append(u.Runners, Runner{
	//      NrSteps:     0,
	//      CurrentNode: *u.Nodes[node.Label],
	//    })
	//  }
	// }
}

func parseDirections(lines []string) string {
	return lines[0]
}

func parseNodes(lines []string) map[string]*Node {
	var nodes = make(map[string]*Node, len(lines))

	// First, read all the labels
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		label := parts[0]
		nodes[label] = &Node{
			Label:     label,
			NextNodes: [2]*Node{},
			IsStart:   string(label[2]) == "A",
			IsEnd:     string(label[2]) == "Z",
		}
	}

	// Now, read all the destinations
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		subparts := strings.Split(parts[1], ", ")

		label := parts[0]
		labelLeft := strings.TrimPrefix(subparts[0], "(")
		labelRight := strings.TrimSuffix(subparts[1], ")")

		nodes[label].NextNodes[0] = nodes[labelLeft]
		nodes[label].NextNodes[1] = nodes[labelRight]
	}

	return nodes
}

func Lcm(int1, int2 int) int {
	return (int1 * int2) / Gcd(int1, int2)
}

func Gcd(numerator, dividor int) int {
	var remainder int = numerator % dividor

	fmt.Printf("solve: numerator / dividor = remainder: %d / %d = %d\n", numerator, dividor, remainder)

	if remainder != 0 {
		return Gcd(dividor, remainder)
	}

	fmt.Printf("gcd: %v\n", dividor)

	return dividor
}
