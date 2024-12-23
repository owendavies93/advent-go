package advent202402

import (
	"fmt"
	"math"
	"os"

	"github.com/owendavies93/advent-go/util"
)

type Day struct {
	nums [][]int
}

func (d *Day) Part1() any {
	count := 0
	for _, line := range d.nums {
		if IsValidLine(line) {
			count++
		}
	}

	return count
}

func (d *Day) Part2() any {
	count := 0
Outer:
	for _, line := range d.nums {
		if IsValidLine(line) {
			count++
			continue
		}

		for i := 0; i < len(line)-1; i++ {
			new := make([]int, len(line))
			copy(new, line)
			splice := SpliceOne(i, new)
			if IsValidLine(splice) {
				count++
				continue Outer
			}
		}

		new := make([]int, len(line))
		copy(new, line)
		endSplice := SpliceOne(len(line)-1, new)
		if IsValidLine(endSplice) {
			count++
		}
	}

	return count
}

func (d *Day) ParseInput(input string) {
	nums, err := util.ReadIntsFromLines(input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	d.nums = nums
}

func IsValidLine(line []int) bool {
	is_descending := line[0] > line[1]
	valid := true
	for i := 0; i < len(line)-1; i++ {
		n := line[i]
		m := line[i+1]
		if is_descending && n <= m {
			valid = false
			break
		} else if !is_descending && n >= m {
			valid = false
			break
		} else if math.Abs(float64(n-m)) > 3 {
			valid = false
			break
		}
	}

	return valid
}

func SpliceOne(offset int, items []int) []int {
	return append(items[:offset], items[offset+1:]...)
}
