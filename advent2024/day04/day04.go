package advent202404

import (
	"fmt"
	"os"

	"github.com/owendavies93/advent-go/util"
)

type Day struct {
	grid [][]rune
}

func (d *Day) Part1() any {
	sum := 0
	for y, row := range d.grid {
		for x := range row {
			for _, dir := range util.D8() {
				dy := dir[0]
				dx := dir[1]
				str := ""
				for i := 0; i < 4; i++ {
					ny := y + dy*i
					nx := x + dx*i
					if ny < 0 || ny >= len(d.grid) || nx < 0 || nx >= len(row) {
						break
					}
					str += string(d.grid[ny][nx])
				}
				if str == "XMAS" {
					sum++
				}
			}
		}
	}

	return sum
}

func (d *Day) Part2() any {
	sum := 0
	for y := 1; y < len(d.grid)-1; y++ {
		for x := 1; x < len(d.grid[y])-1; x++ {
			if d.grid[y][x] == 'A' {
				tl := d.grid[y-1][x-1]
				tr := d.grid[y-1][x+1]
				bl := d.grid[y+1][x-1]
				br := d.grid[y+1][x+1]

				if tl == 'M' && tr == 'M' && bl == 'S' && br == 'S' {
					sum++
				}
				if tl == 'S' && tr == 'S' && bl == 'M' && br == 'M' {
					sum++
				}
				if tl == 'S' && tr == 'M' && bl == 'S' && br == 'M' {
					sum++
				}
				if tl == 'M' && tr == 'S' && bl == 'M' && br == 'S' {
					sum++
				}
			}
		}
	}

	return sum
}

func (d *Day) ParseInput(input string) {
	grid, err := util.ReadStringsToGrid(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}
	d.grid = grid
}
