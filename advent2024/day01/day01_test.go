package advent202401

import (
	"testing"

	"github.com/owendavies93/advent-go/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetExampleInputForDayOrExit(2024, 1))
	assert.EqualValues(t, 11, d.Part1())
}

func TestPart2(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetExampleInputForDayOrExit(2024, 1))
	assert.EqualValues(t, 31, d.Part2())
}
