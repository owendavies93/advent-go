package advent202405

import (
	"testing"

	"github.com/owendavies93/advent-go/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetExampleInputForDayOrExit(2024, 5))
	assert.EqualValues(t, 143, d.Part1())
}

func TestPart2(t *testing.T) {
	d := Day{}
	d.ParseInput(util.GetExampleInputForDayOrExit(2024, 5))
	assert.EqualValues(t, 123, d.Part2())
}
