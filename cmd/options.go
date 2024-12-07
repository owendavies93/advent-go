package cmd

import (
	"fmt"
	"os"
	"strconv"
)

type Options struct {
	Year  int
	Day   int
	Part1 bool
	Part2 bool
}

func GetOptions() *Options {
	opts := Options{Part1: true, Part2: true}

	if len(os.Args[1]) == 4 {
		year, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid year")
			os.Exit(1)
		}
		opts.Year = year
	} else {
		fmt.Println("Invalid year")
		os.Exit(1)
	}

	if len(os.Args) >= 3 && len(os.Args[2]) <= 2 {
		day, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid day")
			os.Exit(1)
		}
		opts.Day = day
	} else if len(os.Args) >= 3 {
		fmt.Println("Invalid day")
		os.Exit(1)
	}

	if opts.Part1 && opts.Part2 && len(os.Args) >= 4 {
		switch os.Args[3] {
		case "1":
			opts.Part2 = false
		case "2":
			opts.Part1 = false
		default:
			fmt.Println("Invalid part")
			os.Exit(1)
		}
	}

	return &opts
}
