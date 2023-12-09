package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/ewoutquax/advent-of-code-2023/internal/day-01-trebuchet"
	_ "github.com/ewoutquax/advent-of-code-2023/internal/day-02-cube-conundrum"
	_ "github.com/ewoutquax/advent-of-code-2023/internal/day-03-gear-ratios"
	_ "github.com/ewoutquax/advent-of-code-2023/internal/day-04-scratchcards"
	_ "github.com/ewoutquax/advent-of-code-2023/internal/day-05-if-you-gave-a-seed-a-fertilizer"
	_ "github.com/ewoutquax/advent-of-code-2023/internal/day-06-wait-for-it"
	_ "github.com/ewoutquax/advent-of-code-2023/internal/day-07-camel-cards"
	_ "github.com/ewoutquax/advent-of-code-2023/internal/day-08-haunted-wasteland"
	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
)

func main() {
	for _, puzzle := range getPuzzles() {
		register.ExecDay(puzzle)
	}
}

func getPuzzles() (puzzles []string) {
	var allPuzzles []string = register.GetAllDays()

	selection := readUserInput(fmt.Sprintf("Which puzzle to run %s:\n", allPuzzles))
	switch selection {
	case "":
		latestPuzzle := allPuzzles[len(allPuzzles)-1]
		fmt.Printf("Running latest puzzle: %s\n\n", latestPuzzle)
		puzzles = []string{latestPuzzle}
	case "all":
		fmt.Printf("Running all puzzles\n\n")
		puzzles = allPuzzles
	default:
		fmt.Printf("Running selected puzzle: '%s'\n\n", selection)
		puzzles = []string{selection}
	}

	return
}

func readUserInput(question string) string {
	if len(os.Args) == 2 {
		return os.Args[1]
	}

	fmt.Printf("%s", question)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return strings.Trim(text, "\n")
}
