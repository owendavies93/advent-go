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

type Point struct {
	X, Y int
}

type Direction struct {
	Dx, Dy int
}

var UP Direction = Direction{0, -1}
var DOWN Direction = Direction{0, 1}
var LEFT Direction = Direction{-1, 0}
var RIGHT Direction = Direction{1, 0}

type PointWithDirection struct {
	Point Point
	Dir   Direction
}

func (pd PointWithDirection) RotateRight() PointWithDirection {
	d := pd.Dir
	switch pd.Dir {
	case UP:
		d = RIGHT
	case DOWN:
		d = LEFT
	case LEFT:
		d = UP
	case RIGHT:
		d = DOWN
	}
	return PointWithDirection{pd.Point, d}
}

func (pd PointWithDirection) MoveInDirection() PointWithDirection {
	return PointWithDirection{Point{pd.Point.X + pd.Dir.Dx, pd.Point.Y + pd.Dir.Dy}, pd.Dir}
}
