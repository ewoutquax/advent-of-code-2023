package day01trebuchet_test

import (
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-01-trebuchet"
	"github.com/stretchr/testify/assert"
)

func TestSumCalibrationValues(t *testing.T) {
	lines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	assert.Equal(t, 142, SumCalibrations(lines, NumbersBase()))
}

func TestCalibrationValueExtended(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(29, ExtractCalibration("two1nine", MatchingsExtend()))
	assert.Equal(14, ExtractCalibration("zoneight234", MatchingsExtend()))
}

func TestCalibrationValue(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(12, ExtractCalibration("1abc2", NumbersBase()))
	assert.Equal(38, ExtractCalibration("pqr3stu8vwx", NumbersBase()))
}
