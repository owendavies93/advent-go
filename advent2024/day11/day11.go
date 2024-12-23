package advent202411

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/owendavies93/advent-go/util"
)

type Day struct {
	stones []int
	cache  map[[2]int]int
}

func (d *Day) Part1() any {
	d.cache = make(map[[2]int]int)
	total := 0
	for _, stone := range d.stones {
		total += d.blink(stone, 25)
	}
	return total
}

func (d *Day) Part2() any {
	d.cache = make(map[[2]int]int)
	total := 0
	for _, stone := range d.stones {
		total += d.blink(stone, 75)
	}
	return total
}

func (d *Day) blink(count int, times int) int {
	if val, ok := d.cache[[2]int{count, times}]; ok {
		return val
	}

	if times == 0 {
		return 1
	}

	if count == 0 {
		return d.blink(1, times-1)
	}

	if len(strconv.Itoa(count))%2 == 0 {
		l := int(math.Log10(float64(count))) + 1
		left := count / int(math.Pow(10, float64(l/2)))
		right := count % int(math.Pow(10, float64(l/2)))
		val := d.blink(left, times-1) + d.blink(right, times-1)
		d.cache[[2]int{count, times}] = val
		return val
	} else {
		val := d.blink(count*2024, times-1)
		d.cache[[2]int{count, times}] = val
		return val
	}
}

func (d *Day) ParseInput(input string) {
	stones, err := util.ReadIntsFromSingleLine(input, " ")
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}
	d.stones = stones
}
