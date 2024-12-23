package advent202411

import (
	"testing"

	"github.com/owendavies93/advent-go/util"
	"github.com/stretchr/testify/assert"
)

func TestDay(t *testing.T) {
	day := &Day{}
	day.ParseInput(util.GetExampleInputForDayOrExit(2024, 11))
	assert.EqualValues(t, 55312, day.Part1())
}

func TestDay2(t *testing.T) {
	day := &Day{}
	day.ParseInput(util.GetExampleInputForDayOrExit(2024, 11))
	assert.EqualValues(t, 65601038650482, day.Part2())
}
