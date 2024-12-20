package cmd

import (
	"fmt"
)

func (opts *Options) Run() {
	if opts.Day == 0 {
		fmt.Println("Running all days")
		for day := 1; day <= 7; day++ {
			fmt.Println("Day", day)
			day := DayFor(opts.Year, day)
			day.ParseInput()

			fmt.Println("Part 1")
			fmt.Println(day.Part1())

			fmt.Println("Part 2")
			fmt.Println(day.Part2())

			fmt.Println("---")
		}
	} else {
		fmt.Printf("Running %d day %d\n", opts.Year, opts.Day)

		day := DayFor(opts.Year, opts.Day)
		day.ParseInput()

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
