package filereader

import (
	"bufio"
	"os"
)

// ReadLines reads all lines from a specified file
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for any errors while reading the file
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
