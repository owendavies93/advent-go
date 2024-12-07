package day01

import (
	"fmt"
	"os"
	"sort"

	"github.com/owendavies93/advent-go/util"
)

type Day struct {
	nums [][]int
}

func (d *Day) Part1() any {
	sort.Ints(d.nums[0])
	sort.Ints(d.nums[1])

	score := 0
	for i, n := range d.nums[0] {
		if n > d.nums[1][i] {
			score += n - d.nums[1][i]
		} else {
			score += d.nums[1][i] - n
		}
	}

	return score
}

func (d *Day) Part2() any {
	freq := make(map[int]int)
	for _, n := range d.nums[1] {
		freq[n]++
	}

	score := 0
	for _, n := range d.nums[0] {
		if val, ok := freq[n]; ok {
			score += n * val
		}
	}

	return score
}

func (d *Day) ParseInput() {
	nums, err := util.ReadTwoIntColumns("inputs/2024/01")
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	d.nums = nums
}
