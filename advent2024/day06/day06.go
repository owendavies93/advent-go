package advent202406

import (
	"fmt"
	"os"

	"github.com/owendavies93/advent-go/util"
)

type Day struct{}

func (d *Day) Part1() {
	grid, start, err := parseInput()
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}

	seen := map[util.Point]bool{}
	seen[start.Point] = true

	for {
		next := start.MoveInDirection()

		if grid.IsInBounds(next.Point) {
			if grid.Get(next.Point) == '#' {
				start = start.RotateRight()
			} else {
				start = next
				seen[next.Point] = true
			}
		} else {
			break
		}
	}

	fmt.Println(len(seen))
}

func (d *Day) Part2() {
	grid, start, err := parseInput()
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}

	total := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if grid.Get(util.Point{x, y}) == '#' {
				continue
			}

			grid.Set(util.Point{x, y}, '#')

			seen := map[util.PointWithDirection]bool{}
			seen[start] = true
			pos := start

			for {
				next := pos.MoveInDirection()
				if grid.IsInBounds(next.Point) {
					if grid.Get(next.Point) == '#' {
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
			grid.Set(util.Point{x, y}, '.')
		}
	}

	fmt.Println(total)
}

func parseInput() (util.Grid, util.PointWithDirection, error) {
	grid, err := util.ReadStringsToGrid("inputs/2024/06")
	if err != nil {
		return util.Grid{}, util.PointWithDirection{}, err
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
	return util.Grid{grid, height, width}, util.PointWithDirection{util.Point{startx, starty}, util.UP}, nil
}
