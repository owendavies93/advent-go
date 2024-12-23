package advent202409

import (
	"testing"

	"github.com/owendavies93/advent-go/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetExampleInputForDayOrExit(2024, 9))
	assert.EqualValues(t, 1928, d.Part1())
}

func TestPart2(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetExampleInputForDayOrExit(2024, 9))
	assert.EqualValues(t, 2858, d.Part2())
}
