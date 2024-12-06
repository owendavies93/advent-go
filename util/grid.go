package util

import "fmt"

func PrintGrid(grid [][]rune) {
	if len(grid) == 0 {
		return
	}

	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}
