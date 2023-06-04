package shared

import (
	"bufio"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	lines := make([]string, 0)

	f, err := os.Open(filePath)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
