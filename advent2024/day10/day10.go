package advent202410

import (
	"fmt"
	"os"

	"github.com/owendavies93/advent-go/util"
)

type Day struct {
	grid util.Grid
}

func (d *Day) Part1() any {
	total := 0
	for y, row := range d.grid.Grid {
		for x := range row {
			if d.grid.Get(util.Point{X: x, Y: y}) != '0' {
				continue
			}

			ends := util.BFSManyEnds(d.grid, util.Point{X: x, Y: y}, d.isTop, d.getNeighbours)

			endsMap := make(map[util.BFSResult]struct{})
			for _, e := range ends {
				endsMap[e] = struct{}{}
			}

			total += len(endsMap)
		}
	}

	return total
}

func (d *Day) Part2() any {
	total := 0
	for y, row := range d.grid.Grid {
		for x := range row {
			if d.grid.Get(util.Point{X: x, Y: y}) != '0' {
				continue
			}

			ends := util.BFSManyEnds(d.grid, util.Point{X: x, Y: y}, d.isTop, d.getNeighbours)

			total += len(ends)
		}
	}
	return total
}

func (d *Day) isTop(c util.BFSResult) bool {
	return c.Distance == 9
}

func (d *Day) getNeighbours(c util.BFSResult) []util.Point {
	ns := []util.Point{}
	for _, dir := range util.D4() {
		np := c.Point.Add(util.Point{X: dir[0], Y: dir[1]})
		if d.grid.IsInBounds(np) {
			if int(d.grid.Get(np)-'0') == c.Distance+1 {
				ns = append(ns, np)
			}
		}
	}
	return ns
}

func (d *Day) ParseInput(input string) {
	grid, err := util.ReadStringsToGrid(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}

	d.grid = util.Grid{Grid: grid, Height: len(grid), Width: len(grid[0])}
}
