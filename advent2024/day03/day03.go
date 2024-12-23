package advent202403

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/owendavies93/advent-go/util"
)

type Day struct {
	lines []string
}

var re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func calculateMultiplications(input string) int {
	matches := re.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		a, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			os.Exit(1)
		}
		b, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			os.Exit(1)
		}
		sum += a * b
	}
	return sum
}

func (d *Day) Part1() any {
	sum := 0

	for _, line := range d.lines {
		sum += calculateMultiplications(line)
	}

	return sum
}

func (d *Day) Part2() any {
	all := strings.Join(d.lines, "")

	sum := 0
	split := strings.Split(all, "do()")
	for _, i := range split {
		subsplit := strings.Split(i, "don't()")
		do := subsplit[0]
		sum += calculateMultiplications(do)
	}

	return sum
}

func (d *Day) ParseInput(input string) {
	lines, err := util.ReadStrings(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}
	d.lines = lines
}
