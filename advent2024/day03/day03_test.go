package advent202403

import (
	"testing"

	"github.com/owendavies93/advent-go/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetExampleInputForDayOrExit(2024, 3))
	assert.EqualValues(t, 161, d.Part1())
}

func TestPart2(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetNthExampleInputForDayOrExit(2024, 3, 2))
	assert.EqualValues(t, 48, d.Part2())
}
