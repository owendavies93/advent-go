package advent202405

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day struct{}

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

func (d *Day) Part1() {
	rules, updates, err := parseInput()
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}

	total := 0
	for _, update := range updates {
		if isValidUpdate(rules, update) {
			total += update[len(update)/2]
		}
	}

	fmt.Println(total)
}

func (d *Day) Part2() {
	rules, updates, err := parseInput()
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}

	total := 0
	for _, update := range updates {
		if !isValidUpdate(rules, update) {
			sort.Slice(update, func(i, j int) bool {
				for _, from := range rules[update[i]] {
					if update[j] == from {
						return false
					}
				}
				return true
			})
			total += update[len(update)/2]
		}
	}

	fmt.Println(total)
}

func parseInput() (map[int][]int, [][]int, error) {
	f, err := os.Open("inputs/2024/05")
	if err != nil {
		return nil, nil, err
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
			return nil, nil, err
		}

		from, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, nil, err
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
				return nil, nil, err
			}
			update = append(update, n)
		}
		updates = append(updates, update)
	}

	return rules, updates, nil
}
