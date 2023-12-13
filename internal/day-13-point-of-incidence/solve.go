package day13pointofincidence

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "13"

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	u := ParseInput(blocks)
	fmt.Printf("Result of day-%s / part-1: %d\n", Day, u.SumMapScores())
}

func solvePart2(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	u := ParseInput(blocks)
	fmt.Printf("Result of day-%s / part-2: %d\n", Day, u.SumSmudgedMapScores())
}

type Location struct {
	X int
	Y int
}

type Mirror struct {
	AfterHorizontal  int
	BeforeHorizontal int
	AfterVertical    int
	BeforeVertical   int

	isHorizontalFound bool
	isVerticalFound   bool
}

func (m *Mirror) Score() (score int) {
	return m.BeforeHorizontal*100 + m.BeforeVertical
}

type Map struct {
	MaxX      int
	MaxY      int
	Locations []Location

	Mirror *Mirror
}

func (m Map) FindMirror() {
	m.findMirrorHorizontal()
	m.FindMirrorVertical()
}

func (m Map) FindMirrorWithSmudge() {
	m.findMirrorHorizontalWithSmudge()
	m.FindMirrorVerticalWithSmudge()
}

func (m Map) findMirrorHorizontalWithSmudge() {
	var ys = make(map[int][]int, 0)
	var maxY int = 0

	for _, loc := range m.Locations {
		if maxY < loc.Y {
			maxY = loc.Y
		}
		if _, exists := ys[loc.Y]; !exists {
			ys[loc.Y] = []int{}
		}
		ys[loc.Y] = append(ys[loc.Y], loc.X)
	}

	for before := 0; before < maxY; before++ {
		if mirrorsOnHorizontalWithSmudge(before, ys) {
			m.Mirror.AfterHorizontal = before
			m.Mirror.BeforeHorizontal = before + 1
			m.Mirror.isHorizontalFound = true

			return
		}
	}
}

func (m Map) FindMirrorVerticalWithSmudge() {
	var xs = make(map[int][]int, 0)
	var maxX int = 0

	for _, loc := range m.Locations {
		if maxX < loc.X {
			maxX = loc.X
		}
		if _, exists := xs[loc.X]; !exists {
			xs[loc.X] = []int{}
		}
		xs[loc.X] = append(xs[loc.X], loc.Y)
	}

	for before := 0; before < maxX; before++ {
		if mirrorsOnVerticalWithSmudge(before, xs) {
			m.Mirror.AfterVertical = before
			m.Mirror.BeforeVertical = before + 1
			m.Mirror.isVerticalFound = true

			return
		}
	}
}

func (m Map) findMirrorHorizontal() {
	var ys = make(map[int]string, 0)
	var ystrings = make(map[int][]string, 0)
	var maxY int = 0

	for _, loc := range m.Locations {
		if maxY < loc.Y {
			maxY = loc.Y
		}
		if _, exists := ystrings[loc.Y]; !exists {
			ystrings[loc.Y] = []string{}
		}
		ystrings[loc.Y] = append(ystrings[loc.Y], strconv.Itoa(loc.X))
	}

	for y, mystrings := range ystrings {
		ys[y] = strings.Join(mystrings, ",")
	}

	for before := 0; before < maxY; before++ {
		if mirrorsOnHorizontal(before, ys) {
			m.Mirror.AfterHorizontal = before
			m.Mirror.BeforeHorizontal = before + 1
			m.Mirror.isHorizontalFound = true

			return
		}
	}
}

func (m Map) FindMirrorVertical() {
	var xs = make(map[int]string, 0)
	var xstrings = make(map[int][]string, 0)
	var maxX int = 0

	for _, loc := range m.Locations {
		if maxX < loc.X {
			maxX = loc.X
		}
		if _, exists := xstrings[loc.X]; !exists {
			xstrings[loc.X] = []string{}
		}
		xstrings[loc.X] = append(xstrings[loc.X], strconv.Itoa(loc.Y))
	}

	for x, mystrings := range xstrings {
		xs[x] = strings.Join(mystrings, ",")
	}

	for before := 0; before < maxX; before++ {
		if mirrorsOnVertical(before, xs) {
			m.Mirror.AfterVertical = before
			m.Mirror.BeforeVertical = before + 1
			m.Mirror.isVerticalFound = true

			return
		}
	}
}

