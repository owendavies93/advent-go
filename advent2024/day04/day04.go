package day04

import (
	"fmt"
	"os"

	"github.com/owendavies93/advent-go/util"
)

type Day struct{}

func (d *Day) Part1() {
	grid, err := d.ParseInput()
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}

	sum := 0
	for y, row := range grid {
		for x := range row {
			for _, dir := range util.D8() {
				dy := dir[0]
				dx := dir[1]
				str := ""
				for i := 0; i < 4; i++ {
					ny := y + dy*i
					nx := x + dx*i
					if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(row) {
						break
					}
					str += string(grid[ny][nx])
				}
				if str == "XMAS" {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func (d *Day) Part2() {
	grid, err := d.ParseInput()
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}

	sum := 0
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if grid[y][x] == 'A' {
				tl := grid[y-1][x-1]
				tr := grid[y-1][x+1]
				bl := grid[y+1][x-1]
				br := grid[y+1][x+1]

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
	fmt.Println(sum)
}

func (d *Day) ParseInput() ([][]rune, error) {
	return util.ReadStringsToGrid("inputs/2024/04")
}
