package day20pulsepropagation_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/ewoutquax/advent-of-code-2023/internal/day-20-pulse-propagation"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	var universe Universe = ParseInput(testInput())

	assert := assert.New(t)
	assert.Len(universe.Modules, 6)
	assert.Equal(
		"*day20pulsepropagation.ModuleBroadcaster",
		fmt.Sprintf("%s", reflect.TypeOf(universe.Modules["broadcaster"])),
	)
	assert.Equal(
		"*day20pulsepropagation.ModuleFlipFlop",
		fmt.Sprintf("%s", reflect.TypeOf(universe.Modules["a"])),
	)
	assert.Equal(
		"*day20pulsepropagation.ModuleConjunction",
		fmt.Sprintf("%s", reflect.TypeOf(universe.Modules["con"])),
	)
	moduleInv := (universe.Modules["inv"]).(*ModuleConjunction)
	assert.Len(moduleInv.Destinations, 1)
	ptrInvDestination := moduleInv.Destinations[0].(*ModuleFlipFlop)
	assert.Equal(ModuleName("b"), (*ptrInvDestination).Name)
	assert.Len(moduleInv.Inputs, 1)
}

func TestSendPulseToBroadcast(t *testing.T) {
	u := ParseInput(testInput())

	for idx := 0; idx < 1000; idx++ {
		u.PressButton()
	}

	assert.Equal(t, 11687500, u.NrPulses())
}

func testInput() []string {
	return []string{
		"broadcaster -> a",
		"%a -> inv, con",
		"&inv -> b",
		"%b -> con",
		"&con -> output",
	}
}
