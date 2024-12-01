package cmd

import (
	advent202401 "github.com/owendavies93/advent-go/advent2024/day01"
)

type Day interface {
	Part1()
	Part2()
}

func DayFor(year int, day int) Day {
	switch year {
	case 2024:
		switch day {
		case 1:
			return &advent202401.Day{}
		}
	}

	return nil
}
