package util

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func GetInput(year int, day int) (string, error) {
	file := fmt.Sprintf("inputs/%d/%02d", year, day)

	if _, err := os.Stat(file); err == nil {
		return file, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("file does not exist")
	} else {
		return "", fmt.Errorf("error getting input: %w", err)
	}
}

func GetExampleInput(year int, day int) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current directory: %w", err)
	}

	for dir := currentDir; dir != "/"; dir = filepath.Dir(dir) {
		file := filepath.Join(dir, fmt.Sprintf("inputs/%d/%02d-test", year, day))
		if _, err := os.Stat(file); err == nil {
			return file, nil
		} else if !errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("error checking file: %w", err)
		}
	}

	return "", fmt.Errorf("example input does not exist in any parent directory")
}

func GetInputForDayOrExit(year int, day int) string {
	input, err := GetInput(year, day)
	if err != nil {
		fmt.Println("Error getting input:", err)
		os.Exit(1)
	}
	return input
}

func GetExampleInputForDayOrExit(year int, day int) string {
	input, err := GetExampleInput(year, day)
	if err != nil {
		fmt.Println("Error getting example input:", err)
		os.Exit(1)
	}
	return input
}
