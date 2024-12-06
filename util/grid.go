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

func D8() [][]int {
	return [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
}
