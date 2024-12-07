package main

import (
	"fmt"
	"os"

	"github.com/owendavies93/advent-go/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <year> [<day>] [<part>]")
		os.Exit(1)
	}

	opts := cmd.GetOptions()

	opts.Run()
}
