package day01

import (
	"testing"

	"github.com/owendavies93/advent-go/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetExampleInputForDayOrExit(2024, 1))
	assert.EqualValues(t, d.Part1(), 11)
}

func TestPart2(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetExampleInputForDayOrExit(2024, 1))
	assert.EqualValues(t, d.Part2(), 31)
}
