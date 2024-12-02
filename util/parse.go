package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInts(path string) ([]int, error) {
	var nums []int

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func ReadIntsFromLines(path string) ([][]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines [][]int
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		if len(nums) == 0 {
			return nil, fmt.Errorf("empty line")
		}
		numsInt := make([]int, len(nums))
		for i, n := range nums {
			numsInt[i], err = strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
		}
		lines = append(lines, numsInt)
	}

	return lines, nil
}

func ReadTwoIntColumns(path string) ([][]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var column1 []int
	var column2 []int
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		if len(nums) != 2 {
			return nil, fmt.Errorf("expected 2 numbers, got %d", len(nums))
		}
		n1, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, err
		}
		n2, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, err
		}
		column1 = append(column1, n1)
		column2 = append(column2, n2)
	}

	return [][]int{column1, column2}, nil
}
