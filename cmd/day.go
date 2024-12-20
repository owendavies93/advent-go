package cmd

import (
	advent202401 "github.com/owendavies93/advent-go/advent2024/day01"
	advent202402 "github.com/owendavies93/advent-go/advent2024/day02"
	advent202403 "github.com/owendavies93/advent-go/advent2024/day03"
	advent202404 "github.com/owendavies93/advent-go/advent2024/day04"
	advent202405 "github.com/owendavies93/advent-go/advent2024/day05"
	advent202406 "github.com/owendavies93/advent-go/advent2024/day06"
	advent202407 "github.com/owendavies93/advent-go/advent2024/day07"
	advent202408 "github.com/owendavies93/advent-go/advent2024/day08"
	advent202409 "github.com/owendavies93/advent-go/advent2024/day09"
)

type Day interface {
	Part1() any
	Part2() any
	ParseInput()
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
		case 6:
			return &advent202406.Day{}
		case 7:
			return &advent202407.Day{}
		case 8:
			return &advent202408.Day{}
		case 9:
			return &advent202409.Day{}
		}
	}

	return nil
}
