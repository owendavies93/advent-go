package advent202402

import (
	"testing"

	"github.com/owendavies93/advent-go/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetExampleInputForDayOrExit(2024, 2))
	assert.EqualValues(t, 2, d.Part1())
}

func TestPart2(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetExampleInputForDayOrExit(2024, 2))
	assert.EqualValues(t, 4, d.Part2())
}
