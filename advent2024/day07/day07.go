package advent202407

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/owendavies93/advent-go/util"
)

type Day struct {
	eqs []Equation
}

type Equation struct {
	Target int
	Ops    []int
}

func (d *Day) Part1() any {
	total := 0
	for _, eq := range d.eqs {
		if eval(0, eq.Ops, 0, eq.Target) {
			total += eq.Target
		}
	}

	return total
}

func eval(i int, ops []int, val int, target int) bool {
	if val > target {
		return false
	}

	if i == len(ops) {
		return val == target
	}

	return eval(i+1, ops, val+ops[i], target) || eval(i+1, ops, val*ops[i], target)
}

func evalWithConcat(i int, ops []int, val int, target int) bool {
	if val > target {
		return false
	}

	if i == len(ops) {
		return val == target
	}

	if evalWithConcat(i+1, ops, val+ops[i], target) || evalWithConcat(i+1, ops, val*ops[i], target) {
		return true
	}

	a := strconv.Itoa(val) + strconv.Itoa(ops[i])
	b, err := strconv.Atoi(a)
	if err != nil {
		return false
	}

	return evalWithConcat(i+1, ops, b, target)
}

func (d *Day) Part2() any {
	total := 0
	for _, eq := range d.eqs {
		if evalWithConcat(0, eq.Ops, 0, eq.Target) {
			total += eq.Target
		}
	}

	return total
}

func (d *Day) ParseInput(input string) {
	lines, err := util.ReadStrings(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}

	res := []Equation{}
	for _, line := range lines {
		t := strings.Split(line, ": ")[0]
		target, err := strconv.Atoi(t)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			os.Exit(1)
		}

		rest := strings.Split(line, ": ")[1]

		nums := strings.Split(rest, " ")
		numsInt := make([]int, len(nums))
		for i, n := range nums {
			numsInt[i], err = strconv.Atoi(n)
			if err != nil {
				fmt.Println("Error parsing input:", err)
				os.Exit(1)
			}
		}

		res = append(res, Equation{Target: target, Ops: numsInt})
	}

	d.eqs = res
}
