package advent202408

import (
	"fmt"
	"os"

	"github.com/owendavies93/advent-go/util"
)

type Day struct {
	grid     util.Grid
	antennas map[rune][]util.Point
}

func (d *Day) Part1() any {
	antinodes := make(map[util.Point]int)

	for _, antenna := range d.antennas {
		for _, n1 := range antenna {
			for _, n2 := range antenna {
				if n1.X == n2.X && n1.Y == n2.Y {
					continue
				}

				dx, dy := n1.X-n2.X, n1.Y-n2.Y
				mx, my := n1.X+dx, n1.Y+dy

				if d.grid.IsInBounds(util.Point{X: mx, Y: my}) {
					antinodes[util.Point{X: mx, Y: my}]++
				}

				mx, my = n2.X-dx, n2.Y-dy

				if d.grid.IsInBounds(util.Point{X: mx, Y: my}) {
					antinodes[util.Point{X: mx, Y: my}]++
				}
			}
		}
	}

	return len(antinodes)
}

func (d *Day) Part2() any {
	antinodes := make(map[util.Point]int)

	for _, antenna := range d.antennas {
		for _, n1 := range antenna {
			for _, n2 := range antenna {
				if n1.X == n2.X && n1.Y == n2.Y {
					continue
				}

				antinodes[n1]++
				antinodes[n2]++

				dx, dy := n1.X-n2.X, n1.Y-n2.Y
				mx, my := n1.X+dx, n1.Y+dy

				for d.grid.IsInBounds(util.Point{X: mx, Y: my}) {
					antinodes[util.Point{X: mx, Y: my}]++
					mx += dx
					my += dy
				}

				mx, my = n2.X-dx, n2.Y-dy

				for d.grid.IsInBounds(util.Point{X: mx, Y: my}) {
					antinodes[util.Point{X: mx, Y: my}]++
					mx -= dx
					my -= dy
				}
			}
		}
	}

	return len(antinodes)
}

func (d *Day) ParseInput() {
	grid, err := util.ReadStringsToGrid("inputs/2024/08")
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}

	antennas := make(map[rune][]util.Point)
	height := len(grid)
	width := 0
	for j, row := range grid {
		if len(row) > width {
			width = len(row)
		}
		for i, cell := range row {
			if cell != '.' {
				antennas[cell] = append(antennas[cell], util.Point{X: i, Y: j})
			}
		}
	}

	d.grid = util.Grid{Grid: grid, Height: height, Width: width}
	d.antennas = antennas
}
