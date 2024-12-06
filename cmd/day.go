package cmd

import (
	advent202401 "github.com/owendavies93/advent-go/advent2024/day01"
	advent202402 "github.com/owendavies93/advent-go/advent2024/day02"
	advent202403 "github.com/owendavies93/advent-go/advent2024/day03"
	advent202404 "github.com/owendavies93/advent-go/advent2024/day04"
	advent202405 "github.com/owendavies93/advent-go/advent2024/day05"
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
		case 3:
			return &advent202403.Day{}
		case 4:
			return &advent202404.Day{}
		case 5:
			return &advent202405.Day{}
		}
	}

	return nil
}
