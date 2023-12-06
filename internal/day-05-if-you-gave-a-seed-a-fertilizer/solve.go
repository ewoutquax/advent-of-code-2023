package day05ifyougaveaseedafertilizer

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "05"

type Category string

const MaxInt int = int(^uint(0) >> 1)

const (
	CategorySeed        Category = "seed"
	CategorySoil        Category = "soil"
	CategoryFertilizer  Category = "fertilizer"
	CategoryWater       Category = "water"
	CategoryLight       Category = "light"
	CategoryTemperature Category = "temperature"
	CategoryHumidity    Category = "humidity"
	CategoryLocation    Category = "location"
)

type CategoryAmount struct {
	Category Category
	Amount   int
}

type Range struct {
	From   int
	To     int
	Length int
	Delta  int
}

type Conversion struct {
	From   Category
	To     Category
	Ranges []Range
}

type Universe struct {
	Seeds          []int   // Starting values
	SeedsWithRange []Range // For when we use ranges

	Conversions []Conversion // The conversion tables
}

func (u Universe) ConvertCategory(ca CategoryAmount) (CategoryAmount, int) {
	var rangeStopsIn int

	for _, conversion := range u.Conversions {
		if conversion.From == ca.Category {
			targetAmount := ca.Amount
			rangeStopsIn = MaxInt
			for _, r := range conversion.Ranges {
				if ca.Amount < r.From && rangeStopsIn > r.From-ca.Amount {
					rangeStopsIn = r.From - ca.Amount
				}
				if ca.Amount >= r.From && ca.Amount <= r.To {
					rangeStopsIn = r.To - ca.Amount
					targetAmount += r.Delta
				}
			}
			return CategoryAmount{
				Category: conversion.To,
				Amount:   targetAmount,
			}, rangeStopsIn
		}
	}

	panic("No conversion found")
}

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	universe := ParseInput(blocks)

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, FindNearestLocation(universe))
}

func solvePart2(inputFile string) {
	blocks := utils.ReadFileAsBlocks(inputFile)
	universe := ParseInputWithSeedRanges(blocks)

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, FindNearestLocationWithRanges(universe))
}

func FindNearestLocation(u Universe) (nearest int) {
	for _, seed := range u.Seeds {
		location, _ := ConvertSeedToLocation(seed, u)
		if nearest == 0 || nearest > location {
			nearest = location
		}

	}

	return
}

func FindNearestLocationWithRanges(u Universe) (nearest int) {
	var seedCurrent int

	nearest = -1
	for _, r := range u.SeedsWithRange {
		seedCurrent = r.From

		for seedCurrent <= r.To {
			location, rangeStopsIn := ConvertSeedToLocation(seedCurrent, u)

			if rangeStopsIn < 1 {
				seedCurrent++
			} else {
				seedCurrent += rangeStopsIn
			}

			if nearest == -1 || nearest > location {
				nearest = location
			}
		}
	}

	return
}

func ConvertSeedToLocation(seed int, u Universe) (int, int) {
	ca := CategoryAmount{
		Category: CategorySeed,
		Amount:   seed,
	}

	var minRangeStopsIn, rangeStopsIn int
	minRangeStopsIn = ^int(0)
	for ca.Category != CategoryLocation {
		ca, rangeStopsIn = u.ConvertCategory(ca)
		if minRangeStopsIn == -1 || minRangeStopsIn > rangeStopsIn {
			minRangeStopsIn = rangeStopsIn
		}
	}

	return ca.Amount, minRangeStopsIn
}

func ParseInput(blocks [][]string) (u Universe) {
	u.Seeds = parseSeeds(blocks[0])

	for idx := 1; idx < len(blocks); idx++ {
		u.Conversions = append(u.Conversions, parseConversion(blocks[idx]))
	}

	return
}

func ParseInputWithSeedRanges(blocks [][]string) (u Universe) {
	u.SeedsWithRange = parseSeedsWithRanges(blocks[0])

	for idx := 1; idx < len(blocks); idx++ {
		u.Conversions = append(u.Conversions, parseConversion(blocks[idx]))
	}

	return
}

func parseConversion(lines []string) (c Conversion) {
	fromTo := parseFromTo(lines[0])

	c.From = fromTo[0]
	c.To = fromTo[1]

	for idx := 1; idx < len(lines); idx++ {
		nmbrs := strings.Split(lines[idx], " ")
		from := utils.ConvStrToI(nmbrs[1])
		length := utils.ConvStrToI(nmbrs[2])
		r := Range{
			From:   from,
			Length: length,
			To:     from + length - 1,
			Delta:  utils.ConvStrToI(nmbrs[0]) - from,
		}

		c.Ranges = append(c.Ranges, r)
	}

	return
}

func parseFromTo(line string) [2]Category {
	parts := strings.Split(line, " ")
	categoryParts := strings.Split(parts[0], "-to-")

	return [2]Category{
		Category(categoryParts[0]),
		Category(categoryParts[1]),
	}
}

func parseSeeds(lines []string) (seeds []int) {
	line := lines[0]
	nmbrs := strings.Split(line, " ")

	for idx := 1; idx < len(nmbrs); idx++ {
		seeds = append(seeds, utils.ConvStrToI(nmbrs[idx]))
	}

	return seeds
}

func parseSeedsWithRanges(lines []string) (seeds []Range) {
	line := lines[0]
	nmbrs := strings.Split(line, " ")

	for idx := 1; idx < len(nmbrs); idx += 2 {
		from := utils.ConvStrToI(nmbrs[idx])
		length := utils.ConvStrToI(nmbrs[idx+1])
		seeds = append(seeds, Range{
			From:   from,
			To:     from + length - 1,
			Length: length,
			Delta:  0,
		})
	}

	return seeds
}
