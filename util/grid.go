package util

import "fmt"

type Grid struct {
	Grid   [][]rune
	Height int
	Width  int
}

func (g Grid) Get(p Point) rune {
	return g.Grid[p.Y][p.X]
}

func (g Grid) IsInBounds(p Point) bool {
	return p.Y >= 0 && p.Y < g.Height && p.X >= 0 && p.X < g.Width
}

func (g Grid) Print() {
	if len(g.Grid) == 0 {
		return
	}

	for _, row := range g.Grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}

func (g Grid) Set(p Point, r rune) {
	g.Grid[p.Y][p.X] = r
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

func D4() [][]int {
	return [][]int{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}
}
