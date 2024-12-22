package advent202406

import (
	"fmt"
	"os"

	"github.com/owendavies93/advent-go/util"
)

type Day struct {
	grid  util.Grid
	start util.PointWithDirection
}

func (d *Day) Part1() any {
	seen := map[util.Point]bool{}
	temp := d.start
	seen[temp.Point] = true

	for {
		next := temp.MoveInDirection()

		if d.grid.IsInBounds(next.Point) {
			if d.grid.Get(next.Point) == '#' {
				temp = temp.RotateRight()
			} else {
				temp = next
				seen[next.Point] = true
			}
		} else {
			break
		}
	}

	return len(seen)
}

func (d *Day) Part2() any {
	total := 0
	for y := 0; y < d.grid.Height; y++ {
		for x := 0; x < d.grid.Width; x++ {
			if d.grid.Get(util.Point{X: x, Y: y}) == '#' {
				continue
			}

			d.grid.Set(util.Point{X: x, Y: y}, '#')

			seen := map[util.PointWithDirection]bool{}
			seen[d.start] = true
			pos := d.start

			for {
				next := pos.MoveInDirection()
				if d.grid.IsInBounds(next.Point) {
					if d.grid.Get(next.Point) == '#' {
						pos = pos.RotateRight()
					} else {
						pos = next
					}

					if seen[pos] {
						total++
						break
					}
					seen[pos] = true
				} else {
					break
				}
			}
			d.grid.Set(util.Point{X: x, Y: y}, '.')
		}
	}

	return total
}

func (d *Day) ParseInput(input string) {
	grid, err := util.ReadStringsToGrid(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}

	startx, starty := 0, 0
	height := len(grid)
	width := 0
	for j, row := range grid {
		if len(row) > width {
			width = len(row)
		}
		for i, cell := range row {
			if cell == '^' {
				startx = i
				starty = j
			}
		}
	}
	d.grid = util.Grid{Grid: grid, Height: height, Width: width}
	d.start = util.PointWithDirection{Point: util.Point{X: startx, Y: starty}, Dir: util.UP}
}
