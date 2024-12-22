package advent202405

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day struct {
	rules   map[int][]int
	updates [][]int
}

func isValidUpdate(rules map[int][]int, update []int) bool {
	for i := range update {
		for j := 0; j < i; j++ {
			for _, from := range rules[update[i]] {
				if update[j] == from {
					return false
				}
			}
		}
	}
	return true
}

func (d *Day) Part1() any {
	total := 0
	for _, update := range d.updates {
		if isValidUpdate(d.rules, update) {
			total += update[len(update)/2]
		}
	}

	return total
}

func (d *Day) Part2() any {
	total := 0
	for _, update := range d.updates {
		if !isValidUpdate(d.rules, update) {
			sort.Slice(update, func(i, j int) bool {
				for _, from := range d.rules[update[i]] {
					if update[j] == from {
						return false
					}
				}
				return true
			})
			total += update[len(update)/2]
		}
	}

	return total
}

func (d *Day) ParseInput(input string) {
	f, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening input:", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rules := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		split := strings.Split(line, "|")
		to, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("Error parsing input:", err)
			os.Exit(1)
		}

		from, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println("Error parsing input:", err)
			os.Exit(1)
		}

		rules[to] = append(rules[to], from)
	}

	updates := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		update := []int{}
		for _, s := range split {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error parsing input:", err)
				os.Exit(1)
			}
			update = append(update, n)
		}
		updates = append(updates, update)
	}

	d.rules = rules
	d.updates = updates
}