type Universe struct {
	Maps []Map
}

func (u Universe) SumMapScores() (sum int) {
	for _, m := range u.Maps {
		m.FindMirror()
		sum += m.Mirror.Score()
	}
	return
}

func (u Universe) SumSmudgedMapScores() (sum int) {
	for _, m := range u.Maps {
		m.FindMirrorWithSmudge()
		sum += m.Mirror.Score()
	}
	return
}

func ParseInput(blocks [][]string) Universe {
	u := Universe{
		Maps: make([]Map, 0, len(blocks)),
	}

	for _, block := range blocks {
		u.Maps = append(u.Maps, ParseMap(block))
	}

	return u
}

func ParseMap(lines []string) (m Map) {
	m.Locations = make([]Location, 0)
	m.Mirror = &Mirror{
		isHorizontalFound: false,
		isVerticalFound:   false,
	}

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				if m.MaxX < x {
					m.MaxX = x
				}
				if m.MaxY < y {
					m.MaxY = y
				}
				m.Locations = append(m.Locations, Location{x, y})
			}
		}
	}

	return
}

func MirrorsWithSmudge(part1, part2 []int) (bool, bool) {
	var mpart1 = make(map[int]bool, len(part1))
	for _, index := range part1 {
		mpart1[index] = true
	}
	var diff = make([]int, 0)
	for _, index := range part2 {
		if _, exists := mpart1[index]; exists {
			delete(mpart1, index)
		} else {
			diff = append(diff, index)
		}
	}
	mismatches := len(mpart1) + len(diff)
	switch mismatches {
	case 0:
		return true, false
	case 1:
		return true, true
	default:
		return false, false
	}
}

func mirrorsOnVerticalWithSmudge(x int, xs map[int][]int) bool {
	var nrSmudges int = 0
	for offset := 0; offset <= x; offset++ {
		mirrorLeft, existsLeft := xs[x-offset]
		mirrorRight, existsRight := xs[x+offset+1]
		if existsLeft && existsRight {
			matches, withSmudge := MirrorsWithSmudge(mirrorLeft, mirrorRight)
			if withSmudge {
				nrSmudges++
			}
			if !matches || nrSmudges > 1 {
				return false
			}
		}
	}

	return nrSmudges == 1
}

func mirrorsOnHorizontalWithSmudge(y int, ys map[int][]int) bool {
	var nrSmudges int = 0
	for offset := 0; offset <= y; offset++ {
		mirrorUp, existsUp := ys[y-offset]
		mirrorDown, existsDown := ys[y+offset+1]
		if existsUp && existsDown {
			matches, withSmudgge := MirrorsWithSmudge(mirrorUp, mirrorDown)
			if withSmudgge {
				nrSmudges++
			}
			if !matches || nrSmudges > 1 {
				return false
			}
		}
	}

	return nrSmudges == 1
}

func mirrorsOnVertical(x int, xs map[int]string) bool {
	for offset := 0; offset <= x; offset++ {
		mirrorLeft, existsLeft := xs[x-offset]
		mirrorRight, existsRight := xs[x+offset+1]
		if existsLeft && existsRight && mirrorLeft != mirrorRight {
			return false
		}
	}

	return true
}

func mirrorsOnHorizontal(y int, ys map[int]string) bool {
	for offset := 0; offset <= y; offset++ {
		mirrorUp, existsUp := ys[y-offset]
		mirrorDown, existsDown := ys[y+offset+1]
		if existsUp && existsDown && mirrorUp != mirrorDown {
			return false
		}
	}

	return true
}
