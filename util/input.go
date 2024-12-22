package util

import (
	"errors"
	"fmt"
	"os"
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
