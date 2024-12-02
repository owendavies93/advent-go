package cmd

import (
	advent202401 "github.com/owendavies93/advent-go/advent2024/day01"
	advent202402 "github.com/owendavies93/advent-go/advent2024/day02"
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
		case 2:
			return &advent202402.Day{}
		}
	}

	return nil
}
