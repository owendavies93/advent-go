package cmd

import (
	"fmt"

	"github.com/owendavies93/advent-go/util"
)

func (opts *Options) Run() {
	if opts.Day == 0 {
		fmt.Println("Running all days")
		for d := 1; d <= 7; d++ {
			fmt.Println("Day", d)
			day := DayFor(opts.Year, d)
			day.ParseInput(util.GetInputForDayOrExit(opts.Year, d))

			fmt.Println("Part 1")
			fmt.Println(day.Part1())

			fmt.Println("Part 2")
			fmt.Println(day.Part2())

			fmt.Println("---")
		}
	} else {
		fmt.Printf("Running %d day %d\n", opts.Year, opts.Day)

		day := DayFor(opts.Year, opts.Day)
		if !opts.Example {
			day.ParseInput(util.GetInputForDayOrExit(opts.Year, opts.Day))
		} else {
			day.ParseInput(util.GetExampleInputForDayOrExit(opts.Year, opts.Day))
		}

		if opts.Part1 {
			fmt.Println("Part 1")
			fmt.Println(day.Part1())
		}
		if opts.Part2 {
			fmt.Println("Part 2")
			fmt.Println(day.Part2())
		}
	}
}
