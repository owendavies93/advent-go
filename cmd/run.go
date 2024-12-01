package cmd

import (
	"fmt"
)

func (opts *Options) Run() {
	fmt.Printf("Running %d day %d\n", opts.Year, opts.Day)

	day := DayFor(opts.Year, opts.Day)

	if opts.Part1 {
		fmt.Println("Part 1")
		day.Part1()
	}
	if opts.Part2 {
		fmt.Println("Part 2")
		day.Part2()
	}
}
