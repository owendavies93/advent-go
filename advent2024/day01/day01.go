package day01

import (
	"fmt"
	"sort"

	"github.com/owendavies93/advent-go/util"
)

type Day struct{}

func (d *Day) Part1() {
	nums, err := util.ReadTwoIntColumns("inputs/2024/01")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	sort.Ints(nums[0])
	sort.Ints(nums[1])

	score := 0
	for i, n := range nums[0] {
		if n > nums[1][i] {
			score += n - nums[1][i]
		} else {
			score += nums[1][i] - n
		}
	}

	fmt.Println(score)
}

func (d *Day) Part2() {
	nums, err := util.ReadTwoIntColumns("inputs/2024/01")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	freq := make(map[int]int)
	for _, n := range nums[1] {
		freq[n]++
	}

	score := 0
	for _, n := range nums[0] {
		if val, ok := freq[n]; ok {
			score += n * val
		}
	}

	fmt.Println(score)
}
